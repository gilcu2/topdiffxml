package io

import (
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

func Compare_XML_Files(file1 string, file2 string) error {
	var str1, err1 = read_file(file1)
	if err1 != nil {
		return err1
	}
	var str2, err2 = read_file(file2)
	if err2 != nil {
		return err2
	}

	var result, err = xml.Compare(str1, str2)
	if err != nil {
		return err
	}

	return nil
}
