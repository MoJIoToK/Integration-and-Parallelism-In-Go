package main

type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends []*User `json:"friends"`
}

type Network struct {
	network map[int]*User
}

// func (n *Network) Create(w http.ResponseWriter, r *http.Request) {
// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	defer r.Body.Close()
// 	n.network[user.Id] = &user
// 	jsonUser, err := json.Marshal(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonUser)
// }

// func (n *Network) GetAll(w http.ResponseWriter, r *http.Request) {
// 	result := make([]User, 0, len(n.network))
// 	for _, user := range n.network {
// 		result = append(result, *user)
// 	}
// 	resultJson, err := json.Marshal(result)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resultJson)
// }

// func (n *Network) MakeFriends(w http.ResponseWriter, r *http.Request) {
// 	sourceIdStr := chi.URLParam(r, "source")
// 	sourceId, err := strconv.Atoi(sourceIdStr)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}

// 	targetIdStr := chi.URLParam(r, "target")
// 	targetId, err := strconv.Atoi(targetIdStr)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}

// 	if _, ok := n.network[sourceId]; !ok {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	if _, ok := n.network[targetId]; !ok {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	sourceName := n.network[sourceId].Name
// 	targetName := n.network[targetId].Name

// 	n.UpdateFriends(n.network[sourceId], n.network[targetId])

// 	msg := fmt.Sprintf("%v и %v теперь друзья", targetName, sourceName)

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(msg))
// }

// func (n *Network) Delete(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}

// 	if _, ok := n.network[id]; !ok {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	deletedName := n.network[id].Name
// 	msg := fmt.Sprintf("%v удалён!", deletedName)

// 	delete(n.network, id)

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(msg))

// }

// func (n *Network) GetFriends(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}

// 	if _, ok := n.network[id]; !ok {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	resultFriends := make([]*User, 0, len(n.network))

// 	resultFriends = append(resultFriends, n.network[id].Friends...)

// 	resultJson, err := json.Marshal(resultFriends)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resultJson)

// }

// func (n *Network) UpdateAge(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}

// 	ageStr := chi.URLParam(r, "age")
// 	age, err := strconv.Atoi(ageStr)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}

// 	if _, ok := n.network[id]; !ok {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	n.network[id].Age = age
// 	msg := fmt.Sprintf("Возраст пользователя %v успешно обновлен!", n.network[id].Name)

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(msg))
// }

// func (n *Network) UpdateFriends(user1, user2 *User) {
// 	n.network[user1.Id].Friends = append(n.network[user1.Id].Friends, n.network[user2.Id])
// }
