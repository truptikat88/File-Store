package FileHandler

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

type Items struct {
	Elements []Files `json:"items"`
}

func addSum() int {

	a := 2 + 2

	return a

}

func GetFile(w http.ResponseWriter, r *http.Request) {
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

	// path := r.URL.Path
	// fmt.Println(path)

	fileName := r.URL.Query().Get("fileName")
	if len(fileName) == 0 {
		fmt.Println("filters not present")
	}
	//fmt.Println(fileName)

	fileBytes, err := ioutil.ReadFile(filepath.Join(DirectoryName, fileName+".txt"))
	if err != nil {
		data := []byte("File Not Found")
		w.Write(data)
		w.WriteHeader(http.StatusBadRequest)
		//panic(err)
	}

	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page of FileStore....")
}

func AddFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside addFile function...")

	var items Items

	err := json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, file := range items.Elements {
		//strFileName := file.Name
		data := []byte(file.Content)

		//err = os.WriteFile(filepath.Join() DirectoryName+"/"+file.Name+".txt", data, 0644)
		err = os.WriteFile(filepath.Join(DirectoryName, file.Name+".txt"), data, 0644)

		if err != nil {
			log.Fatal(err)
		}

		strFileName := file.Name + " file got created successfully"
		fmt.Fprintf(w, strFileName)
	}

	//fmt.Fprintf(w, "File is created successfully...")
	w.WriteHeader(http.StatusOK)

}

func AddFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside addFile function...")

	var file Files

	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//strFileName := file.Name
	data := []byte(file.Content)

	//err = os.WriteFile(DirectoryName+"/"+file.Name+".txt", data, 0644)
	err = os.WriteFile(filepath.Join(DirectoryName, file.Name+".txt"), data, 0644)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	strFileName := file.Name + " file got created successfully"
	fmt.Fprintf(w, strFileName)

	//fmt.Fprintf(w, "File is created successfully...")
	w.WriteHeader(http.StatusOK)

}

func RemoveFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside removeFile function...")

	fileName := r.URL.Query().Get("fileName")
	if len(fileName) == 0 {
		fmt.Println("filters not present")
	}

	// _, err := ioutil.ReadFile(filepath.Join(DirectoryName, fileName+".txt"))
	// if err != nil {
	// 	data := []byte("File Not Found")
	// 	w.Write(data)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	e := os.Remove(filepath.Join(DirectoryName, fileName+".txt"))
	if e != nil {
		data := []byte("File can not be removed")
		w.Write(data)
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(e)
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "inside update....")

	var file Files

	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//strFileName := file.Name
	dataToUpdate := []byte(file.Content)

	// Read Write Mode
	fileToUpdate, err := os.OpenFile(filepath.Join(DirectoryName, file.Name+".txt"), os.O_RDWR, 0644)

	if err != nil {
		data := []byte("File not found")
		w.Write(data)
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("failed opening file: %s", err)
	}
	defer fileToUpdate.Close()

	data := []byte(dataToUpdate)

	// _, err = fileToUpdate.WriteAt(data, 0) // Write at 0 beginning
	// if err != nil {
	// 	log.Fatalf("failed writing to file: %s", err)
	// }

	_, err = fileToUpdate.WriteString(file.Content) // Write at 0 beginning
	if err != nil {
		data := []byte("File cannot be updated")
		w.Write(data)
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("failed writing to file: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}

// func getFile(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Inside getFile function...")

// 	// Open the file and print the data to cmd line
// 	// f, _ := os.Open(DirectoryName + "\\file.txt")
// 	// // Create a new Scanner for the file.
// 	// scanner := bufio.NewScanner(f)
// 	// // Loop over all lines in the file and print them.
// 	// for scanner.Scan() {
// 	// 	line := scanner.Text()
// 	// 	fmt.Println(line)
// 	// }

// 	// path := r.URL.Path
// 	// fmt.Println(path)

// 	fileName := r.URL.Query().Get("fileName")
// 	if len(fileName) == 0 {
// 		fmt.Println("filters not present")
// 	}
// 	//fmt.Println(fileName)

// 	fileBytes, err := ioutil.ReadFile(filepath.Join(DirectoryName, fileName+".txt"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	//w.Header().Set("Content-Type", "application/octet-stream")
// 	w.Write(fileBytes)

// }

func GetFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside getFiles function...")

	files, err := ioutil.ReadDir(filepath.Join(DirectoryName))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}

	if len(files) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	for _, file := range files {
		fmt.Fprintf(w, file.Name(), file.ModTime(), file.Size())
		fmt.Fprintf(w, "\n")
	}

	w.WriteHeader(http.StatusOK)

}

func CountWord(w http.ResponseWriter, r *http.Request) {

	fileName := r.URL.Query().Get("fileName")
	if len(fileName) == 0 {
		fmt.Println("filters not present")
	}

	fh, err := os.OpenFile(filepath.Join(DirectoryName, fileName+".txt"), os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Could not open file '%v': %v", "File1.txt", err)
		w.WriteHeader(http.StatusBadRequest)
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

	b := new(bytes.Buffer)
	for _, word := range words {
		fmt.Fprintf(b, "word = %s , ocurrences = %d \n", word, counter[word])
	}

	result := []byte(b.String())

	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func WordCount(rdr io.Reader, channel chan (string)) map[string]int {
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

	fmt.Println("\n inside findword... ")

	files, err := ioutil.ReadDir(filepath.Join(DirectoryName))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
		file, fileOpenErr := os.Open(filepath.Join(DirectoryName, file.Name()))
		//fmt.Println("File Name: ", file.Name())
		if fileOpenErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Not able to open the file : ", file.Name())
		}

		go func(file io.Reader) {

			WordCount(bufio.NewReader(file), messages)

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

	b := new(bytes.Buffer)
	for i := 0; i <= 9; i++ {
		fmt.Println("\n Total number of words in file", keys[i])
		fmt.Fprintf(b, "%v \n", keys[i])
	}

	result := []byte(b.String())

	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
