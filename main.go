package main

import (
	"net/http"
	"os"

	"git-lab.boldapps.net/nifty-logix/mvc/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// MainApp ...
type MainApp struct {
	r routes.AppRoute
}

func main() {

	godotenv.Load()
	m := MainApp{}
	a := mux.NewRouter()
	m.r.Routes(a)
	http.ListenAndServe(os.Getenv("APP_PORT"), a)
}
