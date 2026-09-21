package intacct_test

import (
	"encoding/json"
	"log"
	"testing"

	intacct "github.com/omniboost/go-sage-intacct"
)

func TestGetVendors(t *testing.T) {
	req := client.NewGetVendorsRequest()
	content := req.RequestBody().Content()
	content.Function.Query.Select = append(content.Function.Query.Select, []string{
		"VENDORID",
		"NAME",
	}...)
	content.Function.Query.Filters = append(content.Function.Query.Filters, intacct.EqualTo{Field: "NAME", Value: "SilverDoor"})
	req.RequestBody().Operation.Content = content
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
