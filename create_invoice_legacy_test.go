package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCreateInvoiceLegacy(t *testing.T) {
	req := client.NewCreateInvoiceLegacyRequest()
	content := req.RequestBody().Content()
	content.Function.CreateInvoice.DateCreated.Year = "2022"
	content.Function.CreateInvoice.DateCreated.Month = "06"
	content.Function.CreateInvoice.DateCreated.Day = "30"
	content.Function.CreateInvoice.CustomerID = "280"
	content.Function.CreateInvoice.Action = "Draft"
	req.RequestBody().Operation.Content = content

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
