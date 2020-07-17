package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"strings"
)

type Prop struct {
	Name      string
	LowerName string
	T         string
}

type Props []*Prop

type StructInfo struct {
	Name      string
	LowerName string
	Props     Props
	Args      string
}

var (
	errInvalidFilePath = errors.New("invalid file path")
)

func generateModel(filePath string) error {
	if filePath == "" {
		return errInvalidFilePath
	}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	src := string(b)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		return err
	}
	props := make(Props, 0)
	var typeName string
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			spec := x.Specs[0]
			if ts, ok := spec.(*ast.TypeSpec); ok {
				typeName = ts.Name.Name
				if st, ok := ts.Type.(*ast.StructType); ok {
					fl := st.Fields.List
					for _, fi := range fl {
						name := fi.Names[0].Name
						var tp string
						if i, ok := fi.Type.(*ast.Ident); ok {
							tp = i.Name
						}
						props = append(props, &Prop{
							Name:      name,
							LowerName: fmt.Sprintf("%s%s", strings.ToLower(string(name[0])), name[1:]),
							T:         tp,
						})
					}
				}
			}
		}
		return true
	})
	for _, p := range props {
		fmt.Printf("%+v", p)
	}

	t, err := template.New("").Parse(tpl)
	if err != nil {
		return err
	}
	si := struct {
		StructInfo *StructInfo
	}{
		&StructInfo{
			Name:      typeName,
			LowerName: fmt.Sprintf("%s%s", strings.ToLower(string(typeName[0])), typeName[1:]),
			Props:     props,
			Args:      props.genArgs(),
		},
	}
	o := new(bytes.Buffer)
	if err := t.Execute(o, si); err != nil {
		return err
	}
	if err := ioutil.WriteFile("dest/output.go", o.Bytes(), 0644); err != nil {
		return err
	}
	fmt.Println("finished")
	return nil
}

func (ps Props) genArgs() string {
	if len(ps) < 2 {
		return fmt.Sprintf("%s %s,", ps[0].LowerName, ps[0].T)
	}
	nextT := ps[1].T
	sb := strings.Builder{}
	for i, p := range ps {
		sb.WriteString(p.LowerName)
		if p.T != nextT || i == len(ps)-1 {
			sb.WriteString(" ")
			sb.WriteString(p.T)
		}
		sb.WriteString(", ")
		if i < len(ps)-1 {
			nextT = ps[i+1].T
		}
	}
	return sb.String()
}
