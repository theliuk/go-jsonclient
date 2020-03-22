package jsonclient

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckResponse(t *testing.T) {
	t.Run("return error with a status code more than 299", func(t *testing.T) {
		body := ioutil.NopCloser(strings.NewReader(`{"message":"error"}`))

		resp := &http.Response{
			StatusCode: 300,
			Body:       body,
			Request: &http.Request{
				Method: "METHOD",
				URL:    &url.URL{Path: "/request-url"},
			},
		}
		err := checkResponse(resp)
		require.EqualError(t, err, `METHOD /request-url: 300 - {"message":"error"}`, "error checking response")
	})

	t.Run("return error with a status code less than 200", func(t *testing.T) {
		body := ioutil.NopCloser(strings.NewReader(`{"message":"error"}`))

		resp := &http.Response{
			StatusCode: 199,
			Body:       body,
			Request: &http.Request{
				Method: "METHOD",
				URL:    &url.URL{Path: "/request-url"},
			},
		}
		err := checkResponse(resp)
		require.EqualError(t, err, `METHOD /request-url: 199 - {"message":"error"}`, "error checking response")
	})

	t.Run("return nil with status code 200", func(t *testing.T) {
		body := ioutil.NopCloser(strings.NewReader(`{"message":"error"}`))

		resp := &http.Response{
			StatusCode: 200,
			Body:       body,
			Request: &http.Request{
				Method: "METHOD",
				URL:    &url.URL{Path: "/request-url"},
			},
		}
		err := checkResponse(resp)
		require.NoError(t, err, "error checking response")
	})

	t.Run("return nil with status code 299", func(t *testing.T) {
		body := ioutil.NopCloser(strings.NewReader(`{"message":"error"}`))

		resp := &http.Response{
			StatusCode: 299,
			Body:       body,
			Request: &http.Request{
				Method: "METHOD",
				URL:    &url.URL{Path: "/request-url"},
			},
		}
		err := checkResponse(resp)
		require.NoError(t, err, "error checking response")
	})
}