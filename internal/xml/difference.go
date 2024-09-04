package xml

import (
	"fmt"
	"github.com/akedrou/textdiff"
)

type XMLDifference interface {
	GetOutput() []string
}

type StringDifferences struct {
	path    string
	source  string
	changes []textdiff.Edit
}

func (sd StringDifferences) GetOutput() []string {
	var contextBegin int
	var contextEnd = 0
	var r []string
	for i := range len(sd.changes) {
		var change = sd.changes[i]
		var s = fmt.Sprintf("%s[%d:%d]\n", sd.path, change.Start, change.End)
		contextBegin = max(contextEnd, change.Start-10)
		if i < len(sd.changes)-1 {
			contextEnd = min(change.End+10, (change.End+sd.changes[i+1].Start)/2)
		} else {
			contextEnd = min(len(sd.source), change.End+10)
		}
		var leftContext = sd.source[contextBegin:change.Start]
		var rightContext = sd.source[change.End:contextEnd]
		var oldPart = sd.source[change.Start:change.End]
		var newPart = change.New
		s += fmt.Sprintf("...%s --(%s) ++(%s) %s...\n",
			leftContext, oldPart, newPart, rightContext)
		r = append(r, s)
	}
	return r
}

type OtherDifference struct {
	path    string
	oldPart string
	newPart string
}

func (sd OtherDifference) GetOutput() []string {
	var s = fmt.Sprintf("%s\n--(%s) ++(%s)\n", sd.path, sd.oldPart, sd.newPart)
	return []string{s}
}
