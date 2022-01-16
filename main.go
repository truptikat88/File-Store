package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
)

const (
	DirectoryName = "FileStore"
)

var finalCount = map[string]int{}

type Files struct {
	Name    string `json:"fileName"`
	Content string `json:"fileContent"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page of FileStore....")
}

func addFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside addFile function...")

	var file Files

	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//strFileName := file.Name
	data := []byte(file.Content)

	err = os.WriteFile(DirectoryName+"/"+file.Name+".txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	strFileName := file.Name + " file got created successfully"
	fmt.Fprintf(w, strFileName)

	//fmt.Fprintf(w, "File is created successfully...")

}

func removeFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside removeFile function...")

	// vars := mux.Vars(r)
	// fileToRemove, ok := vars["file"]

	// keys, ok := r.URL.Query()["file"]

	// if !ok || len(keys[0]) < 1 {
	// 	fmt.Println("File Name is missing in parameters...")
	// }

	// strFileName := keys[0]

	strFileName := "File3"

	fmt.Println(`Name of the file to remove := `, strFileName)

	e := os.Remove(DirectoryName + "/" + strFileName + ".txt")
	if e != nil {
		log.Fatal(e)
	}
}

func updateFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "inside update....")

	// Read Write Mode
	file, err := os.OpenFile(DirectoryName+"/"+"File1.txt", os.O_RDWR, 0644)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	data := []byte("updated the content")

	_, err = file.WriteAt(data, 0) // Write at 0 beginning
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside getFile function...")

	// Open the file and print the data to cmd line
	// f, _ := os.Open(DirectoryName + "\\file.txt")
	// // Create a new Scanner for the file.
	// scanner := bufio.NewScanner(f)
	// // Loop over all lines in the file and print them.
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// }

	fileBytes, err := ioutil.ReadFile(DirectoryName + "\\File.txt")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)

}

func getFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside getFiles function...")

	files, err := ioutil.ReadDir(DirectoryName + "\\")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Fprintf(w, file.Name(), file.ModTime(), file.Size())
		fmt.Fprintf(w, "\n")
	}
}

func countWord(w http.ResponseWriter, r *http.Request) {
	fh, err := os.OpenFile(DirectoryName+"/"+"File1.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Could not open file '%v': %v", "File1.txt", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(fh)
	counter := make(map[string]int)
	for {
		line, _ := reader.ReadString('\n')
		//fmt.Print(line)
		fields := strings.Fields(line)
		//fmt.Println(fields)
		for _, word := range fields {
			word = strings.ToLower(word)
			counter[word]++
		}
		if line == "" {
			break
		}
	}

	// for word, cnt := range counter {
	//     fmt.Printf("%v %v\n", word, cnt)
	// }

	words := make([]string, 0, len(counter))
	for word := range counter {
		words = append(words, word)
	}
	sort.Slice(words, func(i, j int) bool {
		return counter[words[i]] > counter[words[j]]
	})

	for _, word := range words {
		fmt.Println("\n", word, counter[word])
	}
}

func wordCount(rdr io.Reader, channel chan (string)) map[string]int {
	//counts := map[string]int{}

	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		channel <- word
		//	finalCount[word]++
	}

	return finalCount
}

func FindWordCount(w http.ResponseWriter, r *http.Request) {

	files, err := ioutil.ReadDir(DirectoryName + "\\")

	if err != nil {
		log.Fatal(err)
	}

	// A waitgroup to wait for all go-routines to finish.
	wg := sync.WaitGroup{}

	messages := make(chan string)
	done := make(chan (bool), 1)

	// Read all incoming words from the channel and add them to the dictionary.
	go func() {
		for word := range messages {
			finalCount[word]++
		}

		// Signal the main thread that all the words have entered the dictionary.
		done <- true
	}()

	//var finalCount map[string]int
	for _, file := range files {

		wg.Add(1)

		fmt.Println("\n ************ ", file.Name())
		file, fileOpenErr := os.Open(DirectoryName + "\\" + file.Name())
		//fmt.Println("File Name: ", file.Name())
		if fileOpenErr != nil {
			fmt.Println("Not able to open the file : ", file.Name())
		}

		go func(file io.Reader) {

			wordCount(bufio.NewReader(file), messages)

			//fmt.Println("\n Total number of words in file ", finalCount)

			wg.Done()
		}(file)
	}

	wg.Wait()
	fmt.Println("\n Total number of words in file ", finalCount, len(finalCount))

	type key struct {
		name  string
		count int
	}
	keys := make([]key, 0, len(finalCount))
	var tempData = key{}
	for word, count := range finalCount {
		tempData.name = word
		tempData.count = count
		keys = append(keys, tempData)
		//fmt.Println("\n ** ", key, val)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].count > keys[j].count
	})

	for i := 0; i <= 10; i++ {
		fmt.Println("\n Total number of words in file ", keys[i])
	}

	// fmt.Println("\n Total number of words in file ", keys)

}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/filesAdd", addFile)
	http.HandleFunc("/filesRemove", removeFile)
	http.HandleFunc("/filesUpdate", updateFile)
	http.HandleFunc("/filesGet", getFile)
	http.HandleFunc("/filesGetAll", getFiles)

	http.HandleFunc("/countWord", countWord)
	http.HandleFunc("/countWordFromAllFiles", FindWordCount)

	http.HandleFunc("/maxUsedWords", getFiles)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	//for postman requests
	HandleRequests()

	fmt.Println("Finished..")

}
