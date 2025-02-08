package auth

import "net/http"


type AuthHandler struct{

}

func NewAuthHandler(router *http.ServeMux){
	handler := &AuthHandler{}
	router.HandleFunc("POST /auth/register",handler.Register())
	router.HandleFunc("POST /auth/login",handler.Login())
}

func(a *AuthHandler)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}

func(a *AuthHandler)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {}
}