package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetProjectsRequest() GetProjectsRequest {
	r := GetProjectsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetProjectsRequest struct {
	client      *Client
	queryParams *GetProjectsQueryParams
	pathParams  *GetProjectsPathParams
	method      string
	headers     http.Header
	requestBody GetProjectsRequestBody
}

func (r GetProjectsRequest) NewQueryParams() *GetProjectsQueryParams {
	return &GetProjectsQueryParams{}
}

type GetProjectsQueryParams struct{}

func (p GetProjectsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetProjectsRequest) QueryParams() *GetProjectsQueryParams {
	return r.queryParams
}

func (r GetProjectsRequest) NewPathParams() *GetProjectsPathParams {
	return &GetProjectsPathParams{}
}

type GetProjectsPathParams struct {
}

func (p *GetProjectsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetProjectsRequest) PathParams() *GetProjectsPathParams {
	return r.pathParams
}

func (r *GetProjectsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetProjectsRequest) Method() string {
	return r.method
}

func (r GetProjectsRequest) NewContent() GetProjectsRequestContent {
	content := GetProjectsRequestContent{}
	content.Function.ReadByQuery.Object = "project"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetProjectsRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetProjectsRequestBody struct {
	Request
}

func (r GetProjectsRequestBody) Content() GetProjectsRequestContent {
	data, ok := r.Operation.Content.(GetProjectsRequestContent)
	if ok {
		return data
	}
	return GetProjectsRequestContent{}
}

func (r *GetProjectsRequest) NewRequestBody() GetProjectsRequestBody {
	body := GetProjectsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetProjectsRequest) SetRequestBody(body GetProjectsRequestBody) {
	r.requestBody = body
}

func (r *GetProjectsRequest) RequestBody() *GetProjectsRequestBody {
	return &r.requestBody
}

func (r *GetProjectsRequest) NewResponseBody() *GetProjectsResponseBody {
	body := &GetProjectsResponseBody{
		Response: NewResponse(),
	}

	body.Response.Operation.Result.Data = r.NewResponseData()
	return body
}

type GetProjectsResponseBody struct {
	Response
}

func (r GetProjectsResponseBody) Data() *GetProjectsResponseData {
	data, ok := r.Operation.Result.Data.(*GetProjectsResponseData)
	if ok {
		return data
	}
	return &GetProjectsResponseData{}
}

type GetProjectsResponseData struct {
	ListType     string   `xml:"listtype,attr"`
	Count        int      `xml:"count,attr"`
	TotalCount   int      `xml:"totalcount,attr"`
	NumRemaining int      `xml:"numremaining,attr"`
	ResultID     string   `xml:"resultId,attr"`
	Projects     Projects `xml:"project"`
}

func (r *GetProjectsRequest) NewResponseData() *GetProjectsResponseData {
	return &GetProjectsResponseData{}
}

func (r *GetProjectsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetProjectsRequest) Do() (GetProjectsResponseBody, error) {
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
