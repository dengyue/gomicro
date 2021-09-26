package services

import "errors"

type IUserService interface {
	GetName(userId int) string
	DeleteUser(userId int) error
}

type UserService struct {
}

func (u UserService) GetName(userId int) string {
	if userId == 101 {
		return "dengyue"
	}
	return "guest"
}

func (u UserService) DeleteUser(userId int) error {
	if userId == 101 {
		return errors.New("no rights")
	}
	return nil
}
