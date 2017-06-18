Golog: Basic logging with batteries
===================================

R Pratap Chakravarthy <br/>
prataprc@gmail.com <br/>
https://github.com/prataprc/golog

---

Why logging ?
=============

- *To*develop*, know how the programs fit together.
- *To*debug*, fix problems while taking them to production.
- *To*communicate*, with users letting them know what is happening.
- *To*analyse*, characterize programs under production.

Let log messages be meaningful enough to serve their end.

---

By default use Golang's standard logging
========================================

Use *log* package from golang, as much as possible. Lesser the baggage better the journey :)

What you can already do with golang's log package ?

- Prefix log messages with date,time,microsecond,file,line-no.
- Fatal, Fatalf, Fatalln, Print, Printf, Println.
- Add custom prefix to all messages.
- Set output file (device) for logging, by default it goes to os.Stdout.

---

Golog: Basic logging with batteries
===================================

I use golog only if I need more that what *log* pkg already provides. And I
figured I wanted the following:

- *Log*levels*, used for filtering messages and prefixing the level information with every message.
- *Console*Logs*.
- *Color* attributes for console logs.
- *Configure* via JSON.

These facilities are supported on top of what the log pkg already provides.
I find them sufficient for my case, but if your situation demands additional
feature raise an [issue](http://github.com/prataprc/golog/issues)

---

Log levels
==========

Log levels are listed in decreasing order of *importance*. That is,
if log level is configured as *Info*, all messages logged at level lesser
than Info level shall be filtered out.

- *Ignore*, messages cannot be logged at Ignore level, so all messages logged using golog will be filtered out.
- *Fatal*, will panic after logging the message, messages logged at level Error and below will be filtered out.
- *Error*, means there was a critical error, needs supervisor attention.
- *Warn*, means there was an unexpected situation, but system can recover.
- *Info*, to communicate with user about system progress and its sanity.
- *Verbose*, same as info but more verbose.
- *Debug*, for debugging.
- *Trace*, for involved debugging.

---

Settings and configuration
==========================

Logging is typically initialized during application bootstrap, or via init()
code. Sometimes, it is required to re-configured logging after application
has started, which can be done via HTTP endpoints.

For all these cases, golog provides an API - *SetLogger*.

.code golog.snippets /^func init/,/^}/

---

Console logging
===============

By default log APIs will worry about log-level, prefix format, time-format
etc. Sometimes it become too much of clutter on the screen to communicate simple
messages with user via console. In such cases use the *Consolef* API.

.code golog.snippets /^func showversion/,/^}/

*Consolef* does not print the log time, log level and always outputs to
os.Stdout.

---

Colors for console logging
==========================

While logging to console it is possible to add colors. *golog* uses
[fatih/color](http://github.com/fatih/color) for colorizing outputs.

Default color values are:

    setts = map[string]interface{}{
        "log.colorfatal":   "red",
        "log.colorerror":   "hired",
        "log.colorwarn":    "yellow",
    }

color configuration is available for each and every log level.


Reference

If golog sounds useful please check out the following links.

.link https://github.com/prataprc/golog Project README.
.link https://godoc.org/github.com/prataprc/golog Golog API doc.
.link https://github.com/prataprc/golog/issues Please contribute.
