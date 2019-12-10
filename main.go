package main

import (
	"net/http"
	"os"

	"github.com/hammadtahirch/golang_basic_app/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// AppRoute ...
type AppRoute struct {
	r routes.AppRoute
}

// main ...
func main() {
	godotenv.Load()
	m := AppRoute{}
	a := mux.NewRouter()
	m.r.Routes(a)
	a.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), a)
}
