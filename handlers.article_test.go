package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that a GET request to the home page returnss the home page with
// the HTTp code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// test that the http status code is 200
		statusOk := w.Code == http.StatusOK

		// Test that the home page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "<title> Homepage </title>") > 0

		return statusOk && pageOk
	})
}
