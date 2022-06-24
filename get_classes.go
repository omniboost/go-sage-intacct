package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetClassesRequest() GetClassesRequest {
	r := GetClassesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetClassesRequest struct {
	client      *Client
	queryParams *GetClassesQueryParams
	pathParams  *GetClassesPathParams
	method      string
	headers     http.Header
	requestBody GetClassesRequestBody
}

func (r GetClassesRequest) NewQueryParams() *GetClassesQueryParams {
	return &GetClassesQueryParams{}
}

type GetClassesQueryParams struct{}

func (p GetClassesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetClassesRequest) QueryParams() *GetClassesQueryParams {
	return r.queryParams
}

func (r GetClassesRequest) NewPathParams() *GetClassesPathParams {
	return &GetClassesPathParams{}
}

type GetClassesPathParams struct {
}

func (p *GetClassesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetClassesRequest) PathParams() *GetClassesPathParams {
	return r.pathParams
}

func (r *GetClassesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetClassesRequest) Method() string {
	return r.method
}

func (r GetClassesRequest) NewContent() GetClassesRequestContent {
	content := GetClassesRequestContent{}
	content.Function.ReadByQuery.Object = "class"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetClassesRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetClassesRequestBody struct {
	Request
}

func (r GetClassesRequestBody) Content() GetClassesRequestContent {
	data, ok := r.Operation.Content.(GetClassesRequestContent)
	if ok {
		return data
	}
	return GetClassesRequestContent{}
}

func (r *GetClassesRequest) NewRequestBody() GetClassesRequestBody {
	body := GetClassesRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetClassesRequest) SetRequestBody(body GetClassesRequestBody) {
	r.requestBody = body
}

func (r *GetClassesRequest) RequestBody() *GetClassesRequestBody {
	return &r.requestBody
}

func (r *GetClassesRequest) NewResponseBody() *GetClassesResponseBody {
	body := &GetClassesResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetClassesResponseBody struct {
	Response
}

func (r *GetClassesRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetClassesRequest) Do() (GetClassesResponseBody, error) {
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
