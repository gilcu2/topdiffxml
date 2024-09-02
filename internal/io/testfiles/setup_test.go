package testfiles

import (
	"github.com/gkampitakis/coverage"
	"testing"
)

func TestMain(m *testing.M) {
	coverage.Run(m, 95)
}