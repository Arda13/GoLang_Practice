package main

import (
	"arda/webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
