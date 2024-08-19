package process

import (
	"os"
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

	return nil
}
