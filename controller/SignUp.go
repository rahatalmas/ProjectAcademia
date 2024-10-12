package controller

import (
	"fmt"
	"net/http"
)

func ControllerTest(){
	fmt.Println("package controller");
}

func SignUp(w http.ResponseWriter, r* http.Request){
	cookie, err := r.Cookie("jwt")
    if err != nil {
        if err == http.ErrNoCookie {
            http.Error(w, "No cookie found", http.StatusUnauthorized)
            return
        }
        http.Error(w, "Error retrieving cookie", http.StatusInternalServerError)
        return
    }

    tokenString := cookie.Value
	fmt.Fprintf(w,"SignUp Request "+tokenString);
}