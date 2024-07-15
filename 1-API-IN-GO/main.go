package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// model for course- file
type Course struct {
	CourseId    string  `json: "courseid"`
	CourseName  string  `json: "coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

//model for the author of the courses

type Author struct {
	Fullname string `json: "fullname"`
	Website  string `json: "website"`
}

// fake DB
var courses []Course

// middleware,helper -file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {
	fmt.Println("API IN GOLANG")
}

// controllers- file
// serve home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome guys , iam kinda trying to create an API here !</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json") //Setting the Response Header
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Response) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from the Request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the Response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course) //here we only have to return the maching course

		}
	}
	json.NewEncoder(w).Encode("No course found with given id !")
	return

}
