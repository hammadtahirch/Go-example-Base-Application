package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/hammadtahirch/golang_basic_app/app/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// View ... This struct is realted to the  LoadTemplate func
type View struct {
	Template *template.Template
}

// RespondJSON ... makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, status int, payload interface{}, name string) {

	var err error
	var response []byte
	if name != "-" {
		wrapped := map[string]interface{}{
			name: payload,
		}
		response, err = json.Marshal(wrapped)
	} else {
		response, err = json.Marshal(payload)
	}

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

// ComparePasswords ... this function helps to compare passsword.
func ComparePasswords(hashedPwd string, plainPwd []byte) (bool, error) {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false, err
	}

	return true, err
}

//GenerateJwtToken ... this function helps to take user Crdentails
// and make jwt token
func GenerateJwtToken(muc models.UserCredentials) (models.TokenPayload, error) {

	expirationTime := time.Now().Add(24 * time.Minute)
	claims := models.Claims{
		Username: muc.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRATE")))
	t := models.TokenPayload{
		Token: tokenString,
	}
	return t, err
}
