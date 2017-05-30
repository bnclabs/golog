[![Build Status](https://travis-ci.org/prataprc/golog.png)](https://travis-ci.org/prataprc/golog)
[![Coverage Status](https://coveralls.io/repos/prataprc/golog/badge.png?branch=master&service=github)](https://coveralls.io/github/prataprc/golog?branch=master)
[![GoDoc](https://godoc.org/github.com/prataprc/golog?status.png)](https://godoc.org/github.com/prataprc/golog)

Basic logging with batteries
----------------------------

* APIs to prefix log level in log messages.
* Gobal option to redirect logs to a file.
* Include/Exclude/Format log time.
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

note here that *log* is not an object name, it resolves to the imported *log*
package that has exported methods *Fatalf()* *Warnf()* etc ... For more
information please read the godoc-umentation for *log* package.

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
should never occur during production. If panics become un-avoidable please use
[panic/recover](https://blog.golang.org/defer-panic-and-recover).
