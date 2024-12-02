package main

import (
	"net/http"

	"github.com/abc/bishal/responses"
	"github.com/abc/bishal/routes"
)

func main() {
	print("hello all")
	///initialling all routes
	router := routes.InitALLRoutes()
	//router := mux.NewRouter()
	router.HandleFunc("/api", HandleInitialRoute).Methods("GET")
	//creating server
	print("listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		print("error creating server: ", err.Error())
	}
}

func HandleInitialRoute(w http.ResponseWriter, r *http.Request) {
	print("Welcome to go api")
	//w.Header().Set("Content-Type","application/json")
	//w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{
		"user-list":  []string{"user1", "user2", "user3"},
		"total-size": 100,
		"page":       1,
	}
	responses.SuccessResponse(w, data, "Success fully Connected")
}
