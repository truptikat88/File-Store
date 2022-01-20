package FileHandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func Test_GetFile(t *testing.T) {
	req, err := http.NewRequest("GET", "/fileGet?fileName=File", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	data := "Hi this is file store"
	if rr.Body.String() != data {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), data)
	}
}

func Test_GetFileNotExists(t *testing.T) {
	req, err := http.NewRequest("GET", "/fileGet?fileName=File99", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// // Check the status code is what we expect.
	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }

	// Check the response body is what we expect.
	data := "File Not Found"
	if rr.Body.String() != data {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), data)
	}
}

func Test_GetFiles(t *testing.T) {

	req, err := http.NewRequest("GET", "/filesGetAll", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetFiles)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func Test_FileAdd(t *testing.T) {

	file := Files{
		Name:    "File111",
		Content: "File created from Unit Test case",
	}

	body, _ := json.Marshal(file)

	req, err := http.NewRequest("GET", "/fileAdd", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	data := "File111 file got created successfully"
	if rr.Body.String() != data {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), data)
	}

	//clear the file
	os.Remove(filepath.Join(DirectoryName, "File111.txt"))

}

func Test_FileAddAlreadyExists(t *testing.T) {

	file := Files{
		Name:    "File",
		Content: "File created from Unit Test case",
	}

	body, _ := json.Marshal(file)

	req, err := http.NewRequest("GET", "/fileAdd", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	_ = http.HandlerFunc(AddFile)

	handlerAgain := http.HandlerFunc(AddFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handlerAgain.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	data := "File already exists, please choose different file name"
	if rr.Body.String() != data {
		t.Errorf("handler returned unexpected body: got %v \n want %v",
			rr.Body.String(), data)
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

	req, err := http.NewRequest("GET", "/fileAdd", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddFiles)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	data := "File99 file got created successfullyFile98 file got created successfully"
	if rr.Body.String() != data {
		t.Errorf("handler returned unexpected body: got %v \n want %v",
			rr.Body.String(), data)
	}

	//clear the file
	os.Remove(filepath.Join(DirectoryName, "File99.txt"))

	//clear the file
	os.Remove(filepath.Join(DirectoryName, "File98.txt"))

}

func Test_RemoveFile(t *testing.T) {

	//create file to remove
	data := []byte("temp file created to to remov")
	os.WriteFile(filepath.Join(DirectoryName, "FileToRemove.txt"), data, 0644)

	req, err := http.NewRequest("GET", "/filesRemove?fileName=FileToRemove", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RemoveFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_UpdateFile(t *testing.T) {
	//create file to remove
	data := []byte("temp file created to to remov")
	os.WriteFile(filepath.Join(DirectoryName, "FileToUpdate.txt"), data, 0644)

	file := Files{
		Name:    "FileToUpdate",
		Content: "File updated from Unit Test case",
	}

	body, _ := json.Marshal(file)

	req, err := http.NewRequest("GET", "/filesUpdate", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateFile)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	dataResult := "File updated from Unit Test case"
	if rr.Body.String() != dataResult {
		t.Errorf("handler returned unexpected body: got %v \n want %v",
			rr.Body.String(), dataResult)
	}

	//clear the file
	os.Remove(filepath.Join(DirectoryName, "FileToUpdate.txt"))

}

func Test_CountWord(t *testing.T) {

	req, err := http.NewRequest("GET", "/countWord?fileName=File", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CountWord)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_FindWordCount(t *testing.T) {

	req, err := http.NewRequest("GET", "/countWordFromAllFiles", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FindWordCount)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
