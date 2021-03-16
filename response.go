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
	Status    string      `xml:"status"`
	Function  string      `xml:"function"`
	ControlID string      `xml:"controlid"`
	Data      interface{} `xml:"data"`
	// Data      struct {
	// 	Listtype string `xml:"listtype,attr"`
	// 	Count    int    `xml:"count"`
	// } `xml:"data"`
	ErrorMessage ErrorMessage `xml:"errormessage"`
}
