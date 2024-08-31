package xml

import (
	"github.com/akedrou/textdiff"
	"github.com/stretchr/testify/assert"
	"testing"
	"xmldiff/internal/util"
)

func TestCompare_WhenEqual(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	<contactEmail>foo@example.org</contactEmail>
	<label>WooCommerce</label>
	<oauthConfig>
		<!-- Url for callback -->
		<callbackUrl>https://login.salesforce.com/services/oauth2/callback</callbackUrl>
		<consumerKey required="true">CLIENTID</consumerKey>
		<scopes>Basic</scopes>
		<scopes>Api</scopes>
		<scopes>Web</scopes>
		<scopes>Full</scopes>
	</oauthConfig>
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	util.Assert(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	<contactEmail>foo@example.org</contactEmail>
	<label>WooCommerce</label>
	<oauthConfig>
		<!-- Url for callback -->
		<callbackUrl>https://login.salesforce.com/services/oauth2/callback</callbackUrl>
		<consumerKey required="true">CLIENTID</consumerKey>
		<scopes>Basic</scopes>
		<scopes>Api</scopes>
		<scopes>Web</scopes>
		<scopes>Full</scopes>
	</oauthConfig>
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	util.Assert(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be no differences
	util.Assert(t, len(diffs), 0)
}

func TestCompare_WhenDifferentRootNodeName(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	assert.Equal(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp1 xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp1>`
	var xml2, err2 = Parse(xml2Str)
	util.Assert(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be expected
	util.Assert(t, len(diffs), 1)
	var diff0 = diffs[0].(StringDifferences)
	util.Assert(t, diff0.path, "/0.NAME")
	util.Assert(t, len(diff0.changes), 1)
	var change0 = diff0.changes[0]
	util.Assert(t, change0, textdiff.Edit{12, 12, "1"})
}

func TestCompare_WhenDifferentDataRootNode(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	util.Assert(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 2 different
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	util.Assert(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be expected
	util.Assert(t, len(diffs), 1)
	var diff0 = diffs[0].(StringDifferences)
	util.Assert(t, diff0.path, "/0.ConnectedApp.DATA")
	util.Assert(t, len(diff0.changes), 1)
	var change0 = diff0.changes[0]
	util.Assert(t, change0, textdiff.Edit{5, 6, "2"})
}

func TestCompare_WhenDifferentAttributeValueRootNode(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	util.Assert(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2007/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	util.Assert(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be expected
	util.Assert(t, len(diffs), 1)
	var diff = diffs[0].(StringDifferences)
	util.Assert(t, diff.path, "/0.ConnectedApp.ATTR.xmlns")
	util.Assert(t,len(diff.changes), 1)
	var change0 = diff.changes[0]
	util.Assert(t, change0, textdiff.Edit{26, 27, "7"})
}

func TestCompare_WhenDifferentAttributeNameRootNode(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	util.Assert(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmln="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	util.Assert(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be expected
	util.Assert(t, len(diffs), 1)
	var diff = diffs[0].(StringDifferences)
	util.Assert(t, diff.path, "/0.ConnectedApp.ATTR[0].NAME")
	util.Assert(t,len(diff.changes), 1)
	var change0 = diff.changes[0]
	util.Assert(t, change0, textdiff.Edit{4, 5, ""})
}

func TestCompare_WhenDifferentAttributeNumberRootNode(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	util.Assert(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata" qq="pp">
	Test 1 different
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	util.Assert(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be expected
	util.Assert(t, len(diffs), 1)
	var diff = diffs[0].(OtherDifference)
	util.Assert(t, diff.path, "/0.ConnectedApp.ATTR.LEN")
	util.Assert(t, diff.oldPart, "1")
	util.Assert(t, diff.newPart, "2")
}

func TestCompareChildrenNode(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	<contactEmail>foo@example.org</contactEmail>
	<label>WooCommerce</label>
	<oauthConfig>
		<!-- Url for callback -->
		<callbackUrl>https://login.salesforce.com/services/oauth2/callback</callbackUrl>
		<consumerKey required="true">CLIENTID</consumerKey>
		<scopes>Basic</scopes>
		<scopes>Api</scopes>
		<scopes>Web</scopes>
		<scopes>Full</scopes>
	</oauthConfig>
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	assert.Equal(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	<contactEmail>foo@example.org</contactEmail>
	<label>WooCommerce</label>
	<oauthConfig>
		<!-- Url for callback -->
		<callbackUrl>https://login.salesforce.com/services/oauth2/callback</callbackUrl>
		<consumerKey required="true">CLIENTID</consumerKey>
		<scopes>Basic</scopes>
		<scopes>Api1</scopes>
		<scopes>Web</scopes>
		<scopes>Full</scopes>
	</oauthConfig>
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	util.Assert(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be different
	util.Assert(t, len(diffs), 1)
	var diff = diffs[0].(StringDifferences)
	util.Assert(t, diff.path, "/0.ConnectedApp/2.oauthConfig/3.scopes.DATA")
	util.Assert(t,len(diff.changes), 1)
	var change0 = diff.changes[0]
	util.Assert(t, change0, textdiff.Edit{3, 3, "1"})
}
