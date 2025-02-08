package req

import "github.com/go-playground/validator/v10"




func IsValid[T any](item T)error{
	validate := validator.New()
	err := validate.Struct(item)
	return err
}