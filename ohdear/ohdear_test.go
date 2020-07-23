package ohdear

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testTkn string = "testing_token"

var (
	tMux    *http.ServeMux
	tServer *httptest.Server
	tClient *Client
)

func TestCheckResponse(t *testing.T) {
	cases := []struct {
		name     string
		wantErr  bool
		response *http.Response
	}{
		{
			"200 error nil",
			false,
			&http.Response{
				Status:     http.StatusText(http.StatusOK),
				StatusCode: http.StatusOK,
			},
		},
		{
			"300 error",
			true,
			&http.Response{
				Status:     http.StatusText(http.StatusPermanentRedirect),
				StatusCode: http.StatusPermanentRedirect,
			},
		},
		{
			"400 error",
			true,
			&http.Response{
				Status:     http.StatusText(http.StatusUnauthorized),
				StatusCode: http.StatusUnauthorized,
			},
		},
		{
			"500 error",
			true,
			&http.Response{
				Status:     http.StatusText(http.StatusInternalServerError),
				StatusCode: http.StatusInternalServerError,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			err := CheckResponse(c.response)

			if c.wantErr {
				assert.Contains(tt, err.Error(), c.response.Status)
			} else {
				assert.Nil(tt, err)
			}
		})
	}
}

// package test helpers
func setup() {
	tMux = http.NewServeMux()
	tServer = httptest.NewServer(tMux)
	tClient, _ = NewClient(nil, "", "")
	tClient.BaseURL, _ = url.Parse(tServer.URL + "/")
}

// defer after setup
func tearDown() {
	tServer.Close()
}

func setEnv() {
	os.Setenv(APITokenEnv, testTkn)
}

// defer after setEnv
func unsetEnv() {
	os.Unsetenv(APITokenEnv)
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request Body is %s, want %s", got, want)
	}
}

func testJSONMarshal(t *testing.T, v interface{}, want string) {
	j, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %v", v)
	}

	w := new(bytes.Buffer)
	err = json.Compact(w, []byte(want))
	if err != nil {
		t.Errorf("String is not valid json: %s", want)
	}

	if w.String() != string(j) {
		t.Errorf("json.Marshal(%q) returned %s, want %s", v, j, w)
	}

	// now go the other direction and make sure things unmarshal as expected
	u := reflect.ValueOf(v).Interface()
	if err := json.Unmarshal([]byte(want), u); err != nil {
		t.Errorf("Unable to unmarshal JSON for %v: %v", want, err)
	}

	if !reflect.DeepEqual(v, u) {
		t.Errorf("json.Unmarshal(%q) returned %s, want %s", want, u, v)
	}
}
