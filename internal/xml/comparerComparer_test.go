package xml

import (
	"github.com/akedrou/textdiff"
	"github.com/stretchr/testify/assert"
	"testing"
	"xmldiff/internal/util"
)

func TestCompare(t *testing.T) {
	t.Skip("Skipping testing until explore down nodes")
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
	var diffs, err = Compare(xml1, xml2)
	util.Assert(t, err, nil)

	// Then must be different
	util.Assert(t, len(diffs), 1)
}

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
		<scopes>Api1</scopes>
		<scopes>Web</scopes>
		<scopes>Full</scopes>
	</oauthConfig>
</ConnectedApp>`
	var xml2, err2 = Parse(xml2Str)
	util.Assert(t, err2, nil)

	// When compare
	var diffs, err = Compare(xml1, xml2)

	// Then must be different
	util.Assert(t, err, nil)

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
	var diffs, err = Compare(xml1, xml2)
	util.Assert(t, err, nil)

	// Then must be expected
	util.Assert(t, len(diffs), 1)
	var diff0 = diffs[0]
	util.Assert(t, diff0.path, "/")
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
	var diffs, err = Compare(xml1, xml2)
	util.Assert(t, err, nil)

	// Then must be expected
	util.Assert(t, len(diffs), 1)
	var diff0 = diffs[0]
	util.Assert(t, diff0.path, "/ConnectedApp")
	util.Assert(t, len(diff0.changes), 1)
	var change0 = diff0.changes[0]
	util.Assert(t, change0, textdiff.Edit{5, 6, "2"})
}

func TestCompare_WhenDifferentAttributeRootNode(t *testing.T) {
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
	var diffs, err = Compare(xml1, xml2)
	util.Assert(t, err, nil)

	// Then must be expected
	util.Assert(t, len(diffs), 1)
	var diff = diffs[0]
	util.Assert(t, diff.path, "/")
	var change0 = diff.changes[0]
	util.Assert(t, change0, textdiff.Edit{5, 6, "2"})
}
