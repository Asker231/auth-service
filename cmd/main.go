package main

import (
	"net/http"

	"github.com/Asker231/auth-service.git/config"
	"github.com/Asker231/auth-service.git/internal/auth"
)

func main() {


	//config
	cnf := config.LoadAppConfig()


	//init mux 
	router := http.NewServeMux()


	//init handlers
	auth.NewAuthHandler(router,cnf)

	//init server struct
	server := http.Server{
		Addr: ":8000",
		Handler: router,
	}
	server.ListenAndServe()



}