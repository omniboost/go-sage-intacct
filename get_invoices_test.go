package intacct_test

import (
	"encoding/json"
	"log"
	"testing"

	intacct "github.com/omniboost/go-sage-intacct"
)

func TestGetInvoices(t *testing.T) {
	req := client.NewGetInvoicesRequest()

	content := req.RequestBody().Content()
	content.Function.Query.Select = append(content.Function.Query.Select, []string{
		"RECORDNO",
		"RECORDTYPE",
		"RECORDID",
		"CONTACTTAXGROUP",
		"STATE",
		"RAWSTATE",
		"CUSTOMERID",
		"CUSTOMERNAME",
		"CUSTEMAILOPTIN",
		"TRX_ENTITYDUE",
		"CUSTMESSAGEID",
		"CUSTMESSAGE.MESSAGE",
		"DELIVERY_OPTIONS",
		"DOCNUMBER",
		"DESCRIPTION",
		"DESCRIPTION2",
		"TERMNAME",
		"TERMKEY",
		"TERMVALUE",
		"WHENCREATED",
		"WHENPOSTED",
		"WHENDISCOUNT",
		"WHENDUE",
		"WHENPAID",
		"BASECURR",
		"CURRENCY",
		"EXCH_RATE_DATE",
		"EXCH_RATE_TYPE_ID",
		"EXCHANGE_RATE",
		"TOTALENTERED",
		"TOTALSELECTED",
		"TOTALPAID",
		"TOTALDUE",
		"TRX_TOTALENTERED",
		"TRX_TOTALSELECTED",
		"TRX_TOTALPAID",
		"TRX_TOTALDUE",
		"BILLTOPAYTOCONTACTNAME",
		"SHIPTORETURNTOCONTACTNAME",
		"BILLTOPAYTOKEY",
		"SHIPTORETURNTOKEY",
		"CONTACT.CONTACTNAME",
		"CONTACT.PREFIX",
		"CONTACT.FIRSTNAME",
		"CONTACT.INITIAL",
		"CONTACT.LASTNAME",
		"CONTACT.COMPANYNAME",
		"CONTACT.PRINTAS",
		"CONTACT.PHONE1",
		"CONTACT.PHONE2",
		"CONTACT.CELLPHONE",
		"CONTACT.PAGER",
		"CONTACT.FAX",
		"CONTACT.EMAIL1",
		"CONTACT.EMAIL2",
		"CONTACT.URL1",
		"CONTACT.URL2",
		"CONTACT.VISIBLE",
		"CONTACT.MAILADDRESS.ADDRESS1",
		"CONTACT.MAILADDRESS.ADDRESS2",
		"CONTACT.MAILADDRESS.CITY",
		"CONTACT.MAILADDRESS.STATE",
		"CONTACT.MAILADDRESS.ZIP",
		"CONTACT.MAILADDRESS.COUNTRY",
		"CONTACT.MAILADDRESS.COUNTRYCODE",
		"SHIPTO.CONTACTNAME",
		"SHIPTO.PREFIX",
		"SHIPTO.FIRSTNAME",
		"SHIPTO.INITIAL",
		"SHIPTO.LASTNAME",
		"SHIPTO.COMPANYNAME",
		"SHIPTO.PRINTAS",
		"SHIPTO.PHONE1",
		"SHIPTO.PHONE2",
		"SHIPTO.CELLPHONE",
		"SHIPTO.PAGER",
		"SHIPTO.FAX",
		"SHIPTO.EMAIL1",
		"SHIPTO.EMAIL2",
		"SHIPTO.URL1",
		"SHIPTO.URL2",
		"SHIPTO.VISIBLE",
		"SHIPTO.MAILADDRESS.ADDRESS1",
		"SHIPTO.MAILADDRESS.ADDRESS2",
		"SHIPTO.MAILADDRESS.CITY",
		"SHIPTO.MAILADDRESS.STATE",
		"SHIPTO.MAILADDRESS.ZIP",
		"SHIPTO.MAILADDRESS.COUNTRY",
		"SHIPTO.MAILADDRESS.COUNTRYCODE",
		"BILLTO.CONTACTNAME",
		"BILLTO.PREFIX",
		"BILLTO.FIRSTNAME",
		"BILLTO.INITIAL",
		"BILLTO.LASTNAME",
		"BILLTO.COMPANYNAME",
		"BILLTO.PRINTAS",
		"BILLTO.PHONE1",
		"BILLTO.PHONE2",
		"BILLTO.CELLPHONE",
		"BILLTO.PAGER",
		"BILLTO.FAX",
		"BILLTO.EMAIL1",
		"BILLTO.EMAIL2",
		"BILLTO.URL1",
		"BILLTO.URL2",
		"BILLTO.VISIBLE",
		"BILLTO.MAILADDRESS.ADDRESS1",
		"BILLTO.MAILADDRESS.ADDRESS2",
		"BILLTO.MAILADDRESS.CITY",
		"BILLTO.MAILADDRESS.STATE",
		"BILLTO.MAILADDRESS.ZIP",
		"BILLTO.MAILADDRESS.COUNTRY",
		"BILLTO.MAILADDRESS.COUNTRYCODE",
		"PRBATCH",
		"PRBATCHKEY",
		"MODULEKEY",
		"SCHOPKEY",
		"SYSTEMGENERATED",
		"HASPOSTEDREVREC",
		"BILLBACKTEMPLATEKEY",
		"AUWHENCREATED",
		"WHENMODIFIED",
		"CREATEDBY",
		"MODIFIEDBY",
		"DUE_IN_DAYS",
		"SHIPTO.TAXGROUP.NAME",
		"SHIPTO.TAXGROUP.RECORDNO",
		"SHIPTO.TAXID",
		"TAXSOLUTIONID",
		"RETAINAGEPERCENTAGE",
		"TRX_TOTALRETAINED",
		"TRX_TOTALRELEASED",
		"TOTALRETAINED",
		"SUPDOCID",
		"PROJECTCONTRACTKEY",
		"PROJECTCONTRACTID",
		"DUNNINGCOUNT",
		"MEGAENTITYKEY",
		"MEGAENTITYID",
		"MEGAENTITYNAME",
		"RECORD_URL",
	}...)
	content.Function.Query.Filters = append(
		content.Function.Query.Filters,
		intacct.EqualTo{Field: "RECORDID", Value: "SI000018"},
	)
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
