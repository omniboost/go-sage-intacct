package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetCustomerObjectDefinitionRequest() GetCustomerObjectDefinitionRequest {
	r := GetCustomerObjectDefinitionRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetCustomerObjectDefinitionRequest struct {
	client      *Client
	queryParams *GetCustomerObjectDefinitionQueryParams
	pathParams  *GetCustomerObjectDefinitionPathParams
	method      string
	headers     http.Header
	requestBody GetCustomerObjectDefinitionRequestBody
}

func (r GetCustomerObjectDefinitionRequest) NewQueryParams() *GetCustomerObjectDefinitionQueryParams {
	return &GetCustomerObjectDefinitionQueryParams{}
}

type GetCustomerObjectDefinitionQueryParams struct{}

func (p GetCustomerObjectDefinitionQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomerObjectDefinitionRequest) QueryParams() *GetCustomerObjectDefinitionQueryParams {
	return r.queryParams
}

func (r GetCustomerObjectDefinitionRequest) NewPathParams() *GetCustomerObjectDefinitionPathParams {
	return &GetCustomerObjectDefinitionPathParams{}
}

type GetCustomerObjectDefinitionPathParams struct {
}

func (p *GetCustomerObjectDefinitionPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomerObjectDefinitionRequest) PathParams() *GetCustomerObjectDefinitionPathParams {
	return r.pathParams
}

func (r *GetCustomerObjectDefinitionRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomerObjectDefinitionRequest) Method() string {
	return r.method
}

func (r GetCustomerObjectDefinitionRequest) NewContent() GetCustomerObjectDefinitionRequestContent {
	content := GetCustomerObjectDefinitionRequestContent{}
	content.Function.Lookup.Object = "CUSTOMER"
	return content
}

type GetCustomerObjectDefinitionRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Lookup    struct {
			Object string `xml:"object"`
		} `xml:"lookup"`
	} `xml:"function"`
}

type GetCustomerObjectDefinitionRequestBody struct {
	Request
}

func (r GetCustomerObjectDefinitionRequestBody) Content() GetCustomerObjectDefinitionRequestContent {
	data, ok := r.Operation.Content.(GetCustomerObjectDefinitionRequestContent)
	if ok {
		return data
	}
	return GetCustomerObjectDefinitionRequestContent{}
}

func (r *GetCustomerObjectDefinitionRequest) NewRequestBody() GetCustomerObjectDefinitionRequestBody {
	body := GetCustomerObjectDefinitionRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetCustomerObjectDefinitionRequest) SetRequestBody(body GetCustomerObjectDefinitionRequestBody) {
	r.requestBody = body
}

func (r *GetCustomerObjectDefinitionRequest) RequestBody() *GetCustomerObjectDefinitionRequestBody {
	return &r.requestBody
}

func (r *GetCustomerObjectDefinitionRequest) NewResponseBody() *GetCustomerObjectDefinitionResponseBody {
	body := &GetCustomerObjectDefinitionResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetCustomerObjectDefinitionResponseBody struct {
	Response
}

func (r *GetCustomerObjectDefinitionRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetCustomerObjectDefinitionRequest) Do() (GetCustomerObjectDefinitionResponseBody, error) {
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
