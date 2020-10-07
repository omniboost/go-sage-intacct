package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetDepartments(t *testing.T) {
	req := client.NewGetDepartmentsRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
