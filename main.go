package main

import (
	"log"
	"net/http"
	"os"
)

type taskServer struct {
	store *TaskStore
}

var taskStore TaskStore

func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()
	mux.HandleFunc("POST /task/", server.createTaskHandler)
	mux.HandleFunc("GET /task/", server.getTasksHandler)
	mux.HandleFunc("DELETE /task/{id}", server.deleteTaskHandler)
	mux.HandleFunc("GET /task/{id}", server.getTaskHandler)
	mux.HandleFunc("DELETE /task/", server.deleteTasksHandler)
	mux.HandleFunc("GET /tag/{tag}", server.tagHandler)
	mux.HandleFunc("GET /due/{year}/{month}/{day}/", server.dueHandler)
	
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), mux))
}

func NewTaskServer() *taskServer {
	store := taskStore.New()
	return &taskServer{store: store}
}
