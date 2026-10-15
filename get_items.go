package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetItemsRequest() GetItemsRequest {
	r := GetItemsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetItemsRequest struct {
	client      *Client
	queryParams *GetItemsQueryParams
	pathParams  *GetItemsPathParams
	method      string
	headers     http.Header
	requestBody GetItemsRequestBody
}

func (r GetItemsRequest) NewQueryParams() *GetItemsQueryParams {
	return &GetItemsQueryParams{}
}

type GetItemsQueryParams struct{}

func (p GetItemsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetItemsRequest) QueryParams() *GetItemsQueryParams {
	return r.queryParams
}

func (r GetItemsRequest) NewPathParams() *GetItemsPathParams {
	return &GetItemsPathParams{}
}

type GetItemsPathParams struct {
}

func (p *GetItemsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetItemsRequest) PathParams() *GetItemsPathParams {
	return r.pathParams
}

func (r *GetItemsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetItemsRequest) Method() string {
	return r.method
}

func (r GetItemsRequest) NewContent() GetItemsRequestContent {
	content := GetItemsRequestContent{}
	content.Function.ReadByQuery.Object = "ITEM"
	content.Function.ReadByQuery.Fields = "RECORDNO,ITEMID,NAME,ITEMTYPE,STATUS,TAXABLE"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetItemsRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetItemsRequestBody struct {
	Request
}

func (r GetItemsRequestBody) Content() GetItemsRequestContent {
	data, ok := r.Operation.Content.(GetItemsRequestContent)
	if ok {
		return data
	}
	return GetItemsRequestContent{}
}

func (r *GetItemsRequest) NewRequestBody() GetItemsRequestBody {
	body := GetItemsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetItemsRequest) SetRequestBody(body GetItemsRequestBody) {
	r.requestBody = body
}

func (r *GetItemsRequest) RequestBody() *GetItemsRequestBody {
	return &r.requestBody
}

func (r *GetItemsRequest) NewResponseBody() *GetItemsResponseBody {
	body := &GetItemsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetItemsResponseBody struct {
	Response
}

func (r *GetItemsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetItemsRequest) Do() (GetItemsResponseBody, error) {
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
