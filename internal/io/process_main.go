package io

import (
	"fmt"
	"github.com/gilcu2/topdiffxml/internal/xml"
	"os"
)

func ReadFile(filename string) (string, error) {
	var bytes, err = os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	var str = string(bytes)
	return str, nil
}

func CompareXmlFiles(file1 string, file2 string) ([]string, error) {
	var str1, err1 = ReadFile(file1)
	if err1 != nil {
		return nil, fmt.Errorf("error reading file %s: %s", file1, err1)
	}

	var str2, err2 = ReadFile(file2)
	if err2 != nil {
		return nil, fmt.Errorf("error reading file %s: %s", file2, err2)
	}

	var diffs, err = CompareXmlStrings(str1, str2)

	if err == nil {
		for i, diff := range diffs {
			fmt.Printf("%d- %s\n", i, diff)
		}
	}

	return diffs, err

}

func CompareXmlStrings(str1 string, str2 string) ([]string, error) {

	var xml1, err1 = xml.Parse(str1)
	if err1 != nil {
		return nil, fmt.Errorf("error parsing first xml: %s", err1)
	}

	var xml2, err2 = xml.Parse(str2)
	if err2 != nil {
		return nil, fmt.Errorf("error parsing second xml: %s", err2)
	}

	var differences = xml.Compare(xml1, xml2)

	var printedDiffs []string

	for _, diff := range differences {
		printedDiffs = append(printedDiffs, diff.GetOutput()...)
	}

	return printedDiffs, nil
}
