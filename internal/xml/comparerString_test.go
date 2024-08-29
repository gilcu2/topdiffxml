package xml

import (
	"xmldiff/internal/util"
	"testing"
)

func TestCompareString_WhenEqual(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1"
	var str2 = "str1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	util.Assert(t, len(differences), 0)
}

func TestCompareString_WhenAddedAtEnd(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1"
	var str2 = "str1more"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	util.Assert(t, len(differences), 1)
	var diff0 = differences[0]
	util.Assert(t,diff0.path,"/")
	util.Assert(t,diff0.preamble,"str1")
	util.Assert(t,diff0.postamble,"")
	util.Assert(t,diff0.part1,"")
	util.Assert(t,diff0.part2,"more")
}

func TestCompareString_WhenAddedAtBegin(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1"
	var str2 = "morestr1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	util.Assert(t, len(differences), 1)
	var diff0 = differences[0]
	util.Assert(t,diff0.path,"/")
	util.Assert(t,diff0.preamble,"")
	util.Assert(t,diff0.postamble,"str1")
	util.Assert(t,diff0.part1,"")
	util.Assert(t,diff0.part2,"more")
}

func TestCompareString_WhenAddedAtMiddle(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1str2"
	var str2 = "str1morestr2"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	util.Assert(t, len(differences), 1)
	var diff0 = differences[0]
	util.Assert(t,diff0.path,"/")
	util.Assert(t,diff0.preamble,"str1")
	util.Assert(t,diff0.postamble,"str2")
	util.Assert(t,diff0.part1,"")
	util.Assert(t,diff0.part2,"more")
}

func TestCompareString_WhenRemovedAtEnd(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1more"
	var str2 = "str1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	util.Assert(t, len(differences), 1)
	var diff0 = differences[0]
	util.Assert(t,diff0.path,"/")
	util.Assert(t,diff0.preamble,"str1")
	util.Assert(t,diff0.postamble,"")
	util.Assert(t,diff0.part1,"more")
	util.Assert(t,diff0.part2,"")
}

func TestCompareString_WhenRemovedAtBegin(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "morestr1"
	var str2 = "str1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	util.Assert(t, len(differences), 1)
	var diff0 = differences[0]
	util.Assert(t,diff0.path,"/")
	util.Assert(t,diff0.preamble,"")
	util.Assert(t,diff0.postamble,"str1")
	util.Assert(t,diff0.part1,"more")
	util.Assert(t,diff0.part2,"")
}

func TestCompareString_WhenChangedAtBegin(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "morestr1"
	var str2 = "papistr1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	util.Assert(t, len(differences), 1)
	var diff0 = differences[0]
	util.Assert(t,diff0.path,"/")
	util.Assert(t,diff0.preamble,"")
	util.Assert(t,diff0.postamble,"str1")
	util.Assert(t,diff0.part1,"more")
	util.Assert(t,diff0.part2,"papi")
}

func TestCompareString_WhenChangedAtMiddle(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1morestr2"
	var str2 = "str1papistr2"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	util.Assert(t, len(differences), 1)
	var diff0 = differences[0]
	util.Assert(t,diff0.path,"/")
	util.Assert(t,diff0.preamble,"str1")
	util.Assert(t,diff0.postamble,"str2")
	util.Assert(t,diff0.part1,"more")
	util.Assert(t,diff0.part2,"papi")
}

func TestCompareString_WhenChangedAtEnd(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1more"
	var str2 = "str1papi"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	util.Assert(t, len(differences), 1)
	var diff0 = differences[0]
	util.Assert(t,diff0.path,"/")
	util.Assert(t,diff0.preamble,"str1")
	util.Assert(t,diff0.postamble,"")
	util.Assert(t,diff0.part1,"more")
	util.Assert(t,diff0.part2,"papi")
}
