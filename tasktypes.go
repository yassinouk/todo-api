package main

import (
	"math/rand"
	"time"
)

type CreateTaskRequest struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
type GetTaskRequest struct {
	ID int64 `json:"id"`
}
type UpdateTaskRequest struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
type DeleteTaskRequest struct {
	ID int64 `json:"id"`
}

type Task struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewTask(title, content string) (*Task, error) {
	return &Task{
		Title:     title,
		Content:   content,
		Done:      rand.Intn(2) == 1,
		CreatedAt: time.Now().UTC(),
	}, nil
}
