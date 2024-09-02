package xml

import (
	"github.com/akedrou/textdiff"
	"gotest.tools/v3/assert"
	"testing"
)

func TestCompareString_WhenEqual(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1"
	var str2 = "str1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, differences.path, "/")
	assert.Equal(t, differences.source, str1)
	assert.Equal(t, len(differences.changes), 0)
}

func TestCompareString_WhenAddedAtEnd(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1"
	var str2 = "str1more"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences.changes), 1)
	var change0 = differences.changes[0]
	assert.Equal(t,differences.path,"/")
	assert.Equal(t,change0,textdiff.Edit{4,4,"more"})
}

func TestCompareString_WhenAddedAtBegin(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1"
	var str2 = "morestr1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences.changes), 1)
	var change0 = differences.changes[0]
	assert.Equal(t,differences.path,"/")
	assert.Equal(t,change0,textdiff.Edit{0,0,"more"})
}

func TestCompareString_WhenAddedAtMiddle(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1str2"
	var str2 = "str1morestr2"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences.changes), 1)
	var change0 = differences.changes[0]
	assert.Equal(t,differences.path,"/")
	assert.Equal(t,change0,textdiff.Edit{4,4,"more"})
}

func TestCompareString_WhenRemovedAtEnd(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1more"
	var str2 = "str1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences.changes), 1)
	var change0 = differences.changes[0]
	assert.Equal(t,differences.path,"/")
	assert.Equal(t,change0,textdiff.Edit{4,8,""})
}

func TestCompareString_WhenRemovedAtBegin(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "morestr1"
	var str2 = "str1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences.changes), 1)
	var change0 = differences.changes[0]
	assert.Equal(t,differences.path,"/")
	assert.Equal(t,change0,textdiff.Edit{0,4,""})
}

func TestCompareString_WhenChangedAtBegin(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "morestr1"
	var str2 = "papistr1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences.changes), 1)
	var change0 = differences.changes[0]
	assert.Equal(t,differences.path,"/")
	assert.Equal(t,change0,textdiff.Edit{0,4,"papi"})
}

func TestCompareString_WhenChangedAtMiddle(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1morestr2"
	var str2 = "str1papistr2"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences.changes), 1)
	var change0 = differences.changes[0]
	assert.Equal(t,differences.path,"/")
	assert.Equal(t,change0,textdiff.Edit{4,8,"papi"})
}

func TestCompareString_WhenChangedAtEnd(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1more"
	var str2 = "str1papi"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences.changes), 1)
	var change0 = differences.changes[0]
	assert.Equal(t,change0,textdiff.Edit{4,8,"papi"})
}
