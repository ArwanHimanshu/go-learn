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

type Courses struct {
	CourseId string `json:"courseid"`

	CourseName string `json:"coursename"`

	CoursePrice int `json:"price"`

	Author *Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`

	Website string `json:"website"`
}

var courses []Courses

func (c *Courses) IsEmpty() bool {

	return c.CourseName == ""
}

func main() {

	fmt.Println("building an api")

	r := mux.NewRouter()

	courses = append(courses, Courses{CourseId: "2", CourseName: "react js", CoursePrice: 1, Author: &Author{Name: "Himanshu", Website: "h.com"}})

	courses = append(courses, Courses{CourseId: "3", CourseName: "angular js", CoursePrice: 1, Author: &Author{Name: "Himanshu", Website: "h.com"}})

	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serverHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> welcome to learner</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get all courses")

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	fmt.Println("got params ", params)

	for _, course := range courses {

		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)

			return
		}
	}

	json.NewEncoder(w).Encode("no course found")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")

	w.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("data required")
	}

	var course Courses

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("data required")
		return
	}

	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("update one course")

	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {

		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)

			var course Courses
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]

			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)

			return

		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("delete one course")

	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {

		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted")
			break

		}
	}

}
