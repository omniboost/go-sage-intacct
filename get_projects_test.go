package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetProjects(t *testing.T) {
	req := client.NewGetProjectsRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
