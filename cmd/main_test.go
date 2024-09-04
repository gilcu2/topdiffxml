package main

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Test_realMain(t *testing.T) {
	// When call
	var r=realMain()

	// Then must be the expected
	assert.Equal(t,r,1)
}