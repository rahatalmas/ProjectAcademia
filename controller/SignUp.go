package controller

import (
	"fmt"
	"net/http"
)

func ControllerTest(){
	fmt.Println("package controller");
}

func SignUp(w http.ResponseWriter, r* http.Request){
	fmt.Fprintf(w,"SignUp Request "+r.Method);
}