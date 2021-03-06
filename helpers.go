package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func chooseImage() string {
	category := images[status.code]
	rand.Seed(time.Now().Unix())
	img := category[rand.Intn(len(category))]
	return img
}

func addTeamsToUser(u *User, jobs []Job) {
	for _, job := range jobs {
		u.teams = append(u.teams, job.Name)
	}
}
