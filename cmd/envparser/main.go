package main

import (
	"fmt"

	"github.com/gokultp/envparser/internal/generator"
	"github.com/gokultp/envparser/internal/parser"
)

func main() {
	st := parser.NewType("A")
	st.Parse("/home/gokul/projects/envparser/cmd/test/t.go")

	if err := generator.GenerateCode(st); err != nil {
		fmt.Println(err)
	}
}
