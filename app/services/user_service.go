package services

import (
	"git-lab.boldapps.net/nifty-logix/mvc/app/models"
	"git-lab.boldapps.net/nifty-logix/mvc/app/repositories"
)

// UserService ... This struct helps to inject the dependency
type UserService struct {
	ur *repositories.UserRepository
}

// GetUsersService ... This function helps to getUsers.
func (us *UserService) GetUsersService(filter map[string][]string) ([]models.User, error) {
	res, err := us.ur.GetUsersRepo(filter)
	return res, err
}

// GetUserByIDService ... This function helps to get user by id
func (us *UserService) GetUserByIDService(id int) (models.User, error) {
	res, err := us.ur.GetUserByIDRepo(id)
	return res, err
}

// StoreUserService ... This function helps to save user in storage.
func (us *UserService) StoreUserService(mu models.User) (models.User, error) {
	res, err := us.ur.StoreUserRepo(mu)
	return res, err
}

// UpdateUserService ... This function helps to update user in storage
func (us *UserService) UpdateUserService(mu models.User, id int) (models.User, error) {
	res, err := us.ur.UpdateUserRepo(mu, id)
	return res, err
}

// DestoryUserService ... This function helps to delete user from storege.
func (us *UserService) DestoryUserService(id int) (models.User, error) {
	res, err := us.ur.DestoryUserRepo(id)
	return res, err
}
