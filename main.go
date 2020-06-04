package main

import "fmt"

func main() {
	err := generateModel("ex.go")
	fmt.Printf("err=%v", err)
}
