package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

const (
	fin = "!"
)

type Info struct {
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
	info := Info{}
	fmt.Printf("IFName: ")
	ifName := read()
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
	if err := ioutil.WriteFile("dest/_output.go", o.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func read() string {
	sc.Scan()
	return sc.Text()
}

func funcs() []Func {
	funcs := make([]Func, 0)
	for {
		fmt.Printf("Function Name(if nothing type ! and enter): ")
		funcName := read()
		if funcName == fin {
			return funcs
		}
		args := make([]Arg, 0)
		for {
			fmt.Printf("Arg(if nothing type ! and enter): ")
			arg := read()
			if arg == fin {
				break
			}
			fmt.Printf("Type: ")
			t := read()
			args = append(args, Arg{
				Name:    arg,
				ArgType: t,
			})
		}
		rets := make([]Ret, 0)
		for {
			fmt.Printf("Return(if nothing type ! and enter): ")
			ret := read()
			if ret == fin {
				break
			}
			fmt.Printf("Type: ")
			t := read()
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
