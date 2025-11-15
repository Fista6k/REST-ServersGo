package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (s *taskServer) createTaskHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *taskServer) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling get task at %s\n", r.URL.Path)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task, err := s.store.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	renderJson(w, task)
}

func (s *taskServer) getTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling get tasks at %s\n", r.URL.Path)

	tasks := s.store.GetAllTasks()

	renderJson(w, tasks)
}

func (s *taskServer) tagHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling get tasks/{tag} at %s\n", r.URL.Path)

	tag := r.PathValue("tag")

	tasks := s.store.GetTasksByTag(tag)

	renderJson(w, tasks)
}

func (s *taskServer) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling delete task at %s\n", r.URL.Path)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = s.store.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (s *taskServer) deleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling delete tasks at %s\n", r.URL.Path)

	err := s.store.DeleteAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (s *taskServer) dueHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling get due at %s\n", r.URL.Path)

	year, err := strconv.Atoi(r.PathValue("year"))
	if err != nil {
		http.Error(w, "invalid year", http.StatusBadRequest)
		return
	}
	month, err := time.Parse("January", r.PathValue("month"))
	if err != nil {
		http.Error(w, "invalid month", http.StatusBadRequest)
		return
	}

	day, err := strconv.Atoi(r.PathValue("day"))
	if err != nil {
		http.Error(w, "invalid day", http.StatusBadRequest)
		return
	}
	if day < 1 || day > 31 {
		http.Error(w, "invalid day", http.StatusBadRequest)
		return
	}

	tasks := s.store.GetTasksByDueDate(year, day, month)

	renderJson(w, tasks)
}

func renderJson(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
