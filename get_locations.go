package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetLocationsRequest() GetLocationsRequest {
	r := GetLocationsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetLocationsRequest struct {
	client      *Client
	queryParams *GetLocationsQueryParams
	pathParams  *GetLocationsPathParams
	method      string
	headers     http.Header
	requestBody GetLocationsRequestBody
}

func (r GetLocationsRequest) NewQueryParams() *GetLocationsQueryParams {
	return &GetLocationsQueryParams{}
}

type GetLocationsQueryParams struct{}

func (p GetLocationsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetLocationsRequest) QueryParams() *GetLocationsQueryParams {
	return r.queryParams
}

func (r GetLocationsRequest) NewPathParams() *GetLocationsPathParams {
	return &GetLocationsPathParams{}
}

type GetLocationsPathParams struct {
}

func (p *GetLocationsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetLocationsRequest) PathParams() *GetLocationsPathParams {
	return r.pathParams
}

func (r *GetLocationsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetLocationsRequest) Method() string {
	return r.method
}

func (r GetLocationsRequest) NewContent() GetLocationsRequestContent {
	content := GetLocationsRequestContent{}
	content.Function.ReadByQuery.Object = "LOCATION"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetLocationsRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetLocationsRequestBody struct {
	Request
}

func (r GetLocationsRequestBody) Content() GetLocationsRequestContent {
	data, ok := r.Operation.Content.(GetLocationsRequestContent)
	if ok {
		return data
	}
	return GetLocationsRequestContent{}
}

func (r *GetLocationsRequest) NewRequestBody() GetLocationsRequestBody {
	body := GetLocationsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetLocationsRequest) SetRequestBody(body GetLocationsRequestBody) {
	r.requestBody = body
}

func (r *GetLocationsRequest) RequestBody() *GetLocationsRequestBody {
	return &r.requestBody
}

func (r *GetLocationsRequest) NewResponseBody() *GetLocationsResponseBody {
	body := &GetLocationsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetLocationsResponseBody struct {
	Response
}

func (r *GetLocationsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetLocationsRequest) Do() (GetLocationsResponseBody, error) {
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
