package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var (
	status   = Status{"neutral"}
	baseURL  = os.Getenv("JENKINS_URL")
	token    = os.Getenv("JENKINS_TOKEN")
	username = os.Getenv("JENKINS_USER")
	team     = os.Getenv("JENKINS_TEAM")
	pipeline = os.Getenv("JENKINS_PIPELINE")

	images = map[string][]string{
		"success": []string{"letsgo.jpg", "successkid.jpg", "joe.jpg", "miracles.jpg"},
		"fail":    []string{"simply.jpg", "thisisfine.png", "kubi.jpg", "clouddude.jpg"},
		"neutral": []string{"http://chillestmonkey.com/img/monkey.gif"},
	}

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

//Status represents build status
type Status struct {
	code string
}

//Job is a Jenkins Job Response
type Job struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

//Health is a Job Health Response
type Health struct {
	Description string `json:"description"`
	Score       int    `json:"score"`
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/ws", socketHandler)
	http.HandleFunc("/status", manualStatusChange)
	log.Fatal(http.ListenAndServe(":1234", nil))
}

func manualStatusChange(w http.ResponseWriter, r *http.Request) {
	text := r.PostFormValue("status")
	status.code = text
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	var previousStatus string

	for {
		time.Sleep(1 * time.Second)
		var message []byte
		if previousStatus != status.code {
			message = []byte(chooseImage())
			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		}
		previousStatus = status.code
	}
}
