package services

import (
	"git-lab.boldapps.net/nifty-logix/mvc/app/models"
	"git-lab.boldapps.net/nifty-logix/mvc/app/models/repositories"
	"git-lab.boldapps.net/nifty-logix/mvc/app/utils"
)

// UserService ... This struct helps to inject the dependency
type UserService struct {
	ur *repositories.UserRepository
	l  LogService
}

//SignIn ... This function helps to generate token
func (us *UserService) SignIn(muc models.UserCredentials) (models.TokenPayload, models.Error) {
	res, err := us.ur.CheckUserCridentails(muc)
	cp, err := utils.ComparePasswords(res.Password, []byte(muc.Password))
	if err != nil || cp == false {
		l := us.l.LogError("username or password id incorrent", err.Error(), 422, false)
		return models.TokenPayload{}, l
	}
	if err != nil {
		l := us.l.LogError("Some thing went wrong", err.Error(), 500, true)
		return models.TokenPayload{}, l
	}
	t, err := utils.GenerateJwtToken(muc)
	if err != nil {
		l := us.l.LogError("Some thing went wrong", err.Error(), 500, true)
		return models.TokenPayload{}, l
	}
	return t, models.Error{}
}

// GetUsersService ... This function helps to getUsers.
func (us *UserService) GetUsersService(filter map[string][]string) ([]models.User, models.Error) {
	res, err := us.ur.GetUsersRepo(filter)
	if err != nil {
		l := us.l.LogError("Whoops! Something Wrong", err.Error(), 500, true)
		return res, l
	}
	return res, models.Error{}
}

// GetUserByIDService ... This function helps to get user by id
func (us *UserService) GetUserByIDService(id int64) (models.User, models.Error) {
	res, err := us.ur.GetUserByIDRepo(id)
	if err != nil {
		l := us.l.LogError("Whoops! Something Wrong", err.Error(), 500, true)
		return res, l
	}
	return res, models.Error{}
}

// StoreUserService ... This function helps to save user in storage.
func (us *UserService) StoreUserService(mu models.User) (models.User, models.Error) {
	res, err := us.ur.StoreUserRepo(mu)
	if err != nil {
		l := us.l.LogError("Whoops! Something Wrong", err.Error(), 500, true)
		return res, l
	}
	return res, models.Error{}
}

// UpdateUserService ... This function helps to update user in storage
func (us *UserService) UpdateUserService(mu models.User, id int64) (models.User, models.Error) {
	res, err := us.ur.UpdateUserRepo(mu, id)
	if err != nil {
		l := us.l.LogError("Whoops! Something Wrong", err.Error(), 500, true)
		return res, l
	}
	return res, models.Error{}
}

// DestoryUserService ... This function helps to delete user from storege.
func (us *UserService) DestoryUserService(id int64) (models.User, models.Error) {
	res, err := us.ur.DestoryUserRepo(id)
	if err != nil {
		l := us.l.LogError("Whoops! Something Wrong", err.Error(), 500, true)
		return res, l
	}
	return res, models.Error{}
}

//RecoverPassword ... this func helps to take email and send reset password link
func (us *UserService) RecoverPassword() {
	//todo: Add for Generate new password request
}

// NewPassord ... This func helps to change the password
func (us *UserService) NewPassord() {
	//todo: add code to take new password and store it in storage
}

// Registration ... this func helps to take user information and store in storage
func (us *UserService) Registration() {
	//todo: Add code to take user object and save to storage
}
