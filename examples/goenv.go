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
	GoPath string   `env:"GOPATH"`
	Path   []string `env:"PATH"`
	GoRoot string   `env:"GOROOT"`
}

func main() {
	env := GoEnv{}
	if err := envdecoder.Decode(&env); err != nil {
		panic(err)
	}
	fmt.Printf("%#v", env)
}
