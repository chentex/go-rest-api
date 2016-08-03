package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", Test)
	router.HandleFunc("/hola/{name}", Hola)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Test(w http.ResponseWriter, r *http.Request) {
	response := HipChatResponse{Color: "yellow", Message: "This is a Test", Notify: "false", MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

func Hola(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	response := HipChatResponse{Color: "yellow", Message: "Hola " + name, Notify: "false", MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

type HipChatResponse struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	Notify        string `json:"notify"`
	MessageFormat string `json:"message_format"`
}

type HipChatrequest struct {
	Event string `json:"event"`
	Item  struct {
		Message struct {
			Date time.Time `json:"date"`
			From struct {
				ID          int    `json:"id"`
				MentionName string `json:"mention_name"`
				Name        string `json:"name"`
			} `json:"from"`
			ID       string        `json:"id"`
			Mentions []interface{} `json:"mentions"`
			Message  string        `json:"message"`
			Type     string        `json:"type"`
		} `json:"message"`
		Room struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"room"`
	} `json:"item"`
	WebhookID int `json:"webhook_id"`
}
