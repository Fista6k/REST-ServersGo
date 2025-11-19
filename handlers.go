package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *taskServer) createTaskHandler(c *gin.Context) {
	type RequestTask struct {
		Text string    `json:"text"`
		Due  time.Time `json:"due"`
		Tags []string  `json:"tags"`
	}

	var rt RequestTask
	if err := c.ShouldBindJSON(&rt); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id := s.store.CreateTask(rt.Text, rt.Tags, rt.Due)
	c.JSON(http.StatusOK, gin.H{"Id": id})
}

func (s *taskServer) getTaskHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	task, err := s.store.GetTask(id)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (s *taskServer) getTasksHandler(c *gin.Context) {
	tasks := s.store.GetAllTasks()

	c.JSON(http.StatusOK, tasks)
}

func (s *taskServer) tagHandler(c *gin.Context) {
	tag := c.Params.ByName("tag")

	tasks := s.store.GetTasksByTag(tag)

	c.JSON(http.StatusOK, tasks)
}

func (s *taskServer) deleteTaskHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	err := s.store.DeleteTask(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
}

func (s *taskServer) deleteTasksHandler(c *gin.Context) {
	err := s.store.DeleteAllTasks()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
}

func (s *taskServer) dueHandler(c *gin.Context) {
	badRequestError := func() {
		c.String(http.StatusBadRequest, "expect /due/<year>/<month>/<day>, got %v", c.FullPath())
	}

	year, err := strconv.Atoi(c.Params.ByName("year"))
	if err != nil {
		badRequestError()
		return
	}

	month, err := strconv.Atoi(c.Params.ByName("month"))
	if month < 1 || month > 12 || err != nil {
		badRequestError()
		return
	}

	day, err := strconv.Atoi(c.Params.ByName("day"))
	if day < 1 || day > 31 || err != nil {
		badRequestError()
		return
	}

	tasks := s.store.GetTasksByDueDate(year, day, time.Month(month))

	c.JSON(http.StatusOK, tasks)
}
