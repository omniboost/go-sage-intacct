package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetInvoicesRequest() GetInvoicesRequest {
	r := GetInvoicesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetInvoicesRequest struct {
	client      *Client
	queryParams *GetInvoicesQueryParams
	pathParams  *GetInvoicesPathParams
	method      string
	headers     http.Header
	requestBody GetInvoicesRequestBody
}

func (r GetInvoicesRequest) NewQueryParams() *GetInvoicesQueryParams {
	return &GetInvoicesQueryParams{}
}

type GetInvoicesQueryParams struct{}

func (p GetInvoicesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoicesRequest) QueryParams() *GetInvoicesQueryParams {
	return r.queryParams
}

func (r GetInvoicesRequest) NewPathParams() *GetInvoicesPathParams {
	return &GetInvoicesPathParams{}
}

type GetInvoicesPathParams struct {
}

func (p *GetInvoicesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetInvoicesRequest) PathParams() *GetInvoicesPathParams {
	return r.pathParams
}

func (r *GetInvoicesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoicesRequest) Method() string {
	return r.method
}

func (r GetInvoicesRequest) NewContent() GetInvoicesRequestContent {
	content := GetInvoicesRequestContent{}
	content.Function.Query.Object = "ARINVOICE"
	return content
}

type GetInvoicesRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Query     struct {
			Object  string  `xml:"object"`
			Select  Select  `xml:"select"`
			Filters Filters `xml:"filter"`
			// Select Select `xml:"select"`
		} `xml:"query"`
	} `xml:"function"`
}

type GetInvoicesRequestBody struct {
	Request
}

func (r GetInvoicesRequestBody) Content() GetInvoicesRequestContent {
	data, ok := r.Operation.Content.(GetInvoicesRequestContent)
	if ok {
		return data
	}
	return GetInvoicesRequestContent{}
}

func (r *GetInvoicesRequest) NewRequestBody() GetInvoicesRequestBody {
	body := GetInvoicesRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetInvoicesRequest) SetRequestBody(body GetInvoicesRequestBody) {
	r.requestBody = body
}

func (r *GetInvoicesRequest) RequestBody() *GetInvoicesRequestBody {
	return &r.requestBody
}

func (r *GetInvoicesRequest) NewResponseBody() *GetInvoicesResponseBody {
	body := &GetInvoicesResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetInvoicesResponseBody struct {
	Response
}

func (r *GetInvoicesRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetInvoicesRequest) Do() (GetInvoicesResponseBody, error) {
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
