

package util

import (
	"strconv"
	"testing"
)

func TestToStringInt(t *testing.T) {
	// Given int
	var i=43

	// When convert to str
	var r=ToString(i)

	// Then must be the expected
	Assert(t,r,strconv.Itoa(i))
}

func TestToStringFloat(t *testing.T) {
	// Given int
	var f =43.43

	// When convert to str
	var r=ToString(f)

	// Then must be the expected
	Assert(t,r,strconv.FormatFloat(f, 'f', -1, 64))
}