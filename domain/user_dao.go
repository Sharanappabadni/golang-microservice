package domain

import (
	"fmt"
	"net/http"
	utils "sharan/golang-microservice/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Sharan", LastName: "Badni", Email: "sharan.badni@gmail.com"},
		876: &User{ID: 876, FirstName: "Suresh", LastName: "Vaddar", Email: "suresh@gmail.com"},
	}
)

func GetUser(userID int64) (*User, *utils.ApplicationError) {

	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v not found ", userID),
		StatusCode: http.StatusNotFound,
		Code:       "Not found",
	}
}
