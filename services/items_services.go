package services

import (
	"net/http"
	"sharan/golang-microservice/mvc/domain"
	"sharan/golang-microservice/mvc/utils"
)

type itemService struct{}

var (
	ItemService itemService
)

func (i *itemService) GetItems(itemID string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Not yet implemented",
		StatusCode: http.StatusInternalServerError,
		Code:       "Internal Error",
	}
}
