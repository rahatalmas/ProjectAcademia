package controller

import (
	"fmt"
	"net/http"
	"ServeUser/utility"
	"ServeUser/database"
	"ServeUser/model"
	"ServeUser/view"
)

func UserData(w http.ResponseWriter, r* http.Request){
	_,e := utility.VerifyToken(r)
	if e!=nil{
		fmt.Fprintf(w,"Invalid Token");
		return
	}
	var students []model.Student
	var student model.Student
    result, err := database.GetData("student", student)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(result)
    // Type assert the result back to a slice of User
    /*for _, item := range result {
        user, ok := item.(model.Student)
        if ok {
            students = append(students, user)
        }
    }*/

    // Output the results
    fmt.Println(students)
	//fmt.Fprintf(w,"user data");
	view.JsonView(w,r,result)
	return 
}