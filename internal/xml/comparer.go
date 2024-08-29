package xml

import "github.com/akedrou/textdiff"

type XMLDifference interface {
	getOutput() string
}

type StringDifferences struct {
	path    string
	source  string
	changes []textdiff.Edit
}

func (sd StringDifferences) getOutput() string {
	var s = sd.path + "\n"
	var contextBegin = 0
	var contextEnd = 0
	for i := range len(sd.changes) {
		var change = sd.changes[i]
		contextBegin = max(contextEnd, change.Start-10)
		if i < len(sd.changes)-1 {
			contextEnd = min(change.End+10, sd.changes[i+1].Start-10)
		} else {
			contextEnd = min(len(sd.source), change.End+10)
		}
		var leftContext = sd.source[contextBegin:change.Start]
		var rightContext = sd.source[change.End:contextEnd]
		var old = sd.source[change.Start:change.End]
		var new = change.New
		s += leftContext + " --(" + old + ") ++(" + new + ") " + rightContext + "\n"
	}
	return s
}

type OtherDifference struct {
	path string
	info string
}

func (sd OtherDifference) getOutput() string {
	return sd.path + "\n" + sd.info + "\n"
}

func Compare(xml1 *Node, xml2 *Node) ([]*XMLDifference, error) {
	return compare(xml1, xml2, "/")
}

func compare(xml1 *Node, xml2 *Node, path string) ([]*XMLDifference, error) {
	var differences []*XMLDifference

	if xml1.XMLName.Local != xml2.XMLName.Local {
		differences = append(differences, getStringDifferences(xml1.XMLName.Local, xml2.XMLName.Local, path))
		return differences, nil
	}

	var currentPath = path + xml1.XMLName.Local

	if xml1.Data != xml2.Data {
		differences = append(differences, getStringDifferences(xml1.Data, xml2.Data, currentPath))
	}

	var nAttributes = min(len(xml1.Attributes), len(xml2.Attributes))
	for i := 0; i < nAttributes; i++ {

	}

	return differences, nil
}

func getStringDifferences(str1 string, str2 string, currentPath string) *XMLDifference {
	var changes = textdiff.Strings(str1, str2)
	var stringDiferences = StringDifferences{
		path:    currentPath,
		source:  str1,
		changes: changes,
	}

	return &stringDiferences
}
