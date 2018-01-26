package main

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/docker/docker/integration-cli/checker"
	"github.com/docker/docker/integration-cli/cli"
	"github.com/go-check/check"
	"github.com/gotestyourself/gotestyourself/icmd"
)

func (s *DockerSuite) TestImportDisplay(c *check.C) ***REMOVED***
	testRequires(c, DaemonIsLinux)
	out, _ := dockerCmd(c, "run", "-d", "busybox", "true")
	cleanedContainerID := strings.TrimSpace(out)

	out, err := RunCommandPipelineWithOutput(
		exec.Command(dockerBinary, "export", cleanedContainerID),
		exec.Command(dockerBinary, "import", "-"),
	)
	c.Assert(err, checker.IsNil)

	c.Assert(out, checker.Count, "\n", 1, check.Commentf("display is expected 1 '\\n' but didn't"))

	image := strings.TrimSpace(out)
	out, _ = dockerCmd(c, "run", "--rm", image, "true")
	c.Assert(out, checker.Equals, "", check.Commentf("command output should've been nothing."))
***REMOVED***

func (s *DockerSuite) TestImportBadURL(c *check.C) ***REMOVED***
	out, _, err := dockerCmdWithError("import", "http://nourl/bad")
	c.Assert(err, checker.NotNil, check.Commentf("import was supposed to fail but didn't"))
	// Depending on your system you can get either of these errors
	if !strings.Contains(out, "dial tcp") &&
		!strings.Contains(out, "ApplyLayer exit status 1 stdout:  stderr: archive/tar: invalid tar header") &&
		!strings.Contains(out, "Error processing tar file") ***REMOVED***
		c.Fatalf("expected an error msg but didn't get one.\nErr: %v\nOut: %v", err, out)
	***REMOVED***
***REMOVED***

func (s *DockerSuite) TestImportFile(c *check.C) ***REMOVED***
	testRequires(c, DaemonIsLinux)
	dockerCmd(c, "run", "--name", "test-import", "busybox", "true")

	temporaryFile, err := ioutil.TempFile("", "exportImportTest")
	c.Assert(err, checker.IsNil, check.Commentf("failed to create temporary file"))
	defer os.Remove(temporaryFile.Name())

	icmd.RunCmd(icmd.Cmd***REMOVED***
		Command: []string***REMOVED***dockerBinary, "export", "test-import"***REMOVED***,
		Stdout:  bufio.NewWriter(temporaryFile),
	***REMOVED***).Assert(c, icmd.Success)

	out, _ := dockerCmd(c, "import", temporaryFile.Name())
	c.Assert(out, checker.Count, "\n", 1, check.Commentf("display is expected 1 '\\n' but didn't"))
	image := strings.TrimSpace(out)

	out, _ = dockerCmd(c, "run", "--rm", image, "true")
	c.Assert(out, checker.Equals, "", check.Commentf("command output should've been nothing."))
***REMOVED***

func (s *DockerSuite) TestImportGzipped(c *check.C) ***REMOVED***
	testRequires(c, DaemonIsLinux)
	dockerCmd(c, "run", "--name", "test-import", "busybox", "true")

	temporaryFile, err := ioutil.TempFile("", "exportImportTest")
	c.Assert(err, checker.IsNil, check.Commentf("failed to create temporary file"))
	defer os.Remove(temporaryFile.Name())

	w := gzip.NewWriter(temporaryFile)
	icmd.RunCmd(icmd.Cmd***REMOVED***
		Command: []string***REMOVED***dockerBinary, "export", "test-import"***REMOVED***,
		Stdout:  w,
	***REMOVED***).Assert(c, icmd.Success)
	c.Assert(w.Close(), checker.IsNil, check.Commentf("failed to close gzip writer"))
	temporaryFile.Close()
	out, _ := dockerCmd(c, "import", temporaryFile.Name())
	c.Assert(out, checker.Count, "\n", 1, check.Commentf("display is expected 1 '\\n' but didn't"))
	image := strings.TrimSpace(out)

	out, _ = dockerCmd(c, "run", "--rm", image, "true")
	c.Assert(out, checker.Equals, "", check.Commentf("command output should've been nothing."))
***REMOVED***

func (s *DockerSuite) TestImportFileWithMessage(c *check.C) ***REMOVED***
	testRequires(c, DaemonIsLinux)
	dockerCmd(c, "run", "--name", "test-import", "busybox", "true")

	temporaryFile, err := ioutil.TempFile("", "exportImportTest")
	c.Assert(err, checker.IsNil, check.Commentf("failed to create temporary file"))
	defer os.Remove(temporaryFile.Name())

	icmd.RunCmd(icmd.Cmd***REMOVED***
		Command: []string***REMOVED***dockerBinary, "export", "test-import"***REMOVED***,
		Stdout:  bufio.NewWriter(temporaryFile),
	***REMOVED***).Assert(c, icmd.Success)

	message := "Testing commit message"
	out, _ := dockerCmd(c, "import", "-m", message, temporaryFile.Name())
	c.Assert(out, checker.Count, "\n", 1, check.Commentf("display is expected 1 '\\n' but didn't"))
	image := strings.TrimSpace(out)

	out, _ = dockerCmd(c, "history", image)
	split := strings.Split(out, "\n")

	c.Assert(split, checker.HasLen, 3, check.Commentf("expected 3 lines from image history"))
	r := regexp.MustCompile("[\\s]***REMOVED***2,***REMOVED***")
	split = r.Split(split[1], -1)

	c.Assert(message, checker.Equals, split[3], check.Commentf("didn't get expected value in commit message"))

	out, _ = dockerCmd(c, "run", "--rm", image, "true")
	c.Assert(out, checker.Equals, "", check.Commentf("command output should've been nothing"))
***REMOVED***

func (s *DockerSuite) TestImportFileNonExistentFile(c *check.C) ***REMOVED***
	_, _, err := dockerCmdWithError("import", "example.com/myImage.tar")
	c.Assert(err, checker.NotNil, check.Commentf("import non-existing file must failed"))
***REMOVED***

func (s *DockerSuite) TestImportWithQuotedChanges(c *check.C) ***REMOVED***
	testRequires(c, DaemonIsLinux)
	cli.DockerCmd(c, "run", "--name", "test-import", "busybox", "true")

	temporaryFile, err := ioutil.TempFile("", "exportImportTest")
	c.Assert(err, checker.IsNil, check.Commentf("failed to create temporary file"))
	defer os.Remove(temporaryFile.Name())

	cli.Docker(cli.Args("export", "test-import"), cli.WithStdout(bufio.NewWriter(temporaryFile))).Assert(c, icmd.Success)

	result := cli.DockerCmd(c, "import", "-c", `ENTRYPOINT ["/bin/sh", "-c"]`, temporaryFile.Name())
	image := strings.TrimSpace(result.Stdout())

	result = cli.DockerCmd(c, "run", "--rm", image, "true")
	result.Assert(c, icmd.Expected***REMOVED***Out: icmd.None***REMOVED***)
***REMOVED***
