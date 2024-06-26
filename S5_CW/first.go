package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const addr string = "localhost:8080"

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(addr, nil))

}

func handle(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	response := "1 instance: " + text + "\n"

	if _, err := w.Write([]byte(response)); err != nil {
		log.Fatal(err)
	}

}
