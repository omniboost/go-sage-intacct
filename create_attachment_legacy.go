package intacct

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewCreateAttachmentLegacyRequest() CreateAttachmentLegacyRequest {
	r := CreateAttachmentLegacyRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateAttachmentLegacyRequest struct {
	client      *Client
	queryParams *CreateAttachmentLegacyQueryParams
	pathParams  *CreateAttachmentLegacyPathParams
	method      string
	headers     http.Header
	requestBody CreateAttachmentLegacyRequestBody
}

func (r CreateAttachmentLegacyRequest) NewQueryParams() *CreateAttachmentLegacyQueryParams {
	return &CreateAttachmentLegacyQueryParams{}
}

type CreateAttachmentLegacyQueryParams struct{}

func (p CreateAttachmentLegacyQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateAttachmentLegacyRequest) QueryParams() *CreateAttachmentLegacyQueryParams {
	return r.queryParams
}

func (r CreateAttachmentLegacyRequest) NewPathParams() *CreateAttachmentLegacyPathParams {
	return &CreateAttachmentLegacyPathParams{}
}

type CreateAttachmentLegacyPathParams struct {
}

func (p *CreateAttachmentLegacyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateAttachmentLegacyRequest) PathParams() *CreateAttachmentLegacyPathParams {
	return r.pathParams
}

func (r *CreateAttachmentLegacyRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateAttachmentLegacyRequest) Method() string {
	return r.method
}

func (r CreateAttachmentLegacyRequest) NewContent() CreateAttachmentLegacyRequestContent {
	content := CreateAttachmentLegacyRequestContent{}
	return content
}

type CreateAttachmentLegacyRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`

		CreateAttachment struct {
			XMLName           xml.Name `xml:"create_supdoc"`
			SupDocID          string   `xml:"supdocid"`
			SupDocFolderName  string   `xml:"supdocfoldername"`
			SupDocDescription string   `xml:"supdocdescription"`
			Attachments       []struct {
				AttachmentName string `xml:"attachmentname"`
				AttachmentType string `xml:"attachmenttype"`
				AttachmentData string `xml:"attachmentdata"`
			} `xml:"attachments>attachment"`
		} `xml:"create_supdoc"`
	} `xml:"function"`
}

type CreateAttachmentLegacyRequestBody struct {
	Request
}

func (r CreateAttachmentLegacyRequestBody) Content() *CreateAttachmentLegacyRequestContent {
	data, ok := r.Operation.Content.(CreateAttachmentLegacyRequestContent)
	if ok {
		return &data
	}
	return &CreateAttachmentLegacyRequestContent{}
}

func (r *CreateAttachmentLegacyRequestBody) SetContent(content CreateAttachmentLegacyRequestContent) {
	r.Operation.Content = content
}

func (r *CreateAttachmentLegacyRequest) NewRequestBody() CreateAttachmentLegacyRequestBody {
	body := CreateAttachmentLegacyRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *CreateAttachmentLegacyRequest) SetRequestBody(body CreateAttachmentLegacyRequestBody) {
	r.requestBody = body
}

func (r *CreateAttachmentLegacyRequest) RequestBody() *CreateAttachmentLegacyRequestBody {
	return &r.requestBody
}

func (r *CreateAttachmentLegacyRequest) NewResponseBody() *CreateAttachmentLegacyResponseBody {
	body := &CreateAttachmentLegacyResponseBody{
		Response: NewResponse(),
	}
	return body
}

type CreateAttachmentLegacyResponseBody struct {
	Response
}

func (r *CreateAttachmentLegacyRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *CreateAttachmentLegacyRequest) Do() (CreateAttachmentLegacyResponseBody, error) {
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
