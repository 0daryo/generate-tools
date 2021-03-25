package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/format"
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
	PKGName    string
	IFName     string
	StructName string
	Funcs      []Func
}

type Func struct {
	Name string
	Args []Arg
	Rets []Ret
}

type Arg struct {
	Name    string
	ArgType string
}

type Ret struct {
	Name    string
	RetType string
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var (
		pbcopy = flag.Bool("c", false, "copy to clipboard")
	)
	flag.Parse()
	info := Info{}
	info.PKGName = prompt.Input("pkg name: ", pkgCompleter)
	fmt.Printf("interface name: ")
	ifName := prompt.Input("interface name: ", interfaceCompleter)
	info.IFName = ifName
	info.StructName = strings.ToLower(string(ifName[0])) + ifName[1:]
	info.Funcs = funcs()
	t, err := template.New("").Parse(tpl)
	if err != nil {
		panic(err)
	}
	o := new(bytes.Buffer)
	if err := t.Execute(o, info); err != nil {
		panic(err)
	}
	formatted, err := format.Source(o.Bytes())
	if err != nil {
		panic(err)
	}
	if *pbcopy {
		clipboard.WriteAll(string(formatted))
		return
	}
	fmt.Fprint(os.Stdout, formatted)
}

func read() string {
	sc.Scan()
	return sc.Text()
}

func funcs() []Func {
	funcs := make([]Func, 0)
	for {
		funcName := prompt.Input("Function Name(if nothing type ! and enter): ", funcCompleter)
		if funcName == fin {
			return funcs
		}
		args := make([]Arg, 0)
		for {
			arg := prompt.Input("Arg(if nothing type ! and enter): ", argCompleter)
			if arg == fin {
				break
			}
			t := prompt.Input("Type: ", typeCompleter)
			args = append(args, Arg{
				Name:    arg,
				ArgType: t,
			})
		}
		rets := make([]Ret, 0)
		for {
			ret := prompt.Input("Return(if nothing type ! and enter): ", retCompleter)
			if ret == fin {
				break
			}
			t := prompt.Input("Type: ", typeCompleter)
			rets = append(rets, Ret{
				Name:    ret,
				RetType: t,
			})
		}
		funcs = append(funcs, Func{
			Name: funcName,
			Args: args,
			Rets: rets,
		})
	}
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

func argCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "!"},
		{Text: "ctx"},
		{Text: "id"},
		{Text: "name"},
		{Text: "age"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func retCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "!"},
		{Text: "err"},
		{Text: "user"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func funcCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "!"},
		{Text: "List"},
		{Text: "Get"},
		{Text: "Create"},
		{Text: "Update"},
		{Text: "Delete"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func pkgCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "service"},
		{Text: "usecase"},
		{Text: "handler"},
		{Text: "repository"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func interfaceCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "User"},
		{Text: "Email"},
		{Text: "Book"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}
