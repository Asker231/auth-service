package auth

import (
	"net/http"

	"github.com/Asker231/auth-service.git/config"
)


type AuthHandler struct{
	Cnf *config.AppConfig
}

func NewAuthHandler(router *http.ServeMux,cfg *config.AppConfig){
	handler := &AuthHandler{
		Cnf: cfg,
	}
	router.HandleFunc("POST /auth/register",handler.Register())
	router.HandleFunc("POST /auth/login",handler.Login())
}

func(a *AuthHandler)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//read,decode,validate body

	}
}

func(a *AuthHandler)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
			//read,decode,validate body
			
	}
}