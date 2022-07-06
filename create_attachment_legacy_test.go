package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCreateAttachmentLegacy(t *testing.T) {
	req := client.NewCreateAttachmentLegacyRequest()
	content := req.RequestBody().Content()
	req.RequestBody().Operation.Content = content

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
