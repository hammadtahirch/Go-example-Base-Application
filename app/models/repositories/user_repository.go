package repositories

import (
	"strconv"
	"strings"
	"time"

	"github.com/hammadtahirch/golang_basic_app/app/models"
	"github.com/hammadtahirch/golang_basic_app/app/utils"
	"github.com/hammadtahirch/golang_basic_app/config"
	"github.com/jinzhu/gorm"
)

// UserRepository ... This hepls to maintain ther dependency
type UserRepository struct {
	db *gorm.DB
}

// CheckUserCridentails ... This function helps to getUsers.
func (ur UserRepository) CheckUserCridentails(muc models.UserCredentials) (models.User, error) {
	mu := models.User{}
	ur.db = config.DBConnection()
	er := ur.db.Preload("Role").Model(&mu).Where("email = ?", muc.Username).Find(&mu).Error
	if er != nil {
		return mu, er
	}
	ur.db.Close()
	return mu, er
}

// GetUsersRepo ... This function helps to getUsers.
func (ur UserRepository) GetUsersRepo(filter map[string][]string) ([]models.User, error) {
	var mu []models.User
	var ids string
	limit, page := 5, 1
	if filter["page"] != nil {
		page, _ = strconv.Atoi(filter["page"][0])
		page = page - 1
	}
	if filter["limit"] != nil {
		limit, _ = strconv.Atoi(filter["limit"][0])
	}
	ur.db = config.DBConnection()
	res := ur.db.Limit(limit).Offset(page * limit)
	if filter["ids"] != nil {
		ids = filter["ids"][0]
		res = res.Where("id IN (?)", strings.Split(ids, ","))
	}
	er := res.Preload("Role").Find(&mu).Error
	if er != nil {
		return mu, er
	}
	defer ur.db.Close()
	return mu, er
}

// GetUserByIDRepo ... This function helps to get user by id
func (ur UserRepository) GetUserByIDRepo(id int64) (models.User, error) {
	mu := models.User{}
	ur.db = config.DBConnection()
	er := ur.db.Preload("Role").Where("id = ?", id).Find(&mu).Error
	if er != nil {
		return mu, er
	}
	ur.db.Close()
	return mu, er
}

// StoreUserRepo ... This function helps to save user in storage.
func (ur UserRepository) StoreUserRepo(mu models.User) (models.User, error) {
	ur.db = config.DBConnection()
	ur.db.LogMode(true)
	res := ur.db.Save(&models.User{
		Name:        mu.Name,
		Email:       mu.Email,
		Password:    utils.GeneratePassword(mu.Password),
		PhoneNumber: mu.PhoneNumber,
		RoleID:      mu.RoleID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	er := res.Preload("Role").Find(&mu).Error
	if er != nil {
		return mu, er
	}
	defer ur.db.Close()
	return mu, er
}

// UpdateUserRepo ... This function helps to update user in storage
func (ur UserRepository) UpdateUserRepo(mu models.User, id int64) (models.User, error) {
	ur.db = config.DBConnection()
	res := ur.db.Model(&models.User{}).Where("id = ?", id).UpdateColumns(
		&models.User{
			Name:        mu.Name,
			RoleID:      mu.RoleID,
			PhoneNumber: mu.PhoneNumber,
		})
	er := res.Preload("Role").Model(&models.User{}).Where("id=?", id).First(&mu).Error
	if er != nil {
		return mu, er
	}
	defer ur.db.Close()
	return mu, er
}

// DestoryUserRepo ... This function helps to delete user from storege.
func (ur UserRepository) DestoryUserRepo(id int64) (models.User, error) {
	mu := models.User{}
	ur.db = config.DBConnection()
	res := ur.db.Where("id = ?", id).Delete(&models.User{})
	er := res.Unscoped().Preload("Role").Model(&mu).Where("id=? AND deleted_at IS NOT NULL", id).First(&mu).Error
	if er != nil {
		return mu, er
	}
	defer ur.db.Close()
	return mu, er
}

//RecoverPassword ... this func helps to take email and send reset password link
func (ur UserRepository) RecoverPassword() {
	//todo: Add for Generate new password request
}

// NewPassord ... This func helps to change the password
func (ur UserRepository) NewPassord() {
	//todo: add code to take new password and store it in storage
}

// Registration ... this func helps to take user information and store in storage
func (ur UserRepository) Registration() {
	//todo: Add code to take user object and save to storage
}
