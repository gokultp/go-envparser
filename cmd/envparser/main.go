package main

import (
	"fmt"

	"github.com/gokultp/envparser/pkg/parser"
)

func main() {
	st := parser.NewType("A")
	st.Parse("/home/gokul/projects/envparser/cmd/test/t.go")
	fmt.Printf("%#v", *st)
}
