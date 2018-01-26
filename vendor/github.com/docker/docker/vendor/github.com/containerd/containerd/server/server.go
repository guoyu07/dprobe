package server

import (
	"expvar"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"path/filepath"
	"strings"

	"github.com/boltdb/bolt"
	containers "github.com/containerd/containerd/api/services/containers/v1"
	contentapi "github.com/containerd/containerd/api/services/content/v1"
	diff "github.com/containerd/containerd/api/services/diff/v1"
	eventsapi "github.com/containerd/containerd/api/services/events/v1"
	images "github.com/containerd/containerd/api/services/images/v1"
	introspection "github.com/containerd/containerd/api/services/introspection/v1"
	leasesapi "github.com/containerd/containerd/api/services/leases/v1"
	namespaces "github.com/containerd/containerd/api/services/namespaces/v1"
	snapshotsapi "github.com/containerd/containerd/api/services/snapshots/v1"
	tasks "github.com/containerd/containerd/api/services/tasks/v1"
	version "github.com/containerd/containerd/api/services/version/v1"
	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/content/local"
	"github.com/containerd/containerd/events/exchange"
	"github.com/containerd/containerd/log"
	"github.com/containerd/containerd/metadata"
	"github.com/containerd/containerd/plugin"
	"github.com/containerd/containerd/snapshots"
	metrics "github.com/docker/go-metrics"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// New creates and initializes a new containerd server
func New(ctx context.Context, config *Config) (*Server, error) ***REMOVED***
	switch ***REMOVED***
	case config.Root == "":
		return nil, errors.New("root must be specified")
	case config.State == "":
		return nil, errors.New("state must be specified")
	case config.Root == config.State:
		return nil, errors.New("root and state must be different paths")
	***REMOVED***

	if err := os.MkdirAll(config.Root, 0711); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if err := os.MkdirAll(config.State, 0711); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if err := apply(ctx, config); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	plugins, err := loadPlugins(config)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	rpc := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	var (
		services []plugin.Service
		s        = &Server***REMOVED***
			rpc:    rpc,
			events: exchange.NewExchange(),
		***REMOVED***
		initialized = plugin.NewPluginSet()
	)
	for _, p := range plugins ***REMOVED***
		id := p.URI()
		log.G(ctx).WithField("type", p.Type).Infof("loading plugin %q...", id)

		initContext := plugin.NewContext(
			ctx,
			p,
			initialized,
			config.Root,
			config.State,
		)
		initContext.Events = s.events
		initContext.Address = config.GRPC.Address

		// load the plugin specific configuration if it is provided
		if p.Config != nil ***REMOVED***
			pluginConfig, err := config.Decode(p.ID, p.Config)
			if err != nil ***REMOVED***
				return nil, err
			***REMOVED***
			initContext.Config = pluginConfig
		***REMOVED***
		result := p.Init(initContext)
		if err := initialized.Add(result); err != nil ***REMOVED***
			return nil, errors.Wrapf(err, "could not add plugin result to plugin set")
		***REMOVED***

		instance, err := result.Instance()
		if err != nil ***REMOVED***
			if plugin.IsSkipPlugin(err) ***REMOVED***
				log.G(ctx).WithField("type", p.Type).Infof("skip loading plugin %q...", id)
			***REMOVED*** else ***REMOVED***
				log.G(ctx).WithError(err).Warnf("failed to load plugin %s", id)
			***REMOVED***
			continue
		***REMOVED***
		// check for grpc services that should be registered with the server
		if service, ok := instance.(plugin.Service); ok ***REMOVED***
			services = append(services, service)
		***REMOVED***
	***REMOVED***
	// register services after all plugins have been initialized
	for _, service := range services ***REMOVED***
		if err := service.Register(rpc); err != nil ***REMOVED***
			return nil, err
		***REMOVED***
	***REMOVED***
	return s, nil
***REMOVED***

// Server is the containerd main daemon
type Server struct ***REMOVED***
	rpc    *grpc.Server
	events *exchange.Exchange
***REMOVED***

// ServeGRPC provides the containerd grpc APIs on the provided listener
func (s *Server) ServeGRPC(l net.Listener) error ***REMOVED***
	// before we start serving the grpc API regster the grpc_prometheus metrics
	// handler.  This needs to be the last service registered so that it can collect
	// metrics for every other service
	grpc_prometheus.Register(s.rpc)
	return trapClosedConnErr(s.rpc.Serve(l))
***REMOVED***

// ServeMetrics provides a prometheus endpoint for exposing metrics
func (s *Server) ServeMetrics(l net.Listener) error ***REMOVED***
	m := http.NewServeMux()
	m.Handle("/v1/metrics", metrics.Handler())
	return trapClosedConnErr(http.Serve(l, m))
***REMOVED***

// ServeDebug provides a debug endpoint
func (s *Server) ServeDebug(l net.Listener) error ***REMOVED***
	// don't use the default http server mux to make sure nothing gets registered
	// that we don't want to expose via containerd
	m := http.NewServeMux()
	m.Handle("/debug/vars", expvar.Handler())
	m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	return trapClosedConnErr(http.Serve(l, m))
***REMOVED***

// Stop the containerd server canceling any open connections
func (s *Server) Stop() ***REMOVED***
	s.rpc.Stop()
***REMOVED***

func loadPlugins(config *Config) ([]*plugin.Registration, error) ***REMOVED***
	// load all plugins into containerd
	if err := plugin.Load(filepath.Join(config.Root, "plugins")); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	// load additional plugins that don't automatically register themselves
	plugin.Register(&plugin.Registration***REMOVED***
		Type: plugin.ContentPlugin,
		ID:   "content",
		InitFn: func(ic *plugin.InitContext) (interface***REMOVED******REMOVED***, error) ***REMOVED***
			ic.Meta.Exports["root"] = ic.Root
			return local.NewStore(ic.Root)
		***REMOVED***,
	***REMOVED***)
	plugin.Register(&plugin.Registration***REMOVED***
		Type: plugin.MetadataPlugin,
		ID:   "bolt",
		Requires: []plugin.Type***REMOVED***
			plugin.ContentPlugin,
			plugin.SnapshotPlugin,
		***REMOVED***,
		InitFn: func(ic *plugin.InitContext) (interface***REMOVED******REMOVED***, error) ***REMOVED***
			if err := os.MkdirAll(ic.Root, 0711); err != nil ***REMOVED***
				return nil, err
			***REMOVED***
			cs, err := ic.Get(plugin.ContentPlugin)
			if err != nil ***REMOVED***
				return nil, err
			***REMOVED***

			snapshottersRaw, err := ic.GetByType(plugin.SnapshotPlugin)
			if err != nil ***REMOVED***
				return nil, err
			***REMOVED***

			snapshotters := make(map[string]snapshots.Snapshotter)
			for name, sn := range snapshottersRaw ***REMOVED***
				sn, err := sn.Instance()
				if err != nil ***REMOVED***
					log.G(ic.Context).WithError(err).
						Warnf("could not use snapshotter %v in metadata plugin", name)
					continue
				***REMOVED***
				snapshotters[name] = sn.(snapshots.Snapshotter)
			***REMOVED***

			path := filepath.Join(ic.Root, "meta.db")
			ic.Meta.Exports["path"] = path

			db, err := bolt.Open(path, 0644, nil)
			if err != nil ***REMOVED***
				return nil, err
			***REMOVED***
			mdb := metadata.NewDB(db, cs.(content.Store), snapshotters)
			if err := mdb.Init(ic.Context); err != nil ***REMOVED***
				return nil, err
			***REMOVED***
			return mdb, nil
		***REMOVED***,
	***REMOVED***)

	// return the ordered graph for plugins
	return plugin.Graph(), nil
***REMOVED***

func interceptor(
	ctx context.Context,
	req interface***REMOVED******REMOVED***,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface***REMOVED******REMOVED***, error) ***REMOVED***
	ctx = log.WithModule(ctx, "containerd")
	switch info.Server.(type) ***REMOVED***
	case tasks.TasksServer:
		ctx = log.WithModule(ctx, "tasks")
	case containers.ContainersServer:
		ctx = log.WithModule(ctx, "containers")
	case contentapi.ContentServer:
		ctx = log.WithModule(ctx, "content")
	case images.ImagesServer:
		ctx = log.WithModule(ctx, "images")
	case grpc_health_v1.HealthServer:
		// No need to change the context
	case version.VersionServer:
		ctx = log.WithModule(ctx, "version")
	case snapshotsapi.SnapshotsServer:
		ctx = log.WithModule(ctx, "snapshot")
	case diff.DiffServer:
		ctx = log.WithModule(ctx, "diff")
	case namespaces.NamespacesServer:
		ctx = log.WithModule(ctx, "namespaces")
	case eventsapi.EventsServer:
		ctx = log.WithModule(ctx, "events")
	case introspection.IntrospectionServer:
		ctx = log.WithModule(ctx, "introspection")
	case leasesapi.LeasesServer:
		ctx = log.WithModule(ctx, "leases")
	default:
		log.G(ctx).Warnf("unknown GRPC server type: %#v\n", info.Server)
	***REMOVED***
	return grpc_prometheus.UnaryServerInterceptor(ctx, req, info, handler)
***REMOVED***

func trapClosedConnErr(err error) error ***REMOVED***
	if err == nil ***REMOVED***
		return nil
	***REMOVED***
	if strings.Contains(err.Error(), "use of closed network connection") ***REMOVED***
		return nil
	***REMOVED***
	return err
***REMOVED***
