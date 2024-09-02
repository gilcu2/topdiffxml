package xml

import (
	"github.com/akedrou/textdiff"
	"gotest.tools/v3/assert"
	"testing"
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
		<scopes>Api</scopes>
		<scopes>Web</scopes>
		<scopes>Full</scopes>
	</oauthConfig>
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	assert.Equal(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be no differences
	assert.Equal(t, len(diffs), 0)
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
	assert.Equal(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be expected
	assert.Equal(t, len(diffs), 1)
	var diff0 = diffs[0].(StringDifferences)
	assert.Equal(t, diff0.path, "/0.NAME")
	assert.Equal(t, len(diff0.changes), 1)
	var change0 = diff0.changes[0]
	assert.Equal(t, change0, textdiff.Edit{Start: 12, End: 12, New: "1"})
}

func TestCompare_WhenDifferentDataRootNode(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	assert.Equal(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 2 different
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	assert.Equal(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be expected
	assert.Equal(t, len(diffs), 1)
	var diff0 = diffs[0].(StringDifferences)
	assert.Equal(t, diff0.path, "/0.ConnectedApp.DATA")
	assert.Equal(t, len(diff0.changes), 1)
	var change0 = diff0.changes[0]
	assert.Equal(t, change0, textdiff.Edit{Start: 5, End: 6, New: "2"})
}

func TestCompare_WhenDifferentAttributeValueRootNode(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	assert.Equal(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2007/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	assert.Equal(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be expected
	assert.Equal(t, len(diffs), 1)
	var diff = diffs[0].(StringDifferences)
	assert.Equal(t, diff.path, "/0.ConnectedApp.ATTR.xmlns")
	assert.Equal(t,len(diff.changes), 1)
	var change0 = diff.changes[0]
	assert.Equal(t, change0, textdiff.Edit{Start: 26, End: 27, New: "7"})
}

func TestCompare_WhenDifferentAttributeNameRootNode(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	assert.Equal(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmln="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	assert.Equal(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be expected
	assert.Equal(t, len(diffs), 1)
	var diff = diffs[0].(StringDifferences)
	assert.Equal(t, diff.path, "/0.ConnectedApp.ATTR[0].NAME")
	assert.Equal(t,len(diff.changes), 1)
	var change0 = diff.changes[0]
	assert.Equal(t, change0, textdiff.Edit{Start: 4, End: 5})
}

func TestCompare_WhenDifferentAttributeNumberRootNode(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1, err1 = Parse(xml1Str)
	assert.Equal(t, err1, nil)

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata" qq="pp">
	Test 1 different
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	assert.Equal(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be expected
	assert.Equal(t, len(diffs), 1)
	var diff = diffs[0].(OtherDifference)
	assert.Equal(t, diff.path, "/0.ConnectedApp.ATTR.LEN")
	assert.Equal(t, diff.oldPart, "1")
	assert.Equal(t, diff.newPart, "2")
}

func TestCompareChildrenNodeData(t *testing.T) {
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
	assert.Equal(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be different
	assert.Equal(t, len(diffs), 1)
	var diff = diffs[0].(StringDifferences)
	assert.Equal(t, diff.path, "/0.ConnectedApp/2.oauthConfig/3.scopes.DATA")
	assert.Equal(t,len(diff.changes), 1)
	var change0 = diff.changes[0]
	assert.Equal(t, change0, textdiff.Edit{Start: 3, End: 3, New: "1"})
}


func TestCompareChildrenLen(t *testing.T) {
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
		<scopes>Web</scopes>
		<scopes>Full</scopes>
	</oauthConfig>
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	assert.Equal(t, err2, nil)

	// When compare
	var diffs = Compare(xml1, xml2)

	// Then must be different
	assert.Equal(t, len(diffs), 3)

	var diff0 = diffs[0].(OtherDifference)
	assert.Equal(t, diff0.path, "/0.ConnectedApp/2.oauthConfig.NODES.LEN")
	assert.Equal(t,diff0.oldPart, "6")
	assert.Equal(t,diff0.newPart, "5")

	var diff1 = diffs[1].(StringDifferences)
	assert.Equal(t, diff1.path, "/0.ConnectedApp/2.oauthConfig/3.scopes.DATA")
	assert.Equal(t,len(diff1.changes), 1)
	var change1_0 = diff1.changes[0]
	assert.Equal(t, change1_0, textdiff.Edit{End: 3, New: "Web"})

	var diff2 = diffs[2].(StringDifferences)
	assert.Equal(t, diff2.path, "/0.ConnectedApp/2.oauthConfig/4.scopes.DATA")
	assert.Equal(t,len(diff2.changes), 1)
	var change2_0 = diff2.changes[0]
	assert.Equal(t, change2_0, textdiff.Edit{End: 3, New: "Full"})
}
