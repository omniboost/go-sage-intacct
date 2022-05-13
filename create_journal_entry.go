package intacct

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewCreateJournalEntryRequest() CreateJournalEntryRequest {
	r := CreateJournalEntryRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateJournalEntryRequest struct {
	client      *Client
	queryParams *CreateJournalEntryQueryParams
	pathParams  *CreateJournalEntryPathParams
	method      string
	headers     http.Header
	requestBody CreateJournalEntryRequestBody
}

func (r CreateJournalEntryRequest) NewQueryParams() *CreateJournalEntryQueryParams {
	return &CreateJournalEntryQueryParams{}
}

type CreateJournalEntryQueryParams struct{}

func (p CreateJournalEntryQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateJournalEntryRequest) QueryParams() *CreateJournalEntryQueryParams {
	return r.queryParams
}

func (r CreateJournalEntryRequest) NewPathParams() *CreateJournalEntryPathParams {
	return &CreateJournalEntryPathParams{}
}

type CreateJournalEntryPathParams struct {
}

func (p *CreateJournalEntryPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateJournalEntryRequest) PathParams() *CreateJournalEntryPathParams {
	return r.pathParams
}

func (r *CreateJournalEntryRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateJournalEntryRequest) Method() string {
	return r.method
}

func (r CreateJournalEntryRequest) NewContent() CreateJournalEntryRequestContent {
	content := CreateJournalEntryRequestContent{}
	return content
}

type CreateJournalEntryRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`
		Create    struct {
			GLBatch struct {
				Journal     string `xml:"JOURNAL"`
				BatchDate   string `xml:"BATCH_DATE"`
				ReverseDate string `xml:"REVERSEDATE,omitempty"`
				BatchTitle  string `xml:"BATCH_TITLE"`
				State       string `xml:"STATE,omitempty"`
				Entries     []struct {
					AccountNo      string `xml:"ACCOUNTNO"`
					Department     string `xml:"DEPARTMENT,omitempty"`
					Location       string `xml:"LOCATION,omitempty"`
					ProjectID      string `xml:"PROJECTID"`
					ClassID        string `xml:"CLASSID"`
					Currency       string `xml:"CURRENCY"`
					TrType         int    `xml:"TR_TYPE"`
					Amount         Number `xml:"AMOUNT"`
					ExchRateTypeID string `xml:"EXCH_RATE_TYPE_ID"`
					Description    string `xml:"DESCRIPTION"`
				} `xml:"ENTRIES>GLENTRY"`
			} `xml:"GLBATCH"`
		} `xml:"create"`
	} `xml:"function"`
}

type CreateJournalEntryRequestBody struct {
	Request
}

func (r CreateJournalEntryRequestBody) Content() *CreateJournalEntryRequestContent {
	data, ok := r.Operation.Content.(CreateJournalEntryRequestContent)
	if ok {
		return &data
	}
	return &CreateJournalEntryRequestContent{}
}

func (r *CreateJournalEntryRequestBody) SetContent(content CreateJournalEntryRequestContent) {
	r.Operation.Content = content
}

func (r *CreateJournalEntryRequest) NewRequestBody() CreateJournalEntryRequestBody {
	body := CreateJournalEntryRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *CreateJournalEntryRequest) SetRequestBody(body CreateJournalEntryRequestBody) {
	r.requestBody = body
}

func (r *CreateJournalEntryRequest) RequestBody() *CreateJournalEntryRequestBody {
	return &r.requestBody
}

func (r *CreateJournalEntryRequest) NewResponseBody() *CreateJournalEntryResponseBody {
	body := &CreateJournalEntryResponseBody{
		Response: NewResponse(),
	}

	body.Response.Operation.Result.Data = r.NewResponseData()
	return body
}

type CreateJournalEntryResponseBody struct {
	Response
}

func (r CreateJournalEntryResponseBody) Data() *CreateJournalEntryResponseData {
	data, ok := r.Operation.Result.Data.(*CreateJournalEntryResponseData)
	if ok {
		return data
	}
	return &CreateJournalEntryResponseData{}
}

type CreateJournalEntryResponseData struct {
}

func (r *CreateJournalEntryRequest) NewResponseData() *CreateJournalEntryResponseData {
	return &CreateJournalEntryResponseData{}
}

func (r *CreateJournalEntryRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *CreateJournalEntryRequest) Do() (CreateJournalEntryResponseBody, error) {
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
