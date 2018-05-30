package main

import (
	"fmt"
	"net/http"
	"flag"
	"log"
	"errors"
	"encoding/json"
	"todo_server/mydb"
)

var staticDirectory string
var serverHost string
var serverPort int

const GET string = "GET"
const POST string = "POST"
const PUT string = "PUT"

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	ResponseJson interface{} `json:"data"`
}

type ToDo struct {
	Id int32 `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

type TodoWrapper struct {
	id int32
	todos []ToDo
}

type TodoOperation struct {
	IdTodoWrapper int32 `json:"todoWrapperId"`
	Todo ToDo `json:"todo"`
}

func checkHTTPMethod(r *http.Request, method string) error {
	if r.Method == method {
		return nil
	}
	return errors.New("Wrong request method.")
}

func defineFlagVariables() {
	flag.StringVar(&staticDirectory, "dir", "./static", "Static     assets directory")
	flag.StringVar(&serverHost, "host", "0.0.0.0", "Default ser    ver host")
	flag.IntVar(&serverPort, "port", 3000, "Server port")
	flag.Parse()
}

func createResponseEncapsulation(success bool, message string, data interface{}) Response {
	response := Response {
		Success: success,
		Message: message,
		ResponseJson : data,
	}
	return response
}

func markTODOAsDone(todoOperation TodoOperation) {
	fmt.Printf("CHECK TODO %v AS %v. todoWrapperId: %d / todoId: %d.\n", todoOperation.Todo.Title, todoOperation.Todo.Done, todoOperation.IdTodoWrapper, todoOperation.Todo.Id)

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("TEST\n")
}

func handleTODOList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("LIST ALL TODOS")
}

func handleGetTODOSFromList(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET SPECIFIC TODO")
}

func handleTODOAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ADD TODO\n")
}

func handleTODORem(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("REMOVE TODO\n")
}

func handleTODOMarkDone(w http.ResponseWriter, r *http.Request) {
//TODO Create function to deal with the HTTP Method verification
	err := checkHTTPMethod(r, POST)
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, err.Error(), 500)
		return
	}

//TODO Create function to deal with JSON
	var jsonObj TodoOperation
	err = json.NewDecoder(r.Body).Decode(&jsonObj)
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	markTODOAsDone(jsonObj)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	response := createResponseEncapsulation(true, "Atividade adicionada com sucesso.", nil)

//TODO Create function to deal with the errors
	responseJson, jsErr := json.Marshal(response)
	if jsErr != nil {
		fmt.Println("error:", jsErr)
		http.Error(w, "Unable to respond", http.StatusBadRequest)
		return
	}

	w.Write(responseJson)
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/todo/list", handleTODOList)
	mux.HandleFunc("/todo/add", handleTODOAdd)
	mux.HandleFunc("/todo/rem", handleTODORem)
	mux.HandleFunc("/todo/mark/done", handleTODOMarkDone)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func startDB() {
	mydb.OpenConnection()
}

func main() {
	defineFlagVariables()
	startServer()
	startDB()
}
