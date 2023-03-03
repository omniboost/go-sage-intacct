package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetObjectsRequest() GetObjectsRequest {
	r := GetObjectsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetObjectsRequest struct {
	client      *Client
	queryParams *GetObjectsQueryParams
	pathParams  *GetObjectsPathParams
	method      string
	headers     http.Header
	requestBody GetObjectsRequestBody
}

func (r GetObjectsRequest) NewQueryParams() *GetObjectsQueryParams {
	return &GetObjectsQueryParams{}
}

type GetObjectsQueryParams struct{}

func (p GetObjectsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetObjectsRequest) QueryParams() *GetObjectsQueryParams {
	return r.queryParams
}

func (r GetObjectsRequest) NewPathParams() *GetObjectsPathParams {
	return &GetObjectsPathParams{}
}

type GetObjectsPathParams struct {
}

func (p *GetObjectsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetObjectsRequest) PathParams() *GetObjectsPathParams {
	return r.pathParams
}

func (r *GetObjectsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetObjectsRequest) Method() string {
	return r.method
}

func (r GetObjectsRequest) NewContent() GetObjectsRequestContent {
	content := GetObjectsRequestContent{}
	content.Function.Inspect.Object = "*"
	return content
}

type GetObjectsRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Inspect   struct {
			Object string `xml:"object"`
		} `xml:"inspect"`
	} `xml:"function"`
}

type GetObjectsRequestBody struct {
	Request
}

func (r GetObjectsRequestBody) Content() GetObjectsRequestContent {
	data, ok := r.Operation.Content.(GetObjectsRequestContent)
	if ok {
		return data
	}
	return GetObjectsRequestContent{}
}

func (r *GetObjectsRequest) NewRequestBody() GetObjectsRequestBody {
	body := GetObjectsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetObjectsRequest) SetRequestBody(body GetObjectsRequestBody) {
	r.requestBody = body
}

func (r *GetObjectsRequest) RequestBody() *GetObjectsRequestBody {
	return &r.requestBody
}

func (r *GetObjectsRequest) NewResponseBody() *GetObjectsResponseBody {
	body := &GetObjectsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetObjectsResponseBody struct {
	Response
}

func (r *GetObjectsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetObjectsRequest) Do() (GetObjectsResponseBody, error) {
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
