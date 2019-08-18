package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/gokultp/envparser/internal/parser"
)

func GenerateCode(typeDef *parser.Type) ([]byte, error) {
	funcMap := getFuncMap()
	tmpl, err := template.New("template").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer

	if err := tmpl.Execute(&b, typeDef); err != nil {
		return nil, err
	}
	code := b.Bytes()
	code, err = format.Source(removeEmptyLines(code))
	if err != nil {
		return nil, err
	}
	return code, nil
}

func SaveCode(typeDef *parser.Type, code []byte) error {
	fpath := path.Dir(typeDef.FileName)
	fpath = path.Join(fpath, fmt.Sprintf("%sdecoder.go", strings.ToLower(typeDef.Name)))
	return ioutil.WriteFile(fpath, code, 0644)
}

func removeEmptyLines(source []byte) []byte {
	reg, err := regexp.Compile(`\n\s*\n`)
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAll(source, []byte{'\n'})
}
