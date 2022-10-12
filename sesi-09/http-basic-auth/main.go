package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const USERNAME = "superman"
const PASSWORD = "secret"

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, v := range students {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}
	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "bourne", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "ethan", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "wick", Grade: 3})
}

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":8081"

	fmt.Println("server started at localhost:8081")
	server.ListenAndServe()
}
