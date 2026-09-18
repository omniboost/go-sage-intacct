package intacct_test

import (
	"encoding/json"
	"log"
	"testing"

	intacct "github.com/omniboost/go-sage-intacct"
)

func TestGetBills(t *testing.T) {
	req := client.NewGetBillsRequest()
	content := req.RequestBody().Content()
	content.Function.Query.Select = append(content.Function.Query.Select, []string{
		"RECORDNO",
		"VENDORID",
		"STATE",
	}...)
	content.Function.Query.Filters = append(content.Function.Query.Filters, intacct.EqualTo{Field: "VENDORID", Value: "280"})
	req.RequestBody().Operation.Content = content
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
