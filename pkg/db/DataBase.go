package db

import (
	"fmt"

	"github.com/Asker231/auth-service.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDataBase(conf *config.DataBaseConfig)*gorm.DB{
	conn,err := gorm.Open(postgres.Open(conf.DNS),&gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
	}
	return conn
}