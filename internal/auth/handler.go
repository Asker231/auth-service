package auth

import (
	"net/http"

	"github.com/Asker231/auth-service.git/config"
	"github.com/Asker231/auth-service.git/pkg/jwt"
	"github.com/Asker231/auth-service.git/pkg/middleware"
	"github.com/Asker231/auth-service.git/pkg/req"
	"github.com/Asker231/auth-service.git/pkg/res"
)


type AuthHandler struct{
	Cnf *config.AppConfig
	Service *AuthService
}

func NewAuthHandler(router *http.ServeMux,cfg *config.AppConfig,service *AuthService){
	handler := &AuthHandler{
		Cnf: cfg,
		Service: service,
	}
	router.HandleFunc("POST /auth/register",handler.Register())
	router.Handle("POST /auth/login",middleware.IsAuth(handler.Login(),cfg))
}

func(a *AuthHandler)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//read,decode,validate body
		body,err := req.HandleBody[RegisterRequest](w,r)
		if err != nil{
			return
		}
		userResult,err := a.Service.Register(body.Name,body.Email,body.Password)

		if err != nil{
			res.Response(w,err.Error(),404)
			return
		}

		token,err := jwt.NewJWT(a.Cnf.Auth.SECRET).Generate(jwt.JWTData{Email: userResult.Email})
		if err != nil{
			res.Response(w,err.Error(),404)
			return
		}
		res.Response(w,token,200)

	}
}

func(a *AuthHandler)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
			//read,decode,validate body
		body,err := req.HandleBody[LoginRequest](w,r)
		if err != nil{
			return
		}
		userResult,err := a.Service.Login(body.Email,body.Password)
		if err != nil{
			res.Response(w,err.Error(),404)
			return
		}
		res.Response(w,userResult,200)

	}
}