package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetCustomersRequest() GetCustomersRequest {
	r := GetCustomersRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetCustomersRequest struct {
	client      *Client
	queryParams *GetCustomersQueryParams
	pathParams  *GetCustomersPathParams
	method      string
	headers     http.Header
	requestBody GetCustomersRequestBody
}

func (r GetCustomersRequest) NewQueryParams() *GetCustomersQueryParams {
	return &GetCustomersQueryParams{}
}

type GetCustomersQueryParams struct{}

func (p GetCustomersQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCustomersRequest) QueryParams() *GetCustomersQueryParams {
	return r.queryParams
}

func (r GetCustomersRequest) NewPathParams() *GetCustomersPathParams {
	return &GetCustomersPathParams{}
}

type GetCustomersPathParams struct {
}

func (p *GetCustomersPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCustomersRequest) PathParams() *GetCustomersPathParams {
	return r.pathParams
}

func (r *GetCustomersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCustomersRequest) Method() string {
	return r.method
}

func (r GetCustomersRequest) NewContent() GetCustomersRequestContent {
	content := GetCustomersRequestContent{}
	content.Function.ReadByQuery.Object = "CUSTOMER"
	content.Function.ReadByQuery.Fields = "*"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetCustomersRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetCustomersRequestBody struct {
	Request
}

func (r GetCustomersRequestBody) Content() GetCustomersRequestContent {
	data, ok := r.Operation.Content.(GetCustomersRequestContent)
	if ok {
		return data
	}
	return GetCustomersRequestContent{}
}

func (r *GetCustomersRequest) NewRequestBody() GetCustomersRequestBody {
	body := GetCustomersRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetCustomersRequest) SetRequestBody(body GetCustomersRequestBody) {
	r.requestBody = body
}

func (r *GetCustomersRequest) RequestBody() *GetCustomersRequestBody {
	return &r.requestBody
}

func (r *GetCustomersRequest) NewResponseBody() *GetCustomersResponseBody {
	body := &GetCustomersResponseBody{
		Response: NewResponse(),
	}

	body.Response.Operation.Result.Data = r.NewResponseData()
	return body
}

type GetCustomersResponseBody struct {
	Response
}

func (r GetCustomersResponseBody) Data() *GetCustomersResponseData {
	data, ok := r.Operation.Result.Data.(*GetCustomersResponseData)
	if ok {
		return data
	}
	return &GetCustomersResponseData{}
}

type GetCustomersResponseData struct {
	ListType     string `xml:"listtype,attr"`
	Count        int    `xml:"count,attr"`
	TotalCount   int    `xml:"totalcount,attr"`
	NumRemaining int    `xml:"numremaining,attr"`
	ResultID     string `xml:"resultId,attr"`
	Customers    []struct {
	} `xml:"customer"`
}

func (r *GetCustomersRequest) NewResponseData() *GetCustomersResponseData {
	return &GetCustomersResponseData{}
}

func (r *GetCustomersRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetCustomersRequest) Do() (GetCustomersResponseBody, error) {
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
