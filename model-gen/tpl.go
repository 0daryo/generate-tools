package main

var tpl = `
{{- $gStructInfo := .StructInfo }}
{{- $gStructName := $gStructInfo.Name }}
package model

type {{$gStructName}} model { {{ range $i, $v := $gStructInfo.Props }}
	{{ $v.Name }} {{ $v.T }}{{ end }}
}
{{ $lastType := "notType" }}
func New{{$gStructName}}( {{ $gStructInfo.Args }}
)(*{{$gStructName}}, error){
	m := &{{$gStructName}} { {{ range $i, $v := $gStructInfo.Props }}
		{{ $v.Name }}: {{ $v.LowerName }},{{ end }}
	}
	return m, nil
}
`
