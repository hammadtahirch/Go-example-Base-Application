package main

import (
	"net/http"
	"os"

	"github.com/hammadtahirch/nifty_logix/routes"
	"github.com/russross/blackfriday"
)

// AppRoute ...
type AppRoute struct {
	r routes.AppRoute
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

// main ...
// func main() {
// 	godotenv.Load()
// 	m := AppRoute{}
// 	a := mux.NewRouter()
// 	m.r.Routes(a)
// 	a.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
// 	http.ListenAndServe(":"+os.Getenv("APP_PORT"), a)
// }
