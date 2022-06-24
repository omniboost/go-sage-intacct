package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetDepartmentsRequest() GetDepartmentsRequest {
	r := GetDepartmentsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetDepartmentsRequest struct {
	client      *Client
	queryParams *GetDepartmentsQueryParams
	pathParams  *GetDepartmentsPathParams
	method      string
	headers     http.Header
	requestBody GetDepartmentsRequestBody
}

func (r GetDepartmentsRequest) NewQueryParams() *GetDepartmentsQueryParams {
	return &GetDepartmentsQueryParams{}
}

type GetDepartmentsQueryParams struct{}

func (p GetDepartmentsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDepartmentsRequest) QueryParams() *GetDepartmentsQueryParams {
	return r.queryParams
}

func (r GetDepartmentsRequest) NewPathParams() *GetDepartmentsPathParams {
	return &GetDepartmentsPathParams{}
}

type GetDepartmentsPathParams struct {
}

func (p *GetDepartmentsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDepartmentsRequest) PathParams() *GetDepartmentsPathParams {
	return r.pathParams
}

func (r *GetDepartmentsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDepartmentsRequest) Method() string {
	return r.method
}

func (r GetDepartmentsRequest) NewContent() GetDepartmentsRequestContent {
	content := GetDepartmentsRequestContent{}
	content.Function.ReadByQuery.Object = "DEPARTMENT"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetDepartmentsRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetDepartmentsRequestBody struct {
	Request
}

func (r GetDepartmentsRequestBody) Content() GetDepartmentsRequestContent {
	data, ok := r.Operation.Content.(GetDepartmentsRequestContent)
	if ok {
		return data
	}
	return GetDepartmentsRequestContent{}
}

func (r *GetDepartmentsRequest) NewRequestBody() GetDepartmentsRequestBody {
	body := GetDepartmentsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetDepartmentsRequest) SetRequestBody(body GetDepartmentsRequestBody) {
	r.requestBody = body
}

func (r *GetDepartmentsRequest) RequestBody() *GetDepartmentsRequestBody {
	return &r.requestBody
}

func (r *GetDepartmentsRequest) NewResponseBody() *GetDepartmentsResponseBody {
	body := &GetDepartmentsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetDepartmentsResponseBody struct {
	Response
}

func (r *GetDepartmentsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetDepartmentsRequest) Do() (GetDepartmentsResponseBody, error) {
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
