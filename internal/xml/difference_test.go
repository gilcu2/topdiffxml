package xml

import (
	"github.com/akedrou/textdiff"
	"testing"
	"gotest.tools/v3/assert"
)

func Test_StringDifferences_GetOutput_1Difference(t *testing.T) {
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
	assert.Equal(t,len(output),1)
	assert.Equal(t,output[0],"/0.ConnectedApp.DATA[3:6]\n...Woo --(Com) ++(Kan) merce...\n")
}

func Test_StringDifferences_GetOutput_2Differences(t *testing.T) {
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
			textdiff.Edit{
				Start: 11,
				End:   11,
				New:   "papa",
			},
		},
	}

	// When get output
	var output=diff.GetOutput()

	// Then it is the expected
	assert.Equal(t,len(output),2)
	assert.Equal(t,output[0],"/0.ConnectedApp.DATA[3:6]\n...Woo --(Com) ++(Kan) me...\n")
	assert.Equal(t,output[1],"/0.ConnectedApp.DATA[11:11]\n...rce --() ++(papa) ...\n")
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
	assert.Equal(t,len(output),1)
	assert.Equal(t,output[0],"/0.ConnectedApp.NODES.LEN\n--(1) ++(2)\n")
}
