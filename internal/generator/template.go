package generator

const tmpl = `
package {{ .Package }}

import(
	"os"
	"strconv"
	"strings"
)

func (t *{{.Name }}) DecodeEnv() error {
	{{range $field := .Fields}} {{template "decoder" $field}}  {{end}}
	return nil
}

{{define "decoder" }}
	if {{ varname .Name -}}Str := os.Getenv("{{- .EnvTag -}}"); {{- varname .Name -}}Str != "" {

	{{if (is_builtin .Type) }}
	{{if or (eq (basetype .Type) "rune") (eq (basetype .Type) "byte")}}
		{{varname .Name -}} := []{{- .Type -}}({{- varname .Name -}}Str)
		{{if .IsArray}}
		t.{{- .Name -}} = {{- varname .Name -}}
		{{else if .IsPointer}}
		t.{{- .Name -}} = &{{- varname .Name -}}[0]
		{{else}}
		t.{{- .Name -}} = {{- varname .Name -}}[0]
		{{end}}
	{{else}}
	{{if .IsArray}}
		{{varname .Name}}Arr := strings.Split({{varname .Name}}Str, ":")
		for _, {{varname .Name}}Str := range {{varname .Name}}Arr{
			{{template "singledecoder" .}}
		}
	{{else}}
	{{template "singledecoder" .}}
	{{end}}
	{{end }}
	{{else}}
	return errors.New("type {{- .Type -}} not supported")
	{{end}}
	}
{{end }}


{{define "singledecoder"}} 
	{{if or (eq (basetype .Type) "int") (eq (basetype .Type) "uint") (eq (basetype .Type) "float") }}
		{{ varname .Name -}}64, err := {{- (parsefunc .Type .Name )}}
		if err != nil {
			return err
		}
		{{varname .Name -}} := {{- .Type -}}({{- varname .Name -}}64)
		{{template "populate" .}}
	{{else if eq (basetype .Type) "bool"}}
		{{ varname .Name -}}, err := {{- (parsefunc .Type .Name )}}
		if err != nil {
			return err
		}
		{{template "populate" .}}
	{{else if eq (basetype .Type) "string" }}
	{{ varname .Name -}} := {{ varname .Name -}}Str
	{{template "populate" .}}	
	{{end}}
{{end}}

{{ define "populate"}}
	{{if .IsPointer}}
		t.{{- .Name -}} = &{{- varname .Name -}}
		{{else if .IsArray}}
		t.{{- .Name -}} = append(t.{{- .Name -}}, {{- varname .Name -}})
		{{else}}
		t.{{- .Name -}} = {{- varname .Name -}}
	{{end}}
{{end}}
`
