package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	req := client.NewGetAccountsRequest()
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
