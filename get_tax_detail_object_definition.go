package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetTaxDetailObjectDefinitionRequest() GetTaxDetailObjectDefinitionRequest {
	r := GetTaxDetailObjectDefinitionRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetTaxDetailObjectDefinitionRequest struct {
	client      *Client
	queryParams *GetTaxDetailObjectDefinitionQueryParams
	pathParams  *GetTaxDetailObjectDefinitionPathParams
	method      string
	headers     http.Header
	requestBody GetTaxDetailObjectDefinitionRequestBody
}

func (r GetTaxDetailObjectDefinitionRequest) NewQueryParams() *GetTaxDetailObjectDefinitionQueryParams {
	return &GetTaxDetailObjectDefinitionQueryParams{}
}

type GetTaxDetailObjectDefinitionQueryParams struct{}

func (p GetTaxDetailObjectDefinitionQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTaxDetailObjectDefinitionRequest) QueryParams() *GetTaxDetailObjectDefinitionQueryParams {
	return r.queryParams
}

func (r GetTaxDetailObjectDefinitionRequest) NewPathParams() *GetTaxDetailObjectDefinitionPathParams {
	return &GetTaxDetailObjectDefinitionPathParams{}
}

type GetTaxDetailObjectDefinitionPathParams struct {
}

func (p *GetTaxDetailObjectDefinitionPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTaxDetailObjectDefinitionRequest) PathParams() *GetTaxDetailObjectDefinitionPathParams {
	return r.pathParams
}

func (r *GetTaxDetailObjectDefinitionRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTaxDetailObjectDefinitionRequest) Method() string {
	return r.method
}

func (r GetTaxDetailObjectDefinitionRequest) NewContent() GetTaxDetailObjectDefinitionRequestContent {
	content := GetTaxDetailObjectDefinitionRequestContent{}
	content.Function.Lookup.Object = "TAXDETAIL"
	return content
}

type GetTaxDetailObjectDefinitionRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Lookup    struct {
			Object string `xml:"object"`
		} `xml:"lookup"`
	} `xml:"function"`
}

type GetTaxDetailObjectDefinitionRequestBody struct {
	Request
}

func (r GetTaxDetailObjectDefinitionRequestBody) Content() GetTaxDetailObjectDefinitionRequestContent {
	data, ok := r.Operation.Content.(GetTaxDetailObjectDefinitionRequestContent)
	if ok {
		return data
	}
	return GetTaxDetailObjectDefinitionRequestContent{}
}

func (r *GetTaxDetailObjectDefinitionRequest) NewRequestBody() GetTaxDetailObjectDefinitionRequestBody {
	body := GetTaxDetailObjectDefinitionRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetTaxDetailObjectDefinitionRequest) SetRequestBody(body GetTaxDetailObjectDefinitionRequestBody) {
	r.requestBody = body
}

func (r *GetTaxDetailObjectDefinitionRequest) RequestBody() *GetTaxDetailObjectDefinitionRequestBody {
	return &r.requestBody
}

func (r *GetTaxDetailObjectDefinitionRequest) NewResponseBody() *GetTaxDetailObjectDefinitionResponseBody {
	body := &GetTaxDetailObjectDefinitionResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetTaxDetailObjectDefinitionResponseBody struct {
	Response
}

func (r *GetTaxDetailObjectDefinitionRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetTaxDetailObjectDefinitionRequest) Do() (GetTaxDetailObjectDefinitionResponseBody, error) {
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
