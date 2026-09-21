package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCreateAPAdjustment(t *testing.T) {
	req := client.NewCreateAPAdjustmentRequest()
	content := req.RequestBody().Content()
	content.Function.CreateAPAdjustment.DateCreated.Year = "2022"
	content.Function.CreateAPAdjustment.DateCreated.Month = "06"
	content.Function.CreateAPAdjustment.DateCreated.Day = "30"
	content.Function.CreateAPAdjustment.VendorID = "280"
	content.Function.CreateAPAdjustment.Action = "Draft"
	req.RequestBody().Operation.Content = content

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
