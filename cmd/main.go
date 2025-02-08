package main

import (
	"net/http"

	"github.com/Asker231/auth-service.git/config"
	"github.com/Asker231/auth-service.git/internal/auth"
	"github.com/Asker231/auth-service.git/internal/user"
	"github.com/Asker231/auth-service.git/pkg/db"
)

func main() {


	//config
	cnf := config.LoadAppConfig()


	//init mux 
	router := http.NewServeMux()

	//init database
	dataBase := db.InitDataBase(&cnf.DataBase)

	//repositoryes
	userRepo := user.NewUserRepository(dataBase)

	//services
	authService := auth.NewAuthService(userRepo)

	//init handlers
	auth.NewAuthHandler(router,cnf,authService)

	//init server struct
	server := http.Server{
		Addr: ":8000",
		Handler: router,
	}
	server.ListenAndServe()



}