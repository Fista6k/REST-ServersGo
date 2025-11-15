package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *taskServer) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling task create at %s\n", r.URL.Path)

	type RequestTask struct {
		Text string    `json"text"`
		Due  time.Time `json"due"`
		Tags []string  `json:"tags"`
	}

	type ResponceId struct {
		Id int `json:"id"`
	}

	contentType := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		http.Error(w, "Expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var rt RequestTask
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := s.store.CreateTask(rt.Text, rt.Tags, rt.Due)
	renderJson(w, ResponceId{Id: id})
}

func (s *taskServer) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling get task at %s\n", r.URL.Path)

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
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

	tag := mux.Vars(r)["tag"]

	tasks := s.store.GetTasksByTag(tag)

	renderJson(w, tasks)
}

func (s *taskServer) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling delete task at %s\n", r.URL.Path)

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := s.store.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (s *taskServer) deleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling delete tasks at %s\n", r.URL.Path)

	err := s.store.DeleteAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (s *taskServer) dueHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling get due at %s\n", r.URL.Path)

	badRequestError := func() {
		http.Error(w, fmt.Sprintf("expect /due/<year>/<month>/<day>, got %v", r.URL.Path), http.StatusBadRequest)
	}

	vars := mux.Vars(r)
	year, _ := strconv.Atoi(vars["year"])
	month, _ := strconv.Atoi(vars["month"])
	if month < 1 || month > 12 {
		badRequestError()
		return
	}
	day, _ := strconv.Atoi(vars["day"])
	if day < 1 || day > 31 {
		badRequestError()
		return
	}

	tasks := s.store.GetTasksByDueDate(year, day, time.Month(month))

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
