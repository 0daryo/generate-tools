package generator

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"path/filepath"
	"strings"
)

var (
	ErrNotValidFileName = errors.New("invalid filename")
	ErrNotValidLayer    = errors.New("specfiy a valid layer with -l flag")
)

func GetFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func HandleImportGenDecl(w io.Writer, d *ast.GenDecl) {
	if d.Tok == token.IMPORT {
		for _, im := range d.Specs {
			switch x := im.(type) {
			case *ast.ImportSpec:
				fmt.Fprintf(w, "	%s\n", x.Path.Value)
			}
		}
	}
}

func HandleIFGenDecl(w io.Writer, d *ast.GenDecl, recName string) {
	if d.Tok == token.TYPE {
		for _, sp := range d.Specs {
			switch x := sp.(type) {
			case *ast.TypeSpec:
				if it, ok := x.Type.(*ast.InterfaceType); ok {
					for _, f := range it.Methods.List {
						methodName := f.Names[0].Name
						result := make([]string, 0)
						args := make([]string, 0)
						t := f.Type
						if ft, ok := t.(*ast.FuncType); ok {
							for _, rFi := range ft.Results.List {
								if ident, ok := rFi.Type.(*ast.Ident); ok {
									result = append(result, ident.Name)
								}
							}
							for _, fi := range ft.Params.List {
								var fiName, fiType string
								fiName = fi.Names[0].Name
								switch x := fi.Type.(type) {
								case *ast.SelectorExpr:
									fiType = fmt.Sprintf("%s.%s", x.X, x.Sel.Name)
								case *ast.StarExpr:
									if ident, ok := x.X.(*ast.Ident); ok {
										fiType = ident.Name
									}
								case *ast.Ident:
									fiType = x.Name
								}
								args = append(args, fmt.Sprintf("%s %s", fiName, fiType))
							}
							resultStr := JoinArrWithoutLast(result, ",")
							if len(result) > 1 {
								resultStr = fmt.Sprintf("(%s)", resultStr)
							}
							fmt.Fprintf(w, "func (s *%s) %s(%s)%s{\n", recName, methodName,
								JoinArrWithoutLast(args, ","), resultStr)
							fmt.Fprintln(w, "}")
						}
					}
				}
			}
		}
	}
}

func JoinArrWithoutLast(arr []string, delim string) string {
	str := strings.Join(arr, delim)
	return strings.TrimRight(str, delim)
}

func HandleFuncDecl(w io.Writer, d *ast.FuncDecl) (string, string, bool) {
	var recName string
	if d.Recv == nil {
		return "", "", false
	}
	f := d.Recv.List[0]
	if se, ok := f.Type.(*ast.StarExpr); ok {
		if idt, ok := se.X.(*ast.Ident); ok {
			recName = idt.Name
		}
	}

	fmt.Fprintf(w, "func Test_%s_%s(t *testing.T){", recName, d.Name)
	return recName, d.Name.String(), true
}

func GetArgsAndRet(d *ast.FuncDecl) ([]string, []string) {
	args := make([]string, 0)
	rets := make([]string, 0)
	ft := d.Type
	for _, rFi := range ft.Results.List {
		if ident, ok := rFi.Type.(*ast.Ident); ok {
			rets = append(rets, ident.Name)
		}
		if at, ok := rFi.Type.(*ast.ArrayType); ok {
			if se, ok := at.Elt.(*ast.StarExpr); ok {
				if selE, ok := se.X.(*ast.SelectorExpr); ok {
					val := selE.Sel.Name
					if ident, ok := selE.X.(*ast.Ident); ok {
						val = fmt.Sprintf("%s.%s", ident.Name, val)
					}
					rets = append(rets, fmt.Sprintf("[]*%s", val))
				}
			}
		}
	}
	for _, fi := range ft.Params.List {
		var fiName, fiType string
		fiName = fi.Names[0].Name
		switch x := fi.Type.(type) {
		case *ast.SelectorExpr:
			fiType = fmt.Sprintf("%s.%s", x.X, x.Sel.Name)
		case *ast.StarExpr:
			if ident, ok := x.X.(*ast.Ident); ok {
				fiType = ident.Name
			}
		case *ast.Ident:
			fiType = x.Name
		}
		args = append(args, fmt.Sprintf("%s %s", fiName, fiType))
	}
	return args, rets
}

type ArgIdent struct {
	Name string
	Type string
}

func GetProperty(d *ast.GenDecl) []*ArgIdent {
	ret := make([]*ArgIdent, 0)
	if d.Tok == token.TYPE {
		for _, sp := range d.Specs {
			switch x := sp.(type) {
			case *ast.TypeSpec:
				if st, ok := x.Type.(*ast.StructType); ok {
					for _, fi := range st.Fields.List {
						n := fi.Names[0].Name
						name := fmt.Sprintf("%s%s", strings.ToLower(n[:1]), n[1:])
						var t string
						if id, ok := fi.Type.(*ast.Ident); ok {
							t = id.Name
						}
						ret = append(ret, &ArgIdent{
							Name: name,
							Type: t,
						})
					}
				}
			}
		}
	}
	return ret
}

func GetRecName(d *ast.GenDecl) string {
	recName := ""
	if d.Tok == token.TYPE {
		for _, sp := range d.Specs {
			switch x := sp.(type) {
			case *ast.TypeSpec:
				recName = x.Name.Name
			}
		}
	}
	return recName
}
