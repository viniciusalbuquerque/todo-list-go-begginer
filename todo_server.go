package main

import (
	"io"
	"fmt"
	"net/http"
	"flag"
	"log"
	"github.com/gorilla/mux"
)

var staticDirectory string
var serverHost string
var serverPort int

type ToDo struct {
	id int32
	activity string
}

type TodoWrapper struct {
	id int32
	todos []ToDo
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

func handleTODOAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ADD TODO\n")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w,`{"success": true,"message": "Atividade adicionada com sucesso."}`)
}

func handleTODORem(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("REMOVE TODO\n")
}

func startServer() {
	router := mux.NewRouter()
	router.HandleFunc("/todo/list", handleTODOList).Methods("GET")
	router.HandleFunc("/todo/add", handleTODOAdd).Methods("PUT")
	router.HandleFunc("/todo/rem", handleTODORem).Methods("POST")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", serverHost, serverPort), router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	defineFlagVariables()
	startServer()
}
