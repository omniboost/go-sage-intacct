package intacct_test

import (
	"encoding/json"
	"log"
	"testing"

	intacct "github.com/omniboost/go-sage-intacct"
)

func TestGetInvoiceItems(t *testing.T) {
	req := client.NewGetInvoiceItemsRequest()

	content := req.RequestBody().Content()
	content.Function.Query.Select = append(content.Function.Query.Select, []string{
		"RECORDNO",
		"RECORDKEY",
		"ACCOUNTKEY",
		"ACCOUNTNO",
		"OFFSETACCOUNTKEY",
		"OFFSETGLACCOUNTNO",
		"OFFSETGLACCOUNTTITLE",
		"ACCOUNTTITLE",
		"ACCOUNTLABELKEY",
		"ACCOUNTLABEL",
		"ENTRY_DATE",
		"AMOUNT",
		"TRX_AMOUNT",
		"DEPARTMENTID",
		"DEPARTMENTNAME",
		"LOCATIONID",
		"LOCATIONNAME",
		"ENTRYDESCRIPTION",
		"EXCH_RATE_DATE",
		"EXCH_RATE_TYPE_ID",
		"EXCHANGE_RATE",
		"ALLOCATIONKEY",
		"ALLOCATION",
		"LINEITEM",
		"LINE_NO",
		"CURRENCY",
		"BASECURR",
		"TOTALPAID",
		"TRX_TOTALPAID",
		"TOTALSELECTED",
		"TRX_TOTALSELECTED",
		"SUBTOTAL",
		"PARENTENTRY",
		"DEFERREVENUE",
		"REVRECTEMPLATEKEY",
		"REVRECTEMPLATE",
		"DEFERREDREVACCTKEY",
		"DEFERREDREVACCTNO",
		"DEFERREDREVACCTTITLE",
		"REVRECSTARTDATE",
		"REVRECENDDATE",
		"BASELOCATION",
		"STATE",
		"RECORDTYPE",
		"DETAILKEY",
		"WHENCREATED",
		"WHENMODIFIED",
		"CREATEDBY",
		"MODIFIEDBY",
		"RETAINAGEPERCENTAGE",
		"TRX_AMOUNTRETAINED",
		"AMOUNTRETAINED",
		"TRX_AMOUNTRELEASED",
		"RETAINAGE_OFFSETGLACCOUNTNO",
		"PROJECTCONTRACTID",
		"PROJECTCONTRACTKEY",
		"PROJECTCONTRACTLINEID",
		"PROJECTCONTRACTLINEKEY",
		"PROJECTDIMKEY",
		"PROJECTID",
		"PROJECTNAME",
		"CUSTOMERDIMKEY",
		"CUSTOMERID",
		"CUSTOMERNAME",
		"VENDORDIMKEY",
		"VENDORID",
		"VENDORNAME",
		"EMPLOYEEDIMKEY",
		"EMPLOYEEID",
		"EMPLOYEENAME",
		"ITEMDIMKEY",
		"ITEMID",
		"ITEMNAME",
		"CLASSDIMKEY",
		"CLASSID",
		"CLASSNAME",
		"RECORD_URL",
	}...)
	content.Function.Query.Filters = append(content.Function.Query.Filters, intacct.EqualTo{Field: "RECORDKEY", Value: "4151"})
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
