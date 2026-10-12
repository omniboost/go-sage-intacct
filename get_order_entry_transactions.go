package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetOrderEntryTransactionsRequest() GetOrderEntryTransactionsRequest {
	r := GetOrderEntryTransactionsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetOrderEntryTransactionsRequest struct {
	client      *Client
	queryParams *GetOrderEntryTransactionsQueryParams
	pathParams  *GetOrderEntryTransactionsPathParams
	method      string
	headers     http.Header
	requestBody GetOrderEntryTransactionsRequestBody
}

func (r GetOrderEntryTransactionsRequest) NewQueryParams() *GetOrderEntryTransactionsQueryParams {
	return &GetOrderEntryTransactionsQueryParams{}
}

type GetOrderEntryTransactionsQueryParams struct{}

func (p GetOrderEntryTransactionsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetOrderEntryTransactionsRequest) QueryParams() *GetOrderEntryTransactionsQueryParams {
	return r.queryParams
}

func (r GetOrderEntryTransactionsRequest) NewPathParams() *GetOrderEntryTransactionsPathParams {
	return &GetOrderEntryTransactionsPathParams{}
}

type GetOrderEntryTransactionsPathParams struct {
}

func (p *GetOrderEntryTransactionsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetOrderEntryTransactionsRequest) PathParams() *GetOrderEntryTransactionsPathParams {
	return r.pathParams
}

func (r *GetOrderEntryTransactionsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetOrderEntryTransactionsRequest) Method() string {
	return r.method
}

func (r GetOrderEntryTransactionsRequest) NewContent() GetOrderEntryTransactionsRequestContent {
	content := GetOrderEntryTransactionsRequestContent{}
	content.Function.ReadByQuery.Object = "SODOCUMENT"
	content.Function.ReadByQuery.Fields = "RECORDNO,DOCID,DOCNO,DOCPARID,PONUMBER,STATE,CUSTVENDID,TOTAL,WHENCREATED"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetOrderEntryTransactionsRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetOrderEntryTransactionsRequestBody struct {
	Request
}

func (r GetOrderEntryTransactionsRequestBody) Content() GetOrderEntryTransactionsRequestContent {
	data, ok := r.Operation.Content.(GetOrderEntryTransactionsRequestContent)
	if ok {
		return data
	}
	return GetOrderEntryTransactionsRequestContent{}
}

func (r *GetOrderEntryTransactionsRequest) NewRequestBody() GetOrderEntryTransactionsRequestBody {
	body := GetOrderEntryTransactionsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetOrderEntryTransactionsRequest) SetRequestBody(body GetOrderEntryTransactionsRequestBody) {
	r.requestBody = body
}

func (r *GetOrderEntryTransactionsRequest) RequestBody() *GetOrderEntryTransactionsRequestBody {
	return &r.requestBody
}

func (r *GetOrderEntryTransactionsRequest) NewResponseBody() *GetOrderEntryTransactionsResponseBody {
	body := &GetOrderEntryTransactionsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetOrderEntryTransactionsResponseBody struct {
	Response
}

func (r *GetOrderEntryTransactionsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetOrderEntryTransactionsRequest) Do() (GetOrderEntryTransactionsResponseBody, error) {
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
