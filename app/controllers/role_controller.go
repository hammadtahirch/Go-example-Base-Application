package controllers

import (
	"net/http"

	"git-lab.boldapps.net/nifty-logix/mvc/app/services"
)

//RoleController ... This helps to defind the dependencies.
type RoleController struct {
	us services.UserService
}

// GetRoleList ... this func helps to get all roles from storage
func (ro *RoleController) GetRoleList(w http.ResponseWriter, r *http.Request) {
	//todo:add code to render Role list
}
