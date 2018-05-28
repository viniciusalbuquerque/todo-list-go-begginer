package main

import (
	"io"
	"fmt"
	"net/http"
	"flag"
	"log"
	"errors"
	"encoding/json"
)

var staticDirectory string
var serverHost string
var serverPort int

const GET string = "GET"
const POST string = "POST"
const PUT string = "PUT"

type ToDo struct {
	id int32
	activity string
}

type TodoWrapper struct {
	id int32
	todos []ToDo
}

type TodoOperation struct {
	IdTodoWrapper int32
	IdTodo int32
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

	err := checkHTTPMethod(r, PUT)
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, err.Error(), 500)
		return
	}


	var jsonObj TodoOperation
	err = json.NewDecoder(r.Body).Decode(&jsonObj)
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w,`{"success": true,"message": "Atividade adicionada com sucesso."}`)
}

func handleTODORem(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("REMOVE TODO\n")
}

/*func startServer() {
	router := mux.NewRouter()
	router.HandleFunc("/todo/list", handleTODOList).Methods("GET")
	router.HandleFunc("/todo/list/:id", handleGetTODOSFromList).Methods("GET")
	router.HandleFunc("/todo/add", handleTODOAdd).Methods("PUT")
	router.HandleFunc("/todo/rem", handleTODORem).Methods("POST")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}*/

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/todo/list", handleTODOList)
	mux.HandleFunc("/todo/add", handleTODOAdd)
	mux.HandleFunc("/todo/rem", handleTODORem)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func main() {
	defineFlagVariables()
	startServer()
}
