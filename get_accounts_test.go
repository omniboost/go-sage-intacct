package intacct_test

import (
	"encoding/json"
	"log"
	"testing"

	intacct "github.com/omniboost/go-sage-intacct"
)

func TestGetAccounts(t *testing.T) {
	req := client.NewGetAccountsRequest()
	c := req.RequestBody().Content()
	c.Function.ReadByQuery.Query = intacct.RawQuery("ACCOUNTNO = 5004760")
	req.RequestBody().Operation.Content = c
	resp, err := req.Do()
	// resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
