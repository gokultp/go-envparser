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
