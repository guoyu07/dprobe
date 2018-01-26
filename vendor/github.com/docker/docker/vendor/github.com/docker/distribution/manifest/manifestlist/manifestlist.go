package manifestlist

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest"
	"github.com/opencontainers/go-digest"
)

// MediaTypeManifestList specifies the mediaType for manifest lists.
const MediaTypeManifestList = "application/vnd.docker.distribution.manifest.list.v2+json"

// SchemaVersion provides a pre-initialized version structure for this
// packages version of the manifest.
var SchemaVersion = manifest.Versioned***REMOVED***
	SchemaVersion: 2,
	MediaType:     MediaTypeManifestList,
***REMOVED***

func init() ***REMOVED***
	manifestListFunc := func(b []byte) (distribution.Manifest, distribution.Descriptor, error) ***REMOVED***
		m := new(DeserializedManifestList)
		err := m.UnmarshalJSON(b)
		if err != nil ***REMOVED***
			return nil, distribution.Descriptor***REMOVED******REMOVED***, err
		***REMOVED***

		dgst := digest.FromBytes(b)
		return m, distribution.Descriptor***REMOVED***Digest: dgst, Size: int64(len(b)), MediaType: MediaTypeManifestList***REMOVED***, err
	***REMOVED***
	err := distribution.RegisterManifestSchema(MediaTypeManifestList, manifestListFunc)
	if err != nil ***REMOVED***
		panic(fmt.Sprintf("Unable to register manifest: %s", err))
	***REMOVED***
***REMOVED***

// PlatformSpec specifies a platform where a particular image manifest is
// applicable.
type PlatformSpec struct ***REMOVED***
	// Architecture field specifies the CPU architecture, for example
	// `amd64` or `ppc64`.
	Architecture string `json:"architecture"`

	// OS specifies the operating system, for example `linux` or `windows`.
	OS string `json:"os"`

	// OSVersion is an optional field specifying the operating system
	// version, for example `10.0.10586`.
	OSVersion string `json:"os.version,omitempty"`

	// OSFeatures is an optional field specifying an array of strings,
	// each listing a required OS feature (for example on Windows `win32k`).
	OSFeatures []string `json:"os.features,omitempty"`

	// Variant is an optional field specifying a variant of the CPU, for
	// example `ppc64le` to specify a little-endian version of a PowerPC CPU.
	Variant string `json:"variant,omitempty"`

	// Features is an optional field specifying an array of strings, each
	// listing a required CPU feature (for example `sse4` or `aes`).
	Features []string `json:"features,omitempty"`
***REMOVED***

// A ManifestDescriptor references a platform-specific manifest.
type ManifestDescriptor struct ***REMOVED***
	distribution.Descriptor

	// Platform specifies which platform the manifest pointed to by the
	// descriptor runs on.
	Platform PlatformSpec `json:"platform"`
***REMOVED***

// ManifestList references manifests for various platforms.
type ManifestList struct ***REMOVED***
	manifest.Versioned

	// Config references the image configuration as a blob.
	Manifests []ManifestDescriptor `json:"manifests"`
***REMOVED***

// References returns the distribution descriptors for the referenced image
// manifests.
func (m ManifestList) References() []distribution.Descriptor ***REMOVED***
	dependencies := make([]distribution.Descriptor, len(m.Manifests))
	for i := range m.Manifests ***REMOVED***
		dependencies[i] = m.Manifests[i].Descriptor
	***REMOVED***

	return dependencies
***REMOVED***

// DeserializedManifestList wraps ManifestList with a copy of the original
// JSON.
type DeserializedManifestList struct ***REMOVED***
	ManifestList

	// canonical is the canonical byte representation of the Manifest.
	canonical []byte
***REMOVED***

// FromDescriptors takes a slice of descriptors, and returns a
// DeserializedManifestList which contains the resulting manifest list
// and its JSON representation.
func FromDescriptors(descriptors []ManifestDescriptor) (*DeserializedManifestList, error) ***REMOVED***
	m := ManifestList***REMOVED***
		Versioned: SchemaVersion,
	***REMOVED***

	m.Manifests = make([]ManifestDescriptor, len(descriptors), len(descriptors))
	copy(m.Manifests, descriptors)

	deserialized := DeserializedManifestList***REMOVED***
		ManifestList: m,
	***REMOVED***

	var err error
	deserialized.canonical, err = json.MarshalIndent(&m, "", "   ")
	return &deserialized, err
***REMOVED***

// UnmarshalJSON populates a new ManifestList struct from JSON data.
func (m *DeserializedManifestList) UnmarshalJSON(b []byte) error ***REMOVED***
	m.canonical = make([]byte, len(b), len(b))
	// store manifest list in canonical
	copy(m.canonical, b)

	// Unmarshal canonical JSON into ManifestList object
	var manifestList ManifestList
	if err := json.Unmarshal(m.canonical, &manifestList); err != nil ***REMOVED***
		return err
	***REMOVED***

	m.ManifestList = manifestList

	return nil
***REMOVED***

// MarshalJSON returns the contents of canonical. If canonical is empty,
// marshals the inner contents.
func (m *DeserializedManifestList) MarshalJSON() ([]byte, error) ***REMOVED***
	if len(m.canonical) > 0 ***REMOVED***
		return m.canonical, nil
	***REMOVED***

	return nil, errors.New("JSON representation not initialized in DeserializedManifestList")
***REMOVED***

// Payload returns the raw content of the manifest list. The contents can be
// used to calculate the content identifier.
func (m DeserializedManifestList) Payload() (string, []byte, error) ***REMOVED***
	return m.MediaType, m.canonical, nil
***REMOVED***
