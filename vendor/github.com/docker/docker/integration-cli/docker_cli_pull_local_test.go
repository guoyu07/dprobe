package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest"
	"github.com/docker/distribution/manifest/manifestlist"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/docker/docker/integration-cli/checker"
	"github.com/docker/docker/integration-cli/cli/build"
	"github.com/go-check/check"
	"github.com/gotestyourself/gotestyourself/icmd"
	"github.com/opencontainers/go-digest"
)

// testPullImageWithAliases pulls a specific image tag and verifies that any aliases (i.e., other
// tags for the same image) are not also pulled down.
//
// Ref: docker/docker#8141
func testPullImageWithAliases(c *check.C) ***REMOVED***
	repoName := fmt.Sprintf("%v/dockercli/busybox", privateRegistryURL)

	repos := []string***REMOVED******REMOVED***
	for _, tag := range []string***REMOVED***"recent", "fresh"***REMOVED*** ***REMOVED***
		repos = append(repos, fmt.Sprintf("%v:%v", repoName, tag))
	***REMOVED***

	// Tag and push the same image multiple times.
	for _, repo := range repos ***REMOVED***
		dockerCmd(c, "tag", "busybox", repo)
		dockerCmd(c, "push", repo)
	***REMOVED***

	// Clear local images store.
	args := append([]string***REMOVED***"rmi"***REMOVED***, repos...)
	dockerCmd(c, args...)

	// Pull a single tag and verify it doesn't bring down all aliases.
	dockerCmd(c, "pull", repos[0])
	dockerCmd(c, "inspect", repos[0])
	for _, repo := range repos[1:] ***REMOVED***
		_, _, err := dockerCmdWithError("inspect", repo)
		c.Assert(err, checker.NotNil, check.Commentf("Image %v shouldn't have been pulled down", repo))
	***REMOVED***
***REMOVED***

func (s *DockerRegistrySuite) TestPullImageWithAliases(c *check.C) ***REMOVED***
	testPullImageWithAliases(c)
***REMOVED***

func (s *DockerSchema1RegistrySuite) TestPullImageWithAliases(c *check.C) ***REMOVED***
	testPullImageWithAliases(c)
***REMOVED***

// testConcurrentPullWholeRepo pulls the same repo concurrently.
func testConcurrentPullWholeRepo(c *check.C) ***REMOVED***
	repoName := fmt.Sprintf("%v/dockercli/busybox", privateRegistryURL)

	repos := []string***REMOVED******REMOVED***
	for _, tag := range []string***REMOVED***"recent", "fresh", "todays"***REMOVED*** ***REMOVED***
		repo := fmt.Sprintf("%v:%v", repoName, tag)
		buildImageSuccessfully(c, repo, build.WithDockerfile(fmt.Sprintf(`
		    FROM busybox
		    ENTRYPOINT ["/bin/echo"]
		    ENV FOO foo
		    ENV BAR bar
		    CMD echo %s
		`, repo)))
		dockerCmd(c, "push", repo)
		repos = append(repos, repo)
	***REMOVED***

	// Clear local images store.
	args := append([]string***REMOVED***"rmi"***REMOVED***, repos...)
	dockerCmd(c, args...)

	// Run multiple re-pulls concurrently
	results := make(chan error)
	numPulls := 3

	for i := 0; i != numPulls; i++ ***REMOVED***
		go func() ***REMOVED***
			result := icmd.RunCommand(dockerBinary, "pull", "-a", repoName)
			results <- result.Error
		***REMOVED***()
	***REMOVED***

	// These checks are separate from the loop above because the check
	// package is not goroutine-safe.
	for i := 0; i != numPulls; i++ ***REMOVED***
		err := <-results
		c.Assert(err, checker.IsNil, check.Commentf("concurrent pull failed with error: %v", err))
	***REMOVED***

	// Ensure all tags were pulled successfully
	for _, repo := range repos ***REMOVED***
		dockerCmd(c, "inspect", repo)
		out, _ := dockerCmd(c, "run", "--rm", repo)
		c.Assert(strings.TrimSpace(out), checker.Equals, "/bin/sh -c echo "+repo)
	***REMOVED***
***REMOVED***

func (s *DockerRegistrySuite) testConcurrentPullWholeRepo(c *check.C) ***REMOVED***
	testConcurrentPullWholeRepo(c)
***REMOVED***

func (s *DockerSchema1RegistrySuite) testConcurrentPullWholeRepo(c *check.C) ***REMOVED***
	testConcurrentPullWholeRepo(c)
***REMOVED***

// testConcurrentFailingPull tries a concurrent pull that doesn't succeed.
func testConcurrentFailingPull(c *check.C) ***REMOVED***
	repoName := fmt.Sprintf("%v/dockercli/busybox", privateRegistryURL)

	// Run multiple pulls concurrently
	results := make(chan error)
	numPulls := 3

	for i := 0; i != numPulls; i++ ***REMOVED***
		go func() ***REMOVED***
			result := icmd.RunCommand(dockerBinary, "pull", repoName+":asdfasdf")
			results <- result.Error
		***REMOVED***()
	***REMOVED***

	// These checks are separate from the loop above because the check
	// package is not goroutine-safe.
	for i := 0; i != numPulls; i++ ***REMOVED***
		err := <-results
		c.Assert(err, checker.NotNil, check.Commentf("expected pull to fail"))
	***REMOVED***
***REMOVED***

func (s *DockerRegistrySuite) testConcurrentFailingPull(c *check.C) ***REMOVED***
	testConcurrentFailingPull(c)
***REMOVED***

func (s *DockerSchema1RegistrySuite) testConcurrentFailingPull(c *check.C) ***REMOVED***
	testConcurrentFailingPull(c)
***REMOVED***

// testConcurrentPullMultipleTags pulls multiple tags from the same repo
// concurrently.
func testConcurrentPullMultipleTags(c *check.C) ***REMOVED***
	repoName := fmt.Sprintf("%v/dockercli/busybox", privateRegistryURL)

	repos := []string***REMOVED******REMOVED***
	for _, tag := range []string***REMOVED***"recent", "fresh", "todays"***REMOVED*** ***REMOVED***
		repo := fmt.Sprintf("%v:%v", repoName, tag)
		buildImageSuccessfully(c, repo, build.WithDockerfile(fmt.Sprintf(`
		    FROM busybox
		    ENTRYPOINT ["/bin/echo"]
		    ENV FOO foo
		    ENV BAR bar
		    CMD echo %s
		`, repo)))
		dockerCmd(c, "push", repo)
		repos = append(repos, repo)
	***REMOVED***

	// Clear local images store.
	args := append([]string***REMOVED***"rmi"***REMOVED***, repos...)
	dockerCmd(c, args...)

	// Re-pull individual tags, in parallel
	results := make(chan error)

	for _, repo := range repos ***REMOVED***
		go func(repo string) ***REMOVED***
			result := icmd.RunCommand(dockerBinary, "pull", repo)
			results <- result.Error
		***REMOVED***(repo)
	***REMOVED***

	// These checks are separate from the loop above because the check
	// package is not goroutine-safe.
	for range repos ***REMOVED***
		err := <-results
		c.Assert(err, checker.IsNil, check.Commentf("concurrent pull failed with error: %v", err))
	***REMOVED***

	// Ensure all tags were pulled successfully
	for _, repo := range repos ***REMOVED***
		dockerCmd(c, "inspect", repo)
		out, _ := dockerCmd(c, "run", "--rm", repo)
		c.Assert(strings.TrimSpace(out), checker.Equals, "/bin/sh -c echo "+repo)
	***REMOVED***
***REMOVED***

func (s *DockerRegistrySuite) TestConcurrentPullMultipleTags(c *check.C) ***REMOVED***
	testConcurrentPullMultipleTags(c)
***REMOVED***

func (s *DockerSchema1RegistrySuite) TestConcurrentPullMultipleTags(c *check.C) ***REMOVED***
	testConcurrentPullMultipleTags(c)
***REMOVED***

// testPullIDStability verifies that pushing an image and pulling it back
// preserves the image ID.
func testPullIDStability(c *check.C) ***REMOVED***
	derivedImage := privateRegistryURL + "/dockercli/id-stability"
	baseImage := "busybox"

	buildImageSuccessfully(c, derivedImage, build.WithDockerfile(fmt.Sprintf(`
	    FROM %s
	    ENV derived true
	    ENV asdf true
	    RUN dd if=/dev/zero of=/file bs=1024 count=1024
	    CMD echo %s
	`, baseImage, derivedImage)))

	originalID := getIDByName(c, derivedImage)
	dockerCmd(c, "push", derivedImage)

	// Pull
	out, _ := dockerCmd(c, "pull", derivedImage)
	if strings.Contains(out, "Pull complete") ***REMOVED***
		c.Fatalf("repull redownloaded a layer: %s", out)
	***REMOVED***

	derivedIDAfterPull := getIDByName(c, derivedImage)

	if derivedIDAfterPull != originalID ***REMOVED***
		c.Fatal("image's ID unexpectedly changed after a repush/repull")
	***REMOVED***

	// Make sure the image runs correctly
	out, _ = dockerCmd(c, "run", "--rm", derivedImage)
	if strings.TrimSpace(out) != derivedImage ***REMOVED***
		c.Fatalf("expected %s; got %s", derivedImage, out)
	***REMOVED***

	// Confirm that repushing and repulling does not change the computed ID
	dockerCmd(c, "push", derivedImage)
	dockerCmd(c, "rmi", derivedImage)
	dockerCmd(c, "pull", derivedImage)

	derivedIDAfterPull = getIDByName(c, derivedImage)

	if derivedIDAfterPull != originalID ***REMOVED***
		c.Fatal("image's ID unexpectedly changed after a repush/repull")
	***REMOVED***

	// Make sure the image still runs
	out, _ = dockerCmd(c, "run", "--rm", derivedImage)
	if strings.TrimSpace(out) != derivedImage ***REMOVED***
		c.Fatalf("expected %s; got %s", derivedImage, out)
	***REMOVED***
***REMOVED***

func (s *DockerRegistrySuite) TestPullIDStability(c *check.C) ***REMOVED***
	testPullIDStability(c)
***REMOVED***

func (s *DockerSchema1RegistrySuite) TestPullIDStability(c *check.C) ***REMOVED***
	testPullIDStability(c)
***REMOVED***

// #21213
func testPullNoLayers(c *check.C) ***REMOVED***
	repoName := fmt.Sprintf("%v/dockercli/scratch", privateRegistryURL)

	buildImageSuccessfully(c, repoName, build.WithDockerfile(`
	FROM scratch
	ENV foo bar`))
	dockerCmd(c, "push", repoName)
	dockerCmd(c, "rmi", repoName)
	dockerCmd(c, "pull", repoName)
***REMOVED***

func (s *DockerRegistrySuite) TestPullNoLayers(c *check.C) ***REMOVED***
	testPullNoLayers(c)
***REMOVED***

func (s *DockerSchema1RegistrySuite) TestPullNoLayers(c *check.C) ***REMOVED***
	testPullNoLayers(c)
***REMOVED***

func (s *DockerRegistrySuite) TestPullManifestList(c *check.C) ***REMOVED***
	testRequires(c, NotArm)
	pushDigest, err := setupImage(c)
	c.Assert(err, checker.IsNil, check.Commentf("error setting up image"))

	// Inject a manifest list into the registry
	manifestList := &manifestlist.ManifestList***REMOVED***
		Versioned: manifest.Versioned***REMOVED***
			SchemaVersion: 2,
			MediaType:     manifestlist.MediaTypeManifestList,
		***REMOVED***,
		Manifests: []manifestlist.ManifestDescriptor***REMOVED***
			***REMOVED***
				Descriptor: distribution.Descriptor***REMOVED***
					Digest:    "sha256:1a9ec845ee94c202b2d5da74a24f0ed2058318bfa9879fa541efaecba272e86b",
					Size:      3253,
					MediaType: schema2.MediaTypeManifest,
				***REMOVED***,
				Platform: manifestlist.PlatformSpec***REMOVED***
					Architecture: "bogus_arch",
					OS:           "bogus_os",
				***REMOVED***,
			***REMOVED***,
			***REMOVED***
				Descriptor: distribution.Descriptor***REMOVED***
					Digest:    pushDigest,
					Size:      3253,
					MediaType: schema2.MediaTypeManifest,
				***REMOVED***,
				Platform: manifestlist.PlatformSpec***REMOVED***
					Architecture: runtime.GOARCH,
					OS:           runtime.GOOS,
				***REMOVED***,
			***REMOVED***,
		***REMOVED***,
	***REMOVED***

	manifestListJSON, err := json.MarshalIndent(manifestList, "", "   ")
	c.Assert(err, checker.IsNil, check.Commentf("error marshalling manifest list"))

	manifestListDigest := digest.FromBytes(manifestListJSON)
	hexDigest := manifestListDigest.Hex()

	registryV2Path := s.reg.Path()

	// Write manifest list to blob store
	blobDir := filepath.Join(registryV2Path, "blobs", "sha256", hexDigest[:2], hexDigest)
	err = os.MkdirAll(blobDir, 0755)
	c.Assert(err, checker.IsNil, check.Commentf("error creating blob dir"))
	blobPath := filepath.Join(blobDir, "data")
	err = ioutil.WriteFile(blobPath, []byte(manifestListJSON), 0644)
	c.Assert(err, checker.IsNil, check.Commentf("error writing manifest list"))

	// Add to revision store
	revisionDir := filepath.Join(registryV2Path, "repositories", remoteRepoName, "_manifests", "revisions", "sha256", hexDigest)
	err = os.Mkdir(revisionDir, 0755)
	c.Assert(err, checker.IsNil, check.Commentf("error creating revision dir"))
	revisionPath := filepath.Join(revisionDir, "link")
	err = ioutil.WriteFile(revisionPath, []byte(manifestListDigest.String()), 0644)
	c.Assert(err, checker.IsNil, check.Commentf("error writing revision link"))

	// Update tag
	tagPath := filepath.Join(registryV2Path, "repositories", remoteRepoName, "_manifests", "tags", "latest", "current", "link")
	err = ioutil.WriteFile(tagPath, []byte(manifestListDigest.String()), 0644)
	c.Assert(err, checker.IsNil, check.Commentf("error writing tag link"))

	// Verify that the image can be pulled through the manifest list.
	out, _ := dockerCmd(c, "pull", repoName)

	// The pull output includes "Digest: <digest>", so find that
	matches := digestRegex.FindStringSubmatch(out)
	c.Assert(matches, checker.HasLen, 2, check.Commentf("unable to parse digest from pull output: %s", out))
	pullDigest := matches[1]

	// Make sure the pushed and pull digests match
	c.Assert(manifestListDigest.String(), checker.Equals, pullDigest)

	// Was the image actually created?
	dockerCmd(c, "inspect", repoName)

	dockerCmd(c, "rmi", repoName)
***REMOVED***

// #23100
func (s *DockerRegistryAuthHtpasswdSuite) TestPullWithExternalAuthLoginWithScheme(c *check.C) ***REMOVED***
	osPath := os.Getenv("PATH")
	defer os.Setenv("PATH", osPath)

	workingDir, err := os.Getwd()
	c.Assert(err, checker.IsNil)
	absolute, err := filepath.Abs(filepath.Join(workingDir, "fixtures", "auth"))
	c.Assert(err, checker.IsNil)
	testPath := fmt.Sprintf("%s%c%s", osPath, filepath.ListSeparator, absolute)

	os.Setenv("PATH", testPath)

	repoName := fmt.Sprintf("%v/dockercli/busybox:authtest", privateRegistryURL)

	tmp, err := ioutil.TempDir("", "integration-cli-")
	c.Assert(err, checker.IsNil)

	externalAuthConfig := `***REMOVED*** "credsStore": "shell-test" ***REMOVED***`

	configPath := filepath.Join(tmp, "config.json")
	err = ioutil.WriteFile(configPath, []byte(externalAuthConfig), 0644)
	c.Assert(err, checker.IsNil)

	dockerCmd(c, "--config", tmp, "login", "-u", s.reg.Username(), "-p", s.reg.Password(), privateRegistryURL)

	b, err := ioutil.ReadFile(configPath)
	c.Assert(err, checker.IsNil)
	c.Assert(string(b), checker.Not(checker.Contains), "\"auth\":")

	dockerCmd(c, "--config", tmp, "tag", "busybox", repoName)
	dockerCmd(c, "--config", tmp, "push", repoName)

	dockerCmd(c, "--config", tmp, "logout", privateRegistryURL)
	dockerCmd(c, "--config", tmp, "login", "-u", s.reg.Username(), "-p", s.reg.Password(), "https://"+privateRegistryURL)
	dockerCmd(c, "--config", tmp, "pull", repoName)

	// likewise push should work
	repoName2 := fmt.Sprintf("%v/dockercli/busybox:nocreds", privateRegistryURL)
	dockerCmd(c, "tag", repoName, repoName2)
	dockerCmd(c, "--config", tmp, "push", repoName2)

	// logout should work w scheme also because it will be stripped
	dockerCmd(c, "--config", tmp, "logout", "https://"+privateRegistryURL)
***REMOVED***

func (s *DockerRegistryAuthHtpasswdSuite) TestPullWithExternalAuth(c *check.C) ***REMOVED***
	osPath := os.Getenv("PATH")
	defer os.Setenv("PATH", osPath)

	workingDir, err := os.Getwd()
	c.Assert(err, checker.IsNil)
	absolute, err := filepath.Abs(filepath.Join(workingDir, "fixtures", "auth"))
	c.Assert(err, checker.IsNil)
	testPath := fmt.Sprintf("%s%c%s", osPath, filepath.ListSeparator, absolute)

	os.Setenv("PATH", testPath)

	repoName := fmt.Sprintf("%v/dockercli/busybox:authtest", privateRegistryURL)

	tmp, err := ioutil.TempDir("", "integration-cli-")
	c.Assert(err, checker.IsNil)

	externalAuthConfig := `***REMOVED*** "credsStore": "shell-test" ***REMOVED***`

	configPath := filepath.Join(tmp, "config.json")
	err = ioutil.WriteFile(configPath, []byte(externalAuthConfig), 0644)
	c.Assert(err, checker.IsNil)

	dockerCmd(c, "--config", tmp, "login", "-u", s.reg.Username(), "-p", s.reg.Password(), privateRegistryURL)

	b, err := ioutil.ReadFile(configPath)
	c.Assert(err, checker.IsNil)
	c.Assert(string(b), checker.Not(checker.Contains), "\"auth\":")

	dockerCmd(c, "--config", tmp, "tag", "busybox", repoName)
	dockerCmd(c, "--config", tmp, "push", repoName)

	dockerCmd(c, "--config", tmp, "pull", repoName)
***REMOVED***

// TestRunImplicitPullWithNoTag should pull implicitly only the default tag (latest)
func (s *DockerRegistrySuite) TestRunImplicitPullWithNoTag(c *check.C) ***REMOVED***
	testRequires(c, DaemonIsLinux)
	repo := fmt.Sprintf("%v/dockercli/busybox", privateRegistryURL)
	repoTag1 := fmt.Sprintf("%v:latest", repo)
	repoTag2 := fmt.Sprintf("%v:t1", repo)
	// tag the image and upload it to the private registry
	dockerCmd(c, "tag", "busybox", repoTag1)
	dockerCmd(c, "tag", "busybox", repoTag2)
	dockerCmd(c, "push", repo)
	dockerCmd(c, "rmi", repoTag1)
	dockerCmd(c, "rmi", repoTag2)

	out, _ := dockerCmd(c, "run", repo)
	c.Assert(out, checker.Contains, fmt.Sprintf("Unable to find image '%s:latest' locally", repo))

	// There should be only one line for repo, the one with repo:latest
	outImageCmd, _ := dockerCmd(c, "images", repo)
	splitOutImageCmd := strings.Split(strings.TrimSpace(outImageCmd), "\n")
	c.Assert(splitOutImageCmd, checker.HasLen, 2)
***REMOVED***
