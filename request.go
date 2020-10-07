package intacct

import "encoding/xml"

func NewRequest() Request {
	return Request{
		Control: RequestControl{
			SenderID:          "",
			SenderPassword:    "",
			ControlID:         "",
			UniqueID:          true,
			DTDVersion:        "3.0",
			IncludeWhitespace: true,
		},
		Operation: RequestOperation{
			Authentication: RequestAuthentication{},
			Content:        nil,
		},
	}
}

type Request struct {
	XMLName   xml.Name         `xml:"request" json:"-"`
	Control   RequestControl   `xml:"control"`
	Operation RequestOperation `xml:"operation"`
}

func (r *Request) SetSessionID(sessionID string) {
	r.Operation.Authentication.SessionID = sessionID
}

type RequestControl struct {
	SenderID          string `xml:"senderid"`
	SenderPassword    string `xml:"password"`
	ControlID         string `xml:"controlid"`
	UniqueID          bool   `xml:"uniqueid"`
	DTDVersion        string `xml:"dtdversion"`
	IncludeWhitespace bool   `xml:"includewhitespace"`
}

type RequestOperation struct {
	Authentication RequestAuthentication `xml:"authentication"`
	Content        interface{}           `xml:"content"`
}

type RequestAuthentication struct {
	SessionID string        `xml:"sessionid,omitempty"`
	Login     *RequestLogin `xml:"login,omitempty"`
}

type RequestLogin struct {
	UserID    string `xml:"userid"`
	CompanyID string `xml:"companyid"`
	Password  string `xml:"password"`
}
