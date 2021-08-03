package repository

import (
	"github.com/cyantosh0/dockerize-go-app/config"
	"github.com/cyantosh0/dockerize-go-app/model"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]model.User) error {
	if err := config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *model.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *model.User, id string) error {
	if err := config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
