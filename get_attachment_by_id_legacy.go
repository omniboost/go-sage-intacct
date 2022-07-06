package intacct

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetAttachmentByIDLegacyRequest() GetAttachmentByIDLegacyRequest {
	r := GetAttachmentByIDLegacyRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetAttachmentByIDLegacyRequest struct {
	client      *Client
	queryParams *GetAttachmentByIDLegacyQueryParams
	pathParams  *GetAttachmentByIDLegacyPathParams
	method      string
	headers     http.Header
	requestBody GetAttachmentByIDLegacyRequestBody
}

func (r GetAttachmentByIDLegacyRequest) NewQueryParams() *GetAttachmentByIDLegacyQueryParams {
	return &GetAttachmentByIDLegacyQueryParams{}
}

type GetAttachmentByIDLegacyQueryParams struct{}

func (p GetAttachmentByIDLegacyQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAttachmentByIDLegacyRequest) QueryParams() *GetAttachmentByIDLegacyQueryParams {
	return r.queryParams
}

func (r GetAttachmentByIDLegacyRequest) NewPathParams() *GetAttachmentByIDLegacyPathParams {
	return &GetAttachmentByIDLegacyPathParams{}
}

type GetAttachmentByIDLegacyPathParams struct {
}

func (p *GetAttachmentByIDLegacyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAttachmentByIDLegacyRequest) PathParams() *GetAttachmentByIDLegacyPathParams {
	return r.pathParams
}

func (r *GetAttachmentByIDLegacyRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAttachmentByIDLegacyRequest) Method() string {
	return r.method
}

func (r GetAttachmentByIDLegacyRequest) NewContent() GetAttachmentByIDLegacyRequestContent {
	content := GetAttachmentByIDLegacyRequestContent{}
	return content
}

type GetAttachmentByIDLegacyRequestContent struct {
	// Function struct {
	// 	ControlID               string   `xml:"controlid,attr"`
	// 	GetAttachmentByIDLegacy struct{} `xml:"getDimensions"`
	// } `xml:"function"`
	Function struct {
		ControlID string `xml:"controlid,attr"`

		Get struct {
			XMLName xml.Name `xml:"get"`
			Object  string   `xml:"object,attr"`
			Key     string   `xml:"key,attr"`
		} `xml:"get"`
	} `xml:"function"`
}

type GetAttachmentByIDLegacyRequestBody struct {
	Request
}

func (r GetAttachmentByIDLegacyRequestBody) Content() GetAttachmentByIDLegacyRequestContent {
	data, ok := r.Operation.Content.(GetAttachmentByIDLegacyRequestContent)
	if ok {
		return data
	}
	return GetAttachmentByIDLegacyRequestContent{}
}

func (r *GetAttachmentByIDLegacyRequest) NewRequestBody() GetAttachmentByIDLegacyRequestBody {
	body := GetAttachmentByIDLegacyRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetAttachmentByIDLegacyRequest) SetRequestBody(body GetAttachmentByIDLegacyRequestBody) {
	r.requestBody = body
}

func (r *GetAttachmentByIDLegacyRequest) RequestBody() *GetAttachmentByIDLegacyRequestBody {
	return &r.requestBody
}

func (r *GetAttachmentByIDLegacyRequest) NewResponseBody() *GetAttachmentByIDLegacyResponseBody {
	body := &GetAttachmentByIDLegacyResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetAttachmentByIDLegacyResponseBody struct {
	Response
}

func (r *GetAttachmentByIDLegacyRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetAttachmentByIDLegacyRequest) Do() (GetAttachmentByIDLegacyResponseBody, error) {
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
