package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetAttachmentFoldersRequest() GetAttachmentFoldersRequest {
	r := GetAttachmentFoldersRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetAttachmentFoldersRequest struct {
	client      *Client
	queryParams *GetAttachmentFoldersQueryParams
	pathParams  *GetAttachmentFoldersPathParams
	method      string
	headers     http.Header
	requestBody GetAttachmentFoldersRequestBody
}

func (r GetAttachmentFoldersRequest) NewQueryParams() *GetAttachmentFoldersQueryParams {
	return &GetAttachmentFoldersQueryParams{}
}

type GetAttachmentFoldersQueryParams struct{}

func (p GetAttachmentFoldersQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAttachmentFoldersRequest) QueryParams() *GetAttachmentFoldersQueryParams {
	return r.queryParams
}

func (r GetAttachmentFoldersRequest) NewPathParams() *GetAttachmentFoldersPathParams {
	return &GetAttachmentFoldersPathParams{}
}

type GetAttachmentFoldersPathParams struct {
}

func (p *GetAttachmentFoldersPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAttachmentFoldersRequest) PathParams() *GetAttachmentFoldersPathParams {
	return r.pathParams
}

func (r *GetAttachmentFoldersRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAttachmentFoldersRequest) Method() string {
	return r.method
}

func (r GetAttachmentFoldersRequest) NewContent() GetAttachmentFoldersRequestContent {
	content := GetAttachmentFoldersRequestContent{}
	content.Function.GetList.Object = "supdocfolder"
	return content
}

type GetAttachmentFoldersRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		GetList   struct {
			Object string `xml:"object,attr"`
		} `xml:"get_list"`
	} `xml:"function"`
}

type GetAttachmentFoldersRequestBody struct {
	Request
}

func (r GetAttachmentFoldersRequestBody) Content() GetAttachmentFoldersRequestContent {
	data, ok := r.Operation.Content.(GetAttachmentFoldersRequestContent)
	if ok {
		return data
	}
	return GetAttachmentFoldersRequestContent{}
}

func (r *GetAttachmentFoldersRequest) NewRequestBody() GetAttachmentFoldersRequestBody {
	body := GetAttachmentFoldersRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetAttachmentFoldersRequest) SetRequestBody(body GetAttachmentFoldersRequestBody) {
	r.requestBody = body
}

func (r *GetAttachmentFoldersRequest) RequestBody() *GetAttachmentFoldersRequestBody {
	return &r.requestBody
}

func (r *GetAttachmentFoldersRequest) NewResponseBody() *GetAttachmentFoldersResponseBody {
	body := &GetAttachmentFoldersResponseBody{
		Response: NewResponse(),
	}
	return body
}

type GetAttachmentFoldersResponseBody struct {
	Response
}

func (r *GetAttachmentFoldersRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetAttachmentFoldersRequest) Do() (GetAttachmentFoldersResponseBody, error) {
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
