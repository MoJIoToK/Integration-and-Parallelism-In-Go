package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User2 struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []*User2 `json:"friends"`
}

type Network2 struct {
	network2 map[int]*User2
}

func main() {

	n := Network2{make(map[int]*User2)}

	nick2 := User2{
		Id:      1,
		Name:    "Nick",
		Age:     27,
		Friends: nil,
	}

	vasya2 := User2{
		Id:      2,
		Name:    "Vasya",
		Age:     37,
		Friends: nil,
	}

	n.network2[nick2.Id] = &nick2

	n.network2[vasya2.Id] = &vasya2

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/create/", n.Create2)
	r.Get("/user/", n.GetAll2)
	r.Post("/make_friends/{source}-{target}", n.MakeFriends2)
	r.Delete("/delete/{id}", n.Delete2)
	r.Get("/friends/{id}", n.GetFriends2)
	r.Put("/update_age/{id}-{age}", n.UpdateAge2)

	http.ListenAndServe("localhost:8081", r)
}

func (n *Network2) Create2(w http.ResponseWriter, r *http.Request) {
	var user User2
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()
	n.network2[user.Id] = &user
	jsonUser, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
}

func (n *Network2) GetAll2(w http.ResponseWriter, r *http.Request) {
	result := make([]User2, 0, len(n.network2))
	for _, user := range n.network2 {
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

func (n *Network2) MakeFriends2(w http.ResponseWriter, r *http.Request) {
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

	if _, ok := n.network2[sourceId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if _, ok := n.network2[targetId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sourceName := n.network2[sourceId].Name
	targetName := n.network2[targetId].Name

	n.UpdateFriends2(n.network2[sourceId], n.network2[targetId])

	msg := fmt.Sprintf("%v и %v теперь друзья", targetName, sourceName)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func (n *Network2) Delete2(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := n.network2[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	deletedName := n.network2[id].Name
	msg := fmt.Sprintf("%v удалён!", deletedName)

	delete(n.network2, id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))

}

func (n *Network2) GetFriends2(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if _, ok := n.network2[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resultFriends := make([]*User2, 0, len(n.network2))

	resultFriends = append(resultFriends, n.network2[id].Friends...)

	resultJson, err := json.Marshal(resultFriends)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)

}

func (n *Network2) UpdateAge2(w http.ResponseWriter, r *http.Request) {
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

	if _, ok := n.network2[id]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	n.network2[id].Age = age
	msg := fmt.Sprintf("Возраст пользователя %v успешно обновлен!", n.network2[id].Name)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func (n *Network2) UpdateFriends2(user1, user2 *User2) {
	n.network2[user1.Id].Friends = append(n.network2[user1.Id].Friends, n.network2[user2.Id])
}
