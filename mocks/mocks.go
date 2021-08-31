package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type RequesterIFace interface {
	Do(r *http.Request) (*http.Response, error)
}

var Requester RequesterIFace // <--- this is of type interface

type MockRequesterUnauthorized struct{}

func (m *MockRequesterUnauthorized) Do(r *http.Request) (*http.Response, error) {
	res := &http.Response{
		StatusCode: 401,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte("Unauthorized"))),
	}
	return res, nil
}
