package main

// db connection string
const (
	//sql queries
	// CreateTaskTable is a SQL query to create the tasks table
	CreateTaskTable = `CREATE TABLE IF NOT EXISTS tasks (
		"id" serial NOT NULL,
		"title" varchar NOT NULL,
		"content" text NOT NULL,
		"done" boolean NOT NULL DEFAULT false,
		"created_at" timestamptz NOT NULL DEFAULT (now())
	  );`
	// CreateTask is a SQL query to create a new task
	CreateTask = `-- name: CreateTask :one
INSERT INTO tasks (
  title, content, done
) VALUES (
  $1, $2, $3
)
RETURNING id, title, content, done, created_at
`
	// DeleteTask is a SQL query to delete a task
	deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1
`
	// GetTask is a SQL query to get a task by ID
	getTask = `-- name: GetTask :one
SELECT id, title, content, done, created_at FROM tasks
WHERE id = $1 LIMIT 1
`
	// ListTasks is a SQL query to list tasks
	listTasks = `-- name: ListTasks :many
SELECT id, title, content, done, created_at FROM tasks
ORDER BY id
`
	// UpdateTask is a SQL query to update a task
	updateTask = `-- name: UpdateTask :one
UPDATE tasks
SET content =$2
WHERE id = $1
RETURNING id, title, content, done, created_at
`
)
