package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetJournalEntriesRequest() GetJournalEntriesRequest {
	r := GetJournalEntriesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetJournalEntriesRequest struct {
	client      *Client
	queryParams *GetJournalEntriesQueryParams
	pathParams  *GetJournalEntriesPathParams
	method      string
	headers     http.Header
	requestBody GetJournalEntriesRequestBody
}

func (r GetJournalEntriesRequest) NewQueryParams() *GetJournalEntriesQueryParams {
	return &GetJournalEntriesQueryParams{}
}

type GetJournalEntriesQueryParams struct{}

func (p GetJournalEntriesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetJournalEntriesRequest) QueryParams() *GetJournalEntriesQueryParams {
	return r.queryParams
}

func (r GetJournalEntriesRequest) NewPathParams() *GetJournalEntriesPathParams {
	return &GetJournalEntriesPathParams{}
}

type GetJournalEntriesPathParams struct {
}

func (p *GetJournalEntriesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetJournalEntriesRequest) PathParams() *GetJournalEntriesPathParams {
	return r.pathParams
}

func (r *GetJournalEntriesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetJournalEntriesRequest) Method() string {
	return r.method
}

func (r GetJournalEntriesRequest) NewContent() GetJournalEntriesRequestContent {
	content := GetJournalEntriesRequestContent{}
	content.Function.ReadByQuery.Object = "GLBATCH"
	content.Function.ReadByQuery.Fields = "*"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetJournalEntriesRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetJournalEntriesRequestBody struct {
	Request
}

func (r GetJournalEntriesRequestBody) Content() GetJournalEntriesRequestContent {
	data, ok := r.Operation.Content.(GetJournalEntriesRequestContent)
	if ok {
		return data
	}
	return GetJournalEntriesRequestContent{}
}

func (r *GetJournalEntriesRequest) NewRequestBody() GetJournalEntriesRequestBody {
	body := GetJournalEntriesRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetJournalEntriesRequest) SetRequestBody(body GetJournalEntriesRequestBody) {
	r.requestBody = body
}

func (r *GetJournalEntriesRequest) RequestBody() *GetJournalEntriesRequestBody {
	return &r.requestBody
}

func (r *GetJournalEntriesRequest) NewResponseBody() *GetJournalEntriesResponseBody {
	body := &GetJournalEntriesResponseBody{
		Response: NewResponse(),
	}

	body.Response.Operation.Result.Data = r.NewResponseData()
	return body
}

type GetJournalEntriesResponseBody struct {
	Response
}

func (r GetJournalEntriesResponseBody) Data() *GetJournalEntriesResponseData {
	data, ok := r.Operation.Result.Data.(*GetJournalEntriesResponseData)
	if ok {
		return data
	}
	return &GetJournalEntriesResponseData{}
}

type GetJournalEntriesResponseData struct {
}

func (r *GetJournalEntriesRequest) NewResponseData() *GetJournalEntriesResponseData {
	return &GetJournalEntriesResponseData{}
}

func (r *GetJournalEntriesRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetJournalEntriesRequest) Do() (GetJournalEntriesResponseBody, error) {
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
