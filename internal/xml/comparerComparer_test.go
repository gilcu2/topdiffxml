package xml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	// Given 2 xml
	var xml1_str = `<?xml version="1.0" encoding="UTF-8"?>
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
	var xml1,err1=Parse(xml1_str)
	assert.Equal(t, err1,nil)

	var xml2_str = `<?xml version="1.0" encoding="UTF-8"?>
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
	var xml2,err2=Parse(xml2_str)
	assert.Equal(t, err2,nil)

	// When compare
	var diffs, err = Compare(xml1, xml2)

	// Then must be different
	assert.Equal(t, err,nil)

	assert.Equal(t, len(diffs),1)
}

func TestCompare_WhenEqual(t *testing.T) {
	// Given 2 xml
	var xml1_str = `<?xml version="1.0" encoding="UTF-8"?>
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
	var xml1,err1=Parse(xml1_str)
	assert.Equal(t, err1,nil)

	var xml2_str = `<?xml version="1.0" encoding="UTF-8"?>
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
	var xml2,err2=Parse(xml2_str)
	assert.Equal(t, err2,nil)

	// When compare
	var diffs, err = Compare(xml1, xml2)

	// Then must be different
	assert.Equal(t, err,nil)

	assert.Equal(t, len(diffs),0)
}

func TestCompare_WhenDifferentRootNodeName(t *testing.T) {
	// Given 2 xml
	var xml1_str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1,err1=Parse(xml1_str)
	assert.Equal(t, err1,nil)

	var xml2_str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp1 xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml2,err2=Parse(xml2_str)
	assert.Equal(t, err2,nil)

	// When compare
	var diffs, err = Compare(xml1, xml2)

	// Then must be different
	assert.Equal(t, err,nil)

	assert.Equal(t, len(diffs),1)

	var diff=diffs[0]
	assert.Equal(t, diff.path,"/")
	assert.Equal(t, diff.preamble,"ConnectedApp")
	assert.Equal(t, diff.postamble,"")
	assert.Equal(t, diff.part1,"1")
	assert.Equal(t, diff.part2,"2")
}

func TestCompare_WhenDifferentDataRootNode(t *testing.T) {
	// Given 2 xml
	var xml1_str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1,err1=Parse(xml1_str)
	assert.Equal(t, err1,nil)

	var xml2_str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 2 different
</ConnectedApp>`
	var xml2,err2=Parse(xml2_str)
	assert.Equal(t, err2,nil)

	// When compare
	var diffs, err = Compare(xml1, xml2)

	// Then must be different
	assert.Equal(t, err,nil)

	assert.Equal(t, len(diffs),1)

	var diff=diffs[0]
	assert.Equal(t, diff.path,"/ConnectedApp")
	assert.Equal(t, diff.preamble,"Test ")
	assert.Equal(t, diff.postamble," different")
	assert.Equal(t, diff.part1,"1")
	assert.Equal(t, diff.part2,"2")
}

func TestCompare_WhenDifferentAttributeRootNode(t *testing.T) {
	// Given 2 xml
	var xml1_str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml1,err1=Parse(xml1_str)
	assert.Equal(t, err1,nil)

	var xml2_str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2007/04/metadata">
	Test 1 different
</ConnectedApp>`
	var xml2,err2=Parse(xml2_str)
	assert.Equal(t, err2,nil)

	// When compare
	var diffs, err = Compare(xml1, xml2)

	// Then must be different
	assert.Equal(t, err,nil)

	assert.Equal(t, len(diffs),1)

	var diff=diffs[0]
	assert.Equal(t, diff.preamble,"Test ")
	assert.Equal(t, diff.postamble," different")
	assert.Equal(t, diff.part1,"1")
	assert.Equal(t, diff.part2,"2")
}

