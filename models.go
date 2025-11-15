package main

import "time"

type Task struct {
	ID   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type TaskStore struct {
	Tasks map[int]Task
}

func (ts *TaskStore) New() *TaskStore {
	tasks := make(map[int]Task)
	return &TaskStore{Tasks: tasks}
}

func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	task := Task{
		
	}
}

func (ts *TaskStore) GetTask(id int) (Task, error) {

}

func (ts *TaskStore) DeleteTask(id int) error {

}

func (ts *TaskStore) DeleteAllTasks() error {

}

func (ts *TaskStore) GetAllTasks() []Task {

}

func (ts *TaskStore) GetTasksByTag(tag string) []Task {

}

func (ts *TaskStore) GetTasksByDueDate(year, day int, month time.Time) []Task {

}
