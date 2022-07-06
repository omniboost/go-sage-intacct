package intacct_test

import (
	"encoding/json"
	"log"
	"testing"

	intacct "github.com/omniboost/go-sage-intacct"
)

func TestCreateInvoice(t *testing.T) {
	req := client.NewCreateInvoiceRequest()
	content := req.RequestBody().Content()
	content.Function.Create.ARInvoice.CustomerID = "280"
	content.Function.Create.ARInvoice.WhenCreated = "06/30/2022"
	content.Function.Create.ARInvoice.WhenDue = "06/30/2022"
	content.Function.Create.ARInvoice.Description = "TEST"
	// content.Function.Create.ARInvoice.RecordID = "TESTID"
	content.Function.Create.ARInvoice.State = "Draft"
	// content.Function.Create.ARInvoice.BaseCurr = "GBP"
	content.Function.Create.ARInvoice.Currency = "GBP"
	content.Function.Create.ARInvoice.InvoiceItems = append(content.Function.Create.ARInvoice.InvoiceItems, intacct.InvoiceItem{
		// BaseCurr:  "GBP",
		Currency:  "GBP",
		Amount:    12.0,
		TrxAmount: 12.0,
	})
	req.RequestBody().Operation.Content = content

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
