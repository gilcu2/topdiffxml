package xml

import (
	"github.com/akedrou/textdiff"
	"testing"
	"xmldiff/internal/util"
)

func Test_StringDifferences_GetOutput(t *testing.T) {
	// Given string diff
	var diff = StringDifferences{
		path:   "/0.ConnectedApp.DATA",
		source: "WooCommerce",
		changes: []textdiff.Edit{
			textdiff.Edit{
				Start: 3,
				End:   6,
				New:   "Kan",
			},
		},
	}

	// When get output
	var output=diff.GetOutput()

	// Then it is the expected
	util.Assert(t,len(output),1)
	util.Assert(t,output[0],"/0.ConnectedApp.DATA[3:6]\n...Woo --(Com) ++(Kan) merce...\n")
}


func Test_OtherDifference_GetOutput(t *testing.T) {
	// Given string diff
	var diff = OtherDifference{
		path:   "/0.ConnectedApp.NODES.LEN",
		oldPart: "1",
		newPart: "2",
	}

	// When get output
	var output=diff.GetOutput()

	// Then it is the expected
	util.Assert(t,len(output),1)
	util.Assert(t,output[0],"/0.ConnectedApp.NODES.LEN\n--(1) ++(2)\n")
}
