package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/0daryo/generate-tools/generator"
)

const (
	sampleInt      = "1"
	sampleInt20    = "20"
	sampleStr      = "sample"
	sampleURL      = "testurl/jpeg"
	sampleCategory = "shirt"
	sampleBoolT    = "true"
)

func main() {
	flag.Parse()
	if err := generateTarget(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := generateTest(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateTarget() error {
	originPath := "target"
	recName := generator.GetFileNameWithoutExt(originPath)
	bytes, err := ioutil.ReadFile("target.go")
	if err != nil {
		return err
	}
	src := string(bytes)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		return err
	}
	structName := ""
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			structName = generator.GetRecName(x)
		}
		return true
	})
	w, err := os.Create(fmt.Sprintf("output.go"))
	if err != nil {
		return err
	}
	var args []*generator.ArgIdent
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			args = generator.GetProperty(x)
		}
		return true
	})
	fmt.Fprintf(w, "package model\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, `import "github.com/omnisinc/queen/pkg/validator"`)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "func New%s(\n", structName)
	writeArgs(w, args)
	fmt.Fprintf(w, ") (*%s, error){\n", structName)
	fmt.Fprintf(w, "\t%s := &%s{\n", recName, structName)
	for _, v := range args {
		ca := fmt.Sprintf("%s%s", strings.ToUpper(v.Name[:1]), v.Name[1:])
		fmt.Fprintf(w, "		%s: %s,\n", ca, v.Name)
	}
	fmt.Fprintf(w, "	}\n")
	fmt.Fprintf(w, "	if err := validator.Validator.Struct(%s);err!=nil{\n", recName)
	fmt.Fprintf(w, "		return err\n")
	fmt.Fprintf(w, "	}\n")
	fmt.Fprintf(w, "\treturn %s, nil\n", recName)
	fmt.Fprint(w, "}\n")
	return nil
}
func writeArgs(w io.Writer, args []*generator.ArgIdent) {
	for i, v := range args {
		if i+1 < len(args) && args[i+1].Type == v.Type {
			fmt.Fprintf(w, "	%s,\n", v.Name)
		} else {
			fmt.Fprintf(w, "	%s	%s,\n", v.Name, v.Type)
		}
	}
}

func generateTest() error {
	originPath := "target"
	recName := generator.GetFileNameWithoutExt(originPath)
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s.go", recName))
	if err != nil {
		return err
	}
	src := string(bytes)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		return err
	}
	structName := ""
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			structName = generator.GetRecName(x)
		}
		return true
	})
	w, err := os.Create(fmt.Sprintf("output_test.go"))
	if err != nil {
		return err
	}
	var args []*generator.ArgIdent
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			args = generator.GetProperty(x)
		}
		return true
	})
	fmt.Fprintf(w, "package model\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintln(w, `import (`)
	fmt.Fprintln(w, `	"testing"`)
	fmt.Fprintln(w, `	"github.com/stretchr/testify/assert"`)
	fmt.Fprintf(w, `)`)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "func Test_Model_New%s(t *testing.T) {\n", structName)
	fmt.Fprintln(w, "	type args struct{")
	writeTestArgs(w, args)
	fmt.Fprint(w, "	}\n")
	fmt.Fprintln(w, "\ttests := []struct {")
	fmt.Fprintln(w, "\t\tname   string")
	fmt.Fprintln(w, "\t\targs  args")
	fmt.Fprintln(w, "\t\twant  want")
	fmt.Fprintln(w, "\t\twantErr  wantErr")
	fmt.Fprintln(w, "\t}{")
	fmt.Fprintln(w, "\t\t{")
	fmt.Fprint(w, "\t\t\t")
	fmt.Fprintln(w, `name: "success",`)
	fmt.Fprint(w, "\t\t\t")
	fmt.Fprintln(w, `			args: args{`)
	writeTestExampleArgs(w, args)
	fmt.Fprintln(w, `			},`)
	fmt.Fprintln(w, `			want: &%s{`, structName)
	writeTestExampleArgs(w, args)
	fmt.Fprintln(w, `			},`)
	fmt.Fprintln(w, `		},`)
	fmt.Fprintln(w, `	for _, tt := range tests{`)
	fmt.Fprintln(w, `		t.Run(tt.name, func(t *testing.T){`)
	fmt.Fprintln(w, `			got, err := New%s(`, structName)
	writeTestExecArgs(w, args)
	fmt.Fprintln(w, "			)")
	fmt.Fprintln(w, "			assert.Equal(t, tt.wantErr, err != nil)")
	fmt.Fprintln(w, "			assert.Equal(t, tt.want, got)")
	fmt.Fprintln(w, "		})")
	fmt.Fprintln(w, "	}")
	fmt.Fprintln(w, "}")
	return nil
}

func writeTestArgs(w io.Writer, args []*generator.ArgIdent) {
	for _, v := range args {
		fmt.Fprintf(w, "	%s	%s\n", v.Name, v.Type)
	}
}

func writeTestExecArgs(w io.Writer, args []*generator.ArgIdent) {
	for _, v := range args {
		fmt.Fprintf(w, "				tt.args.%s\n", v.Name)
	}
}

func writeTestExampleArgs(w io.Writer, args []*generator.ArgIdent) {
	for _, v := range args {
		var t string
		switch v.Type {
		case "string":
			t = sampleStr
			if strings.Contains(strings.ToLower(v.Name), "url") {
				t = sampleURL
			}
			fmt.Fprintf(w, `				%s:	"%s",`, v.Name, t)
			fmt.Fprintln(w, "")
			return
		case "int64":
			t = sampleInt
		case "bool":
			t = sampleBoolT
		default:
			t = "undefined"
		}
		fmt.Fprintf(w, `				%s:	%s,`, v.Name, t)
		fmt.Fprintln(w, "")
	}
}
