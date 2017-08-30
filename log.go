//  Copyright (c) 2014 Couchbase, Inc.

package log

import "os"
import "fmt"
import "time"
import "strings"
import stdlog "log"

import "github.com/prataprc/color"

var timeformat, prefix = "2006-01-02T15:04:05.999Z-07:00", "[%v]"

func init() {
	setts := map[string]interface{}{
		"log.level":        "info",
		"log.flags":        "",
		"log.file":         "",
		"log.timeformat":   timeformat,
		"log.prefix":       prefix,
		"log.colorignore":  "",
		"log.colorfatal":   "red",
		"log.colorerror":   "hired",
		"log.colorwarn":    "yellow",
		"log.colorinfo":    "",
		"log.colorverbose": "",
		"log.colordebug":   "",
		"log.colortrace":   "",
	}
	SetLogger(nil, setts)
}

var log Logger // can be used used off-the-shelf.

// DefaultLogLevel to use if log.level option is missing.
var DefaultLogLevel = "info"

// Logger interface for application logging, applications can
// supply a logger object implementing this interface, otherwise,
// defaultLogger{} will be used.
type Logger interface {
	// SetLogLevel application's global log level, can be one of the
	// following: "ignore", "fatal", "error", "warn", "info", "verbose",
	// "debug", "trace"
	SetLogLevel(string)

	// SetLogFlags format of log prefix, following golang's log:Flag()
	// specification
	SetLogFlags(flags int)

	// SetTimeFormat to use as prefix for all log messages.
	SetTimeFormat(string)

	// SetLogprefix including the log level.
	SetLogprefix(interface{})

	// SetLogcolor sets coloring attributes for specified log level, can be
	// a list of following attributes: "bold", "underline", "blinkslow",
	// "blinkrapid", "crossedout",
	// "red", "green", "yellow", "blue", "magenta", "cyan", "white"
	// "hired", "higreen", "hiyellow", "hiblue", "himagenta", "hicyan",
	// "hiwhite"
	SetLogcolor(level string, attrs []string)

	// Fatalf similar to Printf, will be logged only when log level is set as
	// "fatal" or above.
	Fatalf(format string, v ...interface{})

	// Errorf similar to Printf, will be logged only when log level is set as
	// "error" or above.
	Errorf(format string, v ...interface{})

	// Warnf similar to Printf, will be logged only when log level is set as
	// "warn" or above.
	Warnf(format string, v ...interface{})

	// Infof similar to Printf, will be logged only when log level is set as
	// "info" or above.
	Infof(format string, v ...interface{})

	// Verbosef similar to Printf, will be logged only when log level is set as
	// "verbose" or above.
	Verbosef(format string, v ...interface{})

	// Debugf similar to Printf, will be logged only when log level is set as
	// "debug" or above.
	Debugf(format string, v ...interface{})

	// Tracef similar to Printf, will be logged only when log level is set as
	// "trace" or above.
	Tracef(format string, v ...interface{})

	// Printlf reserved for future extension.
	Printlf(loglevel LogLevel, format string, v ...interface{})
}

// LogLevel defines application log level.
type LogLevel int

const (
	logLevelIgnore LogLevel = iota + 1
	logLevelFatal
	logLevelError
	logLevelWarn
	logLevelInfo
	logLevelVerbose
	logLevelDebug
	logLevelTrace
)

// SetLogger to integrate storage logging with application logging.
// importing this package will initialize the logger with info level
// logging to console.
func SetLogger(logger Logger, setts map[string]interface{}) Logger {
	if logger != nil {
		log = logger
		return log
	}

	var err error

	logfd := os.Stdout
	if logfile, ok := setts["log.file"]; ok {
		filename := logfile.(string)
		if filename != "" {
			logfd, err = os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)
			if err != nil {
				if logfd, err = os.Create(filename); err != nil {
					panic(err)
				}
			}
		}
	}
	stdlog.SetOutput(logfd)

	deflog := &defaultLogger{
		timeformat: timeformat, prefix: prefix,
		colors: make(map[LogLevel]*color.Color),
	}

	level, ok := setts["log.level"]
	if ok == false {
		level = "info"
	}
	deflog.SetLogLevel(level.(string))

	logflags := int(0)
	if flags, ok := setts["log.flags"]; ok {
		for _, flag := range parsecsv(flags.(string)) {
			logflags |= string2flag(flag)
		}
		deflog.SetLogFlags(logflags)
	}

	if logflags == 0 {
		if timeformat, ok := setts["log.timeformat"]; ok {
			deflog.SetTimeFormat(timeformat.(string))
		}
	} else { // if flags are available disable timeformat.
		deflog.SetTimeFormat("")
	}

	if prefix, ok := setts["log.prefix"]; ok {
		deflog.SetLogprefix(prefix)
	}

	// colors
	params := []string{"log.colorignore", "log.colorfatal", "log.colorerror",
		"log.colorwarn", "log.colorinfo", "log.colorverbose", "log.colordebug",
		"log.colortrace"}
	for _, param := range params {
		level := param[9:]
		if val, ok := setts[param]; ok {
			if v1, ok := val.(string); ok {
				deflog.SetLogcolor(level, parsecsv(v1))
			} else if v2, ok := val.([]string); ok {
				deflog.SetLogcolor(level, v2)
			} else {
				fmsg := "invalid type: color parameter %q has %T"
				panic(fmt.Errorf(fmsg, param, val))
			}
		}
	}

	log = deflog
	return log
}

// defaultLogger with default log-file as os.Stdout and,
// default log-level as logLevelInfo. Applications can
// supply a Logger{} object when instantiating the
// Transport.
type defaultLogger struct {
	level      LogLevel
	timeformat string
	prefix     string
	colors     map[LogLevel]*color.Color
}

// SetLogLevel for defaultLogger.
func (l *defaultLogger) SetLogLevel(level string) {
	l.level = string2logLevel(level)
}

// SetLogFlags for defaultLogger.
func (l *defaultLogger) SetLogFlags(flags int) {
	stdlog.SetFlags(flags)
}

// SetTimeFormat for defaultLogger.
func (l *defaultLogger) SetTimeFormat(format string) {
	l.timeformat = format
}

// SetLogprefix for defaultLogger
func (l *defaultLogger) SetLogprefix(prefix interface{}) {
	if val, ok := prefix.(string); ok {
		l.prefix = val
	} else if _, ok = prefix.(bool); ok {
		l.prefix = ""
	} else {
		panic("level-prefix can either be string format, or bool")
	}
}

// SetLogcolor for defaultLogger
func (l *defaultLogger) SetLogcolor(level string, attrs []string) {
	ll := string2logLevel(level)
	attributes := []color.Attribute{}
	for _, attr := range attrs {
		attributes = append(attributes, string2clrattr(attr))
	}
	l.colors[ll] = color.New(attributes...)
}

// Fatalf for defaultLogger
func (l *defaultLogger) Fatalf(format string, v ...interface{}) {
	l.Printlf(logLevelFatal, format, v...)
}

// Errorf for defaultLogger
func (l *defaultLogger) Errorf(format string, v ...interface{}) {
	l.Printlf(logLevelError, format, v...)
}

// Warnf for defaultLogger
func (l *defaultLogger) Warnf(format string, v ...interface{}) {
	l.Printlf(logLevelWarn, format, v...)
}

// Infof for defaultLogger
func (l *defaultLogger) Infof(format string, v ...interface{}) {
	l.Printlf(logLevelInfo, format, v...)
}

// Verbosef for defaultLogger
func (l *defaultLogger) Verbosef(format string, v ...interface{}) {
	l.Printlf(logLevelVerbose, format, v...)
}

// Debugf for defaultLogger
func (l *defaultLogger) Debugf(format string, v ...interface{}) {
	l.Printlf(logLevelDebug, format, v...)
}

// Tracef for defaultLogger
func (l *defaultLogger) Tracef(format string, v ...interface{}) {
	l.Printlf(logLevelTrace, format, v...)
}

// Printlf for defaultLogger
func (l *defaultLogger) Printlf(level LogLevel, frmt string, v ...interface{}) {
	if l.canlog(level) {
		prefix := ""
		if l.timeformat != "" {
			prefix = time.Now().Format(l.timeformat) + " "
		}
		if lstr := level.String(); lstr != "" && l.prefix != "" {
			prefix += fmt.Sprintf(l.prefix, level.String()) + " "
		}
		newv := []interface{}{prefix}
		newv = append(newv, v...)
		frmt := trimformat(frmt) // output appends a newline because of color
		if color, ok := l.colors[level]; ok && color != nil {
			stdlog.Output(3, color.Sprintf("%v"+frmt, newv...))
		} else {
			stdlog.Output(3, fmt.Sprintf("%v"+frmt, newv...))
		}
	}
}

func (l *defaultLogger) canlog(level LogLevel) bool {
	if level <= l.level {
		return true
	}
	return false
}

func (l LogLevel) String() string {
	switch l {
	case logLevelIgnore:
		return "Ignor"
	case logLevelFatal:
		return "Fatal"
	case logLevelError:
		return "Error"
	case logLevelWarn:
		return "Warng"
	case logLevelInfo:
		return "Infom"
	case logLevelVerbose:
		return "Verbs"
	case logLevelDebug:
		return "Debug"
	case logLevelTrace:
		return "Trace"
	}
	panic("unexpected log level") // should never reach here
}

func string2logLevel(s string) LogLevel {
	s = strings.ToLower(s)
	switch s {
	case "ignore":
		return logLevelIgnore
	case "fatal":
		return logLevelFatal
	case "error":
		return logLevelError
	case "warn":
		return logLevelWarn
	case "info":
		return logLevelInfo
	case "verbose":
		return logLevelVerbose
	case "debug":
		return logLevelDebug
	case "trace":
		return logLevelTrace
	}
	panic(fmt.Errorf("unexpected log level: %q", s)) // never reach here
}

func string2clrattr(s string) color.Attribute {
	s = strings.ToLower(s)
	switch s {
	case "bold":
		return color.Bold
	case "underline":
		return color.Underline
	case "blinkslow":
		return color.BlinkSlow
	case "blinkrapid":
		return color.BlinkRapid
	case "crossedout":
		return color.CrossedOut
	case "red":
		return color.FgRed
	case "green":
		return color.FgGreen
	case "yellow":
		return color.FgYellow
	case "blue":
		return color.FgBlue
	case "magenta":
		return color.FgMagenta
	case "cyan":
		return color.FgCyan
	case "white":
		return color.FgWhite
	case "hired":
		return color.FgHiRed
	case "higreen":
		return color.FgHiGreen
	case "hiyellow":
		return color.FgHiYellow
	case "hiblue":
		return color.FgHiBlue
	case "himagenta":
		return color.FgHiMagenta
	case "hicyan":
		return color.FgHiCyan
	case "hiwhite":
		return color.FgHiWhite
	}
	panic(fmt.Errorf("unexpected color attribute %q", s)) // never reach here
}

func string2flag(s string) int {
	s = strings.ToLower(s)
	switch s {
	case "ldate":
		return stdlog.Ldate
	case "ltime":
		return stdlog.Ltime
	case "lmicroseconds":
		return stdlog.Lmicroseconds
	case "llongfile":
		return stdlog.Llongfile
	case "lshortfile":
		return stdlog.Lshortfile
	case "lutc":
		return stdlog.LUTC
	case "lstdflags":
		return stdlog.LstdFlags
	}
	panic(fmt.Errorf("unexpected color attribute %q", s)) // never reach here
}

// Fatalf similar to Printf, will be logged only when log level is set as
// "fatal" or above.
func Fatalf(format string, v ...interface{}) {
	log.Printlf(logLevelFatal, format, v...)
	panic(fmt.Errorf(format, v...))
}

// Errorf similar to Printf, will be logged only when log level is set as
// "error" or above.
func Errorf(format string, v ...interface{}) {
	log.Printlf(logLevelError, format, v...)
}

// Warnf similar to Printf, will be logged only when log level is set as
// "warn" or above.
func Warnf(format string, v ...interface{}) {
	log.Printlf(logLevelWarn, format, v...)
}

// Infof similar to Printf, will be logged only when log level is set as
// "info" or above.
func Infof(format string, v ...interface{}) {
	log.Printlf(logLevelInfo, format, v...)
}

// Verbosef similar to Printf, will be logged only when log level is set as
// "verbose" or above.
func Verbosef(format string, v ...interface{}) {
	log.Printlf(logLevelVerbose, format, v...)
}

// Debugf similar to Printf, will be logged only when log level is set as
// "debug" or above.
func Debugf(format string, v ...interface{}) {
	log.Printlf(logLevelDebug, format, v...)
}

// Tracef similar to Printf, will be logged only when log level is set as
// "trace" or above.
func Tracef(format string, v ...interface{}) {
	log.Printlf(logLevelTrace, format, v...)
}

// Consolef similar to Printf, will log to os.Stdout.
func Consolef(format string, v ...interface{}) {
	fmt.Fprintf(os.Stdout, format, v...)
}

func parsecsv(input string) []string {
	if input == "" {
		return nil
	}
	ss := strings.Split(input, ",")
	outs := make([]string, 0)
	for _, s := range ss {
		s = strings.Trim(s, " \t\r\n")
		if s == "" {
			continue
		}
		outs = append(outs, s)
	}
	return outs
}

func trimformat(frmt string) string {
	if frmt[len(frmt)-1] == '\n' {
		return frmt[:len(frmt)-1]
	}
	return frmt
}
