package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/loganbestwick/js-go/parser"
	"bytes"
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
	fmt.Println(string(buf))
	fmt.Println("--- OUTPUT ---")

	parser.Parse(bytes.NewReader(buf))
}
