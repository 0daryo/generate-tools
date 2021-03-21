package main

var tpl = `
{{- $gStructName := .StructName }}
package dest

type {{ .IFName }} interface {
	{{ range $i, $v := .Funcs }}{{ $v.Name }}(
		{{ range $j, $x := $v.Args }}{{ $x.Name }} {{ $x.ArgType }},{{ end }}
	) (
		{{ range $j, $x := $v.Rets }}{{ $x.Name }} {{ $x.RetType }},{{ end }}
	){{ end }}
}

type {{ .StructName }} struct{}

func New{{ .IFName }}()*{{ .StructName }} {
	return &{{ .StructName }}{}
}

{{ range $i, $v := .Funcs }}func (s *{{ $gStructName }}) {{ $v.Name }}(
	{{ range $j, $x := $v.Args }}{{ $x.Name }} {{ $x.ArgType }},{{ end }}
) (
	{{ range $j, $x := $v.Rets }}{{ $x.Name }} {{ $x.RetType }},{{ end }}
) {
	return
}
{{ end }}
`
