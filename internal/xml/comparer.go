package xml

import (
	"github.com/akedrou/textdiff"
	"strconv"
	"xmldiff/internal/util"
)


func Compare(xml1 *Node, xml2 *Node) []XMLDifference {
	return compare(xml1, xml2, "/0.")
}

func compare(xml1 *Node, xml2 *Node, path string) []XMLDifference {
	var differences []XMLDifference

	if xml1.XMLName.Local != xml2.XMLName.Local {
		differences = append(differences, getStringDifferences(xml1.XMLName.Local, xml2.XMLName.Local,
			path+"NAME"))
		return differences
	}

	var currentPath = path + xml1.XMLName.Local

	if xml1.Data != xml2.Data {
		differences = append(differences, getStringDifferences(xml1.Data, xml2.Data, currentPath+".DATA"))
	}

	var attrDifferences = getAttributesDifferences(xml1, xml2, currentPath)
	differences = append(differences, attrDifferences...)

	childrenDifferences := getChildrenDifferences(xml1, xml2, currentPath)

	differences = append(differences, childrenDifferences...)

	return differences
}

func getAttributesDifferences(xml1 *Node, xml2 *Node, currentPath string) []XMLDifference {
	var differences []XMLDifference

	if len(xml1.Attributes) != len(xml2.Attributes) {
		var difference = OtherDifference{
			path:    currentPath + ".ATTR.LEN",
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
			var attrPath = currentPath + ".ATTR[" + strconv.Itoa(i) + "]" + ".NAME"
			var strDifferences = getStringDifferences(name1, name2, attrPath)
			differences = append(differences, strDifferences)
		} else {
			var value1 = xml1.Attributes[i].Value
			var value2 = xml2.Attributes[i].Value
			if value1 != value2 {
				var attrPath = currentPath + ".ATTR." + name1
				var strDifferences = getStringDifferences(value1, value2, attrPath)
				differences = append(differences, strDifferences)
			}
		}

	}
	return differences
}

func getChildrenDifferences(xml1 *Node, xml2 *Node, currentPath string) []XMLDifference {
	var childrenDifferences []XMLDifference
	if len(xml1.Nodes) != len(xml2.Nodes) {
		var difference = OtherDifference{
			path:    currentPath + ".NODES.LEN",
			oldPart: util.ToString(len(xml1.Nodes)),
			newPart: util.ToString(len(xml2.Nodes)),
		}
		childrenDifferences = append(childrenDifferences, difference)
	}

	currentPath += "/"

	var nChildren = min(len(xml1.Nodes), len(xml2.Nodes))
	for i := 0; i < nChildren; i++ {
		var childPath = currentPath + util.ToString(i) + "."
		var childDifferences = compare(xml1.Nodes[i], xml2.Nodes[i], childPath)
		childrenDifferences = append(childrenDifferences, childDifferences...)
	}
	return childrenDifferences
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
