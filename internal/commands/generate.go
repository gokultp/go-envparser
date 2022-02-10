package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/gokultp/go-envparser/internal/generator"
	"github.com/gokultp/go-envparser/internal/parser"
)

// Generate is the command doing code generation
type Generate struct {
	flags      flag.FlagSet
	filepath   string
	structName string
}

// NewGenerate return a new instance of Generate
func NewGenerate() *Generate {
	return &Generate{}
}

// InitFlags will initialize all flags
func (c *Generate) InitFlags() {
	c.flags.StringVar(&c.structName, "t", "", "structure name")
	c.flags.StringVar(&c.structName, "type", "", "structure name")
	c.flags.StringVar(&c.filepath, "f", "", "file path")
	c.flags.StringVar(&c.filepath, "file", "", "file path")
}

// ParseFlags will parse given flags
func (c *Generate) ParseFlags(args []string) {
	c.flags.Parse(args)
}

// Help prints the help message
func (Generate) Help() {
	helpText := `

	envparser generate -t [--type] <struct name> -f [--file] <file path>
	
	It generates functions that implemets envdecode.Decoder interface for the given structure.
	`
	fmt.Println(helpText)
}

// Exec will execute the core command functionality, here it generates and saves the code
func (c *Generate) Exec() error {
	if c.structName == "" || c.filepath == "" {
		c.Help()
		os.Exit(2)
	}
	st := parser.NewType(c.structName)
	if err := st.Parse(c.filepath); err != nil {
		return err
	}
	code, err := generator.GenerateCode(st)
	if err != nil {
		return err
	}
	return generator.SaveCode(st, code)
}
