package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/wolfeidau/assembly/assembly"
	"github.com/wolfeidau/assembly/datastore"
)

func init() {
	serveMux.Handle("/", Handler(nil))
}

var (
	serveMux   = http.NewServeMux()
	httpClient = http.Client{Transport: (*muxTransport)(serveMux)}
	apiClient  = assembly.NewClient(&httpClient)
)

func setup() {
	store = datastore.NewMockDatastore()
}

type muxTransport http.ServeMux

// RoundTrip is a custom http.RoundTripper for testing API requests/responses.
// It intercepts all HTTP requests during testing to serve up a local/internal
// response instead of dialing out to the Host specified in the Client's BaseURL.
func (t *muxTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rw := httptest.NewRecorder()
	rw.Body = new(bytes.Buffer)
	(*http.ServeMux)(t).ServeHTTP(rw, req)
	return &http.Response{
		StatusCode:    rw.Code,
		Status:        http.StatusText(rw.Code),
		Header:        rw.HeaderMap,
		Body:          ioutil.NopCloser(rw.Body),
		ContentLength: int64(rw.Body.Len()),
		Request:       req,
	}, nil
}
