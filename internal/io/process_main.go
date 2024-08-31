package io

import (
	"errors"
	"fmt"
	"os"
	"xmldiff/internal/xml"
)

func read_file(filename string) (string, error) {
	var bytes, err = os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	var str = string(bytes)
	return str, nil
}

func Compare_XML_Files(file1 string, file2 string) (chan string, error) {

	var xml1, err1 = readAndParseXML(file1)
	if err1 != nil {
		return nil, err1
	}

	var xml2, err2 = readAndParseXML(file1)
	if err2 != nil {
		return nil, err2
	}

	var differences = xml.Compare(xml1, xml2)

	var channel = make(chan string)

	go func() {
		for _, diff := range differences {
			channel <- diff.GetOutput()
		}
	}()

	return channel, nil
}

func readAndParseXML(filename string) (*xml.Node, error) {
	var str, err = read_file(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading file %s: %s", filename, err))
	}

	var xml, parserErr = xml.Parse(str)
	if parserErr != nil {
		return nil, errors.New(fmt.Sprintf("Error parsing file %s as xml: %s", filename, parserErr))
	}

	return xml, nil
}
