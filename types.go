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
