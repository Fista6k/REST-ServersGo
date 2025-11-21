package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type taskServer struct {
	store *TaskStore
}

var taskStore TaskStore

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	server := NewTaskServer()
	router.POST("/task/", server.createTaskHandler)
	router.GET("/task/", server.getTasksHandler)
	router.DELETE("/task/:id", server.deleteTaskHandler)
	router.GET("/task/:id", server.getTaskHandler)
	router.DELETE("/task/", server.deleteTasksHandler)
	router.GET("/tag/:tag", server.tagHandler)
	router.GET("/due/:year/:month/:day/", server.dueHandler)

	port := os.Getenv("SERVERPORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("start listening on %s", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}

func NewTaskServer() *taskServer {
	store := New()
	return &taskServer{store: store}
}
