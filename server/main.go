package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message  string `json:message`
	Datetime string `json:datetime`
}

func message(w http.ResponseWriter, req *http.Request) {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	currentTime := time.Now().In(loc)
	timeFormatted := currentTime.Format("2006-01-02 3:4:5")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := &Response{
		Message:  "Hello world!",
		Datetime: timeFormatted,
	}
	j, _ := json.Marshal(resp)
	w.Write(j)
}

func main() {
	http.HandleFunc("/", message)
	http.ListenAndServe(":8082", nil)
	log.Print("Server is listening on port :8082")
}
