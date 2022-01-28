package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetDimensionsRequest() GetDimensionsRequest {
	r := GetDimensionsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetDimensionsRequest struct {
	client      *Client
	queryParams *GetDimensionsQueryParams
	pathParams  *GetDimensionsPathParams
	method      string
	headers     http.Header
	requestBody GetDimensionsRequestBody
}

func (r GetDimensionsRequest) NewQueryParams() *GetDimensionsQueryParams {
	return &GetDimensionsQueryParams{}
}

type GetDimensionsQueryParams struct{}

func (p GetDimensionsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDimensionsRequest) QueryParams() *GetDimensionsQueryParams {
	return r.queryParams
}

func (r GetDimensionsRequest) NewPathParams() *GetDimensionsPathParams {
	return &GetDimensionsPathParams{}
}

type GetDimensionsPathParams struct {
}

func (p *GetDimensionsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDimensionsRequest) PathParams() *GetDimensionsPathParams {
	return r.pathParams
}

func (r *GetDimensionsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetDimensionsRequest) Method() string {
	return r.method
}

func (r GetDimensionsRequest) NewContent() GetDimensionsRequestContent {
	content := GetDimensionsRequestContent{}
	return content
}

type GetDimensionsRequestContent struct {
	Function struct {
		ControlID     string   `xml:"controlid,attr"`
		GetDimensions struct{} `xml:"getDimensions"`
	} `xml:"function"`
}

type GetDimensionsRequestBody struct {
	Request
}

func (r GetDimensionsRequestBody) Content() GetDimensionsRequestContent {
	data, ok := r.Operation.Content.(GetDimensionsRequestContent)
	if ok {
		return data
	}
	return GetDimensionsRequestContent{}
}

func (r *GetDimensionsRequest) NewRequestBody() GetDimensionsRequestBody {
	body := GetDimensionsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetDimensionsRequest) SetRequestBody(body GetDimensionsRequestBody) {
	r.requestBody = body
}

func (r *GetDimensionsRequest) RequestBody() *GetDimensionsRequestBody {
	return &r.requestBody
}

func (r *GetDimensionsRequest) NewResponseBody() *GetDimensionsResponseBody {
	body := &GetDimensionsResponseBody{
		Response: NewResponse(),
	}

	body.Response.Operation.Result.Data = r.NewResponseData()
	return body
}

type GetDimensionsResponseBody struct {
	Response
}

func (r GetDimensionsResponseBody) Data() *GetDimensionsResponseData {
	data, ok := r.Operation.Result.Data.(*GetDimensionsResponseData)
	if ok {
		return data
	}
	return &GetDimensionsResponseData{}
}

type GetDimensionsResponseData struct {
}

func (r *GetDimensionsRequest) NewResponseData() *GetDimensionsResponseData {
	return &GetDimensionsResponseData{}
}

func (r *GetDimensionsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetDimensionsRequest) Do() (GetDimensionsResponseBody, error) {
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
