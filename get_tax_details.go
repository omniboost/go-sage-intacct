package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetTaxDetailsRequest() GetTaxDetailsRequest {
	r := GetTaxDetailsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetTaxDetailsRequest struct {
	client      *Client
	queryParams *GetTaxDetailsQueryParams
	pathParams  *GetTaxDetailsPathParams
	method      string
	headers     http.Header
	requestBody GetTaxDetailsRequestBody
}

func (r GetTaxDetailsRequest) NewQueryParams() *GetTaxDetailsQueryParams {
	return &GetTaxDetailsQueryParams{}
}

type GetTaxDetailsQueryParams struct{}

func (p GetTaxDetailsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTaxDetailsRequest) QueryParams() *GetTaxDetailsQueryParams {
	return r.queryParams
}

func (r GetTaxDetailsRequest) NewPathParams() *GetTaxDetailsPathParams {
	return &GetTaxDetailsPathParams{}
}

type GetTaxDetailsPathParams struct {
}

func (p *GetTaxDetailsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTaxDetailsRequest) PathParams() *GetTaxDetailsPathParams {
	return r.pathParams
}

func (r *GetTaxDetailsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTaxDetailsRequest) Method() string {
	return r.method
}

func (r GetTaxDetailsRequest) NewContent() GetTaxDetailsRequestContent {
	content := GetTaxDetailsRequestContent{}
	content.Function.Query.Object = "TAXDETAIL"
	return content
}

type GetTaxDetailsRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Query     struct {
			Object  string  `xml:"object"`
			Select  Select  `xml:"select"`
			Filters Filters `xml:"filter,omitempty"`
			// Select Select `xml:"select"`
		} `xml:"query"`
	} `xml:"function"`
}

type GetTaxDetailsRequestBody struct {
	Request
}

func (r GetTaxDetailsRequestBody) Content() GetTaxDetailsRequestContent {
	data, ok := r.Operation.Content.(GetTaxDetailsRequestContent)
	if ok {
		return data
	}
	return GetTaxDetailsRequestContent{}
}

func (r *GetTaxDetailsRequest) NewRequestBody() GetTaxDetailsRequestBody {
	body := GetTaxDetailsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetTaxDetailsRequest) SetRequestBody(body GetTaxDetailsRequestBody) {
	r.requestBody = body
}

func (r *GetTaxDetailsRequest) RequestBody() *GetTaxDetailsRequestBody {
	return &r.requestBody
}

func (r *GetTaxDetailsRequest) NewResponseBody() *GetTaxDetailsResponseBody {
	body := &GetTaxDetailsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetTaxDetailsResponseBody struct {
	Response
}

func (r *GetTaxDetailsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetTaxDetailsRequest) Do() (GetTaxDetailsResponseBody, error) {
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
