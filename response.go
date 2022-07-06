package intacct

import (
	"encoding/xml"
	"time"
)

func NewResponse() Response {
	return Response{
		Control: ResponseControl{
			ControlID:  "",
			UniqueID:   true,
			DTDVersion: "3.0",
		},
		Operation: ResponseOperation{
			Authentication: ResponseAuthentication{},
			Result:         ResponseResult{},
		},
	}
}

type Response struct {
	XMLName   xml.Name          `xml:"response" json:"-"`
	Control   ResponseControl   `xml:"control"`
	Operation ResponseOperation `xml:"operation"`
}

type ResponseControl struct {
	Status     string `xml:"status"`
	SenderID   string `xml:"senderid"`
	ControlID  string `xml:"controlid"`
	UniqueID   bool   `xml:"uniqueid"`
	DTDVersion string `xml:"dtdversion"`
}

type ResponseOperation struct {
	Authentication ResponseAuthentication `xml:"authentication"`
	Result         ResponseResult         `xml:"result"`
}

type ResponseAuthentication struct {
	Status           string    `xml:"status"`
	UserID           string    `xml:"userid"`
	CompanyID        string    `xml:"companyid"`
	LocationID       string    `xml:"locationid"`
	SessionTimestamp time.Time `xml:"sessiontimestamp"`
	SessionTimeout   time.Time `xml:"sessiontimeout"`
}

type ResponseResult struct {
	Status    string `xml:"status"`
	Function  string `xml:"function"`
	ControlID string `xml:"controlid"`
	// Data      interface{} `xml:"data"`
	Data struct {
		Listtype     string `xml:"listtype,attr"`
		Count        int    `xml:"count"`
		TotalCount   int    `xml:"totalcount,attr"`
		NumRemaining int    `xml:"numremaining,attr"`
		ResultID     string `xml:"resultId,attr"`

		API struct {
			SessionID  string `xml:"sessionid"`
			Endpoint   string `xml:"endpoint"`
			LocationID string `xml:"locationid"`
		} `xml:"api"`
		GLAccounts   GLAccounts           `xml:"GLACCOUNT"`
		Classes      Classes              `xml:"class"`
		Customers    Customers            `xml:"CUSTOMER"`
		Departments  Departments          `xml:"department"`
		Type         ObjectDefinitionType `xml:"Type"`
		Locations    Locations            `xml:"location"`
		Projects     Projects             `xml:"project"`
		TaxDetails   TaxDetails           `xml:"TAXDETAIL"`
		Invoices     Invoices             `xml:"ARINVOICE"`
		InvoiceItems InvoiceItems         `xml:"ARINVOICEITEM"`
		SupDoc       SupDoc               `xml:"supdoc"`
	} `xml:"data"`
	ErrorMessage ErrorMessage `xml:"errormessage"`
}
