package ohdear

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// BaseClient interface describe an oh-dear API implementation.
type BaseClient interface {
	NewAPIRequest(method, uri string, body interface{}) (req *http.Request, err error)
	Do(ctx context.Context, req *http.Request) (res *Response, err error)
}

type srv struct {
	client *Client
}

// Client is the main API caller.
type Client struct {
	BaseURL *url.URL
	client  *http.Client
	common  srv // Reuse a single struct instead of allocating one for each service on the heap.
	token   string
	// Services
	Sites *SitesSrv
}

// NewAPIRequest is a wrapper around the http.NewRequest function.
//
// It will setup the authentication headers/parameters according to the client config.
func (c *Client) NewAPIRequest(method string, uri string, body interface{}) (req *http.Request, err error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, ErrInvalidBaseURL
	}

	u, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add(AuthHeader, strings.Join([]string{TokenType, c.token}, " "))
	req.Header.Set("Content-Type", ContentExchangeType)
	req.Header.Set("Accept", ContentExchangeType)

	return
}

// Do sends an API request and returns the API response or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Response  wraps the standard http.Response returned from oh-dear
// and provides non-blocking access to the request content.
type Response struct {
	*http.Response
	content []byte
}

// response constructor.
// it takes a core http response pointer and transforms it into
// the enriched response struct.
func newResponse(r *http.Response) *Response {
	var res Response
	c, err := ioutil.ReadAll(r.Body)
	if err == nil {
		res.content = c
	}
	json.NewDecoder(r.Body).Decode(&res)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(c))
	res.Response = r
	return &res
}

// NewClient returns a new Oh-Dear HTTP API client.
// You can pass a previously build http client, if none is provided then
// http.DefaultClient will be used.
//
// NewClient will lookup the environment for values to assign to the
// API token (`OHDEAR_API_TOKEN`) to be used as authentication.
func NewClient(baseClient *http.Client, baseURL, apiToken string) (dear *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	var u *url.URL
	{
		if baseURL != "" {
			u, err = url.Parse(baseURL)
			if err != nil {
				return
			}
		} else {
			u, _ = url.Parse(BaseURL)
		}
	}

	dear = &Client{
		BaseURL: u,
		client:  baseClient,
	}

	dear.common.client = dear

	// services for resources
	dear.Sites = (*SitesSrv)(&dear.common)

	// Parse authorization from environment
	// or user provided string.
	if tkn, ok := os.LookupEnv(APITokenEnv); ok {
		dear.token = tkn
	} else {
		if apiToken == "" {
			return nil, ErrEmptyAPIToken
		}

		dear.token = apiToken
	}

	return
}
