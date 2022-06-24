package intacct_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetTaxDetails(t *testing.T) {
	req := client.NewGetTaxDetailsRequest()

	content := req.RequestBody().Content()
	content.Function.Query.Select = append(content.Function.Query.Select, []string{
		"RECORDNO",
		"DETAILID",
		"TAXUID",
		"DESCRIPTION",
		"TAXTYPE",
		"VALUE",
		"MINTAXABLE",
		"MAXTAXABLE",
		"INCLUDE",
		"MINTAX",
		"MAXTAX",
		"GLACCOUNT",
		"TAXAUTHORITY",
		"STATUS",
		"SYSGENERATED",
		"REVERSECHARGE",
		"TAXRATE",
		"TAXSOLUTIONID",
		"USEEXPENSEACCT",
		"MEGAENTITYKEY",
		"MEGAENTITYID",
		"MEGAENTITYNAME",
		"RECORD_URL",
	}...)
	// content.Function.Query.Filters = append(content.Function.Query.Filters, intacct.EqualTo{Field: "DOCNUMBER", Value: "SI000018"})
	// content.Function.Query.Filters = append(content.Function.Query.Filters, intacct.Like{Field: "DOCNUMBER", Value: "%0%"})
	// content.Function.Query.Filters = append(content.Function.Query.Filters, intacct.GreaterThan{Field: "WHENCREATED", Value: "01/01/2022"})
	// content.Function.ReadByQuery.Select = []string{"JOURNAL"}
	req.RequestBody().Operation.Content = content
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
