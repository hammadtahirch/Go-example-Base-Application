package controllers

import (
	"net/http"

	"github.com/hammadtahirch/nifty_logix/app/utils"
)

// WelcomeData ... This is the demo structure for new learners
type WelcomeData struct {
	Title       string
	Description string
}

// Welcome ... this function helps to load the welcome view
// This func also show that how to pass data in view
func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeView := utils.LoadTemplate("resources/views/welcome.gohtml")
	err := homeView.Template.Execute(w, &WelcomeData{Title: "Golang MVC", Description: "Welcome To Golang"})
	if err != nil {
		panic(err)
	}
}

// ReactProduction ... this function helps to load the welcome view
func ReactProduction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeView := utils.LoadStaticTemplate("resources/views/react_build.gohtml")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
