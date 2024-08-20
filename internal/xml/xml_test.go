package xml

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// Given xml
	var xml_str = `<?xml version="1.0" encoding="UTF-8"?>
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

	// When parsed
	var parsed,err=Parse(xml_str)

	// Then it is expected
	assert.Nil(t, err)

	var topNode=parsed
	assert.Equal(t, topNode.XMLName.Local,"ConnectedApp")
	assert.Equal(t, topNode.Attributes[0].Name.Local,"xmlns")
	assert.Equal(t, topNode.Attributes[0].Value,"http://soap.sforce.com/2006/04/metadata")
}

func TestCompare(t *testing.T) {
	// Given 2 xml
	var xml1 = `<?xml version="1.0" encoding="UTF-8"?>
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

	var xml2 = `<?xml version="1.0" encoding="UTF-8"?>
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
	var diff, err = Compare(xml1, xml2)

	// Then must be different
	if err != nil {
		t.Error(err)
	}

	if len(diff) == 0 {
		t.Error("xml compare fail")
	}

}
