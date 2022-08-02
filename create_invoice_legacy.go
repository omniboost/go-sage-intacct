package intacct

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewCreateInvoiceLegacyRequest() CreateInvoiceLegacyRequest {
	r := CreateInvoiceLegacyRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateInvoiceLegacyRequest struct {
	client      *Client
	queryParams *CreateInvoiceLegacyQueryParams
	pathParams  *CreateInvoiceLegacyPathParams
	method      string
	headers     http.Header
	requestBody CreateInvoiceLegacyRequestBody
}

func (r CreateInvoiceLegacyRequest) NewQueryParams() *CreateInvoiceLegacyQueryParams {
	return &CreateInvoiceLegacyQueryParams{}
}

type CreateInvoiceLegacyQueryParams struct{}

func (p CreateInvoiceLegacyQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateInvoiceLegacyRequest) QueryParams() *CreateInvoiceLegacyQueryParams {
	return r.queryParams
}

func (r CreateInvoiceLegacyRequest) NewPathParams() *CreateInvoiceLegacyPathParams {
	return &CreateInvoiceLegacyPathParams{}
}

type CreateInvoiceLegacyPathParams struct {
}

func (p *CreateInvoiceLegacyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateInvoiceLegacyRequest) PathParams() *CreateInvoiceLegacyPathParams {
	return r.pathParams
}

func (r *CreateInvoiceLegacyRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateInvoiceLegacyRequest) Method() string {
	return r.method
}

func (r CreateInvoiceLegacyRequest) NewContent() CreateInvoiceLegacyRequestContent {
	content := CreateInvoiceLegacyRequestContent{}
	return content
}

type CreateInvoiceLegacyRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`

		CreateInvoice struct {
			XMLName     xml.Name `xml:"create_invoice"`
			CustomerID  string   `xml:"customerid"`
			DateCreated struct {
				Year  string `xml:"year"`
				Month string `xml:"month"`
				Day   string `xml:"day"`
			} `xml:"datecreated"`
			DatePosted struct {
				Year  string `xml:"year"`
				Month string `xml:"month"`
				Day   string `xml:"day"`
			} `xml:"dateposted"`
			DateDue struct {
				Year  string `xml:"year"`
				Month string `xml:"month"`
				Day   string `xml:"day"`
			} `xml:"datedue"`
			TermName    string `xml:"termname"`
			BatchKey    string `xml:"batchkey"`
			Action      string `xml:"action"`
			InvoiceNo   string `xml:"invoiceno"`
			PoNumber    string `xml:"ponumber"`
			Description string `xml:"description"`
			ExternalID  string `xml:"externalid"`
			BillTo      struct {
				Contactname string `xml:"contactname"`
			} `xml:"billto"`
			ShipTo struct {
				Contactname string `xml:"contactname"`
			} `xml:"shipto"`
			BaseCurr string `xml:"basecurr"`
			Currency string `xml:"currency"`
			// ExchRateDate struct {
			// 	Year  string `xml:"year"`
			// 	Month string `xml:"month"`
			// 	Day   string `xml:"day"`
			// } `xml:"exchratedate"`
			ExchRateType string `xml:"exchratetype"`
			NoGL         string `xml:"nogl"`
			SupDocID     string `xml:"supdocid"`
			// CustomFields struct {
			// 	CustomField struct {
			// 		CustomFieldName  string `xml:"customfieldname"`
			// 		CustomFieldValue string `xml:"customfieldvalue"`
			// 	} `xml:"customfield"`
			// } `xml:"customfields"`
			TaxSolutionID string `xml:"taxsolutionid,omitempty"`
			InvoiceItems  struct {
				LineItem InvoiceLineItems `xml:"lineitem,omitempty"`
			} `xml:"invoiceitems,omitempty"`
		} `xml:"create_invoice"`
	} `xml:"function"`
}

type CreateInvoiceLegacyRequestBody struct {
	Request
}

func (r CreateInvoiceLegacyRequestBody) Content() *CreateInvoiceLegacyRequestContent {
	data, ok := r.Operation.Content.(CreateInvoiceLegacyRequestContent)
	if ok {
		return &data
	}
	return &CreateInvoiceLegacyRequestContent{}
}

func (r *CreateInvoiceLegacyRequestBody) SetContent(content CreateInvoiceLegacyRequestContent) {
	r.Operation.Content = content
}

func (r *CreateInvoiceLegacyRequest) NewRequestBody() CreateInvoiceLegacyRequestBody {
	body := CreateInvoiceLegacyRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *CreateInvoiceLegacyRequest) SetRequestBody(body CreateInvoiceLegacyRequestBody) {
	r.requestBody = body
}

func (r *CreateInvoiceLegacyRequest) RequestBody() *CreateInvoiceLegacyRequestBody {
	return &r.requestBody
}

func (r *CreateInvoiceLegacyRequest) NewResponseBody() *CreateInvoiceLegacyResponseBody {
	body := &CreateInvoiceLegacyResponseBody{
		Response: NewResponse(),
	}
	return body
}

type CreateInvoiceLegacyResponseBody struct {
	Response
}

func (r *CreateInvoiceLegacyRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *CreateInvoiceLegacyRequest) Do() (CreateInvoiceLegacyResponseBody, error) {
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
