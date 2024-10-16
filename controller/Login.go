package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"ServeUser/view"
	"ServeUser/model"
    "ServeUser/utility"
	"time"
)

type ErrorResponse struct {
    Error string `json:"message"`
}


func Login(w http.ResponseWriter, r* http.Request){
	contentType := r.Header.Get("Content-Type");
	fmt.Println("loing Req..., Content Type: "+contentType);
	var user model.StudentLog;
    var email string
	var password string
    var token string = ""
	//check content types of data
	switch contentType{
	case "application/json":
		defer r.Body.Close();
		//decode json data
		if err := json.NewDecoder(r.Body).Decode(&user);err!=nil{
			fmt.Println("json err ",err);
		    e:= ErrorResponse{Error:"json Content Types Error"}
			view.JsonView(w,r,e);
			return
		}
		email = user.Email
		password = user.Password

	case "application/x-www-form-urlencoded":
		err := r.ParseForm();
		if err != nil{
			fmt.Fprintf(w,"parsing form failed");
			return
		}
        email = r.FormValue("email");
		password = r.FormValue("password");

	default:
		e:= ErrorResponse{Error:"Unknown Content Types"}
		view.JsonView(w,r,e);
		return
	}

	//validate required data to process
	if email == "" || password == ""{
		e:= ErrorResponse{Error:"Not Enough Data"}
		view.JsonView(w,r,e);
	return
	}

	//generate access and refresh tokens
	token,err := utility.CreateToken(email)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(token)
    http.SetCookie(w, &http.Cookie{
        Name:     "jwt",
        Value:    token,
        Path:     "/",
        Expires:  time.Now().Add(1 * time.Hour),
        HttpOnly: true,
        Secure:   false,
    })

	//return successful login with tokens
	responseData := model.LoginResponse{
		Token:token,
	}
	view.JsonView(w,r,responseData);
}