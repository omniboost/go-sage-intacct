package intacct

import (
	"encoding/xml"
	"fmt"
)

type ReadByQuery struct {
	Object   string `xml:"object"`
	Fields   string `xml:"fields"`
	Query    Query  `xml:"query"`
	PageSize int    `xml:"pagesize"`
}

type Query interface {
	Query() string
}

// func (q *Query) AddFilter(filter Filter) {
// }

// func (q *Query) AddCondition(condition Condition) {
// }

// type Filter struct {
// }

// type Condition struct {
// }

type NoQuery struct{}

func (r NoQuery) Query() string {
	return ""
}

type RawQuery string

func (r RawQuery) Query() string {
	return string(r)
}

type Number float64

func (n Number) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.2f", float64(n))
	return e.EncodeElement(s, start)
}

type Select []string

func (s Select) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	typ := struct {
		XMLName xml.Name `xml:"select"`
		Fields  []string `xml:"field"`
	}{Fields: s}
	return e.Encode(typ)
}

type Filters []Filter

func (f Filters) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias Filters
	typ := struct {
		XMLName xml.Name `xml:"filter"`
		Filters alias
	}{Filters: alias(f)}
	return e.Encode(typ)
}

type Filter interface {
	IsFilter() bool
}

type EqualTo struct {
	Field string `xml:"field"`
	Value string `xml:"value"`
}

func (et EqualTo) IsFilter() bool {
	return true
}

func (et EqualTo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	ee := struct {
		XMLName xml.Name `xml:"equalto"`

		Field string `xml:"field"`
		Value string `xml:"value"`
	}{Field: et.Field, Value: et.Value}
	return e.Encode(ee)
}

type Dimension struct {
	RecordNo                          string `xml:"RECORDNO"`
	Name                              string `xml:"NAME"`
	ParentID                          string `xml:"PARENTID"`
	SupervisorName                    string `xml:"SUPERVISORNAME"`
	SupervisorID                      int    `xml:"SUPERVISORID"`
	ContactinfoContactName            string `xml:"CONTACTINFO.CONTACTNAME"`
	ContactinfoPrintAs                string `xml:"CONTACTINFO.PRINTAS"`
	ContactinfoPhone1                 string `xml:"CONTACTINFO.PHONE1"`
	ContactinfoPhon2                  string `xml:"CONTACTINFO.PHONE2"`
	ContactinfoEmail1                 string `xml:"CONTACTINFO.EMAIL1"`
	ContactinfoEmail2                 string `xml:"CONTACTINFO.EMAIL2"`
	ContactinfoFax                    string `xml:"CONTACTINFO.FAX"`
	ContactinfoMailaddressAddress1    string `xml:"CONTACTINFO.MAILADDRESS.ADDRESS1"`
	ContactinfoMailaddressAddress2    string `xml:"CONTACTINFO.MAILADDRESS.ADDRESS2"`
	ContactinfoMailaddressCity        string `xml:"CONTACTINFO.MAILADDRESS.CITY"`
	ContactinfoMailaddressState       string `xml:"CONTACTINFO.MAILADDRESS.STATE"`
	CONTACTINFOMAILADDRESSZIP         string `xml:"CONTACTINFO.MAILADDRESS.ZIP"`
	CONTACTINFOMAILADDRESSCOUNTRY     string `xml:"CONTACTINFO.MAILADDRESS.COUNTRY"`
	CONTACTINFOMAILADDRESSCOUNTRYCODE string `xml:"CONTACTINFO.MAILADDRESS.COUNTRYCODE"`
	STARTDATE                         string `xml:"STARTDATE"`
	ENDDATE                           string `xml:"ENDDATE"`
	SHIPTOCONTACTNAME                 string `xml:"SHIPTO.CONTACTNAME"`
	SHIPTOPHONE1                      string `xml:"SHIPTO.PHONE1"`
	SHIPTOPHONE2                      string `xml:"SHIPTO.PHONE2"`
	SHIPTOMAILADDRESSADDRESS1         string `xml:"SHIPTO.MAILADDRESS.ADDRESS1"`
	SHIPTOMAILADDRESSADDRESS2         string `xml:"SHIPTO.MAILADDRESS.ADDRESS2"`
	SHIPTOMAILADDRESSCITY             string `xml:"SHIPTO.MAILADDRESS.CITY"`
	SHIPTOMAILADDRESSSTATE            string `xml:"SHIPTO.MAILADDRESS.STATE"`
	SHIPTOMAILADDRESSZIP              string `xml:"SHIPTO.MAILADDRESS.ZIP"`
	SHIPTOMAILADDRESSCOUNTRY          string `xml:"SHIPTO.MAILADDRESS.COUNTRY"`
	SHIPTOMAILADDRESSCOUNTRYCODE      string `xml:"SHIPTO.MAILADDRESS.COUNTRYCODE"`
	STATUS                            string `xml:"STATUS"`
	WHENCREATED                       string `xml:"WHENCREATED"`
	WHENMODIFIED                      string `xml:"WHENMODIFIED"`
	FederalID                         string `xml:"FEDERALID"`
	FirstMonth                        string `xml:"FIRSTMONTH"`
	WeekStart                         string `xml:"WEEKSTART"`
	IEPAYABLEACCOUNT                  string `xml:"IEPAYABLE.ACCOUNT"`
	IEPAYABLENUMBER                   string `xml:"IEPAYABLE.NUMBER"`
	IERECEIVABLEACCOUNT               string `xml:"IERECEIVABLE.ACCOUNT"`
	IERECEIVABLENUMBER                string `xml:"IERECEIVABLE.NUMBER"`
	MessageText                       string `xml:"MESSAGE_TEXT"`
	MarketingText                     string `xml:"MARKETING_TEXT"`
	FOOTNOTETEXT                      string `xml:"FOOTNOTETEXT"`
	REPORTPRINTAS                     string `xml:"REPORTPRINTAS"`
	IsRoot                            string `xml:"ISROOT"`
	RESERVEAMT                        string `xml:"RESERVEAMT"`
	VendorName                        string `xml:"VENDORNAME"`
	VendorID                          int    `xml:"VENDORID"`
	CustomerID                        int    `xml:"CUSTOMERID"`
	CustomerName                      string `xml:"CUSTOMERNAME"`
	Currency                          string `xml:"CURRENCY"`
	Entity                            string `xml:"ENTITY"`
	ENTITYRECORDNO                    string `xml:"ENTITYRECORDNO"`
	HASIERELATION                     string `xml:"HAS_IE_RELATION"`
	CustTitle                         string `xml:"CUSTTITLE"`
	BusinessDays                      string `xml:"BUSINESSDAYS"`
	Weekend                           string `xml:"WEEKENDS"`
	FIRSTMONTHTAX                     string `xml:"FIRSTMONTHTAX"`
	ContactKey                        string `xml:"CONTACTKEY"`
	SUPERVISORKEY                     string `xml:"SUPERVISORKEY"`
	ParentKey                         string `xml:"PARENTKEY"`
	SHIPTOKEY                         string `xml:"SHIPTOKEY"`
	IEPAYABLEACCTKEY                  string `xml:"IEPAYABLEACCTKEY"`
	IERECEIVABLEACCTKEY               string `xml:"IERECEIVABLEACCTKEY"`
	VENDENTITY                        string `xml:"VENDENTITY"`
	CUSTENTITY                        string `xml:"CUSTENTITY"`
	TaxID                             string `xml:"TAXID"`
	CreatedBy                         string `xml:"CREATEDBY"`
	ModifiedBy                        string `xml:"MODIFIEDBY"`
	ADDRESSCOUNTRYDEFAULT             string `xml:"ADDRESSCOUNTRYDEFAULT"`
}

type Classes []Class

type Class struct {
	RecordNo       string `xml:"RECORDNO"`
	ClassID        string `xml:"CLASSID"`
	Name           string `xml:"NAME"`
	Description    string `xml:"DESCRIPTION"`
	Status         string `xml:"STATUS"`
	ParentKey      string `xml:"PARENTKEY"`
	ParentID       string `xml:"PARENTID"`
	ParentName     string `xml:"PARENTNAME"`
	WhenCreated    string `xml:"WHENCREATED"`
	WhenModified   string `xml:"WHENMODIFIED"`
	CreatedBy      string `xml:"CREATEDBY"`
	ModifiedBy     string `xml:"MODIFIEDBY"`
	MegaEntityKey  string `xml:"MEGAENTITYKEY"`
	MegaEntityID   string `xml:"MEGAENTITYID"`
	MegaEntityName string `xml:"MEGAENTITYNAME"`
}

// func (c Class) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
// 	return e.EncodeElement(c.Dimension, start)
// }

// func (c Class) MarshalJSON() ([]byte, error) {
// 	m1 := structs.Map(c)
// 	m2 := structs.Map(c.Dimension)
// 	delete(m1, "Dimension")

// 	for k, v := range m2 {
// 		m1[k] = v
// 	}

// 	return json.Marshal(m1)
// }

type Projects []Project

type Project struct {
	ProjectID string `xml:"projectID"`
	Dimension
}
