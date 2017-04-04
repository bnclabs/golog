[![Build Status](https://travis-ci.org/prataprc/golog.png)](https://travis-ci.org/prataprc/golog)
[![Coverage Status](https://coveralls.io/repos/prataprc/golog/badge.png?branch=master&service=github)](https://coveralls.io/github/prataprc/golog?branch=master)
[![GoDoc](https://godoc.org/github.com/prataprc/golog?status.png)](https://godoc.org/github.com/prataprc/golog)

Packages can import log and use its methods:

.. code-block:: go

    import github.com/prataprc/golog

    func myfunc() {
        ..
        log.Fatalf(...)
        ..
        log.Warnf(...)
        ..
        log.Debugf(...)
    }

note here that *log* is not an object name, it resolves to the imported *log*
package that has exported methods *Fatalf()* *Warnf()* etc ... For more
information please read the godoc-umentation for *log* package.

By default, importing the package will initialize the logger to
default-logger that shall log to standard output. To use custom logger
use the following initializer function in your package or application:

.. code-block:: go

    import github.com/prataprc/golog

    var mylogger = newmylogger()

    func init() {
        setts := map[string]interface{}{
            "log.level": "info",
            "log.file":  "",
        }
        SetLogger(mylogger, setts)
    }

*mylogger* should implement the *log.Logger* interface{}.
