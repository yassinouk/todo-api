package main

import (
	"fmt"
	"log"
)

func seedAccount(store TaskStorage, title, content string) *Task {
	task, err := NewTask(title, content)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := store.CreateTask(task); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new task created => ", task.ID)

	return task
}

func seedTasks(s TaskStorage) {
	seedAccount(s, RandomTaskTitle(), RandomTaskContent())
}
func main() {
	db, err := NewPostgresDb()
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	seedTasks(db)
	fmt.Printf("store: %+v\n", db)
	server := NewAPIServer(":8081", db)
	server.Run()
}
