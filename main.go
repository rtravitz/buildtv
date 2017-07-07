package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/success", success)
	r.Get("/fail", fail)
	http.ListenAndServe(":8080", r)
}

func success(w http.ResponseWriter, r *http.Request) {
	i := image{"http://chillestmonkey.com/img/monkey.gif"}
	t, _ := makeTemplate(i)
	respondWithTemplate(w, t, i)
}

func fail(w http.ResponseWriter, r *http.Request) {
	i := image{"https://thenib.imgix.net/usq/1d97429f-4a64-4d52-bfdb-c36172c05228/this-is-not-fine-001-dae9d5.png?auto=compress,format&_=dae9d5fc0800f12f5c720be598b6bea6"}
	t, _ := makeTemplate(i)
	respondWithTemplate(w, t, i)
}

func makeTemplate(i image) (*template.Template, error) {
	t := template.New("maindisplay.html")
	t, err := t.Parse(`
		<head>
			<style>
				.image {
					height: 100%;
					background: url("{{.Address}}") no-repeat center center fixed;
					background-size: cover;
				}
			</style>
		</head>
		<body>
			<div class='image'></div>
		</body>
		`)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func respondWithTemplate(w http.ResponseWriter, t *template.Template, i image) {
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, i)
}

type image struct {
	Address string
}
