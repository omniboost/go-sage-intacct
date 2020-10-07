package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetLocationsRequest() GetLocationsRequest {
	r := GetLocationsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetLocationsRequest struct {
	client      *Client
	queryParams *GetLocationsQueryParams
	pathParams  *GetLocationsPathParams
	method      string
	headers     http.Header
	requestBody GetLocationsRequestBody
}

func (r GetLocationsRequest) NewQueryParams() *GetLocationsQueryParams {
	return &GetLocationsQueryParams{}
}

type GetLocationsQueryParams struct{}

func (p GetLocationsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetLocationsRequest) QueryParams() *GetLocationsQueryParams {
	return r.queryParams
}

func (r GetLocationsRequest) NewPathParams() *GetLocationsPathParams {
	return &GetLocationsPathParams{}
}

type GetLocationsPathParams struct {
}

func (p *GetLocationsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetLocationsRequest) PathParams() *GetLocationsPathParams {
	return r.pathParams
}

func (r *GetLocationsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetLocationsRequest) Method() string {
	return r.method
}

func (r GetLocationsRequest) NewContent() GetLocationsRequestContent {
	content := GetLocationsRequestContent{}
	content.Function.ReadByQuery.Object = "LOCATION"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetLocationsRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetLocationsRequestBody struct {
	Request
}

func (r GetLocationsRequestBody) Content() GetLocationsRequestContent {
	data, ok := r.Operation.Content.(GetLocationsRequestContent)
	if ok {
		return data
	}
	return GetLocationsRequestContent{}
}

func (r *GetLocationsRequest) NewRequestBody() GetLocationsRequestBody {
	body := GetLocationsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetLocationsRequest) SetRequestBody(body GetLocationsRequestBody) {
	r.requestBody = body
}

func (r *GetLocationsRequest) RequestBody() *GetLocationsRequestBody {
	return &r.requestBody
}

func (r *GetLocationsRequest) NewResponseBody() *GetLocationsResponseBody {
	body := &GetLocationsResponseBody{
		Response: NewResponse(),
	}

	body.Response.Operation.Result.Data = r.NewResponseData()
	return body
}

type GetLocationsResponseBody struct {
	Response
}

func (r GetLocationsResponseBody) Data() *GetLocationsResponseData {
	data, ok := r.Operation.Result.Data.(*GetLocationsResponseData)
	if ok {
		return data
	}
	return &GetLocationsResponseData{}
}

type GetLocationsResponseData struct {
	ListType     string `xml:"listtype,attr"`
	Count        int    `xml:"count,attr"`
	TotalCount   int    `xml:"totalcount,attr"`
	NumRemaining int    `xml:"numremaining,attr"`
	ResultID     string `xml:"resultId,attr"`
	Locations    []struct {
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
	} `xml:"location"`
}

func (r *GetLocationsRequest) NewResponseData() *GetLocationsResponseData {
	return &GetLocationsResponseData{}
}

func (r *GetLocationsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetLocationsRequest) Do() (GetLocationsResponseBody, error) {
	sessionID, err := r.client.SessionID()
	if err != nil {
		return *r.NewResponseBody(), err
	}

	r.RequestBody().SetSessionID(sessionID)

	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody().Request)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, &responseBody.Response)
	return *responseBody, err
}
