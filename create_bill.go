package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewCreateBillRequest() CreateBillRequest {
	r := CreateBillRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateBillRequest struct {
	client      *Client
	queryParams *CreateBillQueryParams
	pathParams  *CreateBillPathParams
	method      string
	headers     http.Header
	requestBody CreateBillRequestBody
}

func (r CreateBillRequest) NewQueryParams() *CreateBillQueryParams {
	return &CreateBillQueryParams{}
}

type CreateBillQueryParams struct{}

func (p CreateBillQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateBillRequest) QueryParams() *CreateBillQueryParams {
	return r.queryParams
}

func (r CreateBillRequest) NewPathParams() *CreateBillPathParams {
	return &CreateBillPathParams{}
}

type CreateBillPathParams struct {
}

func (p *CreateBillPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateBillRequest) PathParams() *CreateBillPathParams {
	return r.pathParams
}

func (r *CreateBillRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateBillRequest) Method() string {
	return r.method
}

func (r CreateBillRequest) NewContent() CreateBillRequestContent {
	content := CreateBillRequestContent{}
	return content
}

type CreateBillRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Create    struct {
			APBill Bill `xml:"APBILL"`
		} `xml:"create"`
	} `xml:"function"`
}

type CreateBillRequestBody struct {
	Request
}

func (r CreateBillRequestBody) Content() *CreateBillRequestContent {
	data, ok := r.Operation.Content.(CreateBillRequestContent)
	if ok {
		return &data
	}
	return &CreateBillRequestContent{}
}

func (r *CreateBillRequestBody) SetContent(content CreateBillRequestContent) {
	r.Operation.Content = content
}

func (r *CreateBillRequest) NewRequestBody() CreateBillRequestBody {
	body := CreateBillRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *CreateBillRequest) SetRequestBody(body CreateBillRequestBody) {
	r.requestBody = body
}

func (r *CreateBillRequest) RequestBody() *CreateBillRequestBody {
	return &r.requestBody
}

func (r *CreateBillRequest) NewResponseBody() *CreateBillResponseBody {
	body := &CreateBillResponseBody{
		Response: NewResponse(),
	}
	return body
}

type CreateBillResponseBody struct {
	Response
}

func (r *CreateBillRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *CreateBillRequest) Do() (CreateBillResponseBody, error) {
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
