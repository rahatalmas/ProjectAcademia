package model

import (
	"fmt"
)

type StudentLog struct{
	Email string
	Password string
	Token string 
}

type LoginResponse struct{
	Email string `json:"email"`
	Token string  `json:"accesstoken"`
}

func Hello(){
	fmt.Println("hello i am a new student");
}