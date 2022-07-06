package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetInvoiceItemsRequest() GetInvoiceItemsRequest {
	r := GetInvoiceItemsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetInvoiceItemsRequest struct {
	client      *Client
	queryParams *GetInvoiceItemsQueryParams
	pathParams  *GetInvoiceItemsPathParams
	method      string
	headers     http.Header
	requestBody GetInvoiceItemsRequestBody
}

func (r GetInvoiceItemsRequest) NewQueryParams() *GetInvoiceItemsQueryParams {
	return &GetInvoiceItemsQueryParams{}
}

type GetInvoiceItemsQueryParams struct{}

func (p GetInvoiceItemsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoiceItemsRequest) QueryParams() *GetInvoiceItemsQueryParams {
	return r.queryParams
}

func (r GetInvoiceItemsRequest) NewPathParams() *GetInvoiceItemsPathParams {
	return &GetInvoiceItemsPathParams{}
}

type GetInvoiceItemsPathParams struct {
}

func (p *GetInvoiceItemsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetInvoiceItemsRequest) PathParams() *GetInvoiceItemsPathParams {
	return r.pathParams
}

func (r *GetInvoiceItemsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoiceItemsRequest) Method() string {
	return r.method
}

func (r GetInvoiceItemsRequest) NewContent() GetInvoiceItemsRequestContent {
	content := GetInvoiceItemsRequestContent{}
	content.Function.Query.Object = "ARINVOICEITEM"
	return content
}

type GetInvoiceItemsRequestContent struct {
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

type GetInvoiceItemsRequestBody struct {
	Request
}

func (r GetInvoiceItemsRequestBody) Content() GetInvoiceItemsRequestContent {
	data, ok := r.Operation.Content.(GetInvoiceItemsRequestContent)
	if ok {
		return data
	}
	return GetInvoiceItemsRequestContent{}
}

func (r *GetInvoiceItemsRequest) NewRequestBody() GetInvoiceItemsRequestBody {
	body := GetInvoiceItemsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetInvoiceItemsRequest) SetRequestBody(body GetInvoiceItemsRequestBody) {
	r.requestBody = body
}

func (r *GetInvoiceItemsRequest) RequestBody() *GetInvoiceItemsRequestBody {
	return &r.requestBody
}

func (r *GetInvoiceItemsRequest) NewResponseBody() *GetInvoiceItemsResponseBody {
	body := &GetInvoiceItemsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetInvoiceItemsResponseBody struct {
	Response
}

func (r *GetInvoiceItemsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetInvoiceItemsRequest) Do() (GetInvoiceItemsResponseBody, error) {
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
