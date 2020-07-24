# ohdear
--
    import "github.com/VictorAvelar/goh-dear/ohdear"

Package ohdear is a Golang SDK to interact with oh-dear REST api.

The Oh Dear API lets you configure everything about our application through a
simple, structured Application Programming Interface (API). Everything you see
in your dashboard can be controlled with the API. And as a bonus, all changes
you make with the API will be visible in realtime on your dashboard.

The full api documentation can be found in
https://ohdear.app/docs/general/welcome

When instantiating a new client, you can provide and API token using
OHDEAR_API_TOKEN which is the default environment variable and it will be looked
up by the library automatically. This is strongly recommended as it is the most
secure way to deal with your key.

If the library cannot resolve your API token from the environment you can
provide it when instantiating a new ohdear client using the NewClient
constructor.

## Usage

```go
const (
	BaseURL             string = "https://ohdear.app/api/"
	TokenType           string = "Bearer"
	APITokenEnv         string = "OHDEAR_API_TOKEN"
	ContentExchangeType string = "application/json"
	AuthHeader          string = "Authorization"
)
```
Oh-dear package level constants

```go
var (
	ErrEmtpyAPIToken  error = fmt.Errorf("your api token is empty, please provide a non-empty token")
	ErrInvalidBaseURL error = fmt.Errorf("your base url must contain a trailing slash")
)
```
Oh-dear package level errors

#### func  CheckResponse

```go
func CheckResponse(r *http.Response) error
```
CheckResponse checks the API response for errors, and returns them if present. A
response is considered an error if it has a status code outside the 200 range.
API's error responses must have either no response body, or a JSON response
body.

#### type BaseClient

```go
type BaseClient interface {
	NewAPIRequest(method, uri string, body interface{}) (req *http.Request, err error)
	Do(ctx context.Context, req *http.Request) (res *Response, err error)
}
```

BaseClient interface describe an oh-dear API implementation.

#### type Client

```go
type Client struct {
	BaseURL *url.URL
}
```

Client is the main API caller.

#### func  NewClient

```go
func NewClient(baseClient *http.Client, baseURL, apiToken string) (dear *Client, err error)
```
NewClient returns a new Oh-Dear HTTP API client. You can pass a previously build
http client, if none is provided then http.DefaultClient will be used.

NewClient will lookup the environment for values to assign to the API token
(`OHDEAR_API_TOKEN`) to be used as authentication.

#### func (*Client) Do

```go
func (c *Client) Do(req *http.Request) (*Response, error)
```
Do sends an API request and returns the API response or returned as an error if
an API error has occurred.

#### func (*Client) NewAPIRequest

```go
func (c *Client) NewAPIRequest(method string, uri string, body interface{}) (req *http.Request, err error)
```
NewAPIRequest is a wrapper around the http.NewRequest function.

It will setup the authentication headers/parameters according to the client
config.

#### type Error

```go
type Error struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Response *http.Response `json:"response"` // the full response that produced the error
}
```

Error maps a standard error to a more useful data structure which is enriched
with the failing request pointer.

#### func (*Error) Error

```go
func (e *Error) Error() string
```
Error function complies with the error interface

#### type Response

```go
type Response struct {
	*http.Response
}
```

Response wraps the standard http.Response returned from oh-dear and provides
non-blocking access to the request content.
