package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetAPISessionRequest() GetAPISessionRequest {
	r := GetAPISessionRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewGetAPISessionQueryParams()
	r.pathParams = r.NewGetAPISessionPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetAPISessionRequest struct {
	client      *Client
	queryParams *GetAPISessionQueryParams
	pathParams  *GetAPISessionPathParams
	method      string
	headers     http.Header
	requestBody GetAPISessionRequestBody
}

func (r GetAPISessionRequest) NewGetAPISessionQueryParams() *GetAPISessionQueryParams {
	return &GetAPISessionQueryParams{}
}

type GetAPISessionQueryParams struct{}

func (p GetAPISessionQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAPISessionRequest) QueryParams() *GetAPISessionQueryParams {
	return r.queryParams
}

func (r GetAPISessionRequest) NewGetAPISessionPathParams() *GetAPISessionPathParams {
	return &GetAPISessionPathParams{}
}

type GetAPISessionPathParams struct {
}

func (p *GetAPISessionPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAPISessionRequest) PathParams() *GetAPISessionPathParams {
	return r.pathParams
}

func (r *GetAPISessionRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAPISessionRequest) Method() string {
	return r.method
}

func (r GetAPISessionRequest) NewGetAPISessionRequestContent() GetAPISessionRequestContent {
	return GetAPISessionRequestContent{}
}

type GetAPISessionRequestContent struct {
	Function struct {
		ControlID     string `xml:"controlid,attr"`
		GetAPISession struct {
			LocationID string `xml:"locationid,omitempty"`
		} `xml:"getAPISession"`
	} `xml:"function"`
}

func (r *GetAPISessionRequest) RequestBody() *GetAPISessionRequestBody {
	return &r.requestBody
}

type GetAPISessionRequestBody struct {
	Request
}

func (r GetAPISessionRequestBody) Content() GetAPISessionRequestContent {
	data, ok := r.Operation.Content.(GetAPISessionRequestContent)
	if ok {
		return data
	}
	return GetAPISessionRequestContent{}
}

func (r *GetAPISessionRequest) NewRequestBody() GetAPISessionRequestBody {
	body := GetAPISessionRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Authentication.Login = &RequestLogin{
		UserID:    r.client.UserID(),
		Password:  r.client.UserPassword(),
		CompanyID: r.client.CompanyID(),
	}
	content := r.NewGetAPISessionRequestContent()
	content.Function.GetAPISession.LocationID = r.client.LocationID()
	content.Function.ControlID = r.client.GenerateControlID()
	body.Operation.Content = content
	return body
}

func (r *GetAPISessionRequest) SetRequestBody(body GetAPISessionRequestBody) {
	r.requestBody = body
}

func (r *GetAPISessionRequest) NewResponseBody() *GetAPISessionResponseBody {
	body := &GetAPISessionResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetAPISessionResponseBody struct {
	Response
}

func (r *GetAPISessionRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetAPISessionRequest) Do() (GetAPISessionResponseBody, error) {
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
