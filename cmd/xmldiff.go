package main

import (
	"flag"
	"fmt"
	"os"

	"xmldiff/internal/process"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	var file1, file2 string

	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		fmt.Println("xmldiff <file1> <file2>")
		os.Exit(1)
	}
	file1 = flag.Arg(0)
	file2 = flag.Arg(1)

	var result = process.Compare_XML_Files(file1, file2)

	os.Exit(result)
}
