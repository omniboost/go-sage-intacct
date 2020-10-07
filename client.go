package intacct

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"text/template"

	"github.com/gofrs/uuid"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-sage-intacct/" + libraryVersion
	mediaType      = "application/xml"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "api.intacct.com",
		Path:   "/ia/xml/xmlgw.phtml",
	}
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client, senderID, senderPassword, userID, userPassword, companyID string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetSenderID(senderID)
	client.SetSenderPassword(senderPassword)
	client.SetUserID(userID)
	client.SetUserPassword(userPassword)
	client.SetCompanyID(companyID)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	senderID       string
	senderPassword string
	userID         string
	userPassword   string
	companyID      string

	sessionID string

	// User agent for client
	userAgent string

	mediaType string
	charset   string

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c Client) SenderID() string {
	return c.senderID
}

func (c *Client) SetSenderID(senderID string) {
	c.senderID = senderID
}

func (c Client) SenderPassword() string {
	return c.senderPassword
}

func (c *Client) SetSenderPassword(senderPassword string) {
	c.senderPassword = senderPassword
}

func (c Client) UserID() string {
	return c.userID
}

func (c *Client) SetUserID(userID string) {
	c.userID = userID
}

func (c Client) UserPassword() string {
	return c.userPassword
}

func (c *Client) SetUserPassword(userPassword string) {
	c.userPassword = userPassword
}

func (c Client) CompanyID() string {
	return c.companyID
}

func (c *Client) SetCompanyID(companyID string) {
	c.companyID = companyID
}

func (c Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) GetEndpointURL(path string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()
	clientURL.Path = clientURL.Path + path

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, method string, URL url.URL, body Request) (*http.Request, error) {
	body.Control.SenderID = c.SenderID()
	body.Control.SenderPassword = c.SenderPassword()
	body.Control.ControlID = c.GenerateControlID()

	// convert body struct to json
	buf := new(bytes.Buffer)
	err := xml.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}

	// create new http request
	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// set other headers
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())

	return req, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody *Response) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	err = c.Unmarshal(httpResp.Body, &responseBody)
	if err != nil {
		return httpResp, err
	}

	// if len(errorResponse.Messages) > 0 {
	// 	return httpResp, errorResponse
	// }

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	wg := sync.WaitGroup{}
	wg.Add(len(vv))
	errs := []error{}
	writers := make([]io.Writer, len(vv))

	for i, v := range vv {
		pr, pw := io.Pipe()
		writers[i] = pw

		go func(i int, v interface{}, pr *io.PipeReader, pw *io.PipeWriter) {
			dec := xml.NewDecoder(pr)
			err := dec.Decode(v)
			if err != nil {
				errs = append(errs, err)
			}

			// mark routine as done
			wg.Done()

			// Drain reader
			io.Copy(ioutil.Discard, pr)

			// close reader
			// pr.CloseWithError(err)
			pr.Close()
		}(i, v, pr, pw)
	}

	// copy the data in a multiwriter
	mw := io.MultiWriter(writers...)
	_, err := io.Copy(mw, r)
	if err != nil {
		return err
	}

	wg.Wait()
	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}
	return nil
}

func (c *Client) GenerateControlID() string {
	return uuid.Must(uuid.NewV4()).String()
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r)
	if err != nil {
		if string(data) != "" {
			errorResponse.Errors = append(errorResponse.Errors, errors.New(string(data)))
		} else {
			errorResponse.Errors = append(errorResponse.Errors, errors.New(r.Status))
		}
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		errorResponse.Errors = append(errorResponse.Errors, err)
		return errorResponse
	}

	return errorResponse
}

func (c *Client) SessionID() (string, error) {
	if c.sessionID == "" {
		var err error
		c.sessionID, err = c.NewSessionID()
		if err != nil {
			return "", err
		}
	}

	return c.sessionID, nil
}

func (c *Client) NewSessionID() (string, error) {
	req := c.NewGetAPISessionRequest()
	resp, err := req.Do()
	if err != nil {
		return "", err
	}

	return resp.Data().API.SessionID, nil

}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Errors []error
}

type Error struct {
	Message       string `json:"message"`
	MessageDetail string `json:"MessageDetail"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.MessageDetail)
}

func (r *ErrorResponse) UnmarshalJSON(data []byte) error {
	e := Error{}
	err := json.Unmarshal(data, &e)
	if err != nil {
		return err
	}

	r.Errors = append(r.Errors, e)

	return nil
}

func (r ErrorResponse) Error() string {
	if len(r.Errors) > 0 {
		str := []string{}
		for _, err := range r.Errors {
			str = append(str, err.Error())
		}
		return strings.Join(str, ", ")
	}

	return r.Errors[0].Error()
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

type PathParams interface {
	Params() map[string]string
}
