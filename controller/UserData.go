package controller

import (
	"fmt"
	"net/http"
	"ServeUser/utility"
	"ServeUser/database"
	"ServeUser/model"
)

func UserData(w http.ResponseWriter, r* http.Request){
	_,e := utility.VerifyToken(r)
	if e!=nil{
		fmt.Fprintf(w,"Invalid Token");
		return
	}
	var student model.Student
	database.GetData("student",student)
	fmt.Fprintf(w,"user data");
	return 
}