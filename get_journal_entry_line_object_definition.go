package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewGetJournalEntryLineObjectDefinitionRequest() GetJournalEntryLineObjectDefinitionRequest {
	r := GetJournalEntryLineObjectDefinitionRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetJournalEntryLineObjectDefinitionRequest struct {
	client      *Client
	queryParams *GetJournalEntryLineObjectDefinitionQueryParams
	pathParams  *GetJournalEntryLineObjectDefinitionPathParams
	method      string
	headers     http.Header
	requestBody GetJournalEntryLineObjectDefinitionRequestBody
}

func (r GetJournalEntryLineObjectDefinitionRequest) NewQueryParams() *GetJournalEntryLineObjectDefinitionQueryParams {
	return &GetJournalEntryLineObjectDefinitionQueryParams{}
}

type GetJournalEntryLineObjectDefinitionQueryParams struct{}

func (p GetJournalEntryLineObjectDefinitionQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetJournalEntryLineObjectDefinitionRequest) QueryParams() *GetJournalEntryLineObjectDefinitionQueryParams {
	return r.queryParams
}

func (r GetJournalEntryLineObjectDefinitionRequest) NewPathParams() *GetJournalEntryLineObjectDefinitionPathParams {
	return &GetJournalEntryLineObjectDefinitionPathParams{}
}

type GetJournalEntryLineObjectDefinitionPathParams struct {
}

func (p *GetJournalEntryLineObjectDefinitionPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetJournalEntryLineObjectDefinitionRequest) PathParams() *GetJournalEntryLineObjectDefinitionPathParams {
	return r.pathParams
}

func (r *GetJournalEntryLineObjectDefinitionRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetJournalEntryLineObjectDefinitionRequest) Method() string {
	return r.method
}

func (r GetJournalEntryLineObjectDefinitionRequest) NewContent() GetJournalEntryLineObjectDefinitionRequestContent {
	content := GetJournalEntryLineObjectDefinitionRequestContent{}
	content.Function.Lookup.Object = "GLBATCH"
	return content
}

type GetJournalEntryLineObjectDefinitionRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Lookup    struct {
			Object string `xml:"object"`
		} `xml:"lookup"`
	} `xml:"function"`
}

type GetJournalEntryLineObjectDefinitionRequestBody struct {
	Request
}

func (r GetJournalEntryLineObjectDefinitionRequestBody) Content() GetJournalEntryLineObjectDefinitionRequestContent {
	data, ok := r.Operation.Content.(GetJournalEntryLineObjectDefinitionRequestContent)
	if ok {
		return data
	}
	return GetJournalEntryLineObjectDefinitionRequestContent{}
}

func (r *GetJournalEntryLineObjectDefinitionRequest) NewRequestBody() GetJournalEntryLineObjectDefinitionRequestBody {
	body := GetJournalEntryLineObjectDefinitionRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *GetJournalEntryLineObjectDefinitionRequest) SetRequestBody(body GetJournalEntryLineObjectDefinitionRequestBody) {
	r.requestBody = body
}

func (r *GetJournalEntryLineObjectDefinitionRequest) RequestBody() *GetJournalEntryLineObjectDefinitionRequestBody {
	return &r.requestBody
}

func (r *GetJournalEntryLineObjectDefinitionRequest) NewResponseBody() *GetJournalEntryLineObjectDefinitionResponseBody {
	body := &GetJournalEntryLineObjectDefinitionResponseBody{
		Response: NewResponse(),
	}

	body.Response.Operation.Result.Data = r.NewResponseData()
	return body
}

type GetJournalEntryLineObjectDefinitionResponseBody struct {
	Response
}

func (r GetJournalEntryLineObjectDefinitionResponseBody) Data() *GetJournalEntryLineObjectDefinitionResponseData {
	data, ok := r.Operation.Result.Data.(*GetJournalEntryLineObjectDefinitionResponseData)
	if ok {
		return data
	}
	return &GetJournalEntryLineObjectDefinitionResponseData{}
}

type GetJournalEntryLineObjectDefinitionResponseData struct {
	Listtype string `xml:"listtype,attr"`
	Count    string `xml:"count,attr"`
	Type     struct {
		Name         string `xml:"Name,attr"`
		DocumentType string `xml:"DocumentType,attr"`
		Fields       struct {
			Field []struct {
				ID          string `xml:"ID"`
				LABEL       string `xml:"LABEL"`
				DESCRIPTION string `xml:"DESCRIPTION"`
				REQUIRED    string `xml:"REQUIRED"`
				READONLY    string `xml:"READONLY"`
				DATATYPE    string `xml:"DATATYPE"`
				ISCUSTOM    string `xml:"ISCUSTOM"`
				VALIDVALUES struct {
					VALIDVALUE []string `xml:"VALIDVALUE"`
				} `xml:"VALIDVALUES"`
			} `xml:"Field"`
		} `xml:"Fields"`
		Relationships struct {
			Relationship []struct {
				OBJECTPATH       string `xml:"OBJECTPATH"`
				OBJECTNAME       string `xml:"OBJECTNAME"`
				LABEL            string `xml:"LABEL"`
				RELATIONSHIPTYPE string `xml:"RELATIONSHIPTYPE"`
				RELATEDBY        string `xml:"RELATEDBY"`
			} `xml:"Relationship"`
		} `xml:"Relationships"`
	} `xml:"Type"`
}

func (r *GetJournalEntryLineObjectDefinitionRequest) NewResponseData() *GetJournalEntryLineObjectDefinitionResponseData {
	return &GetJournalEntryLineObjectDefinitionResponseData{}
}

func (r *GetJournalEntryLineObjectDefinitionRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *GetJournalEntryLineObjectDefinitionRequest) Do() (GetJournalEntryLineObjectDefinitionResponseBody, error) {
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
