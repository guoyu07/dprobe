// +build linux

package mount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	/* 36 35 98:0 /mnt1 /mnt2 rw,noatime master:1 - ext3 /dev/root rw,errors=continue
	   (1)(2)(3)   (4)   (5)      (6)      (7)   (8) (9)   (10)         (11)

	   (1) mount ID:  unique identifier of the mount (may be reused after umount)
	   (2) parent ID:  ID of parent (or of self for the top of the mount tree)
	   (3) major:minor:  value of st_dev for files on filesystem
	   (4) root:  root of the mount within the filesystem
	   (5) mount point:  mount point relative to the process's root
	   (6) mount options:  per mount options
	   (7) optional fields:  zero or more fields of the form "tag[:value]"
	   (8) separator:  marks the end of the optional fields
	   (9) filesystem type:  name of filesystem of the form "type[.subtype]"
	   (10) mount source:  filesystem specific information or "none"
	   (11) super options:  per super block options*/
	mountinfoFormat = "%d %d %d:%d %s %s %s %s"
)

// Self retrieves a list of mounts for the current running process.
func Self() ([]Info, error) ***REMOVED***
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	defer f.Close()

	return parseInfoFile(f)
***REMOVED***

func parseInfoFile(r io.Reader) ([]Info, error) ***REMOVED***
	var (
		s   = bufio.NewScanner(r)
		out = []Info***REMOVED******REMOVED***
	)

	for s.Scan() ***REMOVED***
		if err := s.Err(); err != nil ***REMOVED***
			return nil, err
		***REMOVED***

		var (
			p              = Info***REMOVED******REMOVED***
			text           = s.Text()
			optionalFields string
		)

		if _, err := fmt.Sscanf(text, mountinfoFormat,
			&p.ID, &p.Parent, &p.Major, &p.Minor,
			&p.Root, &p.Mountpoint, &p.Options, &optionalFields); err != nil ***REMOVED***
			return nil, fmt.Errorf("Scanning '%s' failed: %s", text, err)
		***REMOVED***
		// Safe as mountinfo encodes mountpoints with spaces as \040.
		index := strings.Index(text, " - ")
		postSeparatorFields := strings.Fields(text[index+3:])
		if len(postSeparatorFields) < 3 ***REMOVED***
			return nil, fmt.Errorf("Error found less than 3 fields post '-' in %q", text)
		***REMOVED***

		if optionalFields != "-" ***REMOVED***
			p.Optional = optionalFields
		***REMOVED***

		p.FSType = postSeparatorFields[0]
		p.Source = postSeparatorFields[1]
		p.VFSOptions = strings.Join(postSeparatorFields[2:], " ")
		out = append(out, p)
	***REMOVED***
	return out, nil
***REMOVED***

// PID collects the mounts for a specific process ID. If the process
// ID is unknown, it is better to use `Self` which will inspect
// "/proc/self/mountinfo" instead.
func PID(pid int) ([]Info, error) ***REMOVED***
	f, err := os.Open(fmt.Sprintf("/proc/%d/mountinfo", pid))
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	defer f.Close()

	return parseInfoFile(f)
***REMOVED***
