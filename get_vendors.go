package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetVendorsRequest() GetVendorsRequest {
	r := GetVendorsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetVendorsRequest struct {
	client      *Client
	queryParams *GetVendorsQueryParams
	pathParams  *GetVendorsPathParams
	method      string
	headers     http.Header
	requestBody GetVendorsRequestBody
}

func (r GetVendorsRequest) NewQueryParams() *GetVendorsQueryParams {
	return &GetVendorsQueryParams{}
}

type GetVendorsQueryParams struct{}

func (p GetVendorsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetVendorsRequest) QueryParams() *GetVendorsQueryParams {
	return r.queryParams
}

func (r GetVendorsRequest) NewPathParams() *GetVendorsPathParams {
	return &GetVendorsPathParams{}
}

type GetVendorsPathParams struct {
}

func (p *GetVendorsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetVendorsRequest) PathParams() *GetVendorsPathParams {
	return r.pathParams
}

func (r *GetVendorsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetVendorsRequest) Method() string {
	return r.method
}

func (r GetVendorsRequest) NewContent() GetVendorsRequestContent {
	content := GetVendorsRequestContent{}
	content.Function.Query.Object = "VENDOR"
	return content
}

type GetVendorsRequestContent struct {
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

type GetVendorsRequestBody struct {
	Request
}

func (r GetVendorsRequestBody) Content() GetVendorsRequestContent {
	data, ok := r.Operation.Content.(GetVendorsRequestContent)
	if ok {
		return data
	}
	return GetVendorsRequestContent{}
}

func (r *GetVendorsRequest) NewRequestBody() GetVendorsRequestBody {
	body := GetVendorsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetVendorsRequest) SetRequestBody(body GetVendorsRequestBody) {
	r.requestBody = body
}

func (r *GetVendorsRequest) RequestBody() *GetVendorsRequestBody {
	return &r.requestBody
}

func (r *GetVendorsRequest) NewResponseBody() *GetVendorsResponseBody {
	body := &GetVendorsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetVendorsResponseBody struct {
	Response
}

func (r *GetVendorsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetVendorsRequest) Do() (GetVendorsResponseBody, error) {
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
