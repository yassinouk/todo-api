package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
	taskstore  TaskStorage
}

func NewAPIServer(listenAddr string, taskstore TaskStorage) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		taskstore:  taskstore,
	}

}
func (s *ApiServer) Run() {
	router := mux.NewRouter()
	log.Println("API Server started and running on port", s.listenAddr)
	router.HandleFunc("/task", s.handleListTasks()).Methods("GET")
	router.HandleFunc("/task/{id}", s.handleGetTaskByID()).Methods("GET")
	router.HandleFunc("/task/", s.handleCreateTask()).Methods("POST")
	router.HandleFunc("/task/{id}", s.handleUpdateTask()).Methods("PUT")
	router.HandleFunc("/task/{id}", s.handleDeleteTask()).Methods("DELETE")
	http.ListenAndServe(s.listenAddr, router)

}
func (s *ApiServer) handleListTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := s.taskstore.ListTasks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := WriteJSON(w, http.StatusOK, tasks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *ApiServer) handleGetTaskByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task, err := s.taskstore.GetTaskByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := WriteJSON(w, http.StatusOK, task); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *ApiServer) handleCreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// updated comment: the think is i should decode the request body and extract the necessary information to create the task
		task, err := s.taskstore.CreateTask(&Task{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := WriteJSON(w, http.StatusOK, task); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *ApiServer) handleDeleteTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := s.taskstore.DeleteTask(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := WriteJSON(w, http.StatusOK, map[string]int64{"deleted": id}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *ApiServer) handleUpdateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// i should be able to update the task with the id

		id, err := getID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// updated comment: the think is i should decode the request body with the a similar logic to creat task and extract the necessary information to update the task
		task, err := s.taskstore.UpdateTask(&Task{
			ID:        id,
			CreatedAt: time.Time{},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := WriteJSON(w, http.StatusOK, task); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
func getID(r *http.Request) (int64, error) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, fmt.Errorf("id parameter not found in URL")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}
