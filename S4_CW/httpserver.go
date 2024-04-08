package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) toString() string {
	return fmt.Sprintf("name is %s and age is %d \n", u.Name, u.Age)
}

type service struct {
	store map[string]*User
}

func main() {
	mux := http.NewServeMux()

	srv := service{make(map[string]*User)}
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.GetAll)
	mux.HandleFunc("/add", srv.Add)

	http.ListenAndServe("localhost:8080", mux)

	//srv.Add()

}

func (s *service) Add(w http.ResponseWriter, r *http.Request) {
	data := []byte(`{"name":"bar", "age":50}`)
	req := bytes.NewReader(data)
	resp, err := http.Post("localhost:8080/create", "application/json", req)
	if err != nil {
		log.Fatal(err)
	}
	//req, _ := http.NewRequest("POST", "http://localhost:8080/create", r)
	//
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()

	fmt.Println(resp.Status)
}

func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		//splittedContent := strings.Split(string(content), " ")
		s.store[u.Name] = &u

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created " + u.Name))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := ""
		for _, user := range s.store {
			response += user.toString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}
	w.WriteHeader(http.StatusBadRequest)
}
