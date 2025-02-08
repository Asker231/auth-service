package user

import "gorm.io/gorm"


type UserRepository struct{
	Db *gorm.DB
}

func NewUserRepository(dataBase *gorm.DB)*UserRepository{
	return &UserRepository{
		Db: dataBase,
	}
}

//create user
func(repo *UserRepository)CreateUser(user UserModel)(*UserModel,error){
	result := repo.Db.Create(&user)
	if result.Error != nil{
		return nil,result.Error
	}
	return &user,nil
}	

//finduser
func(repo *UserRepository)FindByEmail(email string)(*UserModel){
	var payload UserModel
	result := repo.Db.First(&payload,"email = ?",email)

	if result.Error != nil{
		return nil
	}
	return &payload
}