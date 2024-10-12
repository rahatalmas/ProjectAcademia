package controller

import (
	"fmt"
	"net/http"
	"ServeUser/utility"
	"ServeUser/database"
)

func UserData(w http.ResponseWriter, r* http.Request){
	_,e := utility.VerifyToken(r)
	if e!=nil{
		fmt.Fprintf(w,"Invalid Token");
		return
	}
	conn := database.GetDB()
	query := "SELECT registration_year FROM student"
	rows,err := conn.Query(query)
	if err != nil{
		fmt.Println(err)
		return
	}
    for rows.Next(){
		var registration_year string
		if err := rows.Scan(&registration_year); err!=nil{
			fmt.Println(err)

		}
		fmt.Println(registration_year)
	}
	fmt.Fprintf(w,"user data");
	return 
}