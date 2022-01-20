package filehandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func Test_addSum(t *testing.T) {
// 	result := addSum()

// 	if result != 4 {
// 		t.Error("result should be 4")
// 	}

// }

func Test_AddFile(t *testing.T) {
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

func Test_AddFiles(t *testing.T) {
	filesData := Items{
		Elements: []Files{
			{
				Name:    "File99",
				Content: "File99 created from Unit Test case",
			},
			{
				Name:    "File98",
				Content: "File98 created from Unit Test case",
			},
		},
	}

	body, _ := json.Marshal(filesData)

	_, request := initializeHTTPRequest("POST", body, "")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddFiles)

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
	FileName := "File1"

	_, request := initializeHTTPRequest("GET", nil, FileName)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, request)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func Test_GetFiles(t *testing.T) {

	_, request := initializeHTTPRequest("GET", nil, "")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetFiles)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, request)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func Test_RemoveFile(t *testing.T) {
	FileName := "File1"

	_, request := initializeHTTPRequest("DELETE", nil, FileName)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RemoveFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, request)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func Test_UpdateFile(t *testing.T) {
	file := Files{
		Name:    "File111",
		Content: "File created from Unit Test case",
	}

	body, _ := json.Marshal(file)

	_, request := initializeHTTPRequest("PATCH", body, "")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, request)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_CountWord(t *testing.T) {

	_, request := initializeHTTPRequest("GET", nil, "")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CountWord)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, request)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_FindWordCount(t *testing.T) {

	_, request := initializeHTTPRequest("GET", nil, "")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FindWordCount)

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
