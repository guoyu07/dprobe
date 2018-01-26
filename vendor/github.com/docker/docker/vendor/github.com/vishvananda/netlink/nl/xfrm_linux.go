package nl

import (
	"bytes"
	"net"
	"unsafe"
)

// Infinity for packet and byte counts
const (
	XFRM_INF = ^uint64(0)
)

type XfrmMsgType uint8

type XfrmMsg interface ***REMOVED***
	Type() XfrmMsgType
***REMOVED***

// Message Types
const (
	XFRM_MSG_BASE        XfrmMsgType = 0x10
	XFRM_MSG_NEWSA                   = 0x10
	XFRM_MSG_DELSA                   = 0x11
	XFRM_MSG_GETSA                   = 0x12
	XFRM_MSG_NEWPOLICY               = 0x13
	XFRM_MSG_DELPOLICY               = 0x14
	XFRM_MSG_GETPOLICY               = 0x15
	XFRM_MSG_ALLOCSPI                = 0x16
	XFRM_MSG_ACQUIRE                 = 0x17
	XFRM_MSG_EXPIRE                  = 0x18
	XFRM_MSG_UPDPOLICY               = 0x19
	XFRM_MSG_UPDSA                   = 0x1a
	XFRM_MSG_POLEXPIRE               = 0x1b
	XFRM_MSG_FLUSHSA                 = 0x1c
	XFRM_MSG_FLUSHPOLICY             = 0x1d
	XFRM_MSG_NEWAE                   = 0x1e
	XFRM_MSG_GETAE                   = 0x1f
	XFRM_MSG_REPORT                  = 0x20
	XFRM_MSG_MIGRATE                 = 0x21
	XFRM_MSG_NEWSADINFO              = 0x22
	XFRM_MSG_GETSADINFO              = 0x23
	XFRM_MSG_NEWSPDINFO              = 0x24
	XFRM_MSG_GETSPDINFO              = 0x25
	XFRM_MSG_MAPPING                 = 0x26
	XFRM_MSG_MAX                     = 0x26
	XFRM_NR_MSGTYPES                 = 0x17
)

// Attribute types
const (
	/* Netlink message attributes.  */
	XFRMA_UNSPEC         = 0x00
	XFRMA_ALG_AUTH       = 0x01 /* struct xfrm_algo */
	XFRMA_ALG_CRYPT      = 0x02 /* struct xfrm_algo */
	XFRMA_ALG_COMP       = 0x03 /* struct xfrm_algo */
	XFRMA_ENCAP          = 0x04 /* struct xfrm_algo + struct xfrm_encap_tmpl */
	XFRMA_TMPL           = 0x05 /* 1 or more struct xfrm_user_tmpl */
	XFRMA_SA             = 0x06 /* struct xfrm_usersa_info  */
	XFRMA_POLICY         = 0x07 /* struct xfrm_userpolicy_info */
	XFRMA_SEC_CTX        = 0x08 /* struct xfrm_sec_ctx */
	XFRMA_LTIME_VAL      = 0x09
	XFRMA_REPLAY_VAL     = 0x0a
	XFRMA_REPLAY_THRESH  = 0x0b
	XFRMA_ETIMER_THRESH  = 0x0c
	XFRMA_SRCADDR        = 0x0d /* xfrm_address_t */
	XFRMA_COADDR         = 0x0e /* xfrm_address_t */
	XFRMA_LASTUSED       = 0x0f /* unsigned long  */
	XFRMA_POLICY_TYPE    = 0x10 /* struct xfrm_userpolicy_type */
	XFRMA_MIGRATE        = 0x11
	XFRMA_ALG_AEAD       = 0x12 /* struct xfrm_algo_aead */
	XFRMA_KMADDRESS      = 0x13 /* struct xfrm_user_kmaddress */
	XFRMA_ALG_AUTH_TRUNC = 0x14 /* struct xfrm_algo_auth */
	XFRMA_MARK           = 0x15 /* struct xfrm_mark */
	XFRMA_TFCPAD         = 0x16 /* __u32 */
	XFRMA_REPLAY_ESN_VAL = 0x17 /* struct xfrm_replay_esn */
	XFRMA_SA_EXTRA_FLAGS = 0x18 /* __u32 */
	XFRMA_MAX            = 0x18
)

const (
	SizeofXfrmAddress     = 0x10
	SizeofXfrmSelector    = 0x38
	SizeofXfrmLifetimeCfg = 0x40
	SizeofXfrmLifetimeCur = 0x20
	SizeofXfrmId          = 0x18
	SizeofXfrmMark        = 0x08
)

// Netlink groups
const (
	XFRMNLGRP_NONE    = 0x0
	XFRMNLGRP_ACQUIRE = 0x1
	XFRMNLGRP_EXPIRE  = 0x2
	XFRMNLGRP_SA      = 0x3
	XFRMNLGRP_POLICY  = 0x4
	XFRMNLGRP_AEVENTS = 0x5
	XFRMNLGRP_REPORT  = 0x6
	XFRMNLGRP_MIGRATE = 0x7
	XFRMNLGRP_MAPPING = 0x8
	__XFRMNLGRP_MAX   = 0x9
)

// typedef union ***REMOVED***
//   __be32    a4;
//   __be32    a6[4];
// ***REMOVED*** xfrm_address_t;

type XfrmAddress [SizeofXfrmAddress]byte

func (x *XfrmAddress) ToIP() net.IP ***REMOVED***
	var empty = [12]byte***REMOVED******REMOVED***
	ip := make(net.IP, net.IPv6len)
	if bytes.Equal(x[4:16], empty[:]) ***REMOVED***
		ip[10] = 0xff
		ip[11] = 0xff
		copy(ip[12:16], x[0:4])
	***REMOVED*** else ***REMOVED***
		copy(ip[:], x[:])
	***REMOVED***
	return ip
***REMOVED***

func (x *XfrmAddress) ToIPNet(prefixlen uint8) *net.IPNet ***REMOVED***
	ip := x.ToIP()
	if GetIPFamily(ip) == FAMILY_V4 ***REMOVED***
		return &net.IPNet***REMOVED***IP: ip, Mask: net.CIDRMask(int(prefixlen), 32)***REMOVED***
	***REMOVED***
	return &net.IPNet***REMOVED***IP: ip, Mask: net.CIDRMask(int(prefixlen), 128)***REMOVED***
***REMOVED***

func (x *XfrmAddress) FromIP(ip net.IP) ***REMOVED***
	var empty = [16]byte***REMOVED******REMOVED***
	if len(ip) < net.IPv4len ***REMOVED***
		copy(x[4:16], empty[:])
	***REMOVED*** else if GetIPFamily(ip) == FAMILY_V4 ***REMOVED***
		copy(x[0:4], ip.To4()[0:4])
		copy(x[4:16], empty[:12])
	***REMOVED*** else ***REMOVED***
		copy(x[0:16], ip.To16()[0:16])
	***REMOVED***
***REMOVED***

func DeserializeXfrmAddress(b []byte) *XfrmAddress ***REMOVED***
	return (*XfrmAddress)(unsafe.Pointer(&b[0:SizeofXfrmAddress][0]))
***REMOVED***

func (x *XfrmAddress) Serialize() []byte ***REMOVED***
	return (*(*[SizeofXfrmAddress]byte)(unsafe.Pointer(x)))[:]
***REMOVED***

// struct xfrm_selector ***REMOVED***
//   xfrm_address_t  daddr;
//   xfrm_address_t  saddr;
//   __be16  dport;
//   __be16  dport_mask;
//   __be16  sport;
//   __be16  sport_mask;
//   __u16 family;
//   __u8  prefixlen_d;
//   __u8  prefixlen_s;
//   __u8  proto;
//   int ifindex;
//   __kernel_uid32_t  user;
// ***REMOVED***;

type XfrmSelector struct ***REMOVED***
	Daddr      XfrmAddress
	Saddr      XfrmAddress
	Dport      uint16 // big endian
	DportMask  uint16 // big endian
	Sport      uint16 // big endian
	SportMask  uint16 // big endian
	Family     uint16
	PrefixlenD uint8
	PrefixlenS uint8
	Proto      uint8
	Pad        [3]byte
	Ifindex    int32
	User       uint32
***REMOVED***

func (msg *XfrmSelector) Len() int ***REMOVED***
	return SizeofXfrmSelector
***REMOVED***

func DeserializeXfrmSelector(b []byte) *XfrmSelector ***REMOVED***
	return (*XfrmSelector)(unsafe.Pointer(&b[0:SizeofXfrmSelector][0]))
***REMOVED***

func (msg *XfrmSelector) Serialize() []byte ***REMOVED***
	return (*(*[SizeofXfrmSelector]byte)(unsafe.Pointer(msg)))[:]
***REMOVED***

// struct xfrm_lifetime_cfg ***REMOVED***
//   __u64 soft_byte_limit;
//   __u64 hard_byte_limit;
//   __u64 soft_packet_limit;
//   __u64 hard_packet_limit;
//   __u64 soft_add_expires_seconds;
//   __u64 hard_add_expires_seconds;
//   __u64 soft_use_expires_seconds;
//   __u64 hard_use_expires_seconds;
// ***REMOVED***;
//

type XfrmLifetimeCfg struct ***REMOVED***
	SoftByteLimit         uint64
	HardByteLimit         uint64
	SoftPacketLimit       uint64
	HardPacketLimit       uint64
	SoftAddExpiresSeconds uint64
	HardAddExpiresSeconds uint64
	SoftUseExpiresSeconds uint64
	HardUseExpiresSeconds uint64
***REMOVED***

func (msg *XfrmLifetimeCfg) Len() int ***REMOVED***
	return SizeofXfrmLifetimeCfg
***REMOVED***

func DeserializeXfrmLifetimeCfg(b []byte) *XfrmLifetimeCfg ***REMOVED***
	return (*XfrmLifetimeCfg)(unsafe.Pointer(&b[0:SizeofXfrmLifetimeCfg][0]))
***REMOVED***

func (msg *XfrmLifetimeCfg) Serialize() []byte ***REMOVED***
	return (*(*[SizeofXfrmLifetimeCfg]byte)(unsafe.Pointer(msg)))[:]
***REMOVED***

// struct xfrm_lifetime_cur ***REMOVED***
//   __u64 bytes;
//   __u64 packets;
//   __u64 add_time;
//   __u64 use_time;
// ***REMOVED***;

type XfrmLifetimeCur struct ***REMOVED***
	Bytes   uint64
	Packets uint64
	AddTime uint64
	UseTime uint64
***REMOVED***

func (msg *XfrmLifetimeCur) Len() int ***REMOVED***
	return SizeofXfrmLifetimeCur
***REMOVED***

func DeserializeXfrmLifetimeCur(b []byte) *XfrmLifetimeCur ***REMOVED***
	return (*XfrmLifetimeCur)(unsafe.Pointer(&b[0:SizeofXfrmLifetimeCur][0]))
***REMOVED***

func (msg *XfrmLifetimeCur) Serialize() []byte ***REMOVED***
	return (*(*[SizeofXfrmLifetimeCur]byte)(unsafe.Pointer(msg)))[:]
***REMOVED***

// struct xfrm_id ***REMOVED***
//   xfrm_address_t  daddr;
//   __be32    spi;
//   __u8    proto;
// ***REMOVED***;

type XfrmId struct ***REMOVED***
	Daddr XfrmAddress
	Spi   uint32 // big endian
	Proto uint8
	Pad   [3]byte
***REMOVED***

func (msg *XfrmId) Len() int ***REMOVED***
	return SizeofXfrmId
***REMOVED***

func DeserializeXfrmId(b []byte) *XfrmId ***REMOVED***
	return (*XfrmId)(unsafe.Pointer(&b[0:SizeofXfrmId][0]))
***REMOVED***

func (msg *XfrmId) Serialize() []byte ***REMOVED***
	return (*(*[SizeofXfrmId]byte)(unsafe.Pointer(msg)))[:]
***REMOVED***

type XfrmMark struct ***REMOVED***
	Value uint32
	Mask  uint32
***REMOVED***

func (msg *XfrmMark) Len() int ***REMOVED***
	return SizeofXfrmMark
***REMOVED***

func DeserializeXfrmMark(b []byte) *XfrmMark ***REMOVED***
	return (*XfrmMark)(unsafe.Pointer(&b[0:SizeofXfrmMark][0]))
***REMOVED***

func (msg *XfrmMark) Serialize() []byte ***REMOVED***
	return (*(*[SizeofXfrmMark]byte)(unsafe.Pointer(msg)))[:]
***REMOVED***
