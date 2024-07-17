package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
func main() {
	fmt.Println("API IN GOLANG")
	r := mux.NewRouter()

	//adding data into the slice
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{Fullname: "Mrudul Mohan", Website: "www.youtube.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN STACK", CoursePrice: 199, Author: &Author{Fullname: "Mrudul Mohan", Website: "www.udemy.com"}})

	//routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
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

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from the Request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the Response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course) //here we only have to return the maching course
			break

		}
	}
	json.NewEncoder(w).Encode("No course found with given id !")
	return

}

// create a course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if:body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("PLease send some data")
	}

	// what about the data like this -->  {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("NO DATA INSIDE JSON YOU ARE SENTING")
		return
	}
	// Loop through courses to check for duplicate CourseName
	for _, existingCourse := range courses {
		if existingCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course name already exists")
			return
		}
	}

	// generate a unique id and convert them to string
	// append the new course
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	course.CourseId = strconv.Itoa(random.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// update course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//first- grab id from Request
	params := mux.Vars(r)

	//loop through the value , once we get the value we remove the id
	//and add with my Id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) //courses[:index] creates a slice from the beginning of cour                                                             ses up to (but not including) index.
			//courses[index+1:] creates a slice from the element just after index to the end of the courses slice.
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//loop through the slice and when we find the id delete It
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}

	}

}
