package ohdear

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/VictorAvelar/goh-dear/testdata"
)

func TestSitesSrv_Get(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		tearDown()
		unsetEnv()
	}()

	cases := []struct {
		id       uint
		name     string
		status   int
		respBody string
		wantErr  bool
		err      error
	}{
		{
			1,
			"successful request/response cycle",
			http.StatusOK,
			testdata.SingleSiteResponse,
			false,
			nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tMux.HandleFunc(fmt.Sprintf("/sites/%d", c.id), func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, fmt.Sprintf("%s %s", TokenType, testTkn))
				testMethod(t, r, http.MethodGet)

				w.WriteHeader(http.StatusOK)
				_, _ = fmt.Fprint(w, c.respBody)
			})

			got, err := tClient.Sites.Get(c.id)
			if err != nil {
				if c.wantErr {
					assert.EqualError(t, err, c.err.Error())
				} else {
					t.Fatal(err)
				}
			}

			assert.Equal(t, c.id, got.ID)
			fmt.Printf("%+v", got)
		})
	}
}
