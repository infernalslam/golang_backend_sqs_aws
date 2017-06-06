package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
    "encoding/json"
    "fmt"
)

    // "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/sqs"


// type stuctrue
type Person struct {
    ID  string   `json:"id"`
}

type Hello struct {
    HELLO  string  `json:"hello"`
}

var people []Person
var say []Hello

func root(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(say)
}

func GetPeopleEndpoint (w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

func main() {
    people = append(people, Person{ID: "1"})
    say = append(say, Hello {HELLO: "selsuki"})

	router := mux.NewRouter()
	router.HandleFunc("/", root).Methods("GET")
    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")

    fmt.Println("start on -> http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}