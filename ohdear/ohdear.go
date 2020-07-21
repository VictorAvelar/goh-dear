package ohdear

import (
	"fmt"
	"net/http"
)

// Oh-dear package level constants
const (
	BaseURL             string = "https://ohdear.app/api/"
	TokenType           string = "Bearer"
	APITokenEnv         string = "OHDEAR_API_TOKEN"
	ContentExchangeType string = "application/json"
	AuthHeader          string = "Authorization"
)

// Oh-dear package level errors
var (
	ErrEmtpyAPIToken  error = fmt.Errorf("your api token is empty, please provide a non-empty token")
	ErrInvalidBaseURL error = fmt.Errorf("your base url must contain a trailing slash")
)

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body.
func CheckResponse(r *http.Response) error {
	if r.StatusCode >= http.StatusMultipleChoices {
		return newError(r)
	}
	return nil
}

// Error maps a standard error to a more useful
// data structure which is enriched with the
// failing request pointer.
type Error struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Response *http.Response `json:"response"` // the full response that produced the error
}

// Error function complies with the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("response failed with status %v|%v", e.Code, e.Message)
}

// Error constructor
func newError(r *http.Response) *Error {
	var e Error
	e.Response = r
	e.Code = r.StatusCode
	e.Message = r.Status
	return &e
}