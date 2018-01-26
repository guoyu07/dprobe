package configs

import "fmt"

// blockIODevice holds major:minor format supported in blkio cgroup
type blockIODevice struct ***REMOVED***
	// Major is the device's major number
	Major int64 `json:"major"`
	// Minor is the device's minor number
	Minor int64 `json:"minor"`
***REMOVED***

// WeightDevice struct holds a `major:minor weight`|`major:minor leaf_weight` pair
type WeightDevice struct ***REMOVED***
	blockIODevice
	// Weight is the bandwidth rate for the device, range is from 10 to 1000
	Weight uint16 `json:"weight"`
	// LeafWeight is the bandwidth rate for the device while competing with the cgroup's child cgroups, range is from 10 to 1000, cfq scheduler only
	LeafWeight uint16 `json:"leafWeight"`
***REMOVED***

// NewWeightDevice returns a configured WeightDevice pointer
func NewWeightDevice(major, minor int64, weight, leafWeight uint16) *WeightDevice ***REMOVED***
	wd := &WeightDevice***REMOVED******REMOVED***
	wd.Major = major
	wd.Minor = minor
	wd.Weight = weight
	wd.LeafWeight = leafWeight
	return wd
***REMOVED***

// WeightString formats the struct to be writable to the cgroup specific file
func (wd *WeightDevice) WeightString() string ***REMOVED***
	return fmt.Sprintf("%d:%d %d", wd.Major, wd.Minor, wd.Weight)
***REMOVED***

// LeafWeightString formats the struct to be writable to the cgroup specific file
func (wd *WeightDevice) LeafWeightString() string ***REMOVED***
	return fmt.Sprintf("%d:%d %d", wd.Major, wd.Minor, wd.LeafWeight)
***REMOVED***

// ThrottleDevice struct holds a `major:minor rate_per_second` pair
type ThrottleDevice struct ***REMOVED***
	blockIODevice
	// Rate is the IO rate limit per cgroup per device
	Rate uint64 `json:"rate"`
***REMOVED***

// NewThrottleDevice returns a configured ThrottleDevice pointer
func NewThrottleDevice(major, minor int64, rate uint64) *ThrottleDevice ***REMOVED***
	td := &ThrottleDevice***REMOVED******REMOVED***
	td.Major = major
	td.Minor = minor
	td.Rate = rate
	return td
***REMOVED***

// String formats the struct to be writable to the cgroup specific file
func (td *ThrottleDevice) String() string ***REMOVED***
	return fmt.Sprintf("%d:%d %d", td.Major, td.Minor, td.Rate)
***REMOVED***
