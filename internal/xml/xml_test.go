package xml

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func assertNode(node *Node, name string, data string,
	attributesName []string, attributesValue []string,
	childrens int, t *testing.T) {
	assert.Equal(t, node.XMLName.Local,name)
	assert.Equal(t, node.Data,data)

	assert.Equal(t, len(node.Attributes),len(attributesName))
	for i := 0; i < len(attributesName); i++ {
		assert.Equal(t, node.Attributes[i].Name.Local, attributesName[i])
		assert.Equal(t, node.Attributes[i].Value, attributesValue[i])
	}

	assert.Equal(t, len(node.Nodes),childrens)
}

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
	assertNode(topNode,"ConnectedApp","",
		[]string{"xmlns"},[]string{"http://soap.sforce.com/2006/04/metadata"},
		3,t,
	)

	var child0=topNode.Nodes[0]
	assertNode(child0,"contactEmail","foo@example.org",
		[]string{},[]string{},
		0,t,
	)

	var child1=topNode.Nodes[1]
	assertNode(child1,"label","WooCommerce",
		[]string{},[]string{},
		0,t,
	)

	var child2=topNode.Nodes[2]
	assertNode(child2,"oauthConfig","",
		[]string{},[]string{},
		6,t,
	)

	var child2_0=child2.Nodes[0]
	assertNode(child2_0,"callbackUrl","https://login.salesforce.com/services/oauth2/callback",
		[]string{},[]string{},
		0,t,
	)
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
