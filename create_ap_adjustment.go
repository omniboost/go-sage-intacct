package intacct

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewCreateAPAdjustmentRequest() CreateAPAdjustmentRequest {
	r := CreateAPAdjustmentRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateAPAdjustmentRequest struct {
	client      *Client
	queryParams *CreateAPAdjustmentQueryParams
	pathParams  *CreateAPAdjustmentPathParams
	method      string
	headers     http.Header
	requestBody CreateAPAdjustmentRequestBody
}

func (r CreateAPAdjustmentRequest) NewQueryParams() *CreateAPAdjustmentQueryParams {
	return &CreateAPAdjustmentQueryParams{}
}

type CreateAPAdjustmentQueryParams struct{}

func (p CreateAPAdjustmentQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateAPAdjustmentRequest) QueryParams() *CreateAPAdjustmentQueryParams {
	return r.queryParams
}

func (r CreateAPAdjustmentRequest) NewPathParams() *CreateAPAdjustmentPathParams {
	return &CreateAPAdjustmentPathParams{}
}

type CreateAPAdjustmentPathParams struct {
}

func (p *CreateAPAdjustmentPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateAPAdjustmentRequest) PathParams() *CreateAPAdjustmentPathParams {
	return r.pathParams
}

func (r *CreateAPAdjustmentRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateAPAdjustmentRequest) Method() string {
	return r.method
}

func (r CreateAPAdjustmentRequest) NewContent() CreateAPAdjustmentRequestContent {
	content := CreateAPAdjustmentRequestContent{}
	return content
}

type CreateAPAdjustmentRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`

		CreateAPAdjustment struct {
			XMLName     xml.Name `xml:"create_apadjustment"`
			VendorID    string   `xml:"vendorid"`
			DateCreated struct {
				Year  string `xml:"year"`
				Month string `xml:"month"`
				Day   string `xml:"day"`
			} `xml:"datecreated"`
			DatePosted struct {
				Year  string `xml:"year"`
				Month string `xml:"month"`
				Day   string `xml:"day"`
			} `xml:"dateposted,omitempty"`
			BatchKey     string `xml:"batchkey,omitempty"`
			AdjustmentNo string `xml:"adjustmentno,omitempty"`
			Action       string `xml:"action,omitempty"`
			BillNo       string `xml:"billno,omitempty"`
			Description  string `xml:"description,omitempty"`
			ExternalID   string `xml:"externalid,omitempty"`
			BaseCurr     string `xml:"basecurr,omitempty"`
			Currency     string `xml:"currency,omitempty"`
			ExchRateDate struct {
				Year  string `xml:"year"`
				Month string `xml:"month"`
				Day   string `xml:"day"`
			} `xml:"exchratedate,omitempty"`
			ExchRateType      string `xml:"exchratetype,omitempty"`
			ExchRate          string `xml:"exchrate,omitempty"`
			NoGL              string `xml:"nogl,omitempty"`
			InclusiveTax      string `xml:"inclusivetax,omitempty"`
			TaxSolutionID     string `xml:"taxsolutionid,omitempty"`
			APAdjustmentItems struct {
				LineItem APAdjustmentItems `xml:"lineitem,omitempty"`
			} `xml:"apadjustmentitems,omitempty"`
		} `xml:"create_apadjustment"`
	} `xml:"function"`
}

type CreateAPAdjustmentRequestBody struct {
	Request
}

func (r CreateAPAdjustmentRequestBody) Content() *CreateAPAdjustmentRequestContent {
	data, ok := r.Operation.Content.(CreateAPAdjustmentRequestContent)
	if ok {
		return &data
	}
	return &CreateAPAdjustmentRequestContent{}
}

func (r *CreateAPAdjustmentRequestBody) SetContent(content CreateAPAdjustmentRequestContent) {
	r.Operation.Content = content
}

func (r *CreateAPAdjustmentRequest) NewRequestBody() CreateAPAdjustmentRequestBody {
	body := CreateAPAdjustmentRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *CreateAPAdjustmentRequest) SetRequestBody(body CreateAPAdjustmentRequestBody) {
	r.requestBody = body
}

func (r *CreateAPAdjustmentRequest) RequestBody() *CreateAPAdjustmentRequestBody {
	return &r.requestBody
}

func (r *CreateAPAdjustmentRequest) NewResponseBody() *CreateAPAdjustmentResponseBody {
	body := &CreateAPAdjustmentResponseBody{
		Response: NewResponse(),
	}
	return body
}

type CreateAPAdjustmentResponseBody struct {
	Response
}

func (r *CreateAPAdjustmentRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *CreateAPAdjustmentRequest) Do() (CreateAPAdjustmentResponseBody, error) {
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
