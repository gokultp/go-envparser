package main

import (
	"fmt"

	"github.com/gokultp/envparser/pkg/envdecoder"
)

//go:generate envparser generate -t GoEnv -f $GOFILE
//go:generate goimports -w goenvdecoder.go
type GoEnv struct {
	GoPath string   `env:"GOPATH"`
	Path   []string `env:"PATH"`
	GoRoot string   `env:"PATH"`
}

func main() {
	env := GoEnv{}
	if err := envdecoder.Decode(&env); err != nil {
		panic(err)
	}
	fmt.Printf("%#v", env)
}
