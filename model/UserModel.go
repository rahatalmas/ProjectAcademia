package model

import (
	"fmt"
)

type Student struct {
    ID                    int       `json:"id"`
    StudentID            string    `json:"student_id"`
    FullName             string    `json:"full_name"`
    UserName             string    `json:"user_name"`
    Email                string    `json:"email"`
    Password             string    `json:"password"`
    ProfilePicture       string    `json:"profile_picture"`
    RegistrationYear      string `json:"registration_year"`
    RegistrationSemester string       `json:"registration_semester"`
    Department           string    `json:"department"`
    BatchNumber          string    `json:"batch_number"`
    CurrentSection       string    `json:"current_section"`
    Advisor              string    `json:"advisor"`
    SuperAdvisor         string    `json:"super_advisor"`
    CurrentCGPA          string   `json:"current_cgpa"`
    FinalProjectTitle    string    `json:"final_project_title"`
    FinalCGPA            string   `json:"final_cgpa"`
    RefreshToken         string    `json:"refresh_token"`
}

type StudentLog struct{
	Email string
	Password string
	Token string 
}

type LoginResponse struct{
	Token string  `json:"accesstoken"`
}

func Hello(){
	fmt.Println("hello i am a new student");
}