package xml

import (
	"encoding/xml"
)

type Node struct {
	XMLName    xml.Name
	Attributes []xml.Attr
	Data       string
	Nodes      []*Node
}

func Parse(str string) (*Node, error) {
	var nodes []*Node
	var err = xml.Unmarshal([]byte(str), &nodes)
	return nodes[0], err
}

func Compare(str1 string, str2 string) (string, error) {

	return "", nil
}
