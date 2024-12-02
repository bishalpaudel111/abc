package controllers

import (
	"net/http"

	"github.com/abc/bishal/responses"
)

func HandleSccessRequest(w http.ResponseWriter,r*http.Request){
	data :=map[string]interface{}{
		"'user":[]string{"user1","user2","user3"},
		"total-size":100,
		"page":1,
	}
	responses.SuccessResponse(w,data,"successfully fatched user data")
}
 func HandleErrorRequest(w http.ResponseWriter,r*http.Request){

	responses.ErrorResponse(w,http.StatusUnauthorized,"Not authorizes to access this rote")
 }