package xml

import (
	"github.com/akedrou/textdiff"
	"strconv"
	"xmldiff/internal/util"
)

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
		var oldPart = sd.source[change.Start:change.End]
		var newPart = change.New
		s += leftContext + " --(" + oldPart + ") ++(" + newPart + ") " + rightContext + "\n"
	}
	return s
}

type OtherDifference struct {
	path    string
	oldPart string
	newPart string
}

func (sd OtherDifference) getOutput() string {
	return sd.path + "\n" + " --(" + sd.oldPart + ") ++(" + sd.newPart + ") " + "\n"
}

func Compare(xml1 *Node, xml2 *Node) ([]XMLDifference, error) {
	return compare(xml1, xml2, "/")
}

func compare(xml1 *Node, xml2 *Node, path string) ([]XMLDifference, error) {
	var differences []XMLDifference

	if xml1.XMLName.Local != xml2.XMLName.Local {
		differences = append(differences, getStringDifferences(xml1.XMLName.Local, xml2.XMLName.Local, path))
		return differences, nil
	}

	var currentPath = path + xml1.XMLName.Local

	if xml1.Data != xml2.Data {
		differences = append(differences, getStringDifferences(xml1.Data, xml2.Data, currentPath))
	}

	if len(xml1.Attributes) != len(xml2.Attributes) {
		var difference = OtherDifference{
			path:    currentPath + ".attr.len",
			oldPart: util.ToString(len(xml1.Attributes)),
			newPart: util.ToString(len(xml2.Attributes)),
		}
		differences = append(differences, difference)
	}

	var nAttributes = min(len(xml1.Attributes), len(xml2.Attributes))
	for i := 0; i < nAttributes; i++ {

		var name1 = xml1.Attributes[i].Name.Local
		var name2 = xml2.Attributes[i].Name.Local
		if name1 != name2 {
			var attrPath=currentPath+".attr["+strconv.Itoa(i)+"]"+".name"
			var strDifferences = getStringDifferences(name1, name2, attrPath)
			differences = append(differences, strDifferences)
		} else {
			var value1 = xml1.Attributes[i].Value
			var value2 = xml2.Attributes[i].Value
			if value1 != value2 {
				var attrPath = currentPath + ".attr." + name1
				var strDifferences = getStringDifferences(value1, value2, attrPath)
				differences = append(differences, strDifferences)
			}
		}

	}

	return differences, nil
}

func getStringDifferences(str1 string, str2 string, currentPath string) StringDifferences {
	var changes = textdiff.Strings(str1, str2)
	var stringDifferences = StringDifferences{
		path:    currentPath,
		source:  str1,
		changes: changes,
	}

	return stringDifferences
}

