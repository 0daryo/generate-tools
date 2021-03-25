package main

var tpl = `
type {{ .StructName }} struct{ {{ range $j, $x := .Fields }}
	{{ $x.Name }} {{ $x.FieldType }}{{ end }}
}

func New{{ .StructName }}(
	{{ range $j, $x := .Fields }}{{ $x.ArgName }} {{ $x.FieldType }}, {{ end }}
){{ .StructName }}{
	return {{ .StructName }}{ {{ range $j, $x := .Fields }}
		{{ $x.Name }}: {{ $x.ArgName }},{{ end }}
  }
}
`
