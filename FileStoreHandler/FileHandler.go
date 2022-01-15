package fileStoreHandler

import (
	"fmt"
	"net/http"
)

func AddFile(r http.ResponseWriter, w *http.Request) {
	fmt.Println("Inside addFile function...")

	// strFileName := "File"
	// //tempData := "heloo, this is my file!"
	// data := []byte("adhdu")

	// os.WriteFile(DirectoryName+"/"+strFileName+".txt", data, 0644)
}
