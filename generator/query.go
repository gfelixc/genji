package generator

const querySelectorTmpl = `
{{ define "query-selector" }}

{{ template "query-selector-Struct" . }}
{{ template "query-selector-New" . }}
{{ template "query-selector-Table" . }}
{{ template "query-selector-All" . }}
{{ end }}
`

const querySelectorStructTmpl = `
{{ define "Struct" }}
// {{$structName}}QuerySelector provides helpers for selecting fields from the {{$structName}} structure.
type {{$structName}}QuerySelector struct{
{{- range $i, $a := .Fields }}
	{{- if eq .Type "string"}}
		{{$a.Name}} query.StrField
	{{- else if eq .Type "int64"}}
		{{$a.Name}} query.Int64Field
	{{- end}}
{{- end}}
}
{{ end }}
`

const querySelectorNewTmpl = `
{{ define "New" }}
{{- if .IsExported }}
// New{{$structName}}QuerySelector creates a {{$structName}}QuerySelector.
func New{{$structName}}QuerySelector() {{$structName}}QuerySelector {
{{- else}}
// new{{$structName}}QuerySelector creates a {{$structName}}QuerySelector.
func new{{.ExportedName}}QuerySelector() {{$structName}}QuerySelector {
{{- end}}
	return {{$structName}}QuerySelector{
		{{- range $i, $a := .Fields }}
			{{- if eq .Type "string"}}
				{{$a.Name}}: query.NewStrField("{{$a.Name}}"),
			{{- else if eq .Type "int64"}}
				{{$a.Name}}: query.NewInt64Field("{{$a.Name}}"),
			{{- end}}
		{{- end}}
	}
}
{{ end }}
`

const querySelectorTableTmpl = `
{{ define "Table" }}
// Table returns a query.TableSelector for {{$structName}}.
func (*{{$structName}}QuerySelector) Table() query.TableSelector {
	return query.Table("{{$structName}}")
}
{{ end }}
`

const querySelectorAllTmpl = `
{{ define "All" }}
// All returns a list of all selectors for {{$structName}}.
func (s *{{$structName}}QuerySelector) All() []query.FieldSelector {
	return []query.FieldSelector{
		{{- range $i, $a := .Fields }}
		s.{{$a.Name}},
		{{- end}}
	}
}
{{ end }}
`