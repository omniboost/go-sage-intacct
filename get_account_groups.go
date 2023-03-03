package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetAccountGroupsRequest() GetAccountGroupsRequest {
	r := GetAccountGroupsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetAccountGroupsRequest struct {
	client      *Client
	queryParams *GetAccountGroupsQueryParams
	pathParams  *GetAccountGroupsPathParams
	method      string
	headers     http.Header
	requestBody GetAccountGroupsRequestBody
}

func (r GetAccountGroupsRequest) NewQueryParams() *GetAccountGroupsQueryParams {
	return &GetAccountGroupsQueryParams{}
}

type GetAccountGroupsQueryParams struct{}

func (p GetAccountGroupsQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAccountGroupsRequest) QueryParams() *GetAccountGroupsQueryParams {
	return r.queryParams
}

func (r GetAccountGroupsRequest) NewPathParams() *GetAccountGroupsPathParams {
	return &GetAccountGroupsPathParams{}
}

type GetAccountGroupsPathParams struct {
}

func (p *GetAccountGroupsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAccountGroupsRequest) PathParams() *GetAccountGroupsPathParams {
	return r.pathParams
}

func (r *GetAccountGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAccountGroupsRequest) Method() string {
	return r.method
}

func (r GetAccountGroupsRequest) NewContent() GetAccountGroupsRequestContent {
	content := GetAccountGroupsRequestContent{}
	content.Function.ReadByQuery.Object = "GLACCTGRP"
	content.Function.ReadByQuery.Query = NoQuery{}
	return content
}

type GetAccountGroupsRequestContent struct {
	Function struct {
		ControlID   string      `xml:"controlid,attr"`
		ReadByQuery ReadByQuery `xml:"readByQuery"`
	} `xml:"function"`
}

type GetAccountGroupsRequestBody struct {
	Request
}

func (r GetAccountGroupsRequestBody) Content() GetAccountGroupsRequestContent {
	data, ok := r.Operation.Content.(GetAccountGroupsRequestContent)
	if ok {
		return data
	}
	return GetAccountGroupsRequestContent{}
}

func (r *GetAccountGroupsRequest) NewRequestBody() GetAccountGroupsRequestBody {
	body := GetAccountGroupsRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetAccountGroupsRequest) SetRequestBody(body GetAccountGroupsRequestBody) {
	r.requestBody = body
}

func (r *GetAccountGroupsRequest) RequestBody() *GetAccountGroupsRequestBody {
	return &r.requestBody
}

func (r *GetAccountGroupsRequest) NewResponseBody() *GetAccountGroupsResponseBody {
	body := &GetAccountGroupsResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetAccountGroupsResponseBody struct {
	Response
}

func (r *GetAccountGroupsRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetAccountGroupsRequest) Do() (GetAccountGroupsResponseBody, error) {
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
