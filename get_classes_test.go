package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetClasses(t *testing.T) {
	req := client.NewGetClassesRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
