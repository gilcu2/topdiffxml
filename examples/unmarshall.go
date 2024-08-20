package main

import (
	"encoding/xml"
	"fmt"
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
			e.UnmarshalXML(d, t)
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

func (e *Node) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	start.Name = e.XMLName
	start.Attr = e.Attributes
	return enc.EncodeElement(struct {
		Data  string `xml:",chardata"`
		Nodes []*Node
	}{
		Data:  e.Data,
		Nodes: e.Nodes,
	}, start)
}

func main() {
	x := `<Person name="Alice" age="35">
  <Line1>TextLine1</Line1>
  <Line2>TextLine2</Line2>
  <Line3>TextLine3</Line3>
  <ProfilePicture type="round" center="true">http://myImageUrl/thumbnail</ProfilePicture>
</Person>`

	fmt.Printf("In:\n%s\n", x)

	// Unmarshal.
	n := &Node{}
	xml.Unmarshal([]byte(x), &n)
	fmt.Printf("Unmarshal:\n%v\n", n)

	// Marshal.
	out, err := xml.MarshalIndent(n, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling")
		return
	}
	fmt.Printf("Marshal:\n%s\n", string(out))
}
