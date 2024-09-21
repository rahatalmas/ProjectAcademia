package model

import (
	"fmt"
)

type Student struct{
	Id int
	Name string
}

func Hello(){
	fmt.Println("hello i am a new student");
}