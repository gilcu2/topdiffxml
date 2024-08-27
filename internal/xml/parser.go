package xml

import (
	"encoding/xml"
	"strings"
)


type Node struct {
	XMLName    xml.Name
	Attributes []xml.Attr
	Data       string
	Nodes      []*Node
}

func (e *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var nodes []*Node
	var done bool
	for !done {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch t := t.(type) {
		case xml.CharData:
			e.Data = strings.TrimSpace(string(t))
		case xml.StartElement:
			e := &Node{}
			var err=e.UnmarshalXML(d, t)
			if err != nil {
				return err
			}
			nodes = append(nodes, e)
		case xml.EndElement:
			done = true
		}
	}
	e.XMLName = start.Name
	e.Attributes = start.Attr
	e.Nodes = nodes
	return nil
}

func Parse(str string) (*Node, error) {
	n := &Node{}
	var err=xml.Unmarshal([]byte(str), &n)
	return n, err
}

