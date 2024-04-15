package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9000"

var (
	counter            int    = 0
	firstInstanceHost  string = "http://localhost:8080"
	secondInstanceHost string = "http://localhost:8081"
)

func main() {
	http.HandleFunc("/", handleProxe)
	log.Fatal(http.ListenAndServe(proxyAddr, nil))
}

func handleProxe(w http.ResponseWriter, r *http.Request) {
	textBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	text := string(textBytes)

	if counter == 0 {
		resp, err := http.Post(firstInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatal(err)
		}
		counter++

		textBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		fmt.Println(string(textBytes))

		return
	}

	if counter == 1 {
		resp, err := http.Post(secondInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatal(err)
		}
		counter--

		textBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		fmt.Println(string(textBytes))

		return
	}

}
