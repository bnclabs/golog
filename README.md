Basic logging with batteries
----------------------------

[![Build Status](https://travis-ci.org/prataprc/golog.png)](https://travis-ci.org/prataprc/golog)
[![Coverage Status](https://coveralls.io/repos/prataprc/golog/badge.png?branch=master&service=github)](https://coveralls.io/github/prataprc/golog?branch=master)
[![GoDoc](https://godoc.org/github.com/prataprc/golog?status.png)](https://godoc.org/github.com/prataprc/golog)

* APIs to prefix log-level in log messages.
* Global option to redirect logs to a file.
* Include/Exclude/Format log time.
* Colorize log messages for different levels.
* Console logging.

Packages can import log and use its methods

```go

    import github.com/prataprc/golog

    func myfunc() {
        ..
        log.Fatalf(...)
        ..
        log.Warnf(...)
        ..
        log.Debugf(...)
    }
```

Note here that *log* is not an object name, it resolves to the imported *log*
package that has exported methods *Fatalf()* *Warnf()* etc ... For more
information please read the go-documentation for *log* package.

By default, importing the package will initialize the logger to
default-logger that shall log to standard output. To use custom logger
use the following initializer function in your package or application:

```go

    import github.com/prataprc/golog

    var mylogger = newmylogger()

    func init() {
        setts := map[string]interface{}{
            "log.level": "info",
            "log.file":  "",
        }
        SetLogger(mylogger, setts)
    }
```

*mylogger* should implement the *log.Logger* interface{}.

**Order of log levels**

```golang
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
```

Console Logging
---------------

By default log APIs will worry about log-level, prefix format, time-format
sometimes it become too much of clutter on the screen to communicate simple
messages with user via console. In such cases use the ``Consolef`` API.

```go
    log.Consolef("goledger version - goledger%v\n", api.LedgerVersion)
```

``Consolef`` does not print the log time, log level and always outputs to
stdout.

Settings
--------

* **log.level**, filter all messages logged at level greater than the
configured value. Can be one of the following names -
ignore, fatal, error, warn, info, verbose, debug, trace
* **log.file**, if not empty string, all log messages are appended to
configured file.
* **log.timeformat**, format of time string prefixed to log message,
should confirm to ``time.Now().Format()``.
* **log.prefix**, ``fmt.Sprintf`` format string for log level, by
default ``[<leve>]`` format is used.
* **log.colorfatal**, comma separated value of attribute names -
bold, underline, blinkslow, blinkrapid, crossedout, red, green,
yellow, blue, magenta, cyan, white, hired, higreen, hiyellow, hiblue,
himagenta, hicyan, hiwhite. Attribute-settings available for all log
levels.

**Ignore** ignore level can be used to ignore all log messages. Note that
inly log-level can be specified as ``ignore``, no corresponding API
is supported.

Panic cases
-----------

* API ``SetLogger()``

  * if ``log.file`` is not string.
  * if creating or opening ``log.file`` fails.
  * if ``log.level`` is not an allowed log string.
  * if ``log.prefix`` is neither string, nor bool.

* API ``SetLogLevel()``

  * if ``log.level`` is not an allowed log string.

* API ``SetLogprefix()``

  * if ``log.prefix`` is neither string, nor bool.

Typically all the above panic cases needs to be fixed during development, and
should never occur during production. If panics become unavoidable please use
[panic/recover](https://blog.golang.org/defer-panic-and-recover).

How to contribute
-----------------

* Pick an issue, or create an new issue. Provide adequate documentation for
the issue.
* Assign the issue or get it assigned.
* Work on the code, once finished, raise a pull request.
* Golog is written in [golang](https://golang.org/), hence expected to follow the
global guidelines for writing go programs.
