package dns

import "strconv"

const (
	year68 = 1 << 31 // For RFC1982 (Serial Arithmetic) calculations in 32 bits.
	// DefaultMsgSize is the standard default for messages larger than 512 bytes.
	DefaultMsgSize = 4096
	// MinMsgSize is the minimal size of a DNS packet.
	MinMsgSize = 512
	// MaxMsgSize is the largest possible DNS packet.
	MaxMsgSize = 65535
	defaultTtl = 3600 // Default internal TTL.
)

// Error represents a DNS error
type Error struct***REMOVED*** err string ***REMOVED***

func (e *Error) Error() string ***REMOVED***
	if e == nil ***REMOVED***
		return "dns: <nil>"
	***REMOVED***
	return "dns: " + e.err
***REMOVED***

// An RR represents a resource record.
type RR interface ***REMOVED***
	// Header returns the header of an resource record. The header contains
	// everything up to the rdata.
	Header() *RR_Header
	// String returns the text representation of the resource record.
	String() string
	// copy returns a copy of the RR
	copy() RR
	// len returns the length (in octets) of the uncompressed RR in wire format.
	len() int
***REMOVED***

// RR_Header is the header all DNS resource records share.
type RR_Header struct ***REMOVED***
	Name     string `dns:"cdomain-name"`
	Rrtype   uint16
	Class    uint16
	Ttl      uint32
	Rdlength uint16 // length of data after header
***REMOVED***

// Header returns itself. This is here to make RR_Header implement the RR interface.
func (h *RR_Header) Header() *RR_Header ***REMOVED*** return h ***REMOVED***

// Just to imlement the RR interface.
func (h *RR_Header) copy() RR ***REMOVED*** return nil ***REMOVED***

func (h *RR_Header) copyHeader() *RR_Header ***REMOVED***
	r := new(RR_Header)
	r.Name = h.Name
	r.Rrtype = h.Rrtype
	r.Class = h.Class
	r.Ttl = h.Ttl
	r.Rdlength = h.Rdlength
	return r
***REMOVED***

func (h *RR_Header) String() string ***REMOVED***
	var s string

	if h.Rrtype == TypeOPT ***REMOVED***
		s = ";"
		// and maybe other things
	***REMOVED***

	s += sprintName(h.Name) + "\t"
	s += strconv.FormatInt(int64(h.Ttl), 10) + "\t"
	s += Class(h.Class).String() + "\t"
	s += Type(h.Rrtype).String() + "\t"
	return s
***REMOVED***

func (h *RR_Header) len() int ***REMOVED***
	l := len(h.Name) + 1
	l += 10 // rrtype(2) + class(2) + ttl(4) + rdlength(2)
	return l
***REMOVED***

// ToRFC3597 converts a known RR to the unknown RR representation
// from RFC 3597.
func (rr *RFC3597) ToRFC3597(r RR) error ***REMOVED***
	buf := make([]byte, r.len()*2)
	off, err := PackStruct(r, buf, 0)
	if err != nil ***REMOVED***
		return err
	***REMOVED***
	buf = buf[:off]
	rawSetRdlength(buf, 0, off)
	_, err = UnpackStruct(rr, buf, 0)
	if err != nil ***REMOVED***
		return err
	***REMOVED***
	return nil
***REMOVED***
