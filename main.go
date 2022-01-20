package main

import (
	"fmt"
	"log"
	"net/http"

	fileHandler "github.com/truptikat88/File-Store/filehandler"
)

func HandleRequests() {
	http.HandleFunc("/", fileHandler.HomePage)
	http.HandleFunc("/countWord", fileHandler.CountWord)
	http.HandleFunc("/countWordFromAllFiles", fileHandler.FindWordCount)
	http.HandleFunc("/fileAdd", fileHandler.AddFile)
	http.HandleFunc("/filesAdd", fileHandler.AddFiles)
	http.HandleFunc("/filesRemove", fileHandler.RemoveFile)
	http.HandleFunc("/filesUpdate", fileHandler.UpdateFile)
	http.HandleFunc("/fileGet", fileHandler.GetFile)
	http.HandleFunc("/filesGetAll", fileHandler.GetFiles)

	log.Printf("starting service on 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	log.Printf("inside main...")

	//for postman requests
	HandleRequests()

	fmt.Println("Finished..")

}
