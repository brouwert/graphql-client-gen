package gcg

var scalarsTemplateSrc = `package {{.Package}}
// generated by github.com/emicklei/graphql-client-gen/cmd/gcg on {{.Created.Format "Mon, 02 Jan 2006 15:04:05 MST" }} 
// DO NOT EDIT

{{- range .Scalars}}

// {{.Name}} is a Scalar. {{.Comment}}
type {{.Name}} string

// New{{.Name}} returns a pointer to a {{.Name}} value. 
// Use type conversion instead e.g. v := {{.Name}}(s) if you need the non-pointer value.
func New{{.Name}}(s string) *{{.Name}} { v := {{.Name}}(s); return &v }

{{- end}}
`
