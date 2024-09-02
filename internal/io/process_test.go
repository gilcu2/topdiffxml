package io

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Test_CompareXMLStrings_ChildNodeData(t *testing.T) {
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

	// When compare
	var printedDiffs, err = CompareXmlStrings(xml1Str, xml2Str)
	assert.Equal(t, err, nil)

	// Then must be the expected
	assert.Equal(t, len(printedDiffs), 1)
	var diff = printedDiffs[0]
	assert.Equal(t, diff,
		"/0.ConnectedApp/2.oauthConfig/3.scopes.DATA[3:3]\n"+
			"...Api --() ++(1) ...\n")
}

func Test_CompareXMLStrings_ChildNodeAttributeName(t *testing.T) {
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

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	<contactEmail>foo@example.org</contactEmail>
	<label>WooCommerce</label>
	<oauthConfig>
		<!-- Url for callback -->
		<callbackUrl>https://login.salesforce.com/services/oauth2/callback</callbackUrl>
		<consumerKey requiresed="true">CLIENTID</consumerKey>
		<scopes>Basic</scopes>
		<scopes>Api1</scopes>
		<scopes>Web</scopes>
		<scopes>Full</scopes>
	</oauthConfig>
</ConnectedApp>`

	// When compare
	var printedDiffs, err = CompareXmlStrings(xml1Str, xml2Str)
	assert.Equal(t, err, nil)

	// Then must be the expected
	assert.Equal(t, len(printedDiffs), 2)
	assert.Equal(t, printedDiffs[0],
		"/0.ConnectedApp/2.oauthConfig/1.consumerKey.ATTR[0].NAME[6:6]\n"+
			"...requir --() ++(es) ed...\n")
	assert.Equal(t, printedDiffs[1],
		"/0.ConnectedApp/2.oauthConfig/3.scopes.DATA[3:3]\n"+
			"...Api --() ++(1) ...\n")
}

func Test_CompareXMLStrings_InvalidXML(t *testing.T) {
	// Given 2 xml
	var xml1Str = `<?xml version="1.0" encoding="UTF-8"?>
<ConnectedApp xmlns="http://soap.sforce.com/2006/04/metadata">
	<contactEmail>foo@example.org</contactEmail>
	<label>WooCommerce</label>
	<oauthConfig>
		<!-- Url for callback -->
		<callbackUrl>https://login.salesforce.com/services/oauth2/callback</callbackUrl>
		<consumerKey requiresed="true">CLIENTID</consumerKey>
		<scopes>Basic</scopes>
		<scopes>Api1</scopes>
		<scopes>Web</scopes>
		<scopes>Full</scopes>
	</oauthConfig>
</ConnectedApp>`

	var xml2Str = `<?xml version="1.0" encoding="UTF-8"?`

	// When compare
	var printedDiffs, err = CompareXmlStrings(xml1Str, xml2Str)

	// Then must be the expected
	assert.Equal(t, len(printedDiffs), 0)
	assert.ErrorContains(t, err, "Error parsing second xml")
}

func Test_CompareXMLFiles_ChildNode(t *testing.T) {
	// Given 2 xml
	var file1 = "testfiles/basicA.xml"
	var file2 = "testfiles/basicB.xml"

	// When compare
	var printedDiffs, err = CompareXmlFiles(file1, file2)
	assert.Equal(t, err, nil)

	// Then must be the expected
	assert.Equal(t, len(printedDiffs), 1)
	var diff = printedDiffs[0]
	assert.Equal(t, diff,
		"/0.ConnectedApp/2.oauthConfig/3.scopes.DATA[3:3]\n"+
			"...Api --() ++(1) ...\n")
}

func Test_CompareXMLFiles_InvalidFile1(t *testing.T) {
	// Given 2 xml
	var file1 = "testfiles/basicA1.xml"
	var file2 = "testfiles/basicB.xml"

	// When compare
	var printedDiffs, err = CompareXmlFiles(file1, file2)

	// Then must be the expected
	assert.Equal(t, len(printedDiffs), 0)
	assert.ErrorContains(t, err, "Error reading file testfiles/basicA1.xml")
}

func Test_CompareXMLFiles_InvalidFile2(t *testing.T) {
	// Given 2 xml
	var file1 = "testfiles/basicA.xml"
	var file2 = "testfiles/basicB1.xml"

	// When compare
	var printedDiffs, err = CompareXmlFiles(file1, file2)

	// Then must be the expected
	assert.Equal(t, len(printedDiffs), 0)
	assert.ErrorContains(t, err, "Error reading file testfiles/basicB1.xml")
}

func Test_CompareXMLFiles_InvalidFileXML(t *testing.T) {
	// Given 2 xml
	var file1 = "testfiles/invalid.xml"
	var file2 = "testfiles/basicB.xml"

	// When compare
	var printedDiffs, err = CompareXmlFiles(file1, file2)

	// Then must be the expected
	assert.Equal(t, len(printedDiffs), 0)
	assert.ErrorContains(t, err, "Error parsing first xml")
}
