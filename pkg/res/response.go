package res

import (
	"encoding/json"
	"fmt"
	"net/http"
)



func Response(w http.ResponseWriter,data any,STATUScode int){
	err := json.NewEncoder(w).Encode(data)
	if err != nil{
		fmt.Println(err.Error())
	}
	w.WriteHeader(STATUScode)
}