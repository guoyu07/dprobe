package proc

import (
	containerd_types "github.com/containerd/containerd/api/types"
	google_protobuf "github.com/gogo/protobuf/types"
)

// CreateConfig hold task creation configuration
type CreateConfig struct ***REMOVED***
	ID               string
	Bundle           string
	Runtime          string
	Rootfs           []*containerd_types.Mount
	Terminal         bool
	Stdin            string
	Stdout           string
	Stderr           string
	Checkpoint       string
	ParentCheckpoint string
	Options          *google_protobuf.Any
***REMOVED***

// ExecConfig holds exec creation configuration
type ExecConfig struct ***REMOVED***
	ID       string
	Terminal bool
	Stdin    string
	Stdout   string
	Stderr   string
	Spec     *google_protobuf.Any
***REMOVED***

// CheckpointConfig holds task checkpoint configuration
type CheckpointConfig struct ***REMOVED***
	Path    string
	Options *google_protobuf.Any
***REMOVED***
