package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetJournalEntryLinesRequest() GetJournalEntryLinesRequest {
	r := GetJournalEntryLinesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetJournalEntryLinesRequest struct {
	client      *Client
	queryParams *GetJournalEntryLinesQueryParams
	pathParams  *GetJournalEntryLinesPathParams
	method      string
	headers     http.Header
	requestBody GetJournalEntryLinesRequestBody
}

func (r GetJournalEntryLinesRequest) NewQueryParams() *GetJournalEntryLinesQueryParams {
	return &GetJournalEntryLinesQueryParams{}
}

type GetJournalEntryLinesQueryParams struct{}

func (p GetJournalEntryLinesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetJournalEntryLinesRequest) QueryParams() *GetJournalEntryLinesQueryParams {
	return r.queryParams
}

func (r GetJournalEntryLinesRequest) NewPathParams() *GetJournalEntryLinesPathParams {
	return &GetJournalEntryLinesPathParams{}
}

type GetJournalEntryLinesPathParams struct {
}

func (p *GetJournalEntryLinesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetJournalEntryLinesRequest) PathParams() *GetJournalEntryLinesPathParams {
	return r.pathParams
}

func (r *GetJournalEntryLinesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetJournalEntryLinesRequest) Method() string {
	return r.method
}

func (r GetJournalEntryLinesRequest) NewContent() GetJournalEntryLinesRequestContent {
	content := GetJournalEntryLinesRequestContent{}
	content.Function.Query.Object = "GLENTRY"
	return content
}

type GetJournalEntryLinesRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Query     struct {
			Object  string  `xml:"object"`
			Select  Select  `xml:"select"`
			Filters Filters `xml:"filter"`
			// Select Select `xml:"select"`
		} `xml:"query"`
	} `xml:"function"`
}

type GetJournalEntryLinesRequestBody struct {
	Request
}

func (r GetJournalEntryLinesRequestBody) Content() GetJournalEntryLinesRequestContent {
	data, ok := r.Operation.Content.(GetJournalEntryLinesRequestContent)
	if ok {
		return data
	}
	return GetJournalEntryLinesRequestContent{}
}

func (r *GetJournalEntryLinesRequest) NewRequestBody() GetJournalEntryLinesRequestBody {
	body := GetJournalEntryLinesRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetJournalEntryLinesRequest) SetRequestBody(body GetJournalEntryLinesRequestBody) {
	r.requestBody = body
}

func (r *GetJournalEntryLinesRequest) RequestBody() *GetJournalEntryLinesRequestBody {
	return &r.requestBody
}

func (r *GetJournalEntryLinesRequest) NewResponseBody() *GetJournalEntryLinesResponseBody {
	body := &GetJournalEntryLinesResponseBody{
		Response: NewResponse(),
	}

	body.Response.Operation.Result.Data = r.NewResponseData()
	return body
}

type GetJournalEntryLinesResponseBody struct {
	Response
}

func (r GetJournalEntryLinesResponseBody) Data() *GetJournalEntryLinesResponseData {
	data, ok := r.Operation.Result.Data.(*GetJournalEntryLinesResponseData)
	if ok {
		return data
	}
	return &GetJournalEntryLinesResponseData{}
}

type GetJournalEntryLinesResponseData struct {
}

func (r *GetJournalEntryLinesRequest) NewResponseData() *GetJournalEntryLinesResponseData {
	return &GetJournalEntryLinesResponseData{}
}

func (r *GetJournalEntryLinesRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetJournalEntryLinesRequest) Do() (GetJournalEntryLinesResponseBody, error) {
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
