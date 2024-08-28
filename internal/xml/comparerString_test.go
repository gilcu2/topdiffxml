package xml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareString_WhenEqual(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1"
	var str2 = "str1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences), 0)
}

func TestCompareString_WhenAddedAtEnd(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1"
	var str2 = "str1more"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences), 1)
	var diff0 = differences[0]
	assert.Equal(t,diff0.path,"/")
	assert.Equal(t,diff0.preamble,"str1")
	assert.Equal(t,diff0.postamble,"")
	assert.Equal(t,diff0.part1,"")
	assert.Equal(t,diff0.part2,"more")
}

func TestCompareString_WhenRemovedAtEnd(t *testing.T) {
	// Given 2 string where second add more chars
	var str1 = "str1more"
	var str2 = "str1"

	// When compare
	var differences = getStringDifferences(str1, str2, "/")

	// Then must be expected
	assert.Equal(t, len(differences), 1)
	var diff0 = differences[0]
	assert.Equal(t,diff0.path,"/")
	assert.Equal(t,diff0.preamble,"str1")
	assert.Equal(t,diff0.postamble,"")
	assert.Equal(t,diff0.part1,"more")
	assert.Equal(t,diff0.part2,"")
}
