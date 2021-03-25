package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/atotto/clipboard"
	"github.com/c-bata/go-prompt"
)

const (
	fin = "!"
)

type Info struct {
	StructName string
	Fields     []Field
}

type Field struct {
	Name      string
	FieldType string
	ArgName   string
	TestVar   string
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var (
		pbcopy = flag.Bool("c", false, "copy to clipboard")
	)
	flag.Parse()
	info := Info{}
	info.StructName = prompt.Input("struct name: ", structCompleter)
	info.Fields = fields()
	t, err := template.New("").Parse(tpl)
	if err != nil {
		panic(err)
	}
	o := new(bytes.Buffer)
	if err := t.Execute(o, info); err != nil {
		panic(err)
	}
	if *pbcopy {
		clipboard.WriteAll(o.String())
		return
	}
	fmt.Fprint(os.Stdout, o)
}

func read() string {
	sc.Scan()
	return sc.Text()
}

func fields() []Field {
	fields := make([]Field, 0)
	for {
		field := prompt.Input("Field(if nothing type ! and enter): ", fieldCompleter)
		if field == fin {
			break
		}
		t := prompt.Input("Type: ", typeCompleter)
		fields = append(fields, Field{
			Name:      field,
			FieldType: t,
			ArgName:   toArgName(field),
			TestVar:   newTestVar(toArgName(field), t),
		})
	}
	return fields
}

var capitalMap = map[string]bool{
	"ID":  true,
	"URL": true,
}

func toArgName(fieldName string) string {
	if isCapitalized, ok := capitalMap[fieldName]; ok && isCapitalized {
		return strings.ToLower(fieldName)
	}
	a := strings.ToLower(string(fieldName[0])) + fieldName[1:]
	return a
}

func typeCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "context.Context"},
		{Text: "uint8"},
		{Text: "uint16"},
		{Text: "uint32"},
		{Text: "uint64"},
		{Text: "int8"},
		{Text: "int16"},
		{Text: "int32"},
		{Text: "int64"},
		{Text: "float32"},
		{Text: "float64"},
		{Text: "[]byte"},
		{Text: "string"},
		{Text: "bool"},
		{Text: "time.Time"},
		{Text: "error"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func fieldCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "!"},
		{Text: "ID"},
		{Text: "Name"},
		{Text: "Age"},
		{Text: "CreatedAt"},
		{Text: "UpdatedAt"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func structCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "User"},
		{Text: "Email"},
		{Text: "Book"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func newTestVar(arg, t string) string {
	switch t {
	case "string":
		return "\"" + arg + "001\""
	case "bool":
		return "true"
	case "time.Time":
		return "time.Date(2020,time.August,11,22,33,44,55,time.UTC)"
	}
	return "15"
}
