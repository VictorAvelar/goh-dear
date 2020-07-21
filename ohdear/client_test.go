package ohdear

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	type args struct {
		baseClient *http.Client
		baseURL    string
		apiToken   string
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "successful client initialization",
			args: args{
				baseClient: nil,
				baseURL:    "https://apiurl.example.com",
				apiToken:   "auth_token",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "successful client initialization using the default base url",
			args: args{
				baseClient: nil,
				baseURL:    "",
				apiToken:   "auth_token",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "client fails on bad url",
			args: args{
				baseClient: nil,
				baseURL:    ":",
				apiToken:   "auth_token",
			},
			wantErr: true,
			err:     errors.New("parse \":\": missing protocol scheme"),
		},
		{
			name: "client fails on empty api token provided",
			args: args{
				baseClient: nil,
				baseURL:    "",
				apiToken:   "",
			},
			wantErr: true,
			err:     ErrEmtpyAPIToken,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			got, err := NewClient(c.args.baseClient, c.args.baseURL, c.args.apiToken)

			if !c.wantErr {
				assert.Nil(tt, err)
				assert.Equal(tt, c.args.apiToken, got.token)
			} else if c.wantErr {
				assert.EqualError(tt, err, c.err.Error())
			}
		})
	}
}

func TestNewClient_TokenFromEnvironment(t *testing.T) {
	setEnv()
	defer func() {
		unsetEnv()
	}()

	got, err := NewClient(nil, "", "")

	assert.Nil(t, err)
	assert.Equal(t, testTkn, got.token)
}

func TestClient_NewAPIRequest(t *testing.T) {
	setEnv()
	defer tearDown()
	setup()
	defer tearDown()

	tClient.BaseURL, _ = url.Parse(tServer.URL + "/")

	b := []string{"hello", "bye"}
	inURL, outURL := "test", tServer.URL+"/test"
	inBody, outBody := b, `["hello","bye"]`+"\n"
	req, _ := tClient.NewAPIRequest("GET", inURL, inBody)

	testHeader(t, req, "Accept", ContentExchangeType)
	testHeader(t, req, AuthHeader, fmt.Sprintf("%s %s", TokenType, testTkn))
	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}
}

func TestClient_NewAPIRequest_ErrBadBaseURL(t *testing.T) {
	setup()
	defer tearDown()
	c, err := NewClient(nil, tServer.URL, testTkn)
	if err != nil {
		t.Error(err)
	}

	_, err = c.NewAPIRequest(http.MethodGet, "sites", nil)
	if err != nil {
		assert.EqualError(t, ErrInvalidBaseURL, err.Error())
	} else {
		t.Error("nil received when expecting an error")
	}

}

func TestClient_NewAPIRequest_ErrParsingURI(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		tearDown()
		unsetEnv()
	}()

	_, err := tClient.NewAPIRequest(http.MethodGet, ":", nil)

	if err != nil {
		assert.EqualError(t, err, "parse \":\": missing protocol scheme")
	} else {
		t.Error("nil received when expecting an error")
	}
}

func TestClient_NewAPIRequest_NativeHTTPErr(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		tearDown()
		unsetEnv()
	}()

	_, err := tClient.NewAPIRequest("\\\\\\", "test", nil)

	if err != nil {
		assert.EqualError(t, err, "net/http: invalid method \"\\\\\\\\\\\\\"")
	} else {
		t.Fail()
	}
}

func TestClient_NewAPIRequest_ErrJSONBodyMarshaling(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		tearDown()
		unsetEnv()
	}()

	_, err := tClient.NewAPIRequest(http.MethodGet, "test", make(chan int))

	if err != nil {
		assert.EqualError(t, err, "json: unsupported type: chan int")
	} else {
		t.Fail()
	}
}
