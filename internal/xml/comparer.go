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
		var textDiffs = textDifferences[i]
		var difference = &Difference{
			path:      current_path,
			preamble:  str1[:textDiffs.Start],
			postamble: str1[textDiffs.End:],
			part1:     str1[textDiffs.Start:textDiffs.End],
			part2:     str2[textDiffs.Start:textDiffs.End],
		}
		differences = append(differences, difference)
	}
	return differences
}
