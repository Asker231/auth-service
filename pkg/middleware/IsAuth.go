package middleware

import (
	"net/http"

	"github.com/Asker231/auth-service.git/config"
	"github.com/Asker231/auth-service.git/pkg/jwt"
	"github.com/Asker231/auth-service.git/pkg/res"
)


func IsAuth(next http.Handler,conf *config.AppConfig)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Auth")
		isValid,_:= jwt.NewJWT(conf.Auth.SECRET).ParseJWT(token)
		if !isValid{
			res.Response(w,"Не валидные данные",404)
			return
		}
		next.ServeHTTP(w,r)
	})
}