package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hammadtahirch/nifty_logix/routes"
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
	port := os.Getenv("APP_PORT")
	if port == "" && os.Getenv("APP_ENV") == "production" {
		port = os.Getenv("PORT") // heroku does not define static port at the time of deployment so that why i kept if dynamic
	}
	http.ListenAndServe(":"+port, a)
}
