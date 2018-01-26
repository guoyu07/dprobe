// +build linux

// Internal functions for libseccomp Go bindings
// No exported functions

package seccomp

import (
	"fmt"
	"os"
	"syscall"
)

// Unexported C wrapping code - provides the C-Golang interface
// Get the seccomp header in scope
// Need stdlib.h for free() on cstrings

// #cgo pkg-config: libseccomp
/*
#include <stdlib.h>
#include <seccomp.h>

#if SCMP_VER_MAJOR < 2
#error Minimum supported version of Libseccomp is v2.1.0
#elif SCMP_VER_MAJOR == 2 && SCMP_VER_MINOR < 1
#error Minimum supported version of Libseccomp is v2.1.0
#endif

#define ARCH_BAD ~0

const uint32_t C_ARCH_BAD = ARCH_BAD;

#ifndef SCMP_ARCH_AARCH64
#define SCMP_ARCH_AARCH64 ARCH_BAD
#endif

#ifndef SCMP_ARCH_MIPS
#define SCMP_ARCH_MIPS ARCH_BAD
#endif

#ifndef SCMP_ARCH_MIPS64
#define SCMP_ARCH_MIPS64 ARCH_BAD
#endif

#ifndef SCMP_ARCH_MIPS64N32
#define SCMP_ARCH_MIPS64N32 ARCH_BAD
#endif

#ifndef SCMP_ARCH_MIPSEL
#define SCMP_ARCH_MIPSEL ARCH_BAD
#endif

#ifndef SCMP_ARCH_MIPSEL64
#define SCMP_ARCH_MIPSEL64 ARCH_BAD
#endif

#ifndef SCMP_ARCH_MIPSEL64N32
#define SCMP_ARCH_MIPSEL64N32 ARCH_BAD
#endif

#ifndef SCMP_ARCH_PPC
#define SCMP_ARCH_PPC ARCH_BAD
#endif

#ifndef SCMP_ARCH_PPC64
#define SCMP_ARCH_PPC64 ARCH_BAD
#endif

#ifndef SCMP_ARCH_PPC64LE
#define SCMP_ARCH_PPC64LE ARCH_BAD
#endif

#ifndef SCMP_ARCH_S390
#define SCMP_ARCH_S390 ARCH_BAD
#endif

#ifndef SCMP_ARCH_S390X
#define SCMP_ARCH_S390X ARCH_BAD
#endif

const uint32_t C_ARCH_NATIVE       = SCMP_ARCH_NATIVE;
const uint32_t C_ARCH_X86          = SCMP_ARCH_X86;
const uint32_t C_ARCH_X86_64       = SCMP_ARCH_X86_64;
const uint32_t C_ARCH_X32          = SCMP_ARCH_X32;
const uint32_t C_ARCH_ARM          = SCMP_ARCH_ARM;
const uint32_t C_ARCH_AARCH64      = SCMP_ARCH_AARCH64;
const uint32_t C_ARCH_MIPS         = SCMP_ARCH_MIPS;
const uint32_t C_ARCH_MIPS64       = SCMP_ARCH_MIPS64;
const uint32_t C_ARCH_MIPS64N32    = SCMP_ARCH_MIPS64N32;
const uint32_t C_ARCH_MIPSEL       = SCMP_ARCH_MIPSEL;
const uint32_t C_ARCH_MIPSEL64     = SCMP_ARCH_MIPSEL64;
const uint32_t C_ARCH_MIPSEL64N32  = SCMP_ARCH_MIPSEL64N32;
const uint32_t C_ARCH_PPC          = SCMP_ARCH_PPC;
const uint32_t C_ARCH_PPC64        = SCMP_ARCH_PPC64;
const uint32_t C_ARCH_PPC64LE      = SCMP_ARCH_PPC64LE;
const uint32_t C_ARCH_S390         = SCMP_ARCH_S390;
const uint32_t C_ARCH_S390X        = SCMP_ARCH_S390X;

const uint32_t C_ACT_KILL          = SCMP_ACT_KILL;
const uint32_t C_ACT_TRAP          = SCMP_ACT_TRAP;
const uint32_t C_ACT_ERRNO         = SCMP_ACT_ERRNO(0);
const uint32_t C_ACT_TRACE         = SCMP_ACT_TRACE(0);
const uint32_t C_ACT_ALLOW         = SCMP_ACT_ALLOW;

// If TSync is not supported, make sure it doesn't map to a supported filter attribute
// Don't worry about major version < 2, the minimum version checks should catch that case
#if SCMP_VER_MAJOR == 2 && SCMP_VER_MINOR < 2
#define SCMP_FLTATR_CTL_TSYNC _SCMP_CMP_MIN
#endif

const uint32_t C_ATTRIBUTE_DEFAULT = (uint32_t)SCMP_FLTATR_ACT_DEFAULT;
const uint32_t C_ATTRIBUTE_BADARCH = (uint32_t)SCMP_FLTATR_ACT_BADARCH;
const uint32_t C_ATTRIBUTE_NNP     = (uint32_t)SCMP_FLTATR_CTL_NNP;
const uint32_t C_ATTRIBUTE_TSYNC   = (uint32_t)SCMP_FLTATR_CTL_TSYNC;

const int      C_CMP_NE            = (int)SCMP_CMP_NE;
const int      C_CMP_LT            = (int)SCMP_CMP_LT;
const int      C_CMP_LE            = (int)SCMP_CMP_LE;
const int      C_CMP_EQ            = (int)SCMP_CMP_EQ;
const int      C_CMP_GE            = (int)SCMP_CMP_GE;
const int      C_CMP_GT            = (int)SCMP_CMP_GT;
const int      C_CMP_MASKED_EQ     = (int)SCMP_CMP_MASKED_EQ;

const int      C_VERSION_MAJOR     = SCMP_VER_MAJOR;
const int      C_VERSION_MINOR     = SCMP_VER_MINOR;
const int      C_VERSION_MICRO     = SCMP_VER_MICRO;

typedef struct scmp_arg_cmp* scmp_cast_t;

// Wrapper to create an scmp_arg_cmp struct
void*
make_struct_arg_cmp(
                    unsigned int arg,
                    int compare,
                    uint64_t a,
                    uint64_t b
                   )
***REMOVED***
	struct scmp_arg_cmp *s = malloc(sizeof(struct scmp_arg_cmp));

	s->arg = arg;
	s->op = compare;
	s->datum_a = a;
	s->datum_b = b;

	return s;
***REMOVED***
*/
import "C"

// Nonexported types
type scmpFilterAttr uint32

// Nonexported constants

const (
	filterAttrActDefault scmpFilterAttr = iota
	filterAttrActBadArch scmpFilterAttr = iota
	filterAttrNNP        scmpFilterAttr = iota
	filterAttrTsync      scmpFilterAttr = iota
)

const (
	// An error return from certain libseccomp functions
	scmpError C.int = -1
	// Comparison boundaries to check for architecture validity
	archStart ScmpArch = ArchNative
	archEnd   ScmpArch = ArchS390X
	// Comparison boundaries to check for action validity
	actionStart ScmpAction = ActKill
	actionEnd   ScmpAction = ActAllow
	// Comparison boundaries to check for comparison operator validity
	compareOpStart ScmpCompareOp = CompareNotEqual
	compareOpEnd   ScmpCompareOp = CompareMaskedEqual
)

var (
	// Error thrown on bad filter context
	errBadFilter = fmt.Errorf("filter is invalid or uninitialized")
	// Constants representing library major, minor, and micro versions
	verMajor = int(C.C_VERSION_MAJOR)
	verMinor = int(C.C_VERSION_MINOR)
	verMicro = int(C.C_VERSION_MICRO)
)

// Nonexported functions

// Check if library version is greater than or equal to the given one
func checkVersionAbove(major, minor, micro int) bool ***REMOVED***
	return (verMajor > major) ||
		(verMajor == major && verMinor > minor) ||
		(verMajor == major && verMinor == minor && verMicro >= micro)
***REMOVED***

// Init function: Verify library version is appropriate
func init() ***REMOVED***
	if !checkVersionAbove(2, 1, 0) ***REMOVED***
		fmt.Fprintf(os.Stderr, "Libseccomp version too low: minimum supported is 2.1.0, detected %d.%d.%d", C.C_VERSION_MAJOR, C.C_VERSION_MINOR, C.C_VERSION_MICRO)
		os.Exit(-1)
	***REMOVED***
***REMOVED***

// Filter helpers

// Filter finalizer - ensure that kernel context for filters is freed
func filterFinalizer(f *ScmpFilter) ***REMOVED***
	f.Release()
***REMOVED***

// Get a raw filter attribute
func (f *ScmpFilter) getFilterAttr(attr scmpFilterAttr) (C.uint32_t, error) ***REMOVED***
	f.lock.Lock()
	defer f.lock.Unlock()

	if !f.valid ***REMOVED***
		return 0x0, errBadFilter
	***REMOVED***

	if !checkVersionAbove(2, 2, 0) && attr == filterAttrTsync ***REMOVED***
		return 0x0, fmt.Errorf("the thread synchronization attribute is not supported in this version of the library")
	***REMOVED***

	var attribute C.uint32_t

	retCode := C.seccomp_attr_get(f.filterCtx, attr.toNative(), &attribute)
	if retCode != 0 ***REMOVED***
		return 0x0, syscall.Errno(-1 * retCode)
	***REMOVED***

	return attribute, nil
***REMOVED***

// Set a raw filter attribute
func (f *ScmpFilter) setFilterAttr(attr scmpFilterAttr, value C.uint32_t) error ***REMOVED***
	f.lock.Lock()
	defer f.lock.Unlock()

	if !f.valid ***REMOVED***
		return errBadFilter
	***REMOVED***

	if !checkVersionAbove(2, 2, 0) && attr == filterAttrTsync ***REMOVED***
		return fmt.Errorf("the thread synchronization attribute is not supported in this version of the library")
	***REMOVED***

	retCode := C.seccomp_attr_set(f.filterCtx, attr.toNative(), value)
	if retCode != 0 ***REMOVED***
		return syscall.Errno(-1 * retCode)
	***REMOVED***

	return nil
***REMOVED***

// DOES NOT LOCK OR CHECK VALIDITY
// Assumes caller has already done this
// Wrapper for seccomp_rule_add_... functions
func (f *ScmpFilter) addRuleWrapper(call ScmpSyscall, action ScmpAction, exact bool, cond C.scmp_cast_t) error ***REMOVED***
	var length C.uint
	if cond != nil ***REMOVED***
		length = 1
	***REMOVED*** else ***REMOVED***
		length = 0
	***REMOVED***

	var retCode C.int
	if exact ***REMOVED***
		retCode = C.seccomp_rule_add_exact_array(f.filterCtx, action.toNative(), C.int(call), length, cond)
	***REMOVED*** else ***REMOVED***
		retCode = C.seccomp_rule_add_array(f.filterCtx, action.toNative(), C.int(call), length, cond)
	***REMOVED***

	if syscall.Errno(-1*retCode) == syscall.EFAULT ***REMOVED***
		return fmt.Errorf("unrecognized syscall")
	***REMOVED*** else if syscall.Errno(-1*retCode) == syscall.EPERM ***REMOVED***
		return fmt.Errorf("requested action matches default action of filter")
	***REMOVED*** else if retCode != 0 ***REMOVED***
		return syscall.Errno(-1 * retCode)
	***REMOVED***

	return nil
***REMOVED***

// Generic add function for filter rules
func (f *ScmpFilter) addRuleGeneric(call ScmpSyscall, action ScmpAction, exact bool, conds []ScmpCondition) error ***REMOVED***
	f.lock.Lock()
	defer f.lock.Unlock()

	if !f.valid ***REMOVED***
		return errBadFilter
	***REMOVED***

	if len(conds) == 0 ***REMOVED***
		if err := f.addRuleWrapper(call, action, exact, nil); err != nil ***REMOVED***
			return err
		***REMOVED***
	***REMOVED*** else ***REMOVED***
		// We don't support conditional filtering in library version v2.1
		if !checkVersionAbove(2, 2, 1) ***REMOVED***
			return fmt.Errorf("conditional filtering requires libseccomp version >= 2.2.1")
		***REMOVED***

		for _, cond := range conds ***REMOVED***
			cmpStruct := C.make_struct_arg_cmp(C.uint(cond.Argument), cond.Op.toNative(), C.uint64_t(cond.Operand1), C.uint64_t(cond.Operand2))
			defer C.free(cmpStruct)

			if err := f.addRuleWrapper(call, action, exact, C.scmp_cast_t(cmpStruct)); err != nil ***REMOVED***
				return err
			***REMOVED***
		***REMOVED***
	***REMOVED***

	return nil
***REMOVED***

// Generic Helpers

// Helper - Sanitize Arch token input
func sanitizeArch(in ScmpArch) error ***REMOVED***
	if in < archStart || in > archEnd ***REMOVED***
		return fmt.Errorf("unrecognized architecture")
	***REMOVED***

	if in.toNative() == C.C_ARCH_BAD ***REMOVED***
		return fmt.Errorf("architecture is not supported on this version of the library")
	***REMOVED***

	return nil
***REMOVED***

func sanitizeAction(in ScmpAction) error ***REMOVED***
	inTmp := in & 0x0000FFFF
	if inTmp < actionStart || inTmp > actionEnd ***REMOVED***
		return fmt.Errorf("unrecognized action")
	***REMOVED***

	if inTmp != ActTrace && inTmp != ActErrno && (in&0xFFFF0000) != 0 ***REMOVED***
		return fmt.Errorf("highest 16 bits must be zeroed except for Trace and Errno")
	***REMOVED***

	return nil
***REMOVED***

func sanitizeCompareOp(in ScmpCompareOp) error ***REMOVED***
	if in < compareOpStart || in > compareOpEnd ***REMOVED***
		return fmt.Errorf("unrecognized comparison operator")
	***REMOVED***

	return nil
***REMOVED***

func archFromNative(a C.uint32_t) (ScmpArch, error) ***REMOVED***
	switch a ***REMOVED***
	case C.C_ARCH_X86:
		return ArchX86, nil
	case C.C_ARCH_X86_64:
		return ArchAMD64, nil
	case C.C_ARCH_X32:
		return ArchX32, nil
	case C.C_ARCH_ARM:
		return ArchARM, nil
	case C.C_ARCH_NATIVE:
		return ArchNative, nil
	case C.C_ARCH_AARCH64:
		return ArchARM64, nil
	case C.C_ARCH_MIPS:
		return ArchMIPS, nil
	case C.C_ARCH_MIPS64:
		return ArchMIPS64, nil
	case C.C_ARCH_MIPS64N32:
		return ArchMIPS64N32, nil
	case C.C_ARCH_MIPSEL:
		return ArchMIPSEL, nil
	case C.C_ARCH_MIPSEL64:
		return ArchMIPSEL64, nil
	case C.C_ARCH_MIPSEL64N32:
		return ArchMIPSEL64N32, nil
	case C.C_ARCH_PPC:
		return ArchPPC, nil
	case C.C_ARCH_PPC64:
		return ArchPPC64, nil
	case C.C_ARCH_PPC64LE:
		return ArchPPC64LE, nil
	case C.C_ARCH_S390:
		return ArchS390, nil
	case C.C_ARCH_S390X:
		return ArchS390X, nil
	default:
		return 0x0, fmt.Errorf("unrecognized architecture")
	***REMOVED***
***REMOVED***

// Only use with sanitized arches, no error handling
func (a ScmpArch) toNative() C.uint32_t ***REMOVED***
	switch a ***REMOVED***
	case ArchX86:
		return C.C_ARCH_X86
	case ArchAMD64:
		return C.C_ARCH_X86_64
	case ArchX32:
		return C.C_ARCH_X32
	case ArchARM:
		return C.C_ARCH_ARM
	case ArchARM64:
		return C.C_ARCH_AARCH64
	case ArchMIPS:
		return C.C_ARCH_MIPS
	case ArchMIPS64:
		return C.C_ARCH_MIPS64
	case ArchMIPS64N32:
		return C.C_ARCH_MIPS64N32
	case ArchMIPSEL:
		return C.C_ARCH_MIPSEL
	case ArchMIPSEL64:
		return C.C_ARCH_MIPSEL64
	case ArchMIPSEL64N32:
		return C.C_ARCH_MIPSEL64N32
	case ArchPPC:
		return C.C_ARCH_PPC
	case ArchPPC64:
		return C.C_ARCH_PPC64
	case ArchPPC64LE:
		return C.C_ARCH_PPC64LE
	case ArchS390:
		return C.C_ARCH_S390
	case ArchS390X:
		return C.C_ARCH_S390X
	case ArchNative:
		return C.C_ARCH_NATIVE
	default:
		return 0x0
	***REMOVED***
***REMOVED***

// Only use with sanitized ops, no error handling
func (a ScmpCompareOp) toNative() C.int ***REMOVED***
	switch a ***REMOVED***
	case CompareNotEqual:
		return C.C_CMP_NE
	case CompareLess:
		return C.C_CMP_LT
	case CompareLessOrEqual:
		return C.C_CMP_LE
	case CompareEqual:
		return C.C_CMP_EQ
	case CompareGreaterEqual:
		return C.C_CMP_GE
	case CompareGreater:
		return C.C_CMP_GT
	case CompareMaskedEqual:
		return C.C_CMP_MASKED_EQ
	default:
		return 0x0
	***REMOVED***
***REMOVED***

func actionFromNative(a C.uint32_t) (ScmpAction, error) ***REMOVED***
	aTmp := a & 0xFFFF
	switch a & 0xFFFF0000 ***REMOVED***
	case C.C_ACT_KILL:
		return ActKill, nil
	case C.C_ACT_TRAP:
		return ActTrap, nil
	case C.C_ACT_ERRNO:
		return ActErrno.SetReturnCode(int16(aTmp)), nil
	case C.C_ACT_TRACE:
		return ActTrace.SetReturnCode(int16(aTmp)), nil
	case C.C_ACT_ALLOW:
		return ActAllow, nil
	default:
		return 0x0, fmt.Errorf("unrecognized action")
	***REMOVED***
***REMOVED***

// Only use with sanitized actions, no error handling
func (a ScmpAction) toNative() C.uint32_t ***REMOVED***
	switch a & 0xFFFF ***REMOVED***
	case ActKill:
		return C.C_ACT_KILL
	case ActTrap:
		return C.C_ACT_TRAP
	case ActErrno:
		return C.C_ACT_ERRNO | (C.uint32_t(a) >> 16)
	case ActTrace:
		return C.C_ACT_TRACE | (C.uint32_t(a) >> 16)
	case ActAllow:
		return C.C_ACT_ALLOW
	default:
		return 0x0
	***REMOVED***
***REMOVED***

// Internal only, assumes safe attribute
func (a scmpFilterAttr) toNative() uint32 ***REMOVED***
	switch a ***REMOVED***
	case filterAttrActDefault:
		return uint32(C.C_ATTRIBUTE_DEFAULT)
	case filterAttrActBadArch:
		return uint32(C.C_ATTRIBUTE_BADARCH)
	case filterAttrNNP:
		return uint32(C.C_ATTRIBUTE_NNP)
	case filterAttrTsync:
		return uint32(C.C_ATTRIBUTE_TSYNC)
	default:
		return 0x0
	***REMOVED***
***REMOVED***
