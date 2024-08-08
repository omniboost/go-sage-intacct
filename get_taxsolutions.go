package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetTaxsolutionsRequest() GetTaxsolutionsRequest {
	r := GetTaxsolutionsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetTaxsolutionsRequest struct {
	client      *Client
	queryParams *GetTaxsolutionsQueryParams
	pathParams  *GetTaxsolutionsPathParams
	method      string
	headers     http.Header
	requestBody GetTaxsolutionsRequestBody
}

func (r GetTaxsolutionsRequest) NewQueryParams() *GetTaxsolutionsQueryParams {
	return &GetTaxsolutionsQueryParams{}
}

type GetTaxsolutionsQueryParams struct{}

func (p GetTaxsolutionsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTaxsolutionsRequest) QueryParams() *GetTaxsolutionsQueryParams {
	return r.queryParams
}

func (r GetTaxsolutionsRequest) NewPathParams() *GetTaxsolutionsPathParams {
	return &GetTaxsolutionsPathParams{}
}

type GetTaxsolutionsPathParams struct {
}

func (p *GetTaxsolutionsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTaxsolutionsRequest) PathParams() *GetTaxsolutionsPathParams {
	return r.pathParams
}

func (r *GetTaxsolutionsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTaxsolutionsRequest) Method() string {
	return r.method
}

func (r GetTaxsolutionsRequest) NewContent() GetTaxsolutionsRequestContent {
	content := GetTaxsolutionsRequestContent{}
	content.Function.ReadByQuery.Object = "TAXSOLUTION"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetTaxsolutionsRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetTaxsolutionsRequestBody struct {
	Request
}

func (r GetTaxsolutionsRequestBody) Content() GetTaxsolutionsRequestContent {
	data, ok := r.Operation.Content.(GetTaxsolutionsRequestContent)
	if ok {
		return data
	}
	return GetTaxsolutionsRequestContent{}
}

func (r *GetTaxsolutionsRequest) NewRequestBody() GetTaxsolutionsRequestBody {
	body := GetTaxsolutionsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetTaxsolutionsRequest) SetRequestBody(body GetTaxsolutionsRequestBody) {
	r.requestBody = body
}

func (r *GetTaxsolutionsRequest) RequestBody() *GetTaxsolutionsRequestBody {
	return &r.requestBody
}

func (r *GetTaxsolutionsRequest) NewResponseBody() *GetTaxsolutionsResponseBody {
	body := &GetTaxsolutionsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetTaxsolutionsResponseBody struct {
	Response
}

func (r *GetTaxsolutionsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetTaxsolutionsRequest) Do() (GetTaxsolutionsResponseBody, error) {
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

