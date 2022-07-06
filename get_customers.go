package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetCustomersRequest() GetCustomersRequest {
	r := GetCustomersRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetCustomersRequest struct {
	client      *Client
	queryParams *GetCustomersQueryParams
	pathParams  *GetCustomersPathParams
	method      string
	headers     http.Header
	requestBody GetCustomersRequestBody
}

func (r GetCustomersRequest) NewQueryParams() *GetCustomersQueryParams {
	return &GetCustomersQueryParams{}
}

type GetCustomersQueryParams struct{}

func (p GetCustomersQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomersRequest) QueryParams() *GetCustomersQueryParams {
	return r.queryParams
}

func (r GetCustomersRequest) NewPathParams() *GetCustomersPathParams {
	return &GetCustomersPathParams{}
}

type GetCustomersPathParams struct {
}

func (p *GetCustomersPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomersRequest) PathParams() *GetCustomersPathParams {
	return r.pathParams
}

func (r *GetCustomersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomersRequest) Method() string {
	return r.method
}

func (r GetCustomersRequest) NewContent() GetCustomersRequestContent {
	content := GetCustomersRequestContent{}
	content.Function.Query.Object = "CUSTOMER"
	return content
}

type GetCustomersRequestContent struct {
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

type GetCustomersRequestBody struct {
	Request
}

func (r GetCustomersRequestBody) Content() GetCustomersRequestContent {
	data, ok := r.Operation.Content.(GetCustomersRequestContent)
	if ok {
		return data
	}
	return GetCustomersRequestContent{}
}

func (r *GetCustomersRequest) NewRequestBody() GetCustomersRequestBody {
	body := GetCustomersRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetCustomersRequest) SetRequestBody(body GetCustomersRequestBody) {
	r.requestBody = body
}

func (r *GetCustomersRequest) RequestBody() *GetCustomersRequestBody {
	return &r.requestBody
}

func (r *GetCustomersRequest) NewResponseBody() *GetCustomersResponseBody {
	body := &GetCustomersResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetCustomersResponseBody struct {
	Response
}

func (r *GetCustomersRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetCustomersRequest) Do() (GetCustomersResponseBody, error) {
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
