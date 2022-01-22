package client

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// HTTP transport
type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func NewFakeClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestGetWithRoundTripper_Success(t *testing.T) {
	client := NewFakeClient(func(req *http.Request) (*http.Response, error) {
		//return the response we want
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
		}, nil
	})
	api := clientCall{*client}
	url := "https://twitter.com/fillpackart"
	body, err := api.GetValue(url)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.EqualValues(t, http.StatusOK, body.StatusCode)
}

func TestGetWithRoundTripper_No_Match(t *testing.T) {
	client := NewFakeClient(func(req *http.Request) (*http.Response, error) {
		//return the response we want
		return &http.Response{
			StatusCode: 404,               //the real api status code may be 404, 422, 500.
			Header:     make(http.Header), // Must be set to non-nil value or it panics
		}, nil
	})
	api := clientCall{*client}
	url := "https://twitter.com/truecoder34" // user that is not found
	body, err := api.GetValue(url)
	assert.Nil(t, err)
	assert.NotNil(t, body)
	assert.EqualValues(t, http.StatusNotFound, body.StatusCode)
}

func TestGetWithRoundTripper_Failure(t *testing.T) {
	client := NewFakeClient(func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("we couldn't access the url provided") //the response we want
	})
	api := clientCall{*client}
	url := "https://fake/truecoder34" // an invalid url
	body, err := api.GetValue(url)
	assert.NotNil(t, err)
	assert.Nil(t, body)
	assert.EqualValues(t, "Get https://fake/truecoder34: we couldn't access the url provided", err.Error())
}
