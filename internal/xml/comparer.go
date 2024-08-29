package xml

import "github.com/akedrou/textdiff"

type Difference struct {
	path      string
	preamble  string
	postamble string
	part1     string
	part2     string
}

func Compare(xml1 *Node, xml2 *Node) ([]*Difference, error) {
	return compare(xml1, xml2, "/")
}

func compare(xml1 *Node, xml2 *Node, path string) ([]*Difference, error) {
	var differences = []*Difference{}

	if xml1.XMLName.Local != xml2.XMLName.Local {
		differences = append(differences, getStringDifferences(xml1.XMLName.Local, xml2.XMLName.Local, path)...)
		return differences,nil
	}

	var current_path = path + "/" + xml1.XMLName.Local

	if xml1.Data != xml2.Data {
		differences = append(differences, getStringDifferences(xml1.Data, xml2.Data, current_path)...)
	}

	return differences, nil
}

func getStringDifferences(str1 string, str2 string, current_path string) []*Difference {
	var differences = []*Difference{}
	var textDifferences = textdiff.Strings(str1, str2)
	for i := 0; i < len(textDifferences); i++ {
		var textDiff = textDifferences[i]
		var part1=str1[textDiff.Start:textDiff.End]
		var part2=textDiff.New
		var difference = &Difference{
			path:      current_path,
			preamble:  str1[:textDiff.Start],
			postamble: str1[textDiff.End:],
			part1:     part1,
			part2:     part2,
		}
		differences = append(differences, difference)
	}
	return differences
}
