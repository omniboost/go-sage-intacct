package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetAttachmentByIDLegacy(t *testing.T) {
	req := client.NewGetAttachmentByIDLegacyRequest()
	content := req.RequestBody().Content()
	content.Function.Get.Object = "supdoc"
	content.Function.Get.Key = "123.pdf"
	req.RequestBody().Operation.Content = content

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
