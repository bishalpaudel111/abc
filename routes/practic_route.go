package routes

import (
	"github.com/abc/bishal/controllers"
	"github.com/gorilla/mux"
)

func InitAllPracticRoutes(router *mux.Router) {
	router.HandleFunc("/error", controllers.HandleErrorRequest)
	router.HandleFunc("/success",controllers.HandleSccessRequest )

}