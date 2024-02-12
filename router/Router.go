package router

import (
	"net/http"

	"dataentry/controllers"
)

func RouteHandler() http.Handler {

	fileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", controllers.FormHandler)

	return http.DefaultServeMux
}
