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

// Model for course - file
type Course struct {
	CourseId string `json:"courseId"`
	CourseName string `json:"courseName"`
	CoursePrice int `json:"-"`
	Author *Author `json:"author"`

}

// fake DB
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return  c.CourseName == ""
}

type Author struct {
	FullName string `json:"fullName"`
	Website string `json:"website"`

}

func main()  {
	fmt.Println("API")
	
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Karmsetu", Website: "Github"}})

	courses = append(courses, Course{CourseId: "2", CourseName: "Angular", CoursePrice: 109, Author: &Author{FullName: "BOI", Website: "YT"}})

	courses = append(courses, Course{CourseId: "3", CourseName: "Vue", CoursePrice: 169, Author: &Author{FullName: "LOL", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers - file
// serve Home route

func serveHome(w http.ResponseWriter , r *http.Request)  {
	w.Write([]byte("<h1>Welcome to API by Karmsetu</h1>"))
}

func getAllCourses(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab ID from request
	params := mux.Vars(r)

	// loop through course, find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return 
		}
	}

	json.NewEncoder(w).Encode("No course found with id:"+ params["id"])

	return
}

func createOneCourse(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside JSON")
		return
	}

	// generate Unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	// get id from request
	params := mux.Vars(r)
	// loop through the value, get id add whit my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}

// TODO send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove{index, index+1}, 
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index],courses[index+1:]... )

			break
		}
	}
}