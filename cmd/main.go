package main

import (
	"net/http"

	"github.com/Asker231/auth-service.git/internal/auth"
)

func main() {
	//config
	
	//init mux 
	router := http.NewServeMux()
	//init handlers
	auth.NewAuthHandler(router)
}