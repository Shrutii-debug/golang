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

//model for courses - file
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	//injecting the author field here
	//type is a pointer
	Author *Author `json:"author"`
}

type Author struct{
	//We are using first letters as capital because we want to export this
	//but in jon we need all small so we use the `` option
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fake DB
var courses []Course

//middleware, helpers - file

func (c *Course)isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Printf(" APIs")
	r := mux.NewRouter()

	//seeding
courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", 
CoursePrice: 299, Author: &Author{Fullname: "Shruti", Website: "lco.dev"}})


courses = append(courses, Course{CourseId: "4", CourseName: "Mern stack", 
CoursePrice: 1299, Author: &Author{Fullname: "Shruti", Website: "go.dev"}})

//routing
r.HandleFunc("/", serveHome).Methods("GET")
r.HandleFunc("/courses", getAllCourses).Methods("GET")
r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
r.HandleFunc("/course", createOneCourse).Methods("POST")
r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


//listen to a port
log.Fatal(http.ListenAndServe(":4000", r))


}


//controllers - file

//serve Home route

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to APIs </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses) 
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request

	params := mux.Vars(r)

	//loop through courses, find matching id and return the response

	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	} 
	json.NewEncoder(w).Encode("no course found with given id")
	return
}

//injecting something into the database

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("create one courses")
	w.Header().Set("Content-Type", "application/json")

	//what if:body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}
	//what if data is being sent but it is empty - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) //in decode we have to pass a reference

	if course.isEmpty() {
		json.NewEncoder(w).Encode("no data inside JSON")
		return
	}

	//generate a unique id and convert to string, and 
	// then append this course to Courses

	rand.Seed(time.Now().UnixNano())
	// itoa converts whatever value given in int to string
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}
//TODO check only if title is duplicate
//loop through, title matches with course.courseid,send response that exact same exists

//updating a course
func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("update one courses")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through values, once we get the id we work on the remove option
	//add operation again to update, add again with my ID

	for index, course := range courses{
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]... )

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]

			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO send a response when ID is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("delete one courses")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id, remove(index, index + 1)

	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]... )
			//TODO send a confirm or deny response
			break
		}
	}
}


