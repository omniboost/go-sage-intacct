package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewCreateInvoiceRequest() CreateInvoiceRequest {
	r := CreateInvoiceRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateInvoiceRequest struct {
	client      *Client
	queryParams *CreateInvoiceQueryParams
	pathParams  *CreateInvoicePathParams
	method      string
	headers     http.Header
	requestBody CreateInvoiceRequestBody
}

func (r CreateInvoiceRequest) NewQueryParams() *CreateInvoiceQueryParams {
	return &CreateInvoiceQueryParams{}
}

type CreateInvoiceQueryParams struct{}

func (p CreateInvoiceQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateInvoiceRequest) QueryParams() *CreateInvoiceQueryParams {
	return r.queryParams
}

func (r CreateInvoiceRequest) NewPathParams() *CreateInvoicePathParams {
	return &CreateInvoicePathParams{}
}

type CreateInvoicePathParams struct {
}

func (p *CreateInvoicePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateInvoiceRequest) PathParams() *CreateInvoicePathParams {
	return r.pathParams
}

func (r *CreateInvoiceRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateInvoiceRequest) Method() string {
	return r.method
}

func (r CreateInvoiceRequest) NewContent() CreateInvoiceRequestContent {
	content := CreateInvoiceRequestContent{}
	return content
}

type CreateInvoiceRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Create    struct {
			ARInvoice Invoice `xml:"ARINVOICE"`
		} `xml:"create"`
	} `xml:"function"`
}

type CreateInvoiceRequestBody struct {
	Request
}

func (r CreateInvoiceRequestBody) Content() *CreateInvoiceRequestContent {
	data, ok := r.Operation.Content.(CreateInvoiceRequestContent)
	if ok {
		return &data
	}
	return &CreateInvoiceRequestContent{}
}

func (r *CreateInvoiceRequestBody) SetContent(content CreateInvoiceRequestContent) {
	r.Operation.Content = content
}

func (r *CreateInvoiceRequest) NewRequestBody() CreateInvoiceRequestBody {
	body := CreateInvoiceRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *CreateInvoiceRequest) SetRequestBody(body CreateInvoiceRequestBody) {
	r.requestBody = body
}

func (r *CreateInvoiceRequest) RequestBody() *CreateInvoiceRequestBody {
	return &r.requestBody
}

func (r *CreateInvoiceRequest) NewResponseBody() *CreateInvoiceResponseBody {
	body := &CreateInvoiceResponseBody{
		Response: NewResponse(),
	}
	return body
}

type CreateInvoiceResponseBody struct {
	Response
}

func (r *CreateInvoiceRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *CreateInvoiceRequest) Do() (CreateInvoiceResponseBody, error) {
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
