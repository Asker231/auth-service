package main

import (
	"fmt"
	"os"

	"github.com/Asker231/auth-service.git/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Println(err.Error())
	}
	conn,err := gorm.Open(postgres.Open(os.Getenv("DNS")),&gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
	}

	conn.AutoMigrate(user.UserModel{})

}