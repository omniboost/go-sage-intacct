package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetProjects(t *testing.T) {
	req := client.NewGetProjectsRequest()
	content := req.RequestBody().Content()
	content.Function.ReadByQuery.PageSize = 1000
	req.RequestBody().Operation.Content = content
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
