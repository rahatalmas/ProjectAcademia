package main

import(
	"fmt"
	"net/http"
	"ServeUser/controller"
    "ServeUser/database"
)


type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r* http.Request){
    if r.URL.Path == "/api/usr/login" && r.Method == http.MethodPost {
        controller.Login(w,r);
	}else if r.URL.Path == "/api/usr/signup" && r.Method == http.MethodPost{
		controller.SignUp(w,r);
	}else if r.URL.Path == "/api/usr/data" && r.Method == http.MethodGet{
		controller.UserData(w,r);
	}else {
		fmt.Fprintf(w,"Invalid request");
	}
}

func main(){
	if err := database.DBInit(); err!=nil{
		fmt.Println(err)
		return
	}
	server := new(Server)
	http.ListenAndServe(":8080",server);
}