package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"bufio"
	"io/ioutil"
	"log"
)

// method

func writeData(author, entry string) {
	saveFile, err := os.OpenFile("entries.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	defer saveFile.Close()
	
	w := bufio.NewWriter(saveFile)
	if err == nil {
		fmt.Fprintf(w, author + ": " + entry + "\n")
	}
	w.Flush()
}

// api

func timeHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodGet:
			currentTime := time.Now()
    		fmt.Fprintf(w, currentTime.Format("15h04"))
	}
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodPost:
			if err := req.ParseForm(); err != nil {
				fmt.Println("Something went bad")
				fmt.Fprintln(w, "Something went bad")
				return
			}
			author := req.Form.Get("author")
			entry := req.Form.Get("entry")
			if len(author) > 0 && len(entry) > 0 {
				fmt.Fprintf(w, author + ": " + entry)
				writeData(author, entry)
			}
	}
}

func listHandler(w http.ResponseWriter, req *http.Request) {

	content, err := ioutil.ReadFile("entries.txt")

	if err != nil {
		log.Fatal(err)
	}

    fmt.Fprintf(w, string(content))
}

// routing

func main() {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/entries", listHandler)
	http.ListenAndServe(":4567", nil)
}