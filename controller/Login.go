package controller

import (
	"fmt"
	"net/http"
	"ServeUser/view"
	"ServeUser/model"
)

func Login(w http.ResponseWriter, r* http.Request){
	contentType := r.Header.Get("Content-Type");
	fmt.Println("loing Req..., Content Type: "+contentType);
	user := model.Student{Id:1,Name:"Almas"};
    view.JsonView(w,r,user);
}