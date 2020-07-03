package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	services "sharan/golang-microservice/mvc/services"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	log.Printf("About to process user id %v", userID)

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := services.GetUser(userID)

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	JsonValue, _ := json.Marshal(user)

	resp.Write(JsonValue)

}
