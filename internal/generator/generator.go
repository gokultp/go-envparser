package generator

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/gokultp/envparser/internal/parser"
)

func GenerateCode(typeDef *parser.Type) error {
	funcMap := getFuncMap()
	tmpl, err := template.New("template").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return err
	}
	var b bytes.Buffer

	if err := tmpl.Execute(&b, typeDef); err != nil {
		return err
	}
	fmt.Println(b.String())
	return nil
}
