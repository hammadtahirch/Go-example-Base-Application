package repositories

import (
	"strconv"
	"strings"
	"time"

	"git-lab.boldapps.net/nifty-logix/mvc/app/models"
	"git-lab.boldapps.net/nifty-logix/mvc/app/utils"
	"git-lab.boldapps.net/nifty-logix/mvc/config"
	"github.com/jinzhu/gorm"
)

// UserRepository ... This hepls to maintain ther dependency
type UserRepository struct {
	db *gorm.DB
}

// GetUsersRepo ... This function helps to getUsers.
func (ur *UserRepository) GetUsersRepo(filter map[string][]string) ([]models.User, error) {
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

	db := config.DBConnection()
	db.LogMode(true)
	res := db.Limit(limit).Offset(page * limit)
	if filter["ids"] != nil {
		ids = filter["ids"][0]
		res = res.Where("id IN (?)", strings.Split(ids, ","))
	}
	res.Preload("Role").Find(&mu)
	if res.Error != nil {
		return mu, utils.New(res.Error.Error())
	}
	defer db.Close()
	return mu, nil
}

// GetUserByIDRepo ... This function helps to get user by id
func (ur *UserRepository) GetUserByIDRepo(id int) (models.User, error) {
	mu := models.User{}
	db := config.DBConnection()
	res := db.Preload("Role").First(&mu, &models.User{ID: id})
	if res.Error != nil {
		return mu, utils.New(db.Error.Error())
	}
	db.Close()
	return mu, nil
}

// StoreUserRepo ... This function helps to save user in storage.
func (ur *UserRepository) StoreUserRepo(mu models.User) (models.User, error) {
	db := config.DBConnection()
	res := db.Save(&models.User{
		Name:      mu.Name,
		Email:     mu.Email,
		Password:  utils.GeneratePassword(mu.Password),
		RoleID:    mu.RoleID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	res.Preload("Role").Find(&mu)
	if res.Error != nil {
		return mu, utils.New(res.Error.Error())
	}
	defer db.Close()
	return mu, nil
}

// UpdateUserRepo ... This function helps to update user in storage
func (ur *UserRepository) UpdateUserRepo(mu models.User, id int) (models.User, error) {
	db := config.DBConnection()
	res := db.Model(&models.User{}).Where("id = ?", id).UpdateColumns(&models.User{Name: mu.Name, RoleID: mu.RoleID})
	res.Preload("Role").Model(&models.User{}).Where("id=?", id).First(&mu)
	if res.Error != nil {
		return mu, utils.New(db.Error.Error())
	}
	defer db.Close()
	return mu, nil
}

// DestoryUserRepo ... This function helps to delete user from storege.
func (ur *UserRepository) DestoryUserRepo(id int) (models.User, error) {
	mu := models.User{}
	db := config.DBConnection()
	res := db.Where("id = ?", id).Delete(&models.User{})
	res.Unscoped().Preload("Role").Model(&mu).Where("id=? AND deleted_at IS NOT NULL", id).First(&mu)
	if res.Error != nil {
		return mu, utils.New(db.Error.Error())
	}
	defer db.Close()
	return mu, nil
}
