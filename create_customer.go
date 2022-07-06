package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewCreateCustomerRequest() CreateCustomerRequest {
	r := CreateCustomerRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateCustomerRequest struct {
	client      *Client
	queryParams *CreateCustomerQueryParams
	pathParams  *CreateCustomerPathParams
	method      string
	headers     http.Header
	requestBody CreateCustomerRequestBody
}

func (r CreateCustomerRequest) NewQueryParams() *CreateCustomerQueryParams {
	return &CreateCustomerQueryParams{}
}

type CreateCustomerQueryParams struct{}

func (p CreateCustomerQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateCustomerRequest) QueryParams() *CreateCustomerQueryParams {
	return r.queryParams
}

func (r CreateCustomerRequest) NewPathParams() *CreateCustomerPathParams {
	return &CreateCustomerPathParams{}
}

type CreateCustomerPathParams struct {
}

func (p *CreateCustomerPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateCustomerRequest) PathParams() *CreateCustomerPathParams {
	return r.pathParams
}

func (r *CreateCustomerRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateCustomerRequest) Method() string {
	return r.method
}

func (r CreateCustomerRequest) NewContent() CreateCustomerRequestContent {
	content := CreateCustomerRequestContent{}
	return content
}

type CreateCustomerRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Create    struct {
			Customer Customer `xml:"CUSTOMER"`
		} `xml:"create"`
	} `xml:"function"`
}

type CreateCustomerRequestBody struct {
	Request
}

func (r CreateCustomerRequestBody) Content() *CreateCustomerRequestContent {
	data, ok := r.Operation.Content.(CreateCustomerRequestContent)
	if ok {
		return &data
	}
	return &CreateCustomerRequestContent{}
}

func (r *CreateCustomerRequestBody) SetContent(content CreateCustomerRequestContent) {
	r.Operation.Content = content
}

func (r *CreateCustomerRequest) NewRequestBody() CreateCustomerRequestBody {
	body := CreateCustomerRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *CreateCustomerRequest) SetRequestBody(body CreateCustomerRequestBody) {
	r.requestBody = body
}

func (r *CreateCustomerRequest) RequestBody() *CreateCustomerRequestBody {
	return &r.requestBody
}

func (r *CreateCustomerRequest) NewResponseBody() *CreateCustomerResponseBody {
	body := &CreateCustomerResponseBody{
		Response: NewResponse(),
	}
	return body
}

type CreateCustomerResponseBody struct {
	Response
}

func (r *CreateCustomerRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *CreateCustomerRequest) Do() (CreateCustomerResponseBody, error) {
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
