package util

import (
	"github.com/gkampitakis/coverage"
	"testing"
)

func TestMain(m *testing.M) {
	coverage.Run(m, 0)
}