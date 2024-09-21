package controller

import (
	"fmt"
	"net/http"
	"ServeUser/view"
	"ServeUser/model"
)

func Login(w http.ResponseWriter, r* http.Request){
	fmt.Println("loing func call");
	user := model.Student{Id:1,Name:"Almas"};
    view.JsonView(w,r,user);
}