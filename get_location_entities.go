package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetLocationEntitiesRequest() GetLocationEntitiesRequest {
	r := GetLocationEntitiesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetLocationEntitiesRequest struct {
	client      *Client
	queryParams *GetLocationEntitiesQueryParams
	pathParams  *GetLocationEntitiesPathParams
	method      string
	headers     http.Header
	requestBody GetLocationEntitiesRequestBody
}

func (r GetLocationEntitiesRequest) NewQueryParams() *GetLocationEntitiesQueryParams {
	return &GetLocationEntitiesQueryParams{}
}

type GetLocationEntitiesQueryParams struct{}

func (p GetLocationEntitiesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetLocationEntitiesRequest) QueryParams() *GetLocationEntitiesQueryParams {
	return r.queryParams
}

func (r GetLocationEntitiesRequest) NewPathParams() *GetLocationEntitiesPathParams {
	return &GetLocationEntitiesPathParams{}
}

type GetLocationEntitiesPathParams struct {
}

func (p *GetLocationEntitiesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetLocationEntitiesRequest) PathParams() *GetLocationEntitiesPathParams {
	return r.pathParams
}

func (r *GetLocationEntitiesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetLocationEntitiesRequest) Method() string {
	return r.method
}

func (r GetLocationEntitiesRequest) NewContent() GetLocationEntitiesRequestContent {
	content := GetLocationEntitiesRequestContent{}
	content.Function.ReadByQuery.Object = "LOCATIONENTITY"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetLocationEntitiesRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetLocationEntitiesRequestBody struct {
	Request
}

func (r GetLocationEntitiesRequestBody) Content() GetLocationEntitiesRequestContent {
	data, ok := r.Operation.Content.(GetLocationEntitiesRequestContent)
	if ok {
		return data
	}
	return GetLocationEntitiesRequestContent{}
}

func (r *GetLocationEntitiesRequest) NewRequestBody() GetLocationEntitiesRequestBody {
	body := GetLocationEntitiesRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetLocationEntitiesRequest) SetRequestBody(body GetLocationEntitiesRequestBody) {
	r.requestBody = body
}

func (r *GetLocationEntitiesRequest) RequestBody() *GetLocationEntitiesRequestBody {
	return &r.requestBody
}

func (r *GetLocationEntitiesRequest) NewResponseBody() *GetLocationEntitiesResponseBody {
	body := &GetLocationEntitiesResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetLocationEntitiesResponseBody struct {
	Response
}

func (r *GetLocationEntitiesRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetLocationEntitiesRequest) Do() (GetLocationEntitiesResponseBody, error) {
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

