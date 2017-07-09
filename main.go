package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

type Status struct {
	code string
}

var status = Status{"neutral"}

var images = map[string][]string{
	"success": []string{"letsgo.jpg", "successkid.jpg"},
	"fail":    []string{"simply.jpg", "thisisfine.png", "kubi.jpg"},
	"neutral": []string{"http://chillestmonkey.com/img/monkey.gif"},
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	r := chi.NewRouter()
	r.Get("/", serveHome)
	r.Get("/ws", socketHandler)
	r.Post("/status", changeStatus)
	http.ListenAndServe(":8080", r)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
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
