package main

import (
	"flag"
	"fmt"
	"os"

	"xmldiff/internal/io"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	var file1, file2 string
	var exitCode int

	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Println("Find the top level differences between 2 xml files. Usage:")
		fmt.Println("topdiffxml <file1> <file2>")
		exitCode = 1
	} else {
		file1 = flag.Arg(0)
		file2 = flag.Arg(1)

		var _, err = io.CompareXmlFiles(file1, file2)

		if err != nil {
			fmt.Println(err)
			exitCode = 1
		}
	}

	return exitCode
}
