package intacct_test

import (
	"encoding/json"
	"log"
	"testing"

	intacct "github.com/omniboost/go-sage-intacct"
)

func TestGetJournalEntryLines(t *testing.T) {
	req := client.NewGetJournalEntryLinesRequest()

	content := req.RequestBody().Content()
	content.Function.Query.Select = append(content.Function.Query.Select, "AMOUNT", "DEPARTMENT", "LOCATION", "PROJECTID", "CLASSID")
	content.Function.Query.Filters = append(content.Function.Query.Filters, intacct.EqualTo{Field: "BATCHNO", Value: "17"})
	// content.Function.ReadByQuery.Select = []string{"JOURNAL"}
	req.RequestBody().Operation.Content = content
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
