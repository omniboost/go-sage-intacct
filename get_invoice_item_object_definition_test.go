package intacct_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetInvoiceItemObjectDefinition(t *testing.T) {
	req := client.NewGetInvoiceItemObjectDefinitionRequest()

	content := req.RequestBody().Content()
	req.RequestBody().Operation.Content = content
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
