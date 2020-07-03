package app

import (
	"net/http"
	controllers "sharan/golang-microservice/mvc/controllers"
)

//This is application
func StartApp() {
	//This is application
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
