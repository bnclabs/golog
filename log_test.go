package log

import "testing"
import "fmt"
import "os"
import "strings"
import "io/ioutil"
import stdlog "log"

import "github.com/prataprc/color"

func TestSetLogger(t *testing.T) {
	logfile := "setlogger_test.log.file"
	logline := "hello world"
	defer os.Remove(logfile)

	ref := &defaultLogger{level: logLevelIgnore}
	log := SetLogger(ref, nil).(*defaultLogger)
	if log.level != logLevelIgnore {
		t.Errorf("expected %v, got %v", ref, log)
	}

	// test a custom logger
	setts := map[string]interface{}{
		"log.level":      "info",
		"log.file":       logfile,
		"log.flags":      "ldate,lshortfile",
		"log.timeformat": timeformat,
	}
	clog := SetLogger(nil, setts)
	clog.Infof(logline)
	clog.Verbosef(logline)
	clog.Fatalf(logline)
	clog.Errorf(logline)
	clog.Warnf(logline)
	clog.Tracef(logline)
	if data, err := ioutil.ReadFile(logfile); err != nil {
		t.Error(err)
	} else if s := string(data); !strings.Contains(s, "hello world") {
		t.Errorf("expected %v, got %v", logline, s)
	} else if len(strings.Split(s, "\n")) != 5 {
		t.Errorf("expected %v, got %v", logline, s)
	}
}

func TestLogTimeformat(t *testing.T) {
	timeformat := "2006"
	setts := map[string]interface{}{
		"log.level":      "info",
		"log.timeformat": timeformat,
	}
	log := SetLogger(nil, setts).(*defaultLogger)
	if log.timeformat != timeformat {
		t.Errorf("expected %v, got %v", timeformat, log.timeformat)
	}
}

func TestLogColor(t *testing.T) {
	attrs := []string{"red", "blinkslow"}
	setts := map[string]interface{}{
		"log.level":      "info",
		"log.colorfatal": attrs,
	}
	log := SetLogger(nil, setts).(*defaultLogger)
	s := fmt.Sprintf("%T", log.colors[logLevelFatal])
	if s != "*color.Color" {
		t.Errorf("expected *color.Color, %v", s)
	}
}

func TestSetLogPrefix(t *testing.T) {
	setts := map[string]interface{}{"log.prefix": "[%v]"}
	log := SetLogger(nil, setts).(*defaultLogger)
	if log.prefix != "[%v]" {
		t.Errorf("expected %v, got %v", "[%v]", log.prefix)
	}
	log.SetLogprefix(false)
	if log.prefix != "" {
		t.Errorf("expected empty prefix, %v", log.prefix)
	}
}

func TestLogPrefix(t *testing.T) {
	if ref, s := "Ignor", logLevelIgnore.String(); ref != s {
		t.Errorf("expected %v, got %v", ref, s)
	} else if ref, s = "Fatal", logLevelFatal.String(); ref != s {
		t.Errorf("expected %v, got %v", ref, s)
	} else if ref, s = "Error", logLevelError.String(); ref != s {
		t.Errorf("expected %v, got %v", ref, s)
	} else if ref, s = "Warng", logLevelWarn.String(); ref != s {
		t.Errorf("expected %v, got %v", ref, s)
	} else if ref, s = "Infom", logLevelInfo.String(); ref != s {
		t.Errorf("expected %v, got %v", ref, s)
	} else if ref, s = "Verbs", logLevelVerbose.String(); ref != s {
		t.Errorf("expected %v, got %v", ref, s)
	} else if ref, s = "Debug", logLevelDebug.String(); ref != s {
		t.Errorf("expected %v, got %v", ref, s)
	} else if ref, s = "Trace", logLevelTrace.String(); ref != s {
		t.Errorf("expected %v, got %v", ref, s)
	}
}

func TestLogLevelSettings(t *testing.T) {
	if r, l := logLevelIgnore, string2logLevel("ignore"); r != l {
		t.Errorf("expected %v, got %v", r, l)
	} else if r, l = logLevelFatal, string2logLevel("fatal"); r != l {
		t.Errorf("expected %v, got %v", r, l)
	} else if r, l = logLevelError, string2logLevel("error"); r != l {
		t.Errorf("expected %v, got %v", r, l)
	} else if r, l = logLevelWarn, string2logLevel("warn"); r != l {
		t.Errorf("expected %v, got %v", r, l)
	} else if r, l = logLevelInfo, string2logLevel("info"); r != l {
		t.Errorf("expected %v, got %v", r, l)
	} else if r, l = logLevelVerbose, string2logLevel("verbose"); r != l {
		t.Errorf("expected %v, got %v", r, l)
	} else if r, l = logLevelDebug, string2logLevel("debug"); r != l {
		t.Errorf("expected %v, got %v", r, l)
	} else if r, l = logLevelTrace, string2logLevel("trace"); r != l {
		t.Errorf("expected %v, got %v", r, l)
	}
}

func TestColorAttrs(t *testing.T) {
	testcases := [][]interface{}{
		[]interface{}{"bold", color.Bold},
		[]interface{}{"underline", color.Underline},
		[]interface{}{"blinkslow", color.BlinkSlow},
		[]interface{}{"blinkrapid", color.BlinkRapid},
		[]interface{}{"crossedout", color.CrossedOut},
		[]interface{}{"red", color.FgRed},
		[]interface{}{"green", color.FgGreen},
		[]interface{}{"yellow", color.FgYellow},
		[]interface{}{"blue", color.FgBlue},
		[]interface{}{"magenta", color.FgMagenta},
		[]interface{}{"cyan", color.FgCyan},
		[]interface{}{"white", color.FgWhite},
		[]interface{}{"hired", color.FgHiRed},
		[]interface{}{"higreen", color.FgHiGreen},
		[]interface{}{"hiyellow", color.FgHiYellow},
		[]interface{}{"hiblue", color.FgHiBlue},
		[]interface{}{"himagenta", color.FgHiMagenta},
		[]interface{}{"hicyan", color.FgHiCyan},
		[]interface{}{"hiwhite", color.FgHiWhite},
	}
	for _, tc := range testcases {
		s := tc[0].(string)
		if v := string2clrattr(s); v != tc[1].(color.Attribute) {
			t.Errorf("expected %v, got %v", tc[1], v)
		}
	}
}

func TestFlagAttr(t *testing.T) {
	testcases := [][]interface{}{
		[]interface{}{"ldate", stdlog.Ldate},
		[]interface{}{"ltime", stdlog.Ltime},
		[]interface{}{"lmicroseconds", stdlog.Lmicroseconds},
		[]interface{}{"llongfile", stdlog.Llongfile},
		[]interface{}{"lshortfile", stdlog.Lshortfile},
		[]interface{}{"lutc", stdlog.LUTC},
		[]interface{}{"lstdflags", stdlog.LstdFlags},
	}
	for _, tc := range testcases {
		s := tc[0].(string)
		if v := string2flag(s); v != tc[1].(int) {
			t.Errorf("expected %v, got %v", tc[1], v)
		}
	}
}
