package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var users []string

func listUsers(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	users = []string{"Lais", "Asael"}
	j, _ := json.Marshal(users)
	w.Write(j)
}

func main() {
	http.HandleFunc("/list", listUsers)
	http.ListenAndServe(":8082", nil)
	log.Print("Server is listening on port :8082")
}
