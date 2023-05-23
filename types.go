package intacct

import (
	"encoding/xml"
	"fmt"

	"github.com/cydev/zero"
	"github.com/omniboost/go-sage-intacct/omitempty"
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

type Like struct {
	Field string `xml:"field"`
	Value string `xml:"value"`
}

func (l Like) IsFilter() bool {
	return true
}

func (l Like) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	alias := struct {
		XMLName xml.Name `xml:"like"`

		Field string `xml:"field"`
		Value string `xml:"value"`
	}{Field: l.Field, Value: l.Value}
	return e.Encode(alias)
}

type GreaterThan struct {
	Field string `xml:"field"`
	Value string `xml:"value"`
}

func (gt GreaterThan) IsFilter() bool {
	return true
}

func (gt GreaterThan) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	alias := struct {
		XMLName xml.Name `xml:"greaterthan"`

		Field string `xml:"field"`
		Value string `xml:"value"`
	}{Field: gt.Field, Value: gt.Value}
	return e.Encode(alias)
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
	RecordNo               string `xml:"RECORDNO"`
	ProjectID              string `xml:"PROJECTID"`
	Name                   string `xml:"NAME"`
	Description            string `xml:"DESCRIPTION"`
	Currency               string `xml:"CURRENCY"`
	ProjectCategory        string `xml:"PROJECTCATEGORY"`
	ProjectStatusKey       string `xml:"PROJECTSTATUSKEY"`
	ProjectStatus          string `xml:"PROJECTSTATUS"`
	PreventTimesheet       string `xml:"PREVENTTIMESHEET"`
	PreventExpense         string `xml:"PREVENTEXPENSE"`
	PreventAppo            string `xml:"PREVENTAPPO"`
	PreventGenInvoice      string `xml:"PREVENTGENINVOICE"`
	Status                 string `xml:"STATUS"`
	BeginDate              string `xml:"BEGINDATE"`
	EndDate                string `xml:"ENDDATE"`
	BudgetAmount           string `xml:"BUDGETAMOUNT"`
	contractAmount         string `xml:"CONTRACTAMOUNT"`
	ActualAmount           string `xml:"ACTUALAMOUNT"`
	BudgetQty              string `xml:"BUDGETQTY"`
	EstQty                 string `xml:"ESTQTY"`
	ActualQty              string `xml:"ACTUALQTY"`
	ApprovedQty            string `xml:"APPROVEDQTY"`
	RemainingQty           string `xml:"REMAININGQTY"`
	PercentComplete        string `xml:"PERCENTCOMPLETE"`
	ObsPercentComplete     string `xml:"OBSPERCENTCOMPLETE"`
	BillingType            string `xml:"BILLINGTYPE"`
	SONumber               string `xml:"SONUMBER"`
	PONumber               string `xml:"PONUMBER"`
	POAmount               string `xml:"POAMOUNT"`
	PQNumber               string `xml:"PQNUMBER"`
	SFDCKey                string `xml:"SFDCKEY"`
	QarrowKey              string `xml:"QARROWKEY"`
	OAKey                  string `xml:"OAKEY"`
	ParentKey              string `xml:"PARENTKEY"`
	ParentID               string `xml:"PARENTID"`
	ParentName             string `xml:"PARENTNAME"`
	InvoiceWithParent      string `xml:"INVOICEWITHPARENT"`
	CustomerKey            string `xml:"CUSTOMERKEY"`
	CustomerID             string `xml:"CUSTOMERID"`
	CustomerName           string `xml:"CUSTOMERNAME"`
	SalesContactKey        string `xml:"SALESCONTACTKEY"`
	SalesContactID         string `xml:"SALESCONTACTID"`
	SalesContactName       string `xml:"SALESCONTACTNAME"`
	ProjectTypeKey         string `xml:"PROJECTTYPEKEY"`
	ProjectType            string `xml:"PROJECTTYPE"`
	ManagerKey             string `xml:"MANAGERKEY"`
	ManagerID              string `xml:"MANAGERID"`
	ManagerContactName     string `xml:"MANAGERCONTACTNAME"`
	ProjectDeptKey         string `xml:"PROJECTDEPTKEY"`
	DepartmentID           string `xml:"DEPARTMENTID"`
	DepartmentName         string `xml:"DEPARTMENTNAME"`
	ProjectLocationKey     string `xml:"PROJECTLOCATIONKEY"`
	LocationID             string `xml:"LOCATIONID"`
	LocationName           string `xml:"LOCATIONNAME"`
	ContactInfoContactName string `xml:"CONTACTINFO.CONTACTNAME"`
	ShipToContactName      string `xml:"SHIPTO.CONTACTNAME"`
	BillToContactName      string `xml:"BILLTO.CONTACTNAME"`
	TermsKey               string `xml:"TERMSKEY"`
	TermName               string `xml:"TERMNAME"`
	DocNumber              string `xml:"DOCNUMBER"`
	CustUserKey            string `xml:"CUSTUSERKEY"`
	CustUserID             string `xml:"CUSTUSERID"`
	WhenCreated            string `xml:"WHENCREATED"`
	WhenModified           string `xml:"WHENMODIFIED"`
	CreatedBy              string `xml:"CREATEDBY"`
	ModifiedBy             string `xml:"MODIFIEDBY"`
	BudgetCost             string `xml:"BUDGETEDCOST"`
	classID                string `xml:"CLASSID"`
	ClassName              string `xml:"CLASSNAME"`
	ClassKey               string `xml:"CLASSKEY"`
	UserRestrictions       string `xml:"USERRESTRICTIONS"`
	BillableExpDefault     string `xml:"BILLABLEEXPDEFAULT"`
	BillablePODefault      string `xml:"BILLABLEAPPODEFAULT"`
	BudgetID               string `xml:"BUDGETID"`
	BudgetKey              string `xml:"BUDGETKEY"`
	BillingRate            string `xml:"BILLINGRATE"`
	BillingPrice           string `xml:"BILLINGPRICING"`
	ExpenseRate            string `xml:"EXPENSERATE"`
	ExpensePricing         string `xml:"EXPENSEPRICING"`
	POAPRate               string `xml:"POAPRATE"`
	POAPPricing            string `xml:"POAPPRICING"`
	ContactKey             string `xml:"CONTACTKEY"`
	ShipTokey              string `xml:"SHIPTOKEY"`
	BillTokey              string `xml:"BILLTOKEY"`
	InvoiceMessage         string `xml:"INVOICEMESSAGE"`
	InvoiceCurrency        string `xml:"INVOICECURRENCY"`
	BillingOverMax         string `xml:"BILLINGOVERMAX"`
	ExcludeExpenses        string `xml:"EXCLUDEEXPENSES"`
	ContractKey            string `xml:"CONTRACTKEY"`
	ContractID             string `xml:"CONTRACTID"`
	RootParentKey          string `xml:"ROOTPARENTKEY"`
	RootParentID           string `xml:"ROOTPARENTID"`
	RootParentName         string `xml:"ROOTPARENTNAME"`
	MegaEntityKey          string `xml:"MEGAENTITYKEY"`
	MegaEntityID           string `xml:"MEGAENTITYID"`
	MegaEntityName         string `xml:"MEGAENTITYNAME"`
	RCIP                   string `xml:"RCIP"`
	CIPProject             string `xml:"CIP_PROJECT"`
}

type TaxDetails []TaxDetail

type TaxDetail struct {
	// Text           string `xml:",chardata"`
	RECORDNO       string `xml:"RECORDNO"`
	DETAILID       string `xml:"DETAILID"`
	TAXUID         string `xml:"TAXUID"`
	DESCRIPTION    string `xml:"DESCRIPTION"`
	TAXTYPE        string `xml:"TAXTYPE"`
	VALUE          string `xml:"VALUE"`
	MINTAXABLE     string `xml:"MINTAXABLE"`
	MAXTAXABLE     string `xml:"MAXTAXABLE"`
	INCLUDE        string `xml:"INCLUDE"`
	MINTAX         string `xml:"MINTAX"`
	MAXTAX         string `xml:"MAXTAX"`
	GLACCOUNT      string `xml:"GLACCOUNT"`
	TAXAUTHORITY   string `xml:"TAXAUTHORITY"`
	STATUS         string `xml:"STATUS"`
	SYSGENERATED   string `xml:"SYSGENERATED"`
	REVERSECHARGE  string `xml:"REVERSECHARGE"`
	TAXRATE        string `xml:"TAXRATE"`
	TAXSOLUTIONID  string `xml:"TAXSOLUTIONID"`
	USEEXPENSEACCT string `xml:"USEEXPENSEACCT"`
	MEGAENTITYKEY  string `xml:"MEGAENTITYKEY"`
	MEGAENTITYID   string `xml:"MEGAENTITYID"`
	MEGAENTITYNAME string `xml:"MEGAENTITYNAME"`
	RECORDURL      string `xml:"RECORD_URL"`
}

type GLAccounts []GLAccount

type GLAccount struct {
	RecordNo            int         `xml:"RECORDNO"`
	AccountNo           string      `xml:"ACCOUNTNO"`
	Title               string      `xml:"TITLE"`
	AccountType         string      `xml:"ACCOUNTTYPE"`
	NormalBalance       string      `xml:"NORMALBALANCE"`
	ClosingType         string      `xml:"CLOSINGTYPE"`
	ClosingAccountNo    string      `xml:"CLOSINGACCOUNTNO"`
	ClosingAccountTitle string      `xml:"CLOSINGACCOUNTTITLE"`
	Status              string      `xml:"STATUS"`
	RequireDept         bool        `xml:"REQUIREDEPT"`
	RequireLoc          bool        `xml:"REQUIRELOC"`
	Taxable             bool        `xml:"TAXABLE"`
	CategoryKey         string      `xml:"CATEGORYKEY"`
	Category            string      `xml:"CATEGORY"`
	TaxCode             string      `xml:"TAXCODE"`
	MRCCode             string      `xml:"MRCCODE"`
	CloseToAcctKey      string      `xml:"CLOSETOACCTKEY"`
	AlternativeAccount  string      `xml:"ALTERNATIVEACCOUNT"`
	WhenCreated         string      `xml:"WHENCREATED"`
	WhenModified        string      `xml:"WHENMODIFIED"`
	CreatedBy           int         `xml:"CREATEDBY"`
	ModifiedBy          int         `xml:"MODIFIEDBY"`
	SubledgerControlOn  bool        `xml:"SUBLEDGERCONTROLON"`
	MegaEntityKey       interface{} `xml:"MEGAENTITYKEY"`
	MegaEntityID        interface{} `xml:"MEGAENTITYID"`
	MegaEntityName      interface{} `xml:"MEGAENTITYNAME"`
	RequireProject      bool        `xml:"REQUIREPROJECT"`
	RequireCustomer     bool        `xml:"REQUIRECUSTOMER"`
	RequireVendor       bool        `xml:"REQUIREVENDOR"`
	RequireClass        bool        `xml:"REQUIRECLASS"`
}

type Departments []Department

type Department struct {
	DepartmentID   int    `xml:"DEPARTMENTID"`
	RecordNo       int    `xml:"RECORDNO"`
	Title          string `xml:"TITLE"`
	ParentKey      string `xml:"PARENTKEY"`
	ParentID       string `xml:"PARENTID"`
	SupervisorKey  string `xml:"SUPERVISORKEY"`
	SupervisorID   string `xml:"SUPERVISORID"`
	WhenCreated    string `xml:"WHENCREATED"`
	WhenModified   string `xml:"WHENMODIFIED"`
	SupervisorName string `xml:"SUPERVISORNAME"`
	Status         string `xml:"STATUS"`
	CustTitle      string `xml:"CUSTTITLE"`
	CreatedBy      int    `xml:"CREATEDBY"`
	ModifiedBy     int    `xml:"MODIFIEDBY"`
}

type ObjectDefinitionType struct {
	Name         string `xml:"Name,attr"`
	DocumentType string `xml:"DocumentType,attr"`
	Fields       struct {
		Field []struct {
			ID          string `xml:"ID"`
			LABEL       string `xml:"LABEL"`
			DESCRIPTION string `xml:"DESCRIPTION"`
			REQUIRED    string `xml:"REQUIRED"`
			READONLY    string `xml:"READONLY"`
			DATATYPE    string `xml:"DATATYPE"`
			ISCUSTOM    string `xml:"ISCUSTOM"`
			VALIDVALUES struct {
				VALIDVALUE []string `xml:"VALIDVALUE"`
			} `xml:"VALIDVALUES"`
		} `xml:"Field"`
	} `xml:"Fields"`
	Relationships struct {
		Relationship []struct {
			OBJECTPATH       string `xml:"OBJECTPATH"`
			OBJECTNAME       string `xml:"OBJECTNAME"`
			LABEL            string `xml:"LABEL"`
			RELATIONSHIPTYPE string `xml:"RELATIONSHIPTYPE"`
			RELATEDBY        string `xml:"RELATEDBY"`
		} `xml:"Relationship"`
	} `xml:"Relationships"`
}

type Locations []Location

type Location struct {
	LocationID                        string `xml:"LOCATIONID"`
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

type Invoices []Invoice

type Invoice struct {
	RecordNo                      string       `xml:"RECORDNO,omitempty"`
	RecordType                    string       `xml:"RECORDTYPE,omitempty"`
	RecordID                      string       `xml:"RECORDID,omitempty"`
	ContactTaxGroup               string       `xml:"CONTACTTAXGROUP,omitempty"`
	State                         string       `xml:"STATE,omitempty"`
	RawStState                    string       `xml:"RAWSTATE,omitempty"`
	CustomerID                    string       `xml:"CUSTOMERID,omitempty"`
	CustomerName                  string       `xml:"CUSTOMERNAME,omitempty"`
	CustEmailOption               string       `xml:"CUSTEMAILOPTIN,omitempty"`
	TrxEntityDue                  string       `xml:"TRX_ENTITYDUE,omitempty"`
	CustMessageID                 string       `xml:"CUSTMESSAGEID,omitempty"`
	CustMessageMessage            string       `xml:"CUSTMESSAGE.MESSAGE,omitempty"`
	DeliveryOptions               string       `xml:"DELIVERY_OPTIONS,omitempty"`
	DocNumber                     string       `xml:"DOCNUMBER,omitempty"`
	Description                   string       `xml:"DESCRIPTION,omitempty"`
	Description2                  string       `xml:"DESCRIPTION2,omitempty"`
	TermName                      string       `xml:"TERMNAME,omitempty"`
	TermKey                       string       `xml:"TERMKEY,omitempty"`
	TermValue                     string       `xml:"TERMVALUE,omitempty"`
	WhenCreated                   string       `xml:"WHENCREATED,omitempty"`
	WhenPosted                    string       `xml:"WHENPOSTED,omitempty"`
	WhenDiscount                  string       `xml:"WHENDISCOUNT,omitempty"`
	WhenDue                       string       `xml:"WHENDUE,omitempty"`
	WhenPaid                      string       `xml:"WHENPAID,omitempty"`
	BaseCurr                      string       `xml:"BASECURR,omitempty"`
	Currency                      string       `xml:"CURRENCY,omitempty"`
	ExchRateDate                  string       `xml:"EXCH_RATE_DATE,omitempty"`
	ExchRateTypeID                string       `xml:"EXCH_RATE_TYPE_ID,omitempty"`
	ExchangeRate                  string       `xml:"EXCHANGE_RATE,omitempty"`
	TotalEntered                  string       `xml:"TOTALENTERED,omitempty"`
	TotalSelected                 string       `xml:"TOTALSELECTED,omitempty"`
	TotalPaid                     string       `xml:"TOTALPAID,omitempty"`
	TotalDue                      string       `xml:"TOTALDUE,omitempty"`
	TrxTotalEntered               string       `xml:"TRX_TOTALENTERED,omitempty"`
	TrxTotalSelected              string       `xml:"TRX_TOTALSELECTED,omitempty"`
	TrxTotalPaid                  string       `xml:"TRX_TOTALPAID,omitempty"`
	TrxTotalDue                   string       `xml:"TRX_TOTALDUE,omitempty"`
	BillToPayToContactName        string       `xml:"BILLTOPAYTOCONTACTNAME,omitempty"`
	ShipToReturnToContactName     string       `xml:"SHIPTORETURNTOCONTACTNAME,omitempty"`
	BillToPayToKey                string       `xml:"BILLTOPAYTOKEY,omitempty"`
	ShopToReturnToKey             string       `xml:"SHIPTORETURNTOKEY,omitempty"`
	ContactContactName            string       `xml:"CONTACT.CONTACTNAME,omitempty"`
	ContactPrefix                 string       `xml:"CONTACT.PREFIX,omitempty"`
	ContactFirstName              string       `xml:"CONTACT.FIRSTNAME,omitempty"`
	ContactInitial                string       `xml:"CONTACT.INITIAL,omitempty"`
	ContactLastName               string       `xml:"CONTACT.LASTNAME,omitempty"`
	ContactCompanyName            string       `xml:"CONTACT.COMPANYNAME,omitempty"`
	ContactPrintAs                string       `xml:"CONTACT.PRINTAS,omitempty"`
	ContactPhone1                 string       `xml:"CONTACT.PHONE1,omitempty"`
	ContactPhone2                 string       `xml:"CONTACT.PHONE2,omitempty"`
	ContactCellphone              string       `xml:"CONTACT.CELLPHONE,omitempty"`
	ContactPager                  string       `xml:"CONTACT.PAGER,omitempty"`
	ContactFax                    string       `xml:"CONTACT.FAX,omitempty"`
	ContactEmail1                 string       `xml:"CONTACT.EMAIL1,omitempty"`
	ContactEmail2                 string       `xml:"CONTACT.EMAIL2,omitempty"`
	ContactURL1                   string       `xml:"CONTACT.URL1,omitempty"`
	ContactURL2                   string       `xml:"CONTACT.URL2,omitempty"`
	ContactVissible               string       `xml:"CONTACT.VISIBLE,omitempty"`
	ContactMailAddressAddress1    string       `xml:"CONTACT.MAILADDRESS.ADDRESS1,omitempty"`
	ContactMailAddressAddress2    string       `xml:"CONTACT.MAILADDRESS.ADDRESS2,omitempty"`
	contactMailAddresCity         string       `xml:"CONTACT.MAILADDRESS.CITY,omitempty"`
	ContactMailAddressState       string       `xml:"CONTACT.MAILADDRESS.STATE,omitempty"`
	ContactMailAddressZip         string       `xml:"CONTACT.MAILADDRESS.ZIP,omitempty"`
	ContactMailAddressCountry     string       `xml:"CONTACT.MAILADDRESS.COUNTRY,omitempty"`
	ContactMailAddressCountryCode string       `xml:"CONTACT.MAILADDRESS.COUNTRYCODE,omitempty"`
	ShipToContactName             string       `xml:"SHIPTO.CONTACTNAME,omitempty"`
	ShipToPrefix                  string       `xml:"SHIPTO.PREFIX,omitempty"`
	ShipToFirstName               string       `xml:"SHIPTO.FIRSTNAME,omitempty"`
	ShipToInitial                 string       `xml:"SHIPTO.INITIAL,omitempty"`
	ShipToLastName                string       `xml:"SHIPTO.LASTNAME,omitempty"`
	ShipToCompanyName             string       `xml:"SHIPTO.COMPANYNAME,omitempty"`
	ShipToPrintAs                 string       `xml:"SHIPTO.PRINTAS,omitempty"`
	ShipToPhone1                  string       `xml:"SHIPTO.PHONE1,omitempty"`
	ShipToPhone2                  string       `xml:"SHIPTO.PHONE2,omitempty"`
	ShipToCellphone               string       `xml:"SHIPTO.CELLPHONE,omitempty"`
	ShipToPager                   string       `xml:"SHIPTO.PAGER,omitempty"`
	ShipToFax                     string       `xml:"SHIPTO.FAX,omitempty"`
	ShipToEmail1                  string       `xml:"SHIPTO.EMAIL1,omitempty"`
	ShipToEmail2                  string       `xml:"SHIPTO.EMAIL2,omitempty"`
	ShipToURL1                    string       `xml:"SHIPTO.URL1,omitempty"`
	ShipToURL2                    string       `xml:"SHIPTO.URL2,omitempty"`
	ShipToVisible                 string       `xml:"SHIPTO.VISIBLE,omitempty"`
	ShipToMailAddressAddress1     string       `xml:"SHIPTO.MAILADDRESS.ADDRESS1,omitempty"`
	ShipToMailAddressAddress2     string       `xml:"SHIPTO.MAILADDRESS.ADDRESS2,omitempty"`
	ShipToMailAddressCity         string       `xml:"SHIPTO.MAILADDRESS.CITY,omitempty"`
	ShipToMailAddressState        string       `xml:"SHIPTO.MAILADDRESS.STATE,omitempty"`
	ShipToMailAddressZip          string       `xml:"SHIPTO.MAILADDRESS.ZIP,omitempty"`
	ShipToMailAddressCountry      string       `xml:"SHIPTO.MAILADDRESS.COUNTRY,omitempty"`
	ShipToMailAddressCountryCode  string       `xml:"SHIPTO.MAILADDRESS.COUNTRYCODE,omitempty"`
	BillToContactName             string       `xml:"BILLTO.CONTACTNAME,omitempty"`
	BillToPrefix                  string       `xml:"BILLTO.PREFIX,omitempty"`
	BillToFirstName               string       `xml:"BILLTO.FIRSTNAME,omitempty"`
	BillToInitial                 string       `xml:"BILLTO.INITIAL,omitempty"`
	BillToLastName                string       `xml:"BILLTO.LASTNAME,omitempty"`
	BillToCompanyName             string       `xml:"BILLTO.COMPANYNAME,omitempty"`
	BillToPrintAs                 string       `xml:"BILLTO.PRINTAS,omitempty"`
	BillToPhone1                  string       `xml:"BILLTO.PHONE1,omitempty"`
	BillToPhone2                  string       `xml:"BILLTO.PHONE2,omitempty"`
	BillToCellphone               string       `xml:"BILLTO.CELLPHONE,omitempty"`
	BillToPager                   string       `xml:"BILLTO.PAGER,omitempty"`
	BillToFax                     string       `xml:"BILLTO.FAX,omitempty"`
	BillToEmail1                  string       `xml:"BILLTO.EMAIL1,omitempty"`
	BillToEmail2                  string       `xml:"BILLTO.EMAIL2,omitempty"`
	BillToURL1                    string       `xml:"BILLTO.URL1,omitempty"`
	BillToURL2                    string       `xml:"BILLTO.URL2,omitempty"`
	BillToVisible                 string       `xml:"BILLTO.VISIBLE,omitempty"`
	BillToMailAddressAddress1     string       `xml:"BILLTO.MAILADDRESS.ADDRESS1,omitempty"`
	BillToMailAddressAddress2     string       `xml:"BILLTO.MAILADDRESS.ADDRESS2,omitempty"`
	BillToMailAddressCity         string       `xml:"BILLTO.MAILADDRESS.CITY,omitempty"`
	BillToMailAddressState        string       `xml:"BILLTO.MAILADDRESS.STATE,omitempty"`
	BillToMailAddressZip          string       `xml:"BILLTO.MAILADDRESS.ZIP,omitempty"`
	BillToMailAddressCountry      string       `xml:"BILLTO.MAILADDRESS.COUNTRY,omitempty"`
	BillToMailAddressCountryCode  string       `xml:"BILLTO.MAILADDRESS.COUNTRYCODE,omitempty"`
	PRBatch                       string       `xml:"PRBATCH,omitempty"`
	PRBatchKey                    string       `xml:"PRBATCHKEY,omitempty"`
	ModuleKey                     string       `xml:"MODULEKEY,omitempty"`
	SchopKye                      string       `xml:"SCHOPKEY,omitempty"`
	SystemGenerated               string       `xml:"SYSTEMGENERATED,omitempty"`
	HasPostedRevRec               string       `xml:"HASPOSTEDREVREC,omitempty"`
	BillBackTemplateKey           string       `xml:"BILLBACKTEMPLATEKEY,omitempty"`
	AUWhenCreated                 string       `xml:"AUWHENCREATED,omitempty"`
	WhenModified                  string       `xml:"WHENMODIFIED,omitempty"`
	CreatedBy                     string       `xml:"CREATEDBY,omitempty"`
	ModifiedBy                    string       `xml:"MODIFIEDBY,omitempty"`
	DueInDays                     string       `xml:"DUE_IN_DAYS,omitempty"`
	ShopToTaxGroupName            string       `xml:"SHIPTO.TAXGROUP.NAME,omitempty"`
	ShipToTaxGroupRecordNo        string       `xml:"SHIPTO.TAXGROUP.RECORDNO,omitempty"`
	ShipToTaxID                   string       `xml:"SHIPTO.TAXID,omitempty"`
	TaxSolutionID                 string       `xml:"TAXSOLUTIONID,omitempty"`
	RetainagePercentage           string       `xml:"RETAINAGEPERCENTAGE,omitempty"`
	TrxTotalRetained              string       `xml:"TRX_TOTALRETAINED,omitempty"`
	TrxTotalReleased              string       `xml:"TRX_TOTALRELEASED,omitempty"`
	TotalRetained                 string       `xml:"TOTALRETAINED,omitempty"`
	SupDocID                      string       `xml:"SUPDOCID,omitempty"`
	ProjectContractKey            string       `xml:"PROJECTCONTRACTKEY,omitempty"`
	ProjectContractID             string       `xml:"PROJECTCONTRACTID,omitempty"`
	DunningCount                  string       `xml:"DUNNINGCOUNT,omitempty"`
	MegaEntityKey                 string       `xml:"MEGAENTITYKEY,omitempty"`
	MegaEntityID                  string       `xml:"MEGAENTITYID,omitempty"`
	MegaEntityName                string       `xml:"MEGAENTITYNAME,omitempty"`
	RecordURL                     string       `xml:"RECORD_URL,omitempty"`
	InvoiceItems                  InvoiceItems `xml:"INVOICEITEM>LINEITEM"`
}

type InvoiceItems []InvoiceItem

type InvoiceItem struct {
	RecordNo                   string `xml:"RECORDNO,omitempty"`
	RecordKey                  string `xml:"RECORDKEY,omitempty"`
	AccountKey                 string `xml:"ACCOUNTKEY,omitempty"`
	AccountNo                  string `xml:"ACCOUNTNO,omitempty"`
	OffsetAccountKey           string `xml:"OFFSETACCOUNTKEY,omitempty"`
	OffsetGLAccountNo          string `xml:"OFFSETGLACCOUNTNO,omitempty"`
	OffsetGLAccountTitle       string `xml:"OFFSETGLACCOUNTTITLE,omitempty"`
	AccountTitle               string `xml:"ACCOUNTTITLE,omitempty"`
	AccountLabelKey            string `xml:"ACCOUNTLABELKEY,omitempty"`
	AccountLabel               string `xml:"ACCOUNTLABEL,omitempty"`
	EntryDate                  string `xml:"ENTRY_DATE,omitempty"`
	Amount                     Number `xml:"AMOUNT,omitempty"`
	TrxAmount                  Number `xml:"TRX_AMOUNT,omitempty"`
	DepartmentID               string `xml:"DEPARTMENTID,omitempty"`
	DepartmentName             string `xml:"DEPARTMENTNAME,omitempty"`
	LocationID                 string `xml:"LOCATIONID,omitempty"`
	LocationName               string `xml:"LOCATIONNAME,omitempty"`
	EntryDescription           string `xml:"ENTRYDESCRIPTION,omitempty"`
	ExchRateDate               string `xml:"EXCH_RATE_DATE,omitempty"`
	ExchRateTypeID             string `xml:"EXCH_RATE_TYPE_ID,omitempty"`
	ExchangeRate               string `xml:"EXCHANGE_RATE,omitempty"`
	AllocationKey              string `xml:"ALLOCATIONKEY,omitempty"`
	Allocation                 string `xml:"ALLOCATION,omitempty"`
	LineItem                   string `xml:"LINEITEM,omitempty"`
	LineNo                     string `xml:"LINE_NO,omitempty"`
	Currency                   string `xml:"CURRENCY,omitempty"`
	BaseCurr                   string `xml:"BASECURR,omitempty"`
	TotalPaid                  string `xml:"TOTALPAID,omitempty"`
	TrxTotalPaid               string `xml:"TRX_TOTALPAID,omitempty"`
	TotalSelected              string `xml:"TOTALSELECTED,omitempty"`
	TrxTotalSelected           string `xml:"TRX_TOTALSELECTED,omitempty"`
	Subtotal                   string `xml:"SUBTOTAL,omitempty"`
	ParentEntry                string `xml:"PARENTENTRY,omitempty"`
	DeferredRevenue            string `xml:"DEFERREVENUE,omitempty"`
	RevRecTemplateKey          string `xml:"REVRECTEMPLATEKEY,omitempty"`
	RevRecTemplate             string `xml:"REVRECTEMPLATE,omitempty"`
	DeferredRevAcctKey         string `xml:"DEFERREDREVACCTKEY,omitempty"`
	DeferredRevActNo           string `xml:"DEFERREDREVACCTNO,omitempty"`
	DeferredRevActTitle        string `xml:"DEFERREDREVACCTTITLE,omitempty"`
	RevRecStartDate            string `xml:"REVRECSTARTDATE,omitempty"`
	RevRecEndDate              string `xml:"REVRECENDDATE,omitempty"`
	BaseLocation               string `xml:"BASELOCATION,omitempty"`
	State                      string `xml:"STATE,omitempty"`
	RecordType                 string `xml:"RECORDTYPE,omitempty"`
	DetailKey                  string `xml:"DETAILKEY,omitempty"`
	WhenCreated                string `xml:"WHENCREATED,omitempty"`
	WhenModified               string `xml:"WHENMODIFIED,omitempty"`
	CreatedBy                  string `xml:"CREATEDBY,omitempty"`
	ModifiedBy                 string `xml:"MODIFIEDBY,omitempty"`
	RetainagePercentage        string `xml:"RETAINAGEPERCENTAGE,omitempty"`
	TrxAmountRetained          string `xml:"TRX_AMOUNTRETAINED,omitempty"`
	AmountRetained             string `xml:"AMOUNTRETAINED,omitempty"`
	TrxAmountReleased          string `xml:"TRX_AMOUNTRELEASED,omitempty"`
	RetainageOffsetGLAccountNo string `xml:"RETAINAGE_OFFSETGLACCOUNTNO,omitempty"`
	ProjectContractID          string `xml:"PROJECTCONTRACTID,omitempty"`
	ProjectContractKey         string `xml:"PROJECTCONTRACTKEY,omitempty"`
	ProjectContractLineID      string `xml:"PROJECTCONTRACTLINEID,omitempty"`
	ProjectContractLineKey     string `xml:"PROJECTCONTRACTLINEKEY,omitempty"`
	ProjectDimKey              string `xml:"PROJECTDIMKEY,omitempty"`
	ProjectID                  string `xml:"PROJECTID,omitempty"`
	ProjectName                string `xml:"PROJECTNAME,omitempty"`
	CustomerDimKey             string `xml:"CUSTOMERDIMKEY,omitempty"`
	CustomerID                 string `xml:"CUSTOMERID,omitempty"`
	CustomerName               string `xml:"CUSTOMERNAME,omitempty"`
	VendorDimKey               string `xml:"VENDORDIMKEY,omitempty"`
	VendorID                   string `xml:"VENDORID,omitempty"`
	VendorName                 string `xml:"VENDORNAME,omitempty"`
	EmployeeDimKey             string `xml:"EMPLOYEEDIMKEY,omitempty"`
	EmployeeID                 string `xml:"EMPLOYEEID,omitempty"`
	EmployeeName               string `xml:"EMPLOYEENAME,omitempty"`
	ItemDimKey                 string `xml:"ITEMDIMKEY,omitempty"`
	ItemID                     string `xml:"ITEMID,omitempty"`
	ItemName                   string `xml:"ITEMNAME,omitempty"`
	ClassDimKey                string `xml:"CLASSDIMKEY,omitempty"`
	ClassID                    string `xml:"CLASSID,omitempty"`
	ClassName                  string `xml:"CLASSNAME,omitempty"`
	RecordURL                  string `xml:"RECORD_URL,omitempty"`
}

type Customers []Customer

type Customer struct {
	RecordID                             string  `xml:"RECORDNO,omitempty"`
	CustomerID                           string  `xml:"CUSTOMERID,omitempty"`
	Name                                 string  `xml:"NAME,omitempty"`
	Entity                               string  `xml:"ENTITY,omitempty"`
	ParentKey                            string  `xml:"PARENTKEY,omitempty"`
	ParentID                             string  `xml:"PARENTID,omitempty"`
	ParentName                           string  `xml:"PARENTNAME,omitempty"`
	DisplayContactContactName            string  `xml:"DISPLAYCONTACT.CONTACTNAME,omitempty"`
	DisplayContactCompanyName            string  `xml:"DISPLAYCONTACT.COMPANYNAME,omitempty"`
	DisplayContactPrefix                 string  `xml:"DISPLAYCONTACT.PREFIX,omitempty"`
	DisplayContactFirstName              string  `xml:"DISPLAYCONTACT.FIRSTNAME,omitempty"`
	DisplayContactLastName               string  `xml:"DISPLAYCONTACT.LASTNAME,omitempty"`
	DisplayContactInitial                string  `xml:"DISPLAYCONTACT.INITIAL,omitempty"`
	DisplayContactPrintAs                string  `xml:"DISPLAYCONTACT.PRINTAS,omitempty"`
	DisplayContactTaxable                string  `xml:"DISPLAYCONTACT.TAXABLE,omitempty"`
	DisplayContactTaxGroup               string  `xml:"DISPLAYCONTACT.TAXGROUP,omitempty"`
	DisplayContactTaxSolutionKey         string  `xml:"DISPLAYCONTACT.TAXSOLUTIONKEY,omitempty"`
	DisplayContactTaxSolutionID          string  `xml:"DISPLAYCONTACT.TAXSOLUTIONID,omitempty"`
	DisplayContactTaxSchedule            string  `xml:"DISPLAYCONTACT.TAXSCHEDULE,omitempty"`
	DisplayContactTaxID                  string  `xml:"DISPLAYCONTACT.TAXID,omitempty"`
	DisplayContactPhone1                 string  `xml:"DISPLAYCONTACT.PHONE1,omitempty"`
	DisplayContactPhone2                 string  `xml:"DISPLAYCONTACT.PHONE2,omitempty"`
	DisplayContactCellphone              string  `xml:"DISPLAYCONTACT.CELLPHONE,omitempty"`
	DisplayContactPager                  string  `xml:"DISPLAYCONTACT.PAGER,omitempty"`
	DisplayContactFax                    string  `xml:"DISPLAYCONTACT.FAX,omitempty"`
	DisplayContactTaxIDValidationDate    string  `xml:"DISPLAYCONTACT.TAXIDVALIDATIONDATE,omitempty"`
	DisplayContactGSTRegistered          string  `xml:"DISPLAYCONTACT.GSTREGISTERED,omitempty"`
	DisplayContactTaxCompanyName         string  `xml:"DISPLAYCONTACT.TAXCOMPANYNAME,omitempty"`
	DisplayContactTaxAddress             string  `xml:"DISPLAYCONTACT.TAXADDRESS,omitempty"`
	DisplayContactEmail1                 string  `xml:"DISPLAYCONTACT.EMAIL1,omitempty"`
	DisplayContactEmail2                 string  `xml:"DISPLAYCONTACT.EMAIL2,omitempty"`
	DisplayContactURL1                   string  `xml:"DISPLAYCONTACT.URL1,omitempty"`
	DisplayContactURL2                   string  `xml:"DISPLAYCONTACT.URL2,omitempty"`
	DisplayContactVisible                string  `xml:"DISPLAYCONTACT.VISIBLE,omitempty"`
	DisplayContactMailAddressAddress1    string  `xml:"DISPLAYCONTACT.MAILADDRESS.ADDRESS1,omitempty"`
	DisplayContactMailAddressAddress2    string  `xml:"DISPLAYCONTACT.MAILADDRESS.ADDRESS2,omitempty"`
	DisplayContactMailAddresscity        string  `xml:"DISPLAYCONTACT.MAILADDRESS.CITY,omitempty"`
	DisplayContactMailAddressState       string  `xml:"DISPLAYCONTACT.MAILADDRESS.STATE,omitempty"`
	DisplayContactMailAddressZip         string  `xml:"DISPLAYCONTACT.MAILADDRESS.ZIP,omitempty"`
	DisplayContactMailAddressCountry     string  `xml:"DISPLAYCONTACT.MAILADDRESS.COUNTRY,omitempty"`
	DisplayContactMailAddressCountryCode string  `xml:"DISPLAYCONTACT.MAILADDRESS.COUNTRYCODE,omitempty"`
	DisplayContactMailAddressLatitude    string  `xml:"DISPLAYCONTACT.MAILADDRESS.LATITUDE,omitempty"`
	DisplayContactMailAddressLongitude   string  `xml:"DISPLAYCONTACT.MAILADDRESS.LONGITUDE,omitempty"`
	DisplayContactStatus                 string  `xml:"DISPLAYCONTACT.STATUS,omitempty"`
	TermName                             string  `xml:"TERMNAME,omitempty"`
	TermValue                            string  `xml:"TERMVALUE,omitempty"`
	CustRepID                            string  `xml:"CUSTREPID,omitempty"`
	CustRepName                          string  `xml:"CUSTREPNAME,omitempty"`
	ResaleNo                             string  `xml:"RESALENO,omitempty"`
	TaxID                                string  `xml:"TAXID,omitempty"`
	CreditLimit                          string  `xml:"CREDITLIMIT,omitempty"`
	TotalDue                             string  `xml:"TOTALDUE,omitempty"`
	Comments                             string  `xml:"COMMENTS,omitempty"`
	AccountLabel                         string  `xml:"ACCOUNTLABEL,omitempty"`
	ARAccount                            string  `xml:"ARACCOUNT,omitempty"`
	ARAccountTitle                       string  `xml:"ARACCOUNTTITLE,omitempty"`
	LastInvoiceDate                      string  `xml:"LAST_INVOICEDATE,omitempty"`
	LastStatementDate                    string  `xml:"LAST_STATEMENTDATE,omitempty"`
	DeliveryOption                       string  `xml:"DELIVERY_OPTIONS,omitempty"`
	TerrityoryID                         string  `xml:"TERRITORYID,omitempty"`
	TerritoryName                        string  `xml:"TERRITORYNAME,omitempty"`
	ShippingMethod                       string  `xml:"SHIPPINGMETHOD,omitempty"`
	CustType                             string  `xml:"CUSTTYPE,omitempty"`
	GLGRPKey                             string  `xml:"GLGRPKEY,omitempty"`
	GLGroup                              string  `xml:"GLGROUP,omitempty"`
	PriceSchedule                        string  `xml:"PRICESCHEDULE,omitempty"`
	Discount                             string  `xml:"DISCOUNT,omitempty"`
	PriceList                            string  `xml:"PRICELIST,omitempty"`
	VSOEPriceList                        string  `xml:"VSOEPRICELIST,omitempty"`
	Currency                             string  `xml:"CURRENCY,omitempty"`
	ContactInfoContactName               string  `xml:"CONTACTINFO.CONTACTNAME,omitempty"`
	ContactInfoPrefix                    string  `xml:"CONTACTINFO.PREFIX,omitempty"`
	ContactInfoFirstName                 string  `xml:"CONTACTINFO.FIRSTNAME,omitempty"`
	ContactInfoInitial                   string  `xml:"CONTACTINFO.INITIAL,omitempty"`
	ContactInfoLastName                  string  `xml:"CONTACTINFO.LASTNAME,omitempty"`
	ContactInfoCompanyName               string  `xml:"CONTACTINFO.COMPANYNAME,omitempty"`
	ContactinfoPrintAs                   string  `xml:"CONTACTINFO.PRINTAS,omitempty"`
	ContactInfoPhone1                    string  `xml:"CONTACTINFO.PHONE1,omitempty"`
	ContactInfoPhone2                    string  `xml:"CONTACTINFO.PHONE2,omitempty"`
	ContactInfoCellphone                 string  `xml:"CONTACTINFO.CELLPHONE,omitempty"`
	ContactInfoPager                     string  `xml:"CONTACTINFO.PAGER,omitempty"`
	ContactInfoFax                       string  `xml:"CONTACTINFO.FAX,omitempty"`
	ContactInfoEmail1                    string  `xml:"CONTACTINFO.EMAIL1,omitempty"`
	ContactInfoEmail2                    string  `xml:"CONTACTINFO.EMAIL2,omitempty"`
	ContactInfoURL1                      string  `xml:"CONTACTINFO.URL1,omitempty"`
	ContactInfoURL2                      string  `xml:"CONTACTINFO.URL2,omitempty"`
	ContactInfoVisible                   string  `xml:"CONTACTINFO.VISIBLE,omitempty"`
	ContactInfoMailAddressAddress1       string  `xml:"CONTACTINFO.MAILADDRESS.ADDRESS1,omitempty"`
	ContactInfoMailAddressAddress2       string  `xml:"CONTACTINFO.MAILADDRESS.ADDRESS2,omitempty"`
	ContactInfoMailAddressCity           string  `xml:"CONTACTINFO.MAILADDRESS.CITY,omitempty"`
	ContactInfoMailAddressState          string  `xml:"CONTACTINFO.MAILADDRESS.STATE,omitempty"`
	ContactInfoMailAddressZip            string  `xml:"CONTACTINFO.MAILADDRESS.ZIP,omitempty"`
	ContactInfoMailAddressCountry        string  `xml:"CONTACTINFO.MAILADDRESS.COUNTRY,omitempty"`
	ContactInfoMailAddressCountryCode    string  `xml:"CONTACTINFO.MAILADDRESS.COUNTRYCODE,omitempty"`
	ShipToContactName                    string  `xml:"SHIPTO.CONTACTNAME,omitempty"`
	ShipToPrefix                         string  `xml:"SHIPTO.PREFIX,omitempty"`
	ShipToFirstName                      string  `xml:"SHIPTO.FIRSTNAME,omitempty"`
	ShipToInitial                        string  `xml:"SHIPTO.INITIAL,omitempty"`
	ShipToLastName                       string  `xml:"SHIPTO.LASTNAME,omitempty"`
	ShipToCompanyName                    string  `xml:"SHIPTO.COMPANYNAME,omitempty"`
	ShipToPrintAs                        string  `xml:"SHIPTO.PRINTAS,omitempty"`
	ShipToTaxable                        string  `xml:"SHIPTO.TAXABLE,omitempty"`
	ShipToTaxGroup                       string  `xml:"SHIPTO.TAXGROUP,omitempty"`
	ShipToTaxSolutionKey                 string  `xml:"SHIPTO.TAXSOLUTIONKEY,omitempty"`
	ShipToTaxSolutionID                  string  `xml:"SHIPTO.TAXSOLUTIONID,omitempty"`
	ShipToTaxSchedule                    string  `xml:"SHIPTO.TAXSCHEDULE,omitempty"`
	ShipToTaxID                          string  `xml:"SHIPTO.TAXID,omitempty"`
	ShipToPhone1                         string  `xml:"SHIPTO.PHONE1,omitempty"`
	ShipToPhone2                         string  `xml:"SHIPTO.PHONE2,omitempty"`
	ShipToCellphone                      string  `xml:"SHIPTO.CELLPHONE,omitempty"`
	ShipToPager                          string  `xml:"SHIPTO.PAGER,omitempty"`
	ShipToFax                            string  `xml:"SHIPTO.FAX,omitempty"`
	ShipToEmail1                         string  `xml:"SHIPTO.EMAIL1,omitempty"`
	ShipToEmail2                         string  `xml:"SHIPTO.EMAIL2,omitempty"`
	ShipToURL1                           string  `xml:"SHIPTO.URL1,omitempty"`
	ShipToURL2                           string  `xml:"SHIPTO.URL2,omitempty"`
	ShipToVisible                        string  `xml:"SHIPTO.VISIBLE,omitempty"`
	ShipToMailAddressAddress1            string  `xml:"SHIPTO.MAILADDRESS.ADDRESS1,omitempty"`
	ShipToMailAddressAddress2            string  `xml:"SHIPTO.MAILADDRESS.ADDRESS2,omitempty"`
	ShipToMailAddressCity                string  `xml:"SHIPTO.MAILADDRESS.CITY,omitempty"`
	ShipToMailAddressState               string  `xml:"SHIPTO.MAILADDRESS.STATE,omitempty"`
	ShipToMailAddressZip                 string  `xml:"SHIPTO.MAILADDRESS.ZIP,omitempty"`
	ShipToMailAddressCountry             string  `xml:"SHIPTO.MAILADDRESS.COUNTRY,omitempty"`
	ShipToMailAddressCountryCode         string  `xml:"SHIPTO.MAILADDRESS.COUNTRYCODE,omitempty"`
	BillToContactName                    string  `xml:"BILLTO.CONTACTNAME,omitempty"`
	BillToPrefix                         string  `xml:"BILLTO.PREFIX,omitempty"`
	BillToFirstName                      string  `xml:"BILLTO.FIRSTNAME,omitempty"`
	BillToInitial                        string  `xml:"BILLTO.INITIAL,omitempty"`
	BillToLastName                       string  `xml:"BILLTO.LASTNAME,omitempty"`
	BillToCompanyName                    string  `xml:"BILLTO.COMPANYNAME,omitempty"`
	BillToPrintAs                        string  `xml:"BILLTO.PRINTAS,omitempty"`
	BillToTaxable                        string  `xml:"BILLTO.TAXABLE,omitempty"`
	BillToTaxGroup                       string  `xml:"BILLTO.TAXGROUP,omitempty"`
	BillToTaxSolutionKey                 string  `xml:"BILLTO.TAXSOLUTIONKEY,omitempty"`
	BillToTaxSolutionID                  string  `xml:"BILLTO.TAXSOLUTIONID,omitempty"`
	BillToTaxSchedule                    string  `xml:"BILLTO.TAXSCHEDULE,omitempty"`
	BillToPhone1                         string  `xml:"BILLTO.PHONE1,omitempty"`
	BillToPhone2                         string  `xml:"BILLTO.PHONE2,omitempty"`
	BillToCellphone                      string  `xml:"BILLTO.CELLPHONE,omitempty"`
	BillToPager                          string  `xml:"BILLTO.PAGER,omitempty"`
	BillToFax                            string  `xml:"BILLTO.FAX,omitempty"`
	BillToEmail1                         string  `xml:"BILLTO.EMAIL1,omitempty"`
	BillToEmail2                         string  `xml:"BILLTO.EMAIL2,omitempty"`
	BillToURL1                           string  `xml:"BILLTO.URL1,omitempty"`
	BillToURL2                           string  `xml:"BILLTO.URL2,omitempty"`
	BillToVisible                        string  `xml:"BILLTO.VISIBLE,omitempty"`
	BillToMailAddressAddress1            string  `xml:"BILLTO.MAILADDRESS.ADDRESS1,omitempty"`
	BillToMailAddressAddress2            string  `xml:"BILLTO.MAILADDRESS.ADDRESS2,omitempty"`
	BillToMailAddressCity                string  `xml:"BILLTO.MAILADDRESS.CITY,omitempty"`
	BillToMailAddressState               string  `xml:"BILLTO.MAILADDRESS.STATE,omitempty"`
	BillToMailAddressZip                 string  `xml:"BILLTO.MAILADDRESS.ZIP,omitempty"`
	BillToMailAddressCountry             string  `xml:"BILLTO.MAILADDRESS.COUNTRY,omitempty"`
	BillToMailAddressCountryCode         string  `xml:"BILLTO.MAILADDRESS.COUNTRYCODE,omitempty"`
	Status                               string  `xml:"STATUS,omitempty"`
	Onetime                              string  `xml:"ONETIME,omitempty"`
	CustMessageID                        string  `xml:"CUSTMESSAGEID,omitempty"`
	OnHold                               string  `xml:"ONHOLD,omitempty"`
	PRCLSTOverride                       string  `xml:"PRCLST_OVERRIDE,omitempty"`
	OEPRCLSTKey                          string  `xml:"OEPRCLSTKEY,omitempty"`
	OEPriceschedKey                      string  `xml:"OEPRICESCHEDKEY,omitempty"`
	EnableOnlineCardPayment              string  `xml:"ENABLEONLINECARDPAYMENT,omitempty"`
	EnableOnlineCHPayment                string  `xml:"ENABLEONLINEACHPAYMENT,omitempty"`
	VSOEPRCLSTKEY                        string  `xml:"VSOEPRCLSTKEY,omitempty"`
	WhenModified                         string  `xml:"WHENMODIFIED,omitempty"`
	ARInvoicePrintTemplateID             string  `xml:"ARINVOICEPRINTTEMPLATEID,omitempty"`
	OEQuotePrintTemplateID               string  `xml:"OEQUOTEPRINTTEMPLATEID,omitempty"`
	OEOrderPrintTemplateID               string  `xml:"OEORDERPRINTTEMPLATEID,omitempty"`
	OEListPrintTemplateID                string  `xml:"OELISTPRINTTEMPLATEID,omitempty"`
	OEInvoicePrintTemplateID             string  `xml:"OEINVOICEPRINTTEMPLATEID,omitempty"`
	OEAdjPrintTemplateID                 string  `xml:"OEADJPRINTTEMPLATEID,omitempty"`
	OEOtherPrintTemplateID               string  `xml:"OEOTHERPRINTTEMPLATEID,omitempty"`
	WhenCreated                          string  `xml:"WHENCREATED,omitempty"`
	CreatedBy                            string  `xml:"CREATEDBY,omitempty"`
	ModifiedBy                           string  `xml:"MODIFIEDBY,omitempty"`
	ObjectRestriction                    string  `xml:"OBJECTRESTRICTION,omitempty"`
	DisplayContactKey                    string  `xml:"DISPLAYCONTACTKEY,omitempty"`
	ContactKey                           string  `xml:"CONTACTKEY,omitempty"`
	ShipToKey                            string  `xml:"SHIPTOKEY,omitempty"`
	BillToKey                            string  `xml:"BILLTOKEY,omitempty"`
	CustRepKey                           string  `xml:"CUSTREPKEY,omitempty"`
	ShipViaKey                           string  `xml:"SHIPVIAKEY,omitempty"`
	TerritoryKey                         string  `xml:"TERRITORYKEY,omitempty"`
	TermsKey                             string  `xml:"TERMSKEY,omitempty"`
	AccountLabelKey                      string  `xml:"ACCOUNTLABELKEY,omitempty"`
	AccountKey                           string  `xml:"ACCOUNTKEY,omitempty"`
	CustTypeKey                          string  `xml:"CUSTTYPEKEY,omitempty"`
	PriceScheduleKey                     string  `xml:"PRICESCHEDULEKEY,omitempty"`
	OffsetGLAccountNo                    string  `xml:"OFFSETGLACCOUNTNO,omitempty"`
	OffsetGLAccountNoTitle               string  `xml:"OFFSETGLACCOUNTNOTITLE,omitempty"`
	AdvBillBy                            string  `xml:"ADVBILLBY,omitempty"`
	AdvBillyByType                       string  `xml:"ADVBILLBYTYPE,omitempty"`
	SupDocID                             string  `xml:"SUPDOCID,omitempty"`
	RetainagePercentage                  string  `xml:"RETAINAGEPERCENTAGE,omitempty"`
	EmailOption                          string  `xml:"EMAILOPTIN,omitempty"`
	MegaEntityKey                        string  `xml:"MEGAENTITYKEY,omitempty"`
	MegaEntityID                         string  `xml:"MEGAENTITYID,omitempty"`
	MegaEntityName                       string  `xml:"MEGAENTITYNAME,omitempty"`
	HideDisplayContact                   string  `xml:"HIDEDISPLAYCONTACT,omitempty"`
	DisplayContact                       Contact `xml:"DISPLAYCONTACT,omitempty"`
}

type Contact struct {
	PrintAs       string `xml:"PRINTAS,omitempty"`
	companyName   string `xml:"COMPANYNAME,omitempty"`
	Taxable       string `xml:"TAXABLE,omitempty"`
	TaxGroup      string `xml:"TAXGROUP,omitempty"`
	TaxSolutionID string `xml:"TAXSOLUTIONID,omitempty"`
	TaxSchedule   string `xml:"TAXSCHEDULE,omitempty"`
	prefix        string `xml:"PREFIX,omitempty"`
	FirstName     string `xml:"FIRSTNAME,omitempty"`
	Lastname      string `xml:"LASTNAME,omitempty"`
	Initial       string `xml:"INITIAL,omitempty"`
	Phone1        string `xml:"PHONE1,omitempty"`
	Phone2        string `xml:"PHONE2,omitempty"`
	Cellphone     string `xml:"CELLPHONE,omitempty"`
	Pager         string `xml:"PAGER,omitempty"`
	Fax           string `xml:"FAX,omitempty"`
	Email1        string `xml:"EMAIL1,omitempty"`
	Email2        string `xml:"EMAIL2,omitempty"`
	URL1          string `xml:"URL1,omitempty"`
	URL2          string `xml:"URL2,omitempty"`
	MAILADDRESS   struct {
		Address1       string `xml:"ADDRESS1,omitempty"`
		Address2       string `xml:"ADDRESS2,omitempty"`
		City           string `xml:"CITY,omitempty"`
		State          string `xml:"STATE,omitempty"`
		Zip            string `xml:"ZIP,omitempty"`
		Country        string `xml:"COUNTRY,omitempty"`
		ISOCountryCode string `xml:"ISOCOUNTRYCODE,omitempty"`
	}
}

type SupDoc struct {
	RecordNo     string `xml:"recordno"`
	SupDocID     string `xml:"supdocid"`
	SupDocName   string `xml:"supdocname"`
	Folder       string `xml:"folder"`
	Description  string `xml:"description"`
	Creationdate string `xml:"creationdate"`
	Createdby    string `xml:"createdby"`
	Attachments  struct {
		Attachment struct {
			AttachmentName string `xml:"attachmentname"`
			AttachmentType string `xml:"attachmenttype"`
			AttachmentData string `xml:"attachmentdata"`
		} `xml:"attachment"`
	} `xml:"attachments"`
	// Customfields string `xml:"customfields"`
}

type InvoiceLineItems []InvoiceLineItem

type InvoiceLineItem struct {
	GLAccountNo       string `xml:"glaccountno,omitempty"`
	AccountLabel      string `xml:"accountlabel,omitempty"`
	OffsetGLAccountNo string `xml:"offsetglaccountno,omitempty"`
	Amount            Number `xml:"amount"`
	Memo              string `xml:"memo,omitempty"`
	LocationID        string `xml:"locationid,omitempty"`
	DepartmentID      string `xml:"departmentid,omitempty"`
	Key               string `xml:"key,omitempty"`
	TotalPaid         string `xml:"totalpaid,omitempty"`
	TotalDue          string `xml:"totaldue,omitempty"`
	// CustomFields      struct {
	// 	CustomField struct {
	// 		CustomFieldName  string `xml:"customfieldname,omitempty"`
	// 		CustomFieldValue string `xml:"customfieldvalue,omitempty"`
	// 	} `xml:"customfield,omitempty"`
	// } `xml:"customfields,omitempty"`
	// RevRecTemplate  string `xml:"revrectemplate,omitempty"`
	// DefRevAccount   string `xml:"defrevaccount,omitempty"`
	// RevRecStartDate struct {
	// 	Year  string `xml:"year,omitempty"`
	// 	Month string `xml:"month,omitempty"`
	// 	Day   string `xml:"day,omitempty"`
	// } `xml:"revrecstartdate,omitempty"`
	// RevRecEndDate struct {
	// 	Year  string `xml:"year,omitempty"`
	// 	Month string `xml:"month,omitempty"`
	// 	Day   string `xml:"day,omitempty"`
	// } `xml:"revrecenddate"`
	ProjectID   string     `xml:"projectid,omitempty"`
	CustomerID  string     `xml:"customerid,omitempty"`
	VendorID    string     `xml:"vendorid,omitempty"`
	EmployeeID  string     `xml:"employeeid,omitempty"`
	ItemID      string     `xml:"itemid,omitempty"`
	ClassID     string     `xml:"classid,omitempty"`
	WarehouseID string     `xml:"warehouseid,omitempty"`
	TaxEntries  TaxEntries `xml:"taxentries>taxentry,omitempty"`
}

func (i InvoiceLineItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(i, e, start)
}

type TaxEntries []TaxEntry

func (ee TaxEntries) IsEmpty() bool {
	return len(ee) == 0
}

type TaxEntry struct {
	DetailID string `xml:"detailid"`
	TrxTax   Number `xml:"trx_tax"`
}

func (te TaxEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(te, e, start)
}

func (e TaxEntry) IsEmpty() bool {
	return zero.IsZero(e)
}
