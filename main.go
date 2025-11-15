package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type taskServer struct {
	store *TaskStore
}

var taskStore TaskStore

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)
	server := NewTaskServer()
	router.HandleFunc("/task/", server.createTaskHandler).Methods(http.MethodPost)
	router.HandleFunc("/task/", server.getTasksHandler).Methods(http.MethodGet)
	router.HandleFunc("/task/{id:[0-9]+}", server.deleteTaskHandler).Methods(http.MethodDelete)
	router.HandleFunc("/task/{id:[0-9]+}", server.getTaskHandler).Methods(http.MethodGet)
	router.HandleFunc("/task/", server.deleteTasksHandler).Methods(http.MethodDelete)
	router.HandleFunc("/tag/{tag}", server.tagHandler).Methods(http.MethodGet)
	router.HandleFunc("/due/{year:[0-9]+}/{month:[0-9]+}/{day:[0-9]+}/", server.dueHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), router))
}

func NewTaskServer() *taskServer {
	store := New()
	return &taskServer{store: store}
}
