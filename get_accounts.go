package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetAccountsRequest() GetAccountsRequest {
	r := GetAccountsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetAccountsRequest struct {
	client      *Client
	queryParams *GetAccountsQueryParams
	pathParams  *GetAccountsPathParams
	method      string
	headers     http.Header
	requestBody GetAccountsRequestBody
}

func (r GetAccountsRequest) NewQueryParams() *GetAccountsQueryParams {
	return &GetAccountsQueryParams{}
}

type GetAccountsQueryParams struct{}

func (p GetAccountsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAccountsRequest) QueryParams() *GetAccountsQueryParams {
	return r.queryParams
}

func (r GetAccountsRequest) NewPathParams() *GetAccountsPathParams {
	return &GetAccountsPathParams{}
}

type GetAccountsPathParams struct {
}

func (p *GetAccountsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAccountsRequest) PathParams() *GetAccountsPathParams {
	return r.pathParams
}

func (r *GetAccountsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAccountsRequest) Method() string {
	return r.method
}

func (r GetAccountsRequest) NewContent() GetAccountsRequestContent {
	content := GetAccountsRequestContent{}
	content.Function.ReadByQuery.Object = "GLACCOUNT"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetAccountsRequestContent struct {
	Function Function `xml:"function"`
}

type GetAccountsRequestBody struct {
	Request
}

func (r GetAccountsRequestBody) Content() GetAccountsRequestContent {
	data, ok := r.Operation.Content.(GetAccountsRequestContent)
	if ok {
		return data
	}
	return GetAccountsRequestContent{}
}

func (r *GetAccountsRequest) NewRequestBody() GetAccountsRequestBody {
	body := GetAccountsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetAccountsRequest) SetRequestBody(body GetAccountsRequestBody) {
	r.requestBody = body
}

func (r *GetAccountsRequest) RequestBody() *GetAccountsRequestBody {
	return &r.requestBody
}

func (r *GetAccountsRequest) NewResponseBody() *GetAccountsResponseBody {
	body := &GetAccountsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetAccountsResponseBody struct {
	Response
}

func (r *GetAccountsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetAccountsRequest) Do() (GetAccountsResponseBody, error) {
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

func (r *GetAccountsRequest) All() (GetAccountsResponseBody, error) {
	numRemaining := -1

	concat := GetAccountsResponseBody{}

	for numRemaining != 0 {
		resp, err := r.Do()
		if err != nil {
			return resp, err
		}

		concat.Operation.Result.Data.GLAccounts = append(concat.Operation.Result.Data.GLAccounts, resp.Operation.Result.Data.GLAccounts...)

		numRemaining = resp.Operation.Result.Data.NumRemaining

		// I have no clue how to do this better at this point
		r.RequestBody().Operation.Content = GetAccountsRequestContent{
			Function: Function{
				ReadMore: ReadMore{ResultID: resp.Operation.Result.Data.ResultID},
			},
		}
	}

	return concat, nil
}
