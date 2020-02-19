# go-envparser

## Overview [![GoDoc](https://godoc.org/github.com/gokultp/go-envparser?status.svg)](https://godoc.org/github.com/gokultp/go-envparser) [![Code Climate](https://codeclimate.com/github/gokultp/go-envparser/badges/gpa.svg)](https://codeclimate.com/github/gokultp/go-envparser) [![Go Report Card](https://goreportcard.com/badge/github.com/gokultp/go-envparser)](https://goreportcard.com/report/github.com/gokultp/go-envparser)

`go-envparser` generates static `DecodeEnv`  functions for structures  in Go to decode environment variables,  will implement the interface [Decoder](./pkg/envdecoder/idecoder.go). The generated functions reduce the reliance upon runtime reflection to do serialization.


## Getting Started

### Installation

```shell
git clone https://github.com/gokultp/go-envparser.git
cd go-envparser
make

#check if the command is working
envparser version
```
should get an output in the following format
```
Version : V0.1.0
MinVersion : e3a5a007b94f51f46a64853f308e5a24daf98892
BuildTime : 2019-08-19T00:56:29+0530
```
### Commands
```bash
envparser generate -s <structname> -f <filename>
```
It will generate a file with name `<structname>decoder.go`  (Lowercase) consist of a `DecodeEnv` function which will be implementing the interface  [Decoder](./pkg/envdecoder/idecoder.go).
### Example 

```go
package main

import (
	"fmt"

	"github.com/gokultp/go-envparser/pkg/envdecoder"
)

// Add go generate commands
//go:generate envparser generate -t GoEnv -f $GOFILE
// Dont forget to do goimport on generated files.
//go:generate goimports -w goenvdecoder.go
type GoEnv struct {
	Paths  Path
	GoRoot string `env:"GOROOT"`
}

//go:generate envparser generate -t Path -f $GOFILE
// Dont forget to do goimport on generated files.
//go:generate goimports -w pathdecoder.go
type Path struct {
	System []string `env:"PATH"`
	Go     string   `env:"GOPATH"`
}

func main() {
	env := GoEnv{}
	if err := envdecoder.Decode(&env); err != nil {
		panic(err)
	}
	fmt.Printf("%#v", env)
}

```

Here in the above file I have used go generate flags, will execute all the needed commands( envparser generate && goimports) in a single `go generate` command.

So here to generate the code execute
```bash
go generate
```

to run the code execute
``` bash 
go build
./<binary>
```

## License

`go-envparser` is licensed under the [MIT License](./LICENSE)