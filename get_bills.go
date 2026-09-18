package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetBillsRequest() GetBillsRequest {
	r := GetBillsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetBillsRequest struct {
	client      *Client
	queryParams *GetBillsQueryParams
	pathParams  *GetBillsPathParams
	method      string
	headers     http.Header
	requestBody GetBillsRequestBody
}

func (r GetBillsRequest) NewQueryParams() *GetBillsQueryParams {
	return &GetBillsQueryParams{}
}

type GetBillsQueryParams struct{}

func (p GetBillsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetBillsRequest) QueryParams() *GetBillsQueryParams {
	return r.queryParams
}

func (r GetBillsRequest) NewPathParams() *GetBillsPathParams {
	return &GetBillsPathParams{}
}

type GetBillsPathParams struct {
}

func (p *GetBillsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetBillsRequest) PathParams() *GetBillsPathParams {
	return r.pathParams
}

func (r *GetBillsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetBillsRequest) Method() string {
	return r.method
}

func (r GetBillsRequest) NewContent() GetBillsRequestContent {
	content := GetBillsRequestContent{}
	content.Function.Query.Object = "APBILL"
	return content
}

type GetBillsRequestContent struct {
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

type GetBillsRequestBody struct {
	Request
}

func (r GetBillsRequestBody) Content() GetBillsRequestContent {
	data, ok := r.Operation.Content.(GetBillsRequestContent)
	if ok {
		return data
	}
	return GetBillsRequestContent{}
}

func (r *GetBillsRequest) NewRequestBody() GetBillsRequestBody {
	body := GetBillsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetBillsRequest) SetRequestBody(body GetBillsRequestBody) {
	r.requestBody = body
}

func (r *GetBillsRequest) RequestBody() *GetBillsRequestBody {
	return &r.requestBody
}

func (r *GetBillsRequest) NewResponseBody() *GetBillsResponseBody {
	body := &GetBillsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetBillsResponseBody struct {
	Response
}

func (r *GetBillsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetBillsRequest) Do() (GetBillsResponseBody, error) {
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
