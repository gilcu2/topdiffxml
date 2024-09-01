package io

import (
	"fmt"
	"os"
	"testing"
	"xmldiff/internal/util"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	rc := m.Run()

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < 0.8 {
			fmt.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}
	os.Exit(rc)
}

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
	util.Assert(t, err, nil)

	// Then must be the expected
	util.Assert(t, len(printedDiffs), 1)
	var diff = printedDiffs[0]
	util.Assert(t, diff,
		"/0.ConnectedApp/2.oauthConfig/3.scopes.DATA[3:3]\n" +
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
	util.Assert(t, err, nil)

	// Then must be the expected
	util.Assert(t, len(printedDiffs), 2)
	util.Assert(t, printedDiffs[0],
		"/0.ConnectedApp/2.oauthConfig/1.consumerKey.ATTR[0].NAME[6:6]\n" +
		"...requir --() ++(es) ed...\n")
	util.Assert(t, printedDiffs[1],
		"/0.ConnectedApp/2.oauthConfig/3.scopes.DATA[3:3]\n" +
		"...Api --() ++(1) ...\n")
}


func Test_CompareXMLFiles_ChildNode(t *testing.T) {
	// Given 2 xml
	var file1 = "testfiles/basicA.xml"
	var file2 = "testfiles/basicB.xml"

	// When compare
	var printedDiffs, err = CompareXmlFiles(file1, file2)
	util.Assert(t, err, nil)

	// Then must be the expected
	util.Assert(t, len(printedDiffs), 1)
	var diff = printedDiffs[0]
	util.Assert(t, diff,
		"/0.ConnectedApp/2.oauthConfig/3.scopes.DATA[3:3]\n" +
		"...Api --() ++(1) ...\n")
}
