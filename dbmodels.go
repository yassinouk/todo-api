package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// TaskStorage is an interface to interact with the task storage
type TaskStorage interface {
	ListTasks() ([]*Task, error)
	GetTaskByID(int64) (*Task, error)
	CreateTask(*Task) (*Task, error)
	UpdateTask(*Task) (*Task, error)
	DeleteTask(int64) error
}
type PostgresDb struct {
	db *sql.DB
}

func NewPostgresDb() (*PostgresDb, error) {
	// conecting to the database using the postgres driver and the connection string
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	connStr := os.Getenv("CONN_STR")
	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &PostgresDb{
		db: db,
	}, nil
}
func (s *PostgresDb) Init() error {
	return s.createTaskTable()
}
func (s *PostgresDb) createTaskTable() error {
	_, err := s.db.Exec(CreateTaskTable)
	return err
}

func (s *PostgresDb) ListTasks() ([]*Task, error) {
	rows, err := s.db.Query(listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []*Task{}
	for rows.Next() {
		task := &Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.Done, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func (s *PostgresDb) GetTaskByID(id int64) (*Task, error) {
	task := &Task{}
	err := s.db.QueryRow(getTask, id).Scan(&task.ID, &task.Title, &task.Content, &task.Done, &task.CreatedAt)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func (s *PostgresDb) CreateTask(task *Task) (*Task, error) {
	createdTask := &Task{}
	err := s.db.QueryRow(CreateTask, task.Title, task.Content, task.Done).Scan(&createdTask.ID, &createdTask.Title, &createdTask.Content, &createdTask.Done, &createdTask.CreatedAt)
	if err != nil {
		return nil, err
	}
	return createdTask, nil
}

func (s *PostgresDb) UpdateTask(task *Task) (*Task, error) {
	updatedTask := &Task{}
	err := s.db.QueryRow(updateTask, task.ID, task.Content).
		Scan(&updatedTask.ID, &updatedTask.Title, &updatedTask.Content, &updatedTask.Done, &updatedTask.CreatedAt)
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}
func (s *PostgresDb) DeleteTask(id int64) error {
	_, err := s.db.Exec(deleteTask, id)
	if err != nil {
		return err
	}
	return nil
}
