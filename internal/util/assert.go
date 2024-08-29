package util

import (
	oldassert "github.com/stretchr/testify/assert"
	"testing"
)

func Assert	(t *testing.T,  actual interface{}, expected interface{}) {
	oldassert.Equal(t, expected, actual)
}
