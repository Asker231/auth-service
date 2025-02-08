package auth

import (
	"errors"
	"fmt"

	"github.com/Asker231/auth-service.git/internal/user"
	"golang.org/x/crypto/bcrypt"
)


type AuthService struct{
	UserRepo *user.UserRepository
}

func NewAuthService(repo *user.UserRepository)*AuthService{
	return &AuthService{
		UserRepo: repo,
	}
}

//register
func(serv *AuthService)Register(name,email,password string)(*user.UserModel,error){
    
	//find user
	u := serv.UserRepo.FindByEmail(email)
	if u != nil{
		return nil, errors.New("Пользователь уже существует")
	}
	//generate password
	pass,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		fmt.Println(err.Error())
	}
	//create instns user
	user := user.UserModel{
		Name: name,
		Email: email,
		Password: string(pass),
	}
	//user create user method
    us,err := serv.UserRepo.CreateUser(user)
	if err != nil{
		return nil,err
	}
	return us,nil
}

//login
func(serv *AuthService)Login(email,password string)(*user.UserModel,error){
	u := serv.UserRepo.FindByEmail(email)
	if u == nil{
		return nil,errors.New("User not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password)); err != nil{
		return nil,err
	}
	return u,nil
}