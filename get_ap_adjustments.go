package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetAPAdjustmentsRequest() GetAPAdjustmentsRequest {
	r := GetAPAdjustmentsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetAPAdjustmentsRequest struct {
	client      *Client
	queryParams *GetAPAdjustmentsQueryParams
	pathParams  *GetAPAdjustmentsPathParams
	method      string
	headers     http.Header
	requestBody GetAPAdjustmentsRequestBody
}

func (r GetAPAdjustmentsRequest) NewQueryParams() *GetAPAdjustmentsQueryParams {
	return &GetAPAdjustmentsQueryParams{}
}

type GetAPAdjustmentsQueryParams struct{}

func (p GetAPAdjustmentsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAPAdjustmentsRequest) QueryParams() *GetAPAdjustmentsQueryParams {
	return r.queryParams
}

func (r GetAPAdjustmentsRequest) NewPathParams() *GetAPAdjustmentsPathParams {
	return &GetAPAdjustmentsPathParams{}
}

type GetAPAdjustmentsPathParams struct {
}

func (p *GetAPAdjustmentsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAPAdjustmentsRequest) PathParams() *GetAPAdjustmentsPathParams {
	return r.pathParams
}

func (r *GetAPAdjustmentsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAPAdjustmentsRequest) Method() string {
	return r.method
}

func (r GetAPAdjustmentsRequest) NewContent() GetAPAdjustmentsRequestContent {
	content := GetAPAdjustmentsRequestContent{}
	content.Function.Query.Object = "APADJUSTMENT"
	return content
}

type GetAPAdjustmentsRequestContent struct {
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

type GetAPAdjustmentsRequestBody struct {
	Request
}

func (r GetAPAdjustmentsRequestBody) Content() GetAPAdjustmentsRequestContent {
	data, ok := r.Operation.Content.(GetAPAdjustmentsRequestContent)
	if ok {
		return data
	}
	return GetAPAdjustmentsRequestContent{}
}

func (r *GetAPAdjustmentsRequest) NewRequestBody() GetAPAdjustmentsRequestBody {
	body := GetAPAdjustmentsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetAPAdjustmentsRequest) SetRequestBody(body GetAPAdjustmentsRequestBody) {
	r.requestBody = body
}

func (r *GetAPAdjustmentsRequest) RequestBody() *GetAPAdjustmentsRequestBody {
	return &r.requestBody
}

func (r *GetAPAdjustmentsRequest) NewResponseBody() *GetAPAdjustmentsResponseBody {
	body := &GetAPAdjustmentsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetAPAdjustmentsResponseBody struct {
	Response
}

func (r *GetAPAdjustmentsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetAPAdjustmentsRequest) Do() (GetAPAdjustmentsResponseBody, error) {
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
