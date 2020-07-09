package domain

import (
	"fmt"
	"net/http"
	"sharan/golang-microservice/mvc/utils"
)

type userDao struct{}

var (
	UserDao userDao
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Sharan", LastName: "Badni", Email: "sharan.badni@gmail.com"},
		876: &User{ID: 876, FirstName: "Suresh", LastName: "Vaddar", Email: "sharan.badni@gmail.com"},
	},
)


func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {

	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
