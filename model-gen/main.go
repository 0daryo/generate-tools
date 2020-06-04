package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	err := generateModel(flag.Arg(0))
	fmt.Printf("err=%v", err)
}
