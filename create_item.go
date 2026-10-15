package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewCreateItemRequest() CreateItemRequest {
	r := CreateItemRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateItemRequest struct {
	client      *Client
	queryParams *CreateItemQueryParams
	pathParams  *CreateItemPathParams
	method      string
	headers     http.Header
	requestBody CreateItemRequestBody
}

func (r CreateItemRequest) NewQueryParams() *CreateItemQueryParams {
	return &CreateItemQueryParams{}
}

type CreateItemQueryParams struct{}

func (p CreateItemQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateItemRequest) QueryParams() *CreateItemQueryParams {
	return r.queryParams
}

func (r CreateItemRequest) NewPathParams() *CreateItemPathParams {
	return &CreateItemPathParams{}
}

type CreateItemPathParams struct {
}

func (p *CreateItemPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateItemRequest) PathParams() *CreateItemPathParams {
	return r.pathParams
}

func (r *CreateItemRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateItemRequest) Method() string {
	return r.method
}

func (r CreateItemRequest) NewContent() CreateItemRequestContent {
	content := CreateItemRequestContent{}
	return content
}

type CreateItemRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Create    struct {
			Item Item `xml:"ITEM"`
		} `xml:"create"`
	} `xml:"function"`
}

type CreateItemRequestBody struct {
	Request
}

func (r CreateItemRequestBody) Content() *CreateItemRequestContent {
	data, ok := r.Operation.Content.(CreateItemRequestContent)
	if ok {
		return &data
	}
	return &CreateItemRequestContent{}
}

func (r *CreateItemRequestBody) SetContent(content CreateItemRequestContent) {
	r.Operation.Content = content
}

func (r *CreateItemRequest) NewRequestBody() CreateItemRequestBody {
	body := CreateItemRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *CreateItemRequest) SetRequestBody(body CreateItemRequestBody) {
	r.requestBody = body
}

func (r *CreateItemRequest) RequestBody() *CreateItemRequestBody {
	return &r.requestBody
}

func (r *CreateItemRequest) NewResponseBody() *CreateItemResponseBody {
	body := &CreateItemResponseBody{
		Response: NewResponse(),
	}
	return body
}

type CreateItemResponseBody struct {
	Response
}

func (r *CreateItemRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *CreateItemRequest) Do() (CreateItemResponseBody, error) {
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
