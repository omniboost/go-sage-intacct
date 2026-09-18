package intacct_test

import (
	"encoding/json"
	"log"
	"testing"

	intacct "github.com/omniboost/go-sage-intacct"
)

func TestCreateBill(t *testing.T) {
	req := client.NewCreateBillRequest()
	content := req.RequestBody().Content()
	content.Function.Create.APBill.VendorID = "280"
	content.Function.Create.APBill.WhenCreated = "06/30/2022"
	content.Function.Create.APBill.WhenDue = "06/30/2022"
	content.Function.Create.APBill.Description = "TEST"
	content.Function.Create.APBill.Action = "Draft"
	content.Function.Create.APBill.Currency = "GBP"
	content.Function.Create.APBill.BillItems = append(content.Function.Create.APBill.BillItems, intacct.BillItem{
		AccountNo: "1000",
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
