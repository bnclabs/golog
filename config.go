package log

/*
Defaultsettings used on default logger.

log.level: (default "info")
	Level can be one of the following string,
	"ignore", "fatal", "error", "warn", "info", "verbose",
	"debug", "trace".

log.flags: (default "")
    Flags can be comma seperated string values. These flags as exactly
    same as the golang's standard logger.
    "ldate" the date in the local time zone: 2009/01/23
    "ltime" the time in the local time zone: 01:23:23
    "lmicroseconds" microsecond resolution: 01:23:23.123123.
    "llongfile" full file name and line number: /a/b/c/d.go:23
    "lshortfile" final file name element and line number: d.go:23.
    "lutc" if Ldate or Ltime is set, use UTC rather than the local time zone
    "lstdflags" initial values for the standard logger ldate,ltime

log.file: (default os.Stdout)
    Optional log file name to log o/p. Except Consolef all functions
    will o/p to this file if supplied, else to standard output.

log.timeformat: "2006-01-02T15:04:05.999Z-07:00"
	Log line timeformat.

log.prefix: [%v]
	Prefix format for log-level.

log.colorfatal: "red"
	Output color for fatal level.

log.colorerror: "hired"
	Output color for error level.

log.colorwarn: "yellow"
	Output color for warn level.

log.colorinfo: ""
	Output color for info level.

log.colorverbose: "",
	Output color for verbose level.

log.colordebug: "",
	Output color for debug level.

log.colortrace: "",
	Output color for trace level.
*/
func Defaultsettings() map[string]interface{} {
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
	return setts
}
