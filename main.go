package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	fmt.Print(string(buf))
}
