package intacct_test

import (
	"encoding/json"
	"log"
	"testing"

	intacct "github.com/omniboost/go-sage-intacct"
)

func TestGetJournalEntries(t *testing.T) {
	req := client.NewGetJournalEntriesRequest()

	content := req.RequestBody().Content()
	content.Function.ReadByQuery.Query = intacct.RawQuery("JOURNAL = 'JR'")
	// content.Function.ReadByQuery.Select = []string{"JOURNAL"}
	req.RequestBody().Operation.Content = content
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
