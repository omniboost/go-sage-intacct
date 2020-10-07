package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetAPISession(t *testing.T) {
	req := client.NewGetAPISessionRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
