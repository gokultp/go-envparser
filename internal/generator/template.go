package generator

const tmpl = `
package {{- .Package}}

import(
	"os"
	"errors"
)

func (t *{{.Name }}) Decode() error {
	{{range $field := .Fields}} {{template "decoder" $field}}  {{end}}
}

{{define "decoder" }}
{{if (is_builtin .Type) }}
	{{if is_notsupported_type .Type}}
	return errors.New("type {{- .Type -}} not supported")
	{{else}}
	{{if or (eq (basetype .Type) "int") (eq (basetype .Type) "uint")}}
	if {{ varname .Name -}}Str := os.Getenv("{{- .EnvTag -}}"); {{- varname .Name -}}Str != "" {
		{{ varname .Name -}}64, err := {{- (parsefunc .Type .Name ) -}}
		if err != nil {
			return err
		}
		{{- varname .Name -}} := {{- .Type -}}({{- varname .Name -}}64)
		{{if .IsPointer}}
		t.{{- .Name -}} = &{{- varname .Name -}}
		{{else}}
		t.{{- .Name -}} = {{- varname .Name -}}
		{{end}}
	}
	{{else}}
	// will do later
	{{end}}
	{{end}}
	{{end}}
{{end}}
`
