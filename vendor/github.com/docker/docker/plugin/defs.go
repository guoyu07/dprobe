package plugin

import (
	"sync"

	"github.com/docker/docker/pkg/plugins"
	"github.com/docker/docker/plugin/v2"
)

// Store manages the plugin inventory in memory and on-disk
type Store struct ***REMOVED***
	sync.RWMutex
	plugins map[string]*v2.Plugin
	/* handlers are necessary for transition path of legacy plugins
	 * to the new model. Legacy plugins use Handle() for registering an
	 * activation callback.*/
	handlers map[string][]func(string, *plugins.Client)
***REMOVED***

// NewStore creates a Store.
func NewStore() *Store ***REMOVED***
	return &Store***REMOVED***
		plugins:  make(map[string]*v2.Plugin),
		handlers: make(map[string][]func(string, *plugins.Client)),
	***REMOVED***
***REMOVED***

// CreateOpt is used to configure specific plugin details when created
type CreateOpt func(p *v2.Plugin)

// WithSwarmService is a CreateOpt that flags the passed in a plugin as a plugin
// managed by swarm
func WithSwarmService(id string) CreateOpt ***REMOVED***
	return func(p *v2.Plugin) ***REMOVED***
		p.SwarmServiceID = id
	***REMOVED***
***REMOVED***
