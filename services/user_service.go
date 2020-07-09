package services

import (
	domain "sharan/golang-microservice/mvc/domain"
	"sharan/golang-microservice/mvc/utils"
)

type userService struct{}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
