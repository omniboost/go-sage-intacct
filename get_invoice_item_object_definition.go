package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetInvoiceItemObjectDefinitionRequest() GetInvoiceItemObjectDefinitionRequest {
	r := GetInvoiceItemObjectDefinitionRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetInvoiceItemObjectDefinitionRequest struct {
	client      *Client
	queryParams *GetInvoiceItemObjectDefinitionQueryParams
	pathParams  *GetInvoiceItemObjectDefinitionPathParams
	method      string
	headers     http.Header
	requestBody GetInvoiceItemObjectDefinitionRequestBody
}

func (r GetInvoiceItemObjectDefinitionRequest) NewQueryParams() *GetInvoiceItemObjectDefinitionQueryParams {
	return &GetInvoiceItemObjectDefinitionQueryParams{}
}

type GetInvoiceItemObjectDefinitionQueryParams struct{}

func (p GetInvoiceItemObjectDefinitionQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoiceItemObjectDefinitionRequest) QueryParams() *GetInvoiceItemObjectDefinitionQueryParams {
	return r.queryParams
}

func (r GetInvoiceItemObjectDefinitionRequest) NewPathParams() *GetInvoiceItemObjectDefinitionPathParams {
	return &GetInvoiceItemObjectDefinitionPathParams{}
}

type GetInvoiceItemObjectDefinitionPathParams struct {
}

func (p *GetInvoiceItemObjectDefinitionPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetInvoiceItemObjectDefinitionRequest) PathParams() *GetInvoiceItemObjectDefinitionPathParams {
	return r.pathParams
}

func (r *GetInvoiceItemObjectDefinitionRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoiceItemObjectDefinitionRequest) Method() string {
	return r.method
}

func (r GetInvoiceItemObjectDefinitionRequest) NewContent() GetInvoiceItemObjectDefinitionRequestContent {
	content := GetInvoiceItemObjectDefinitionRequestContent{}
	content.Function.Lookup.Object = "ARINVOICEITEM"
	return content
}

type GetInvoiceItemObjectDefinitionRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Lookup    struct {
			Object string `xml:"object"`
		} `xml:"lookup"`
	} `xml:"function"`
}

type GetInvoiceItemObjectDefinitionRequestBody struct {
	Request
}

func (r GetInvoiceItemObjectDefinitionRequestBody) Content() GetInvoiceItemObjectDefinitionRequestContent {
	data, ok := r.Operation.Content.(GetInvoiceItemObjectDefinitionRequestContent)
	if ok {
		return data
	}
	return GetInvoiceItemObjectDefinitionRequestContent{}
}

func (r *GetInvoiceItemObjectDefinitionRequest) NewRequestBody() GetInvoiceItemObjectDefinitionRequestBody {
	body := GetInvoiceItemObjectDefinitionRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetInvoiceItemObjectDefinitionRequest) SetRequestBody(body GetInvoiceItemObjectDefinitionRequestBody) {
	r.requestBody = body
}

func (r *GetInvoiceItemObjectDefinitionRequest) RequestBody() *GetInvoiceItemObjectDefinitionRequestBody {
	return &r.requestBody
}

func (r *GetInvoiceItemObjectDefinitionRequest) NewResponseBody() *GetInvoiceItemObjectDefinitionResponseBody {
	body := &GetInvoiceItemObjectDefinitionResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetInvoiceItemObjectDefinitionResponseBody struct {
	Response
}

func (r *GetInvoiceItemObjectDefinitionRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetInvoiceItemObjectDefinitionRequest) Do() (GetInvoiceItemObjectDefinitionResponseBody, error) {
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
