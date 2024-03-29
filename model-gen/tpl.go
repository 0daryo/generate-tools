package main

var tpl = `
{{- $gStructInfo := .StructInfo }}
{{- $gStructName := $gStructInfo.Name }}
{{- $gStructLName := $gStructInfo.LowerName }}
package dest

type {{$gStructName}} struct { {{ range $i, $v := $gStructInfo.Props }}
	{{ $v.Name }} {{ $v.T }}'xml:"{{ $v.LowerName }}"'{{ end }}
}
{{ $lastType := "notType" }}
func New{{$gStructName}}( {{ $gStructInfo.Args }}
)(*{{$gStructName}}, error){
	m := &{{$gStructName}} { {{ range $i, $v := $gStructInfo.Props }}
		{{ $v.Name }}: {{ $v.LowerName }},{{ end }}
	}
	if err := validate.Validator.Struct(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (o *{{$gStructName}}) toParam()*{{$gStructLName}}{
	p := &{{$gStructLName}} { {{ range $i, $v := $gStructInfo.Props }}
		{{ $v.Name }}: o.{{ $v.Name }},{{ end }}
	}
return p
}

func new{{$gStructName}}(o *{{$gStructLName}})*{{$gStructName}}{
p := &{{$gStructName}} { {{ range $i, $v := $gStructInfo.Props }}
	{{ $v.Name }}: o.{{ $v.Name }},{{ end }}
}
return p
}

func fake{{$gStructName}}()*{{$gStructName}}{
p := &{{$gStructName}} { {{ range $i, $v := $gStructInfo.Props }}
	{{ $v.Name }}: {{ $v.Fake }},{{ end }}
}
return p
}

args
{{ range $i, $v := $gStructInfo.Props }}
	o.{{ $v.Name }},{{ end }}
`

var ParamTpl = `
func (o *{{$gStructName}}) toParam()*{{$gStructLName}}{
	p := &{{$gStructLName}} { {{ range $i, $v := $gStructInfo.Props }}
		{{ $v.Name }}: o.{{ $v.Name }},{{ end }}
	}
return p
}

func new{{$gStructName}}(o *{{$gStructLName}})*{{$gStructName}}{
p := &{{$gStructName}} { {{ range $i, $v := $gStructInfo.Props }}
	{{ $v.Name }}: o.{{ $v.Name }},{{ end }}
}
return p
}

func fake{{$gStructName}}()*{{$gStructName}}{
p := &{{$gStructName}} { {{ range $i, $v := $gStructInfo.Props }}
	{{ $v.Name }}: {{ $v.Fake }},{{ end }}
}
return p
}
`

var a = "unused"
