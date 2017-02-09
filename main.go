package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/loganbestwick/js-go/parser"
)

func main() {
	if len(os.Args) == 1 {
		panic("No file specified")
	}
	filename := os.Args[1]
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("--- INPUT ---")
	fmt.Println(string(buf))
	fmt.Println("--- PARSE ---")

	node := parser.Parse(bytes.NewReader(buf))
	spew.Dump(node)

	fmt.Println("--- OUTPUT ---")
	result, err := node.Eval()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
