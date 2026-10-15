package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetOrderEntryTransactionDefinitionsRequest() GetOrderEntryTransactionDefinitionsRequest {
	r := GetOrderEntryTransactionDefinitionsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetOrderEntryTransactionDefinitionsRequest struct {
	client      *Client
	queryParams *GetOrderEntryTransactionDefinitionsQueryParams
	pathParams  *GetOrderEntryTransactionDefinitionsPathParams
	method      string
	headers     http.Header
	requestBody GetOrderEntryTransactionDefinitionsRequestBody
}

func (r GetOrderEntryTransactionDefinitionsRequest) NewQueryParams() *GetOrderEntryTransactionDefinitionsQueryParams {
	return &GetOrderEntryTransactionDefinitionsQueryParams{}
}

type GetOrderEntryTransactionDefinitionsQueryParams struct{}

func (p GetOrderEntryTransactionDefinitionsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetOrderEntryTransactionDefinitionsRequest) QueryParams() *GetOrderEntryTransactionDefinitionsQueryParams {
	return r.queryParams
}

func (r GetOrderEntryTransactionDefinitionsRequest) NewPathParams() *GetOrderEntryTransactionDefinitionsPathParams {
	return &GetOrderEntryTransactionDefinitionsPathParams{}
}

type GetOrderEntryTransactionDefinitionsPathParams struct {
}

func (p *GetOrderEntryTransactionDefinitionsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetOrderEntryTransactionDefinitionsRequest) PathParams() *GetOrderEntryTransactionDefinitionsPathParams {
	return r.pathParams
}

func (r *GetOrderEntryTransactionDefinitionsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetOrderEntryTransactionDefinitionsRequest) Method() string {
	return r.method
}

func (r GetOrderEntryTransactionDefinitionsRequest) NewContent() GetOrderEntryTransactionDefinitionsRequestContent {
	content := GetOrderEntryTransactionDefinitionsRequestContent{}
	content.Function.ReadByQuery.Object = "SODOCUMENTPARAMS"
	content.Function.ReadByQuery.Fields = "RECORDNO,DOCID,STATUS,ENABLEOVERRIDETAX,SHOW_TOTALS,SUBTOTALTEMPLATE,DISABLEVAT"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetOrderEntryTransactionDefinitionsRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetOrderEntryTransactionDefinitionsRequestBody struct {
	Request
}

func (r GetOrderEntryTransactionDefinitionsRequestBody) Content() GetOrderEntryTransactionDefinitionsRequestContent {
	data, ok := r.Operation.Content.(GetOrderEntryTransactionDefinitionsRequestContent)
	if ok {
		return data
	}
	return GetOrderEntryTransactionDefinitionsRequestContent{}
}

func (r *GetOrderEntryTransactionDefinitionsRequest) NewRequestBody() GetOrderEntryTransactionDefinitionsRequestBody {
	body := GetOrderEntryTransactionDefinitionsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetOrderEntryTransactionDefinitionsRequest) SetRequestBody(body GetOrderEntryTransactionDefinitionsRequestBody) {
	r.requestBody = body
}

func (r *GetOrderEntryTransactionDefinitionsRequest) RequestBody() *GetOrderEntryTransactionDefinitionsRequestBody {
	return &r.requestBody
}

func (r *GetOrderEntryTransactionDefinitionsRequest) NewResponseBody() *GetOrderEntryTransactionDefinitionsResponseBody {
	body := &GetOrderEntryTransactionDefinitionsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetOrderEntryTransactionDefinitionsResponseBody struct {
	Response
}

func (r *GetOrderEntryTransactionDefinitionsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetOrderEntryTransactionDefinitionsRequest) Do() (GetOrderEntryTransactionDefinitionsResponseBody, error) {
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
