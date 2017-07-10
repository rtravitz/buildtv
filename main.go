package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Status struct {
	code string
}

var status = Status{"neutral"}

var images = map[string][]string{
	"success": []string{"letsgo.jpg", "successkid.jpg", "joe.jpg", "miracles.jpg"},
	"fail":    []string{"simply.jpg", "thisisfine.png", "kubi.jpg", "clouddude.jpg"},
	"neutral": []string{"http://chillestmonkey.com/img/monkey.gif"},
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/ws", socketHandler)
	http.HandleFunc("/status", changeStatus)
	http.ListenAndServe(":1234", nil)
}

func changeStatus(w http.ResponseWriter, r *http.Request) {
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

func chooseImage() string {
	category := images[status.code]
	rand.Seed(time.Now().Unix())
	img := category[rand.Intn(len(category))]
	return img
}
