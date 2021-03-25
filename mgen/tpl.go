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

func Test{{ .StructName }}_New{{ .StructName }}(t *testing.T){
	t.Parallel()
	type in struct{ {{ range $j, $x := .Fields }}{{ $x.ArgName }} {{ $x.FieldType }}
		{{ end }}
	}
	tests := []struct{
		name     string
		in       in
		out      {{ .StructName }}
	}{
		{
			name:    "success",
			in:    in{ 
				{{ range $j, $x := .Fields }}{{ $x.ArgName }}: {{ $x.TestVar }},{{ end }}
			},
			out:   {{ .StructName }}{ 
				{{ range $j, $x := .Fields }}{{ $x.Name }}: {{ $x.TestVar }},{{ end }}
			},
		},
	}
	for _,tt :=range tests{
		tt := tt
		t.Run(tt.name, func(t *testing.T){
			t.Parallel()
			out := New{{ .StructName }}(
				{{ range $j, $x := .Fields }}tt.in.{{ $x.ArgName }},{{ end }}
			)
			assert.Equal(t, tt.out, out)
		})
	}
}
`
