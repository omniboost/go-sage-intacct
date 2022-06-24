package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetInvoiceObjectDefinitionRequest() GetInvoiceObjectDefinitionRequest {
	r := GetInvoiceObjectDefinitionRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetInvoiceObjectDefinitionRequest struct {
	client      *Client
	queryParams *GetInvoiceObjectDefinitionQueryParams
	pathParams  *GetInvoiceObjectDefinitionPathParams
	method      string
	headers     http.Header
	requestBody GetInvoiceObjectDefinitionRequestBody
}

func (r GetInvoiceObjectDefinitionRequest) NewQueryParams() *GetInvoiceObjectDefinitionQueryParams {
	return &GetInvoiceObjectDefinitionQueryParams{}
}

type GetInvoiceObjectDefinitionQueryParams struct{}

func (p GetInvoiceObjectDefinitionQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoiceObjectDefinitionRequest) QueryParams() *GetInvoiceObjectDefinitionQueryParams {
	return r.queryParams
}

func (r GetInvoiceObjectDefinitionRequest) NewPathParams() *GetInvoiceObjectDefinitionPathParams {
	return &GetInvoiceObjectDefinitionPathParams{}
}

type GetInvoiceObjectDefinitionPathParams struct {
}

func (p *GetInvoiceObjectDefinitionPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetInvoiceObjectDefinitionRequest) PathParams() *GetInvoiceObjectDefinitionPathParams {
	return r.pathParams
}

func (r *GetInvoiceObjectDefinitionRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoiceObjectDefinitionRequest) Method() string {
	return r.method
}

func (r GetInvoiceObjectDefinitionRequest) NewContent() GetInvoiceObjectDefinitionRequestContent {
	content := GetInvoiceObjectDefinitionRequestContent{}
	content.Function.Lookup.Object = "ARINVOICE"
	return content
}

type GetInvoiceObjectDefinitionRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Lookup    struct {
			Object string `xml:"object"`
		} `xml:"lookup"`
	} `xml:"function"`
}

type GetInvoiceObjectDefinitionRequestBody struct {
	Request
}

func (r GetInvoiceObjectDefinitionRequestBody) Content() GetInvoiceObjectDefinitionRequestContent {
	data, ok := r.Operation.Content.(GetInvoiceObjectDefinitionRequestContent)
	if ok {
		return data
	}
	return GetInvoiceObjectDefinitionRequestContent{}
}

func (r *GetInvoiceObjectDefinitionRequest) NewRequestBody() GetInvoiceObjectDefinitionRequestBody {
	body := GetInvoiceObjectDefinitionRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetInvoiceObjectDefinitionRequest) SetRequestBody(body GetInvoiceObjectDefinitionRequestBody) {
	r.requestBody = body
}

func (r *GetInvoiceObjectDefinitionRequest) RequestBody() *GetInvoiceObjectDefinitionRequestBody {
	return &r.requestBody
}

func (r *GetInvoiceObjectDefinitionRequest) NewResponseBody() *GetInvoiceObjectDefinitionResponseBody {
	body := &GetInvoiceObjectDefinitionResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetInvoiceObjectDefinitionResponseBody struct {
	Response
}

func (r *GetInvoiceObjectDefinitionRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetInvoiceObjectDefinitionRequest) Do() (GetInvoiceObjectDefinitionResponseBody, error) {
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
