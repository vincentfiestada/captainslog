# Captain's Log

[![GoReportCard](https://goreportcard.com/badge/github.com/vincentfiestada/captainslog)](https://goreportcard.com/report/github.com/vincentfiestada/captainslog)
[![GoDoc](https://godoc.org/github.com/vincentfiestada/captainslog?status.svg)](https://godoc.org/github.com/vincentfiestada/captainslog)

A simple logging library for [Go](https://golang.org/).

- Support for multiple logging levels
- Colored output, even on Windows
- Output name of calling function

![Screenshot of captainslog in action](./assets/screenshot.png)

## Usage

```go
package main

import log "github.com/vincentfiestada/captainslog"

func main() {
	log.SetLevel(LogLevelSilly)

	log.Silly("This is %s", "Silly")
	log.Debug("This is %s", "Debug")
	log.Verbose("This is %s", "Verbose")
	log.Info("This is %s", "Info")
	log.Warn("This is %s", "Warning")
	log.Error("This is %s", "Error")
}
```

<small>Gopher artwork by [Ashley McNamara](https://twitter.com/ashleymcnamara) on [Gopherize.me](https://gopherize.me/gopher/5dcbe4dc48ab6fbf903aae352f8742cb59e7099b)</small>
