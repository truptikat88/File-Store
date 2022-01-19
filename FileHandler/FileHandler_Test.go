package FileHandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CreateFile(t *testing.T) {
	file := Files{
		Name:    "File111",
		Content: "File created from Unit Test case",
	}

	body, _ := json.Marshal(file)

	_, request := initializeHTTPRequest("POST", body, "")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, request)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func Test_GetFile(t *testing.T) {
	file := Files{
		Name:    "File111",
		Content: "File created from Unit Test case",
	}

	body, _ := json.Marshal(file)

	_, request := initializeHTTPRequest("POST", body, "")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, request)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func initializeHTTPRequest(httpMethod string, body []byte, queryParam string) (*httptest.ResponseRecorder, *http.Request) {

	w := httptest.NewRecorder()
	var r *http.Request
	switch {
	case body == nil && queryParam == "":
		r = httptest.NewRequest(httpMethod, "http://localhost", nil)
	case body == nil && queryParam != "":
		url := "http://localhost?" + queryParam
		r = httptest.NewRequest(httpMethod, url, nil)
	default:
		r = httptest.NewRequest(httpMethod, "http://localhost", bytes.NewReader(body))
	}

	return w, r
}
