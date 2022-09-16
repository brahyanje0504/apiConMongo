package services

import (
	"basesdedatos/models"
	"basesdedatos/repositories"
)

func Create(user models.User) error {
	err := repositories.Create(user)
	if err != nil {
		return err
	}
	return nil
}
func Read() (models.Users, error) {
	users, err := repositories.Read()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func Update(user models.User, userId string) error {
	err := repositories.Update(user, userId)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userId string) error {
	err := repositories.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
