package netlink

import (
	"fmt"
	"syscall"
)

// Proto is an enum representing an ipsec protocol.
type Proto uint8

const (
	XFRM_PROTO_ROUTE2    Proto = syscall.IPPROTO_ROUTING
	XFRM_PROTO_ESP       Proto = syscall.IPPROTO_ESP
	XFRM_PROTO_AH        Proto = syscall.IPPROTO_AH
	XFRM_PROTO_HAO       Proto = syscall.IPPROTO_DSTOPTS
	XFRM_PROTO_COMP      Proto = 0x6c // NOTE not defined on darwin
	XFRM_PROTO_IPSEC_ANY Proto = syscall.IPPROTO_RAW
)

func (p Proto) String() string ***REMOVED***
	switch p ***REMOVED***
	case XFRM_PROTO_ROUTE2:
		return "route2"
	case XFRM_PROTO_ESP:
		return "esp"
	case XFRM_PROTO_AH:
		return "ah"
	case XFRM_PROTO_HAO:
		return "hao"
	case XFRM_PROTO_COMP:
		return "comp"
	case XFRM_PROTO_IPSEC_ANY:
		return "ipsec-any"
	***REMOVED***
	return fmt.Sprintf("%d", p)
***REMOVED***

// Mode is an enum representing an ipsec transport.
type Mode uint8

const (
	XFRM_MODE_TRANSPORT Mode = iota
	XFRM_MODE_TUNNEL
	XFRM_MODE_ROUTEOPTIMIZATION
	XFRM_MODE_IN_TRIGGER
	XFRM_MODE_BEET
	XFRM_MODE_MAX
)

func (m Mode) String() string ***REMOVED***
	switch m ***REMOVED***
	case XFRM_MODE_TRANSPORT:
		return "transport"
	case XFRM_MODE_TUNNEL:
		return "tunnel"
	case XFRM_MODE_ROUTEOPTIMIZATION:
		return "ro"
	case XFRM_MODE_IN_TRIGGER:
		return "in_trigger"
	case XFRM_MODE_BEET:
		return "beet"
	***REMOVED***
	return fmt.Sprintf("%d", m)
***REMOVED***

// XfrmMark represents the mark associated to the state or policy
type XfrmMark struct ***REMOVED***
	Value uint32
	Mask  uint32
***REMOVED***

func (m *XfrmMark) String() string ***REMOVED***
	return fmt.Sprintf("(0x%x,0x%x)", m.Value, m.Mask)
***REMOVED***
