package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git-lab.boldapps.net/nifty-logix/mvc/app/models"
	"git-lab.boldapps.net/nifty-logix/mvc/app/services"
	"git-lab.boldapps.net/nifty-logix/mvc/app/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

// UserController ... This struct helps to inject the dependency
type UserController struct {
	us services.UserService
}

// IndexHandler ... This function helps to get All records for storage
func (uc *UserController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	parms := r.URL.Query()
	res, ierr := uc.us.GetUsersService(parms)
	if ierr != nil {
		utils.RespondJSON(w, 500, ierr.Error(), "error")
		return
	}
	utils.RespondJSON(w, 200, res, "users")
}

// ShowHandler ... This function helps to get records by id
func (uc *UserController) ShowHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	res, serr := uc.us.GetUserByIDService(id)
	if serr != nil {
		utils.RespondJSON(w, 500, serr.Error(), "error")
		return
	}
	utils.RespondJSON(w, 200, res, "users")
}

// StoreHandler ... This function helps to get store data in storage
func (uc *UserController) StoreHandler(w http.ResponseWriter, r *http.Request) {

	var p struct {
		Payload models.User `json:"user"`
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.RespondJSON(w, 500, "Json error", "error")
		return
	}
	//validation
	_, err := govalidator.ValidateStruct(&p.Payload)
	if err != nil {
		errors := govalidator.ErrorsByField(err)
		utils.RespondJSON(w, 422, errors, "error")
		return
	}
	//end validation
	res, serr := uc.us.StoreUserService(p.Payload)
	if serr != nil {
		utils.RespondJSON(w, 500, serr.Error(), "error")
		return
	}
	utils.RespondJSON(w, 200, res, "user")
}

// UpdateHandler ... This function helps to update data in storage using id
func (uc *UserController) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var p struct {
		Payload models.User `json:"user"`
	}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.RespondJSON(w, 500, "Json error", "error")
	}

	//validation
	_, err := govalidator.ValidateStruct(&p.Payload)
	if err != nil {
		errors := govalidator.ErrorsByField(err)
		utils.RespondJSON(w, 422, errors, "error")
		return
	}
	//end validation

	res, uerr := uc.us.UpdateUserService(p.Payload, id)
	if uerr != nil {
		utils.RespondJSON(w, 500, uerr.Error(), "error")
		return
	}
	utils.RespondJSON(w, 200, res, "user")
}

// DestroyHandler ... This function helps to delete data from storage using id
func (uc *UserController) DestroyHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	res, derr := uc.us.DestoryUserService(id)
	if derr != nil {
		utils.RespondJSON(w, 500, derr.Error(), "error")
		return
	}
	utils.RespondJSON(w, 200, res, "user")
}
