package main

import (
	"flag"

	"github.com/gokultp/go-envparser/internal/commands"
)

func main() {
	flag.Parse()
	cmd := commands.GetCmd(flag.Args())
	if err := cmd.Exec(); err != nil {
		panic(err)
	}
}
