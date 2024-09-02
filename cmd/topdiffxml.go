package main

import (
	"flag"
	"fmt"
	"os"

	"xmldiff/internal/io"
)

// notest

func main() {
	var file1, file2 string

	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Println("Find the top level differences between 2 xml files. Usage:")
		fmt.Println("topdiffxml <file1> <file2>")
		os.Exit(1)
	}
	file1 = flag.Arg(0)
	file2 = flag.Arg(1)

	var _, err = io.CompareXmlFiles(file1, file2)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
