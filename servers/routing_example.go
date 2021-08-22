package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func aboutUs(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "You're on About us page.")
}

func handleUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var data []byte
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln("Failed to read from request body got error: ", err)
			return
		}
		fmt.Fprintln(w, string(data))
	} else {
		fmt.Fprintln(w, "Nothing to Fetch in GET users")
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is index page.")
}

func main () {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", aboutUs)
	http.HandleFunc("/about/", aboutUs)
	http.HandleFunc("/users", handleUser)

	port := 8080
	fmt.Println("Listening on Port ", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}