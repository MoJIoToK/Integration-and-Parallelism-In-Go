package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User1 struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []*User1 `json:"friends"`
}

type Network1 struct {
	network1 map[int]*User1
}

func main() {

	n := Network1{make(map[int]*User1)}

	nickS := User1{
		Id:      1,
		Name:    "Nick",
		Age:     27,
		Friends: nil,
	}

	vasyaS := User1{
		Id:      2,
		Name:    "Vasya",
		Age:     37,
		Friends: nil,
	}

	n.network1[nickS.Id] = &nickS

	n.network1[vasyaS.Id] = &vasyaS

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/create/", n.Create1)
	r.Get("/user/", n.GetAll1)
	r.Post("/make_friends/{source}-{target}", n.MakeFriends1)
	r.Delete("/delete/{id}", n.Delete1)
	r.Get("/friends/{id}", n.GetFriends1)
	r.Put("/update_age/{id}-{age}", n.UpdateAge1)

	http.ListenAndServe("localhost:8080", r)
}

func (n *Network1) Create1(w http.ResponseWriter, r *http.Request) {
	var user User1
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()
	n.network1[user.Id] = &user
	jsonUser, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
}

func (n *Network1) GetAll1(w http.ResponseWriter, r *http.Request) {
	result := make([]User1, 0, len(n.network1))
	for _, user := range n.network1 {
		result = append(result, *user)
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)
}

func (n *Network1) MakeFriends1(w http.ResponseWriter, r *http.Request) {
	sourceIdStr := chi.URLParam(r, "source")
	sourceId, err := strconv.Atoi(sourceIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	targetIdStr := chi.URLParam(r, "target")
	targetId, err := strconv.Atoi(targetIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := n.network1[sourceId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if _, ok := n.network1[targetId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sourceName := n.network1[sourceId].Name
	targetName := n.network1[targetId].Name

	n.UpdateFriends1(n.network1[sourceId], n.network1[targetId])

	msg := fmt.Sprintf("%v и %v теперь друзья", targetName, sourceName)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func (n *Network1) Delete1(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := n.network1[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	deletedName := n.network1[id].Name
	msg := fmt.Sprintf("%v удалён!", deletedName)

	delete(n.network1, id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))

}

func (n *Network1) GetFriends1(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := n.network1[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resultFriends := make([]*User1, 0, len(n.network1))

	resultFriends = append(resultFriends, n.network1[id].Friends...)

	resultJson, err := json.Marshal(resultFriends)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)

}

func (n *Network1) UpdateAge1(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ageStr := chi.URLParam(r, "age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := n.network1[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	n.network1[id].Age = age
	msg := fmt.Sprintf("Возраст пользователя %v успешно обновлен!", n.network1[id].Name)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func (n *Network1) UpdateFriends1(user1, user2 *User1) {
	n.network1[user1.Id].Friends = append(n.network1[user1.Id].Friends, n.network1[user2.Id])
}
