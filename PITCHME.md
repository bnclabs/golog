Golog: Basic logging with batteries
===================================

R Pratap Chakravarthy <br/>
prataprc@gmail.com <br/>
[https://github.com/bnclabs/golog](https://github.com/bnclabs/golog) <br/>
17, Jun 2017

---

Why logging ?
=============

- **To develop**, know how the programs fit together.
- **To debug**, fix problems while taking them to production.
- **To communicate**, with users letting them know what is happening.
- **To analyse**, characterize programs under production.

Let log messages be meaningful enough to serve their end.

---

By default use Golang's standard logging
========================================

Use __log__ package from golang, as much as possible. Lesser the
baggage better the journey :)

What you can already do with golang's log package ?

- Prefix log messages with date,time,microsecond,file,line-no.
- Fatal, Fatalf, Fatalln, Print, Printf, Println.
- Add custom prefix to all messages.
- Set output file (device) for logging, by default it goes to os.Stdout.

---

Golog: Basic logging with batteries
===================================

I use golog only if I need more than what __log__ pkg already provides.
And more than once I needed the following:

- **Log levels**, used for filtering messages and prefixing the level
information with every message.
- **Console Logs**.
- **Color**  attributes for console logs.
- **Configure** via JSON.

These facilities are supported on top of what the log pkg already
provides. I find them sufficient for my case, but if your situation
demands additional feature raise an
[issue](http://github.com/bnclabs/golog/issues)

---

Log levels
==========

Log levels are listed in __decreasing order of importance__. That is,
if log level is configured as __Info__, all messages logged at level lesser
than Info level shall be filtered out.

- **Ignore**, messages cannot be logged at Ignore level, so all messages
logged using golog will be filtered out.
- **Fatal**, will panic after logging the message.
- **Error**, means there was a critical error, needs supervisor attention.
- **Warn**, means there was an unexpected situation, but system can recover.
- **Info**, to communicate with user about system progress and its sanity.
- **Verbose**, same as info but more verbose.
- **Debug**, for debugging.
- **Trace**, for involved debugging.

---

Settings and configuration
==========================

Logging is typically initialized during application bootstrap, or via init()
code. Sometimes, it is required to re-configured logging after application
has started, which can be done via HTTP endpoints.

For all these cases, golog provides an API - **SetLogger**.

```go
func init() {
    setts := map[string]interface{}{
		"log.level":        "info",
		"log.file":         "",
		"log.colorfatal":   "red",
		"log.colorerror":   "hired",
		"log.colorwarn":    "yellow",
	}
    SetLogger(nil /*use default logger*/, setts)
}
```

---

Console logging
===============

By default log APIs will worry about log-level, prefix format, time-format
etc. Sometimes it become too much of clutter on the screen to communicate simple
messages with user via console. In such cases use the **Consolef** API.

```go
func showversion() {
    log.Consolef("goledger version - goledger%v\n", api.LedgerVersion)
}
```

**Consolef** does not print the log time, log level and always outputs to
os.Stdout.

---

Colors for console logging
==========================

While logging to console it is possible to add colors. **golog** uses
[fatih/color](http://github.com/fatih/color) for colorizing outputs.

Default color values are:

```go
    setts = map[string]interface{}{
        "log.colorfatal":   "red",
        "log.colorerror":   "hired",
        "log.colorwarn":    "yellow",
    }
```

Color configuration is available for each and every log level.

---

Tip: Filtering log messages
===========================

Every Unix server is bound to have ``grep`` and ``vim``. Add the following in
your .vimrc file to prune out unnecessary log messages.

```vim
" Filter command for vim-buffers.
command! -nargs=? Filter let @a='' | execute 'g/<args>/y A' | tabnew | setlocal bt=nofile | put! a
" Filterx command for vim-buffers.
command! -nargs=? Filterx let @a='' | execute 'v/<args>/y A' | tabnew | setlocal bt=nofile | put! a
```

Subsequently:

```vim
:Filter Error
```

Will create a new tab in vim, and list messages that contain ``Error``.
``Filterx`` does the opposite, list messages that do not contain ``Error``.
Can use vim's reg-ex pattern as Filter's argument.

---

Tip: syntax highlighting for log files
======================================

Syntax coloring for log messages can be helpful while eyeballing log files.
For Vim:

```text
github.com/vim-scripts/httplog
github.com/vim-scripts/apachelogs.vim
```

Note that, if log files are large, adding syntax highlights can
significantly slow down the editor's rendering.

---

Thank you
=========

If golog sounds useful please check out the following links.

[Project README](https://github.com/bnclabs/golog). <br/>
[Golog API doc](https://godoc.org/github.com/bnclabs/golog). <br/>
[Please contribute](https://github.com/bnclabs/golog/issues). <br/>
