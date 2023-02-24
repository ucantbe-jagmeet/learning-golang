package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for course - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fake DB 
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool{
	return c.CourseId == "" && c.CourseName ==""
}

func main() {
	fmt.Println("Building API in go lang")
}

//controllers - file
// serve home route

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1> Welcome to Api by LearnCodeOnline</h1>"))
}


func getAllcourses( w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "apllication/json")
	json.NewEncoder(w).Encode(courses)
}
