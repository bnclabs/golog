/*
Package golog provides a simple alternative to standard log that
defines a Logger interface allowing custom logger to be used
across the application and its libraries.

To begin with, import the golog package and start using its exported
functions like:

	import "github.com/prataprc/golog"
	...
	log.Printf()
	log.Fatalf()

Default logger will be used in the above case. To configure default
logger:

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
	log.SetLogger(nil, setts)

Default logger will be configured to "info" level. Refer
Defaultsettings() for description on each settings parameter.

To configure a custom logger:

	log.Setlogger(customlogger, nil)

Note that `customlogger` should implement the Logger interface.
*/
package log
