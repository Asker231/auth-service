package req

import (
	"encoding/json"
	"io"
)


func DecodeBody[T any](r io.ReadCloser)(*T,error){
	var payload T
	err := json.NewDecoder(r).Decode(&payload)
	if err != nil{
		return nil,err
	}
	return &payload,nil	
}