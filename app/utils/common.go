package utils

import (
	"encoding/json"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

// View ... This struct is realted to the  LoadTemplate func
type View struct {
	Template *template.Template
}

// RespondJSON ... makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, status int, payload interface{}, name string) {
	wrapped := map[string]interface{}{
		name: payload,
	}
	response, err := json.Marshal(wrapped)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// RespondError ... makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, message, "error")
}

// LoadTemplate ... This function helps to load the templates
func LoadTemplate(files ...string) *View {
	files = append(files,
		"resources/views/layout/footer.gohtml",
		"resources/views/layout/header.gohtml",
		"resources/views/layout/nav_bar.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

// LoadStaticTemplate ... This function helps to load the templates
func LoadStaticTemplate(files ...string) *View {
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

// GeneratePassword ... This function hepls to generate password
func GeneratePassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
