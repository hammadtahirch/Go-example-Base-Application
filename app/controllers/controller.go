package controllers

import (
	"net/http"

	"git-lab.boldapps.net/nifty-logix/mvc/app/utils"
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
	err := homeView.Template.Execute(w, &WelcomeData{Title: "Golang MVC11111", Description: "Welcome To Golang"})
	if err != nil {
		panic(err)
	}
}
