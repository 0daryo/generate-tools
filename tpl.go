package main

var tpl = `
{{- $gStructInfo := .StructInfo }}
{{- $gStructName := $gStructInfo.Name }}

package model

type {{$gStructName}} model {
{{ range $i, $v := $gStructInfo.Props }}
{{ $v.Name }} {{ $v.T }}
{{ end }}
}
`
