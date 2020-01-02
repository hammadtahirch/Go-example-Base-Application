package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/hammadtahirch/nifty_logix/app/models"
	"github.com/hammadtahirch/nifty_logix/app/services"
	"github.com/hammadtahirch/nifty_logix/app/utils"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gorilla/mux"
)

// UserController ... This struct helps to inject the dependency
type UserController struct {
	us services.UserService
}

// SignIn ... This function helps to generate token
func (uc UserController) SignIn(w http.ResponseWriter, r *http.Request) {
	var c struct {
		Payload models.UserCredentials `json:"user"`
	}
	if er := json.NewDecoder(r.Body).Decode(&c); er != nil {
		utils.RespondJSON(w, 500, "Json error", "error")
		return
	}

	//validation
	er := validation.Errors{
		"username": validation.Validate(c.Payload.Username, validation.Required.Error("UH-HO! User is required."), is.Email.Error("UH-HO! Email is not corrct.")),
		"password": validation.Validate(c.Payload.Password, validation.Required.Error("UH-HO! Password is required."), validation.Length(8, 16).Error("UH-HO! Password should be between 8 to 16.")),
	}.Filter()
	if er != nil {
		utils.RespondJSON(w, 422, er, "error")
		return
	}
	//validation

	res, err := uc.us.SignIn(c.Payload)
	if err.Code != 0 {
		utils.RespondJSON(w, err.Code, err, "error")
		return
	}
	utils.RespondJSON(w, 200, res, "-")
}

// IndexHandler ... This function helps to get All records for storage
func (uc UserController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	parms := r.URL.Query()
	res, err := uc.us.GetUsersService(parms)
	if err.Code != 0 {
		utils.RespondJSON(w, 500, err, "error")
		return
	}
	utils.RespondJSON(w, 200, res, "users")
}

// ShowHandler ... This function helps to get records by id
func (uc UserController) ShowHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	res, err := uc.us.GetUserByIDService(id)
	if err.Code != 0 {
		utils.RespondJSON(w, 500, err, "error")
		return
	}
	utils.RespondJSON(w, 200, res, "users")
}

// StoreHandler ... This function helps to get store data in storage
func (uc UserController) StoreHandler(w http.ResponseWriter, r *http.Request) {

	var p struct {
		Payload models.User `json:"user"`
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.RespondJSON(w, 500, "Json error", "error")
		return
	}
	//validation
	er := validation.Errors{
		"name":         validation.Validate(p.Payload.Name, validation.Required.Error("UH-HO! Name is required.")),
		"email":        validation.Validate(p.Payload.Email, validation.Required.Error("UH-HO! Email is required."), is.Email.Error("UH-HO! Email is not required.")),
		"password":     validation.Validate(p.Payload.Password, validation.Required.Error("UH-HO! Password is required."), validation.Length(8, 16).Error("UH-HO! Passord Should be 8 to 16 characters."), is.UTFLetterNumeric.Error("UH-HO! Password mush contain alphanumaric characters.")),
		"phone_number": validation.Validate(p.Payload.PhoneNumber, validation.Required.Error("UH-HO! Phone number is required."), is.E164.Error("UH-HO! Phone number is not correct.")),
		"role_id":      validation.Validate(p.Payload.RoleID, validation.Required.Error("UH-HO! Role is required.")),
	}.Filter()
	if er != nil {
		utils.RespondJSON(w, 422, er, "error")
		return
	}
	//validation
	res, err := uc.us.StoreUserService(p.Payload)
	if err.Code != 0 {
		utils.RespondJSON(w, err.Code, err, "error")
		return
	}
	utils.RespondJSON(w, 200, res, "user")
}

// UpdateHandler ... This function helps to update data in storage using id
func (uc UserController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	var p struct {
		Payload models.User `json:"user"`
	}
	if jerr := json.NewDecoder(r.Body).Decode(&p); jerr != nil {
		utils.RespondJSON(w, 500, "Json error", "error")
	}
	//validation
	er := validation.Errors{
		"name":         validation.Validate(p.Payload.Name, validation.Required.Error("UH-HO! Name is required.")),
		"email":        validation.Validate(p.Payload.Email, validation.Required.Error("UH-HO! Email is required."), is.Email.Error("UH-HO! Email is not correct.")),
		"phone_number": validation.Validate(p.Payload.PhoneNumber, validation.Required.Error("UH-HO! Phone number is required."), is.E164.Error("UH-HO! Phone number is not correct.")),
		"role_id":      validation.Validate(p.Payload.RoleID, validation.Required.Error("UH-HO! Role is required.")),
	}.Filter()
	if er != nil {
		utils.RespondJSON(w, 422, er, "error")
		return
	}
	//validation

	res, err := uc.us.UpdateUserService(p.Payload, id)
	if err.Code != 0 {
		utils.RespondJSON(w, err.Code, err, "error")
		return
	}
	utils.RespondJSON(w, 200, res, "user")
}

// DestroyHandler ... This function helps to delete data from storage using id
func (uc UserController) DestroyHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	res, err := uc.us.DestoryUserService(id)
	if err.Code != 0 {
		utils.RespondJSON(w, err.Code, err, "error")
		return
	}
	utils.RespondJSON(w, 200, res, "user")
}

//RecoverPassword ... this func helps to take email and send reset password link
func (uc UserController) RecoverPassword(w http.ResponseWriter, r *http.Request) {
	//todo: Add for Generate new password request
}

// NewPassord ... This func helps to change the password
func (uc UserController) NewPassord(w http.ResponseWriter, r *http.Request) {
	//todo: add code to take new password and store it in storage
}

// Registration ... this func helps to take user information and store in storage
func (uc UserController) Registration(w http.ResponseWriter, r *http.Request) {

	// var p struct {
	// 	Payload models.User `json:"user"`
	// }

	// if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
	// 	utils.RespondJSON(w, 500, "Json error", "error")
	// 	return
	// }
	// //validation
	// er := validation.Errors{
	// 	"name":         validation.Validate(p.Payload.Name, validation.Required.Error("UH-HO! Name is required.")),
	// 	"email":        validation.Validate(p.Payload.Email, validation.Required.Error("UH-HO! Email is required."), is.Email.Error("UH-HO! Email is not required.")),
	// 	"password":     validation.Validate(p.Payload.Password, validation.Required.Error("UH-HO! Password is required."), validation.Length(8, 16).Error("UH-HO! Passord Should be 8 to 16 characters."), is.UTFLetterNumeric.Error("UH-HO! Password mush contain alphanumaric characters.")),
	// 	"phone_number": validation.Validate(p.Payload.PhoneNumber, validation.Required.Error("UH-HO! Phone number is required."), is.E164.Error("UH-HO! Phone number is not correct.")),
	// 	"role_id":      validation.Validate(p.Payload.RoleID, validation.Required.Error("UH-HO! Role is required.")),
	// }.Filter()
	// if er != nil {
	// 	utils.RespondJSON(w, 422, er, "error")
	// 	return
	// }
	// //validation
	// res, err := uc.us.StoreUserService(p.Payload)
	// if err.Code != 0 {
	// 	utils.RespondJSON(w, err.Code, err, "error")
	// 	return
	// }
	// utils.RespondJSON(w, 200, res, "user")

	//todo: Add code to take user object and save to storage
	var cnf = config.Config{
		Broker:        "redis://redis-container:6379",
		ResultBackend: "redis://redis-container:6379",
	}

	server, err := machinery.NewServer(&cnf)
	if err != nil {
		//todo: add error handling
	}

	signature := &tasks.Signature{
		Name: "SendEmailJob",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: "hammad.tahir.ch@gmail.com",
			},
		},
	}
	eta := time.Now().UTC().Add(time.Second * 5)
	signature.ETA = &eta
	asyncResult, err := server.SendTask(signature)
	if err != nil {
		// failed to send the task
		// do something with the error
	}
	println(asyncResult)
}
