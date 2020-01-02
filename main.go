package main

import (
	"net/http"
	"os"

	"github.com/hammadtahirch/nifty_logix/app/utils"
	"github.com/hammadtahirch/nifty_logix/routes"
)

// AppRoute ...
type AppRoute struct {
	r routes.AppRoute
}

// WelcomeData ... This is the demo structure for new learners
type WelcomeData struct {
	Title       string
	Description string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/markdown", Welcome)
	http.HandleFunc("/", Welcome)
	http.ListenAndServe(":"+port, nil)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeView := utils.LoadTemplate("resources/views/welcome.gohtml")
	err := homeView.Template.Execute(w, &WelcomeData{Title: "Golang MVC", Description: "Welcome To Golang"})
	if err != nil {
		panic(err)
	}
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
