package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sageone-za/utils"
)

func (c *Client) NewGetAccountsRequest() GetAccountsRequest {
	r := GetAccountsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetAccountsQueryParams()
	r.pathParams = r.NewGetAccountsPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetAccountsRequest struct {
	client      *Client
	queryParams *GetAccountsQueryParams
	pathParams  *GetAccountsPathParams
	method      string
	headers     http.Header
	requestBody GetAccountsRequestBody
}

func (r GetAccountsRequest) NewGetAccountsQueryParams() *GetAccountsQueryParams {
	return &GetAccountsQueryParams{}
}

type GetAccountsQueryParams struct {
}

func (p GetAccountsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAccountsRequest) QueryParams() *GetAccountsQueryParams {
	return r.queryParams
}

func (r GetAccountsRequest) NewGetAccountsPathParams() *GetAccountsPathParams {
	return &GetAccountsPathParams{}
}

type GetAccountsPathParams struct {
}

func (p *GetAccountsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAccountsRequest) PathParams() *GetAccountsPathParams {
	return r.pathParams
}

func (r *GetAccountsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAccountsRequest) Method() string {
	return r.method
}

func (r GetAccountsRequest) NewGetAccountsRequestContent() GetAccountsRequestContent {
	content := GetAccountsRequestContent{}
	content.Function.ReadByQuery.Object = "GLACCOUNT"
	return content
}

type GetAccountsRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetAccountsRequestBody struct {
	Request
}

func (r GetAccountsRequestBody) Content() GetAccountsRequestContent {
	data, ok := r.Operation.Content.(GetAccountsRequestContent)
	if ok {
		return data
	}
	return GetAccountsRequestContent{}
}

func (r *GetAccountsRequest) NewRequestBody() GetAccountsRequestBody {
	body := GetAccountsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewGetAccountsRequestContent()
	return body
}

func (r *GetAccountsRequest) SetRequestBody(body GetAccountsRequestBody) {
	r.requestBody = body
}

func (r *GetAccountsRequest) RequestBody() *GetAccountsRequestBody {
	return &r.requestBody
}

func (r *GetAccountsRequest) NewResponseBody() *GetAccountsResponseBody {
	body := &GetAccountsResponseBody{
		Response: NewResponse(),
	}

	body.Response.Operation.Result.Data = r.NewResponseData()
	return body
}

type GetAccountsResponseBody struct {
	Response
}

func (r GetAccountsResponseBody) Data() *GetAccountsResponseData {
	data, ok := r.Operation.Result.Data.(*GetAccountsResponseData)
	if ok {
		return data
	}
	return &GetAccountsResponseData{}
}

type GetAccountsResponseData struct {
	ListType     string `xml:"listtype,attr"`
	Count        int    `xml:"count,attr"`
	TotalCount   int    `xml:"totalcount,attr"`
	NumRemaining int    `xml:"numremaining,attr"`
	ResultID     string `xml:"resultId,attr"`
	GLAccounts   []struct {
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
		Taxable             bool        `json:"TAXABLE"`
		CategoryKey         string      `json:"CATEGORYKEY"`
		Category            string      `json:"CATEGORY"`
		TaxCode             string      `json:"TAXCODE"`
		MRCCode             string      `json:"MRCCODE"`
		CloseToAcctKey      string      `json:"CLOSETOACCTKEY"`
		AlternativeAccount  string      `json:"ALTERNATIVEACCOUNT"`
		WhenCreated         Date        `json:"WHENCREATED"`
		WhenModified        Date        `json:"WHENMODIFIED"`
		CreatedBy           int         `json:"CREATEDBY"`
		ModifiedBy          int         `json:"MODIFIEDBY"`
		SubledgerControlOn  bool        `json:"SUBLEDGERCONTROLON"`
		MegaEntityKey       interface{} `json:"MEGAENTITYKEY"`
		MegaEntityID        interface{} `json:"MEGAENTITYID"`
		MegaEntityName      interface{} `json:"MEGAENTITYNAME"`
		RequireProject      bool        `json:"REQUIREPROJECT"`
		RequireCustomer     bool        `json:"REQUIRECUSTOMER"`
		RequireVendor       bool        `json:"REQUIREVENDOR"`
		RequireClass        bool        `json:"REQUIRECLASS"`
	} `xml:"glaccount"`
}

func (r *GetAccountsRequest) NewResponseData() *GetAccountsResponseData {
	return &GetAccountsResponseData{}
}

func (r *GetAccountsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetAccountsRequest) Do() (GetAccountsResponseBody, error) {
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
