// Package log implements a wrapper around the Go standard library's
// logging package. Clients should set the current log level; only
// messages below that level will actually be logged. For example, if
// Level is set to LevelWarning, only log messages at the Warning,
// Error, and Critical levels will be logged.
package log

import (
	"fmt"
	"log"
	"os"
)

// The following constants represent logging levels in increasing levels of seriousness.
const (
	// LevelDebug is the log level for Debug statements.
	LevelDebug = iota
	// LevelInfo is the log level for Info statements.
	LevelInfo
	// LevelWarning is the log level for Warning statements.
	LevelWarning
	// LevelError is the log level for Error statements.
	LevelError
	// LevelCritical is the log level for Critical statements.
	LevelCritical
	// LevelFatal is the log level for Fatal statements.
	LevelFatal
)

var levelPrefix = [...]string***REMOVED***
	LevelDebug:    "DEBUG",
	LevelInfo:     "INFO",
	LevelWarning:  "WARNING",
	LevelError:    "ERROR",
	LevelCritical: "CRITICAL",
	LevelFatal:    "FATAL",
***REMOVED***

// Level stores the current logging level.
var Level = LevelInfo

// SyslogWriter specifies the necessary methods for an alternate output
// destination passed in via SetLogger.
//
// SyslogWriter is satisfied by *syslog.Writer.
type SyslogWriter interface ***REMOVED***
	Debug(string)
	Info(string)
	Warning(string)
	Err(string)
	Crit(string)
	Emerg(string)
***REMOVED***

// syslogWriter stores the SetLogger() parameter.
var syslogWriter SyslogWriter

// SetLogger sets the output used for output by this package.
// A *syslog.Writer is a good choice for the logger parameter.
// Call with a nil parameter to revert to default behavior.
func SetLogger(logger SyslogWriter) ***REMOVED***
	syslogWriter = logger
***REMOVED***

func print(l int, msg string) ***REMOVED***
	if l >= Level ***REMOVED***
		if syslogWriter != nil ***REMOVED***
			switch l ***REMOVED***
			case LevelDebug:
				syslogWriter.Debug(msg)
			case LevelInfo:
				syslogWriter.Info(msg)
			case LevelWarning:
				syslogWriter.Warning(msg)
			case LevelError:
				syslogWriter.Err(msg)
			case LevelCritical:
				syslogWriter.Crit(msg)
			case LevelFatal:
				syslogWriter.Emerg(msg)
			***REMOVED***
		***REMOVED*** else ***REMOVED***
			log.Printf("[%s] %s", levelPrefix[l], msg)
		***REMOVED***
	***REMOVED***
***REMOVED***

func outputf(l int, format string, v []interface***REMOVED******REMOVED***) ***REMOVED***
	print(l, fmt.Sprintf(format, v...))
***REMOVED***

func output(l int, v []interface***REMOVED******REMOVED***) ***REMOVED***
	print(l, fmt.Sprint(v...))
***REMOVED***

// Fatalf logs a formatted message at the "fatal" level and then exits. The
// arguments are handled in the same manner as fmt.Printf.
func Fatalf(format string, v ...interface***REMOVED******REMOVED***) ***REMOVED***
	outputf(LevelFatal, format, v)
	os.Exit(1)
***REMOVED***

// Fatal logs its arguments at the "fatal" level and then exits.
func Fatal(v ...interface***REMOVED******REMOVED***) ***REMOVED***
	output(LevelFatal, v)
	os.Exit(1)
***REMOVED***

// Criticalf logs a formatted message at the "critical" level. The
// arguments are handled in the same manner as fmt.Printf.
func Criticalf(format string, v ...interface***REMOVED******REMOVED***) ***REMOVED***
	outputf(LevelCritical, format, v)
***REMOVED***

// Critical logs its arguments at the "critical" level.
func Critical(v ...interface***REMOVED******REMOVED***) ***REMOVED***
	output(LevelCritical, v)
***REMOVED***

// Errorf logs a formatted message at the "error" level. The arguments
// are handled in the same manner as fmt.Printf.
func Errorf(format string, v ...interface***REMOVED******REMOVED***) ***REMOVED***
	outputf(LevelError, format, v)
***REMOVED***

// Error logs its arguments at the "error" level.
func Error(v ...interface***REMOVED******REMOVED***) ***REMOVED***
	output(LevelError, v)
***REMOVED***

// Warningf logs a formatted message at the "warning" level. The
// arguments are handled in the same manner as fmt.Printf.
func Warningf(format string, v ...interface***REMOVED******REMOVED***) ***REMOVED***
	outputf(LevelWarning, format, v)
***REMOVED***

// Warning logs its arguments at the "warning" level.
func Warning(v ...interface***REMOVED******REMOVED***) ***REMOVED***
	output(LevelWarning, v)
***REMOVED***

// Infof logs a formatted message at the "info" level. The arguments
// are handled in the same manner as fmt.Printf.
func Infof(format string, v ...interface***REMOVED******REMOVED***) ***REMOVED***
	outputf(LevelInfo, format, v)
***REMOVED***

// Info logs its arguments at the "info" level.
func Info(v ...interface***REMOVED******REMOVED***) ***REMOVED***
	output(LevelInfo, v)
***REMOVED***

// Debugf logs a formatted message at the "debug" level. The arguments
// are handled in the same manner as fmt.Printf.
func Debugf(format string, v ...interface***REMOVED******REMOVED***) ***REMOVED***
	outputf(LevelDebug, format, v)
***REMOVED***

// Debug logs its arguments at the "debug" level.
func Debug(v ...interface***REMOVED******REMOVED***) ***REMOVED***
	output(LevelDebug, v)
***REMOVED***
