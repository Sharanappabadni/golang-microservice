package services

import (
	domain "sharan/golang-microservice/mvc/domain"
	utils "sharan/golang-microservice/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
