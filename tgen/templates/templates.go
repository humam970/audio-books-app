package templates

// gotmpl
const tsInterfaceTemplate = `
export interface {{.Name}} {
{{- range .Fields}}
  {{.Name}}{{if .IsOptional}}?{{end}}: {{tsType .Type}};
{{- end}}
}
`

// gotmpl
const GoStructTemplate = `
type {{.Name}} struct {
{{- range .Fields}}
  {{.Name}}{{if .IsOptional}}?{{end}}: {{tsType .Type}};
{{- end}}
}
`
