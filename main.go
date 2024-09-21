package main

import(
	"fmt"
	"net/http"
	"database/sql"
	"ServeUser/controller"

	_ "github.com/go-sql-driver/mysql"
)

var(
	DB,DbErr = sql.Open("mysql","root:@/academia_db_main");
)

type Server struct{
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r* http.Request){
    if r.URL.Path == "/api/usr/login" && r.Method == http.MethodPost {
        controller.Login(w,r);
	}else if r.URL.Path == "/api/usr/signup" && r.Method == http.MethodPost{
		controller.SignUp(w,r);
	}else {
		fmt.Fprintf(w,"Invalid Url");
	}
}

func main(){
	fmt.Println("hello pretty pretty <3");
	if DbErr!=nil{
		fmt.Println(DbErr);
	}
	fmt.Println(DB);
	http.ListenAndServe(":8080",new(Server));
}