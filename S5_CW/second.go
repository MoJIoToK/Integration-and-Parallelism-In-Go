// package main

// import (
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// const addr2 string = "localhost:8081"

// func main() {
// 	http.HandleFunc("/", handle2)
// 	log.Fatal(http.ListenAndServe(addr2, nil))

// }

// func handle2(w http.ResponseWriter, r *http.Request) {
// 	bodyBytes, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer r.Body.Close()

// 	text := string(bodyBytes)
// 	response := "2 instance: " + text + "\n"

// 	if _, err := w.Write([]byte(response)); err != nil {
// 		log.Fatal(err)
// 	}

// }
