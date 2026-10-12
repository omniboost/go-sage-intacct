package intacct

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-sage-intacct/utils"
)

func (c *Client) NewCreateOrderEntryTransactionLegacyRequest() CreateOrderEntryTransactionLegacyRequest {
	r := CreateOrderEntryTransactionLegacyRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateOrderEntryTransactionLegacyRequest struct {
	client      *Client
	queryParams *CreateOrderEntryTransactionLegacyQueryParams
	pathParams  *CreateOrderEntryTransactionLegacyPathParams
	method      string
	headers     http.Header
	requestBody CreateOrderEntryTransactionLegacyRequestBody
}

func (r CreateOrderEntryTransactionLegacyRequest) NewQueryParams() *CreateOrderEntryTransactionLegacyQueryParams {
	return &CreateOrderEntryTransactionLegacyQueryParams{}
}

type CreateOrderEntryTransactionLegacyQueryParams struct{}

func (p CreateOrderEntryTransactionLegacyQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateOrderEntryTransactionLegacyRequest) QueryParams() *CreateOrderEntryTransactionLegacyQueryParams {
	return r.queryParams
}

func (r CreateOrderEntryTransactionLegacyRequest) NewPathParams() *CreateOrderEntryTransactionLegacyPathParams {
	return &CreateOrderEntryTransactionLegacyPathParams{}
}

type CreateOrderEntryTransactionLegacyPathParams struct {
}

func (p *CreateOrderEntryTransactionLegacyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateOrderEntryTransactionLegacyRequest) PathParams() *CreateOrderEntryTransactionLegacyPathParams {
	return r.pathParams
}

func (r *CreateOrderEntryTransactionLegacyRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateOrderEntryTransactionLegacyRequest) Method() string {
	return r.method
}

func (r CreateOrderEntryTransactionLegacyRequest) NewContent() CreateOrderEntryTransactionLegacyRequestContent {
	content := CreateOrderEntryTransactionLegacyRequestContent{}
	return content
}

type CreateOrderEntryTransactionLegacyRequestContent struct {
	Function struct {
		ControlID string `xml:"controlid,attr"`

		CreateSOTransaction struct {
			XMLName         xml.Name `xml:"create_sotransaction"`
			TransactionType string   `xml:"transactiontype"`
			DateCreated     struct {
				Year  string `xml:"year"`
				Month string `xml:"month"`
				Day   string `xml:"day"`
			} `xml:"datecreated"`
			CustomerID    string                     `xml:"customerid"`
			DocumentNo    string                     `xml:"documentno,omitempty"`
			ReferenceNo   string                     `xml:"referenceno,omitempty"`
			PONumber      string                     `xml:"ponumber,omitempty"`
			TermName      string                     `xml:"termname,omitempty"`
			Message       string                     `xml:"message,omitempty"`
			ExternalID    string                     `xml:"externalid,omitempty"`
			BaseCurr      string                     `xml:"basecurr,omitempty"`
			Currency      string                     `xml:"currency,omitempty"`
			ExchRateType  string                     `xml:"exchratetype,omitempty"`
			State         string                     `xml:"state,omitempty"`
			SupDocID      string                     `xml:"supdocid,omitempty"`
			TaxSolutionID string                     `xml:"taxsolutionid,omitempty"`
			Items         OrderEntryTransactionItems `xml:"sotransitems>sotransitem"`
		} `xml:"create_sotransaction"`
	} `xml:"function"`
}

type OrderEntryTransactionItems []OrderEntryTransactionItem

type OrderEntryTransactionItem struct {
	ItemID          string `xml:"itemid"`
	ItemDesc        string `xml:"itemdesc,omitempty"`
	Quantity        Number `xml:"quantity"`
	Unit            string `xml:"unit,omitempty"`
	Price           Number `xml:"price"`
	DiscountPercent string `xml:"discountpercent,omitempty"`
	LocationID      string `xml:"locationid,omitempty"`
	DepartmentID    string `xml:"departmentid,omitempty"`
	Memo            string `xml:"memo,omitempty"`
	ProjectID       string `xml:"projectid,omitempty"`
	CustomerID      string `xml:"customerid,omitempty"`
	ClassID         string `xml:"classid,omitempty"`
	WarehouseID     string `xml:"warehouseid,omitempty"`
	Taxable         *bool  `xml:"taxable,omitempty"`
	// LineSubTotals are the VAT tax overrides. To override a single tax value
	// you must provide overrides for all taxes on the line item.
	LineSubTotals OrderEntryLineSubTotals `xml:"linesubtotals,omitempty"`
}

// OrderEntryLineSubTotals is omitted entirely when empty, Intacct's schema
// rejects an empty linesubtotals element.
type OrderEntryLineSubTotals []OrderEntryLineSubTotal

func (l OrderEntryLineSubTotals) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(l) == 0 {
		return nil
	}
	typ := struct {
		XMLName xml.Name                 `xml:"linesubtotals"`
		Items   []OrderEntryLineSubTotal `xml:"linesubtotal"`
	}{Items: l}
	return e.Encode(typ)
}

type OrderEntryLineSubTotal struct {
	TrxTax           Number `xml:"trx_tax"`
	OverrideDetailID string `xml:"overridedetailid"`
}

type CreateOrderEntryTransactionLegacyRequestBody struct {
	Request
}

func (r CreateOrderEntryTransactionLegacyRequestBody) Content() *CreateOrderEntryTransactionLegacyRequestContent {
	data, ok := r.Operation.Content.(CreateOrderEntryTransactionLegacyRequestContent)
	if ok {
		return &data
	}
	return &CreateOrderEntryTransactionLegacyRequestContent{}
}

func (r *CreateOrderEntryTransactionLegacyRequestBody) SetContent(content CreateOrderEntryTransactionLegacyRequestContent) {
	r.Operation.Content = content
}

func (r *CreateOrderEntryTransactionLegacyRequest) NewRequestBody() CreateOrderEntryTransactionLegacyRequestBody {
	body := CreateOrderEntryTransactionLegacyRequestBody{
		Request: NewRequest(),
	}
	body.Operation.Content = r.NewContent()
	return body
}

func (r *CreateOrderEntryTransactionLegacyRequest) SetRequestBody(body CreateOrderEntryTransactionLegacyRequestBody) {
	r.requestBody = body
}

func (r *CreateOrderEntryTransactionLegacyRequest) RequestBody() *CreateOrderEntryTransactionLegacyRequestBody {
	return &r.requestBody
}

func (r *CreateOrderEntryTransactionLegacyRequest) NewResponseBody() *CreateOrderEntryTransactionLegacyResponseBody {
	body := &CreateOrderEntryTransactionLegacyResponseBody{
		Response: NewResponse(),
	}
	return body
}

type CreateOrderEntryTransactionLegacyResponseBody struct {
	Response
}

func (r *CreateOrderEntryTransactionLegacyRequest) URL() url.URL {
	return r.client.GetEndpointURL("", r.PathParams())
}

func (r *CreateOrderEntryTransactionLegacyRequest) Do() (CreateOrderEntryTransactionLegacyResponseBody, error) {
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
