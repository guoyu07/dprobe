package network

import (
	"net"

	networktypes "github.com/docker/docker/api/types/network"
	clustertypes "github.com/docker/docker/daemon/cluster/provider"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
)

// Settings stores configuration details about the daemon network config
// TODO Windows. Many of these fields can be factored out.,
type Settings struct ***REMOVED***
	Bridge                 string
	SandboxID              string
	HairpinMode            bool
	LinkLocalIPv6Address   string
	LinkLocalIPv6PrefixLen int
	Networks               map[string]*EndpointSettings
	Service                *clustertypes.ServiceConfig
	Ports                  nat.PortMap
	SandboxKey             string
	SecondaryIPAddresses   []networktypes.Address
	SecondaryIPv6Addresses []networktypes.Address
	IsAnonymousEndpoint    bool
	HasSwarmEndpoint       bool
***REMOVED***

// EndpointSettings is a package local wrapper for
// networktypes.EndpointSettings which stores Endpoint state that
// needs to be persisted to disk but not exposed in the api.
type EndpointSettings struct ***REMOVED***
	*networktypes.EndpointSettings
	IPAMOperational bool
***REMOVED***

// AttachmentStore stores the load balancer IP address for a network id.
type AttachmentStore struct ***REMOVED***
	//key: networkd id
	//value: load balancer ip address
	networkToNodeLBIP map[string]net.IP
***REMOVED***

// ResetAttachments clears any existing load balancer IP to network mapping and
// sets the mapping to the given attachments.
func (store *AttachmentStore) ResetAttachments(attachments map[string]string) error ***REMOVED***
	store.ClearAttachments()
	for nid, nodeIP := range attachments ***REMOVED***
		ip, _, err := net.ParseCIDR(nodeIP)
		if err != nil ***REMOVED***
			store.networkToNodeLBIP = make(map[string]net.IP)
			return errors.Wrapf(err, "Failed to parse load balancer address %s", nodeIP)
		***REMOVED***
		store.networkToNodeLBIP[nid] = ip
	***REMOVED***
	return nil
***REMOVED***

// ClearAttachments clears all the mappings of network to load balancer IP Address.
func (store *AttachmentStore) ClearAttachments() ***REMOVED***
	store.networkToNodeLBIP = make(map[string]net.IP)
***REMOVED***

// GetIPForNetwork return the load balancer IP address for the given network.
func (store *AttachmentStore) GetIPForNetwork(networkID string) (net.IP, bool) ***REMOVED***
	ip, exists := store.networkToNodeLBIP[networkID]
	return ip, exists
***REMOVED***
