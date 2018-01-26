package daemon

import (
	"github.com/docker/docker/container"
	"github.com/docker/docker/pkg/archive"
)

func (daemon *Daemon) tarCopyOptions(container *container.Container, noOverwriteDirNonDir bool) (*archive.TarOptions, error) ***REMOVED***
	return daemon.defaultTarCopyOptions(noOverwriteDirNonDir), nil
***REMOVED***
