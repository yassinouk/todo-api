package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandlerfunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			log.Println(err)
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type ApiServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
	}

}
func (s *ApiServer) Run() {
	router := mux.NewRouter()
	log.Println("API Server started and running on port", s.listenAddr)
	router.HandleFunc("/task", makeHTTPHandlerfunc(s.handleTask))
	http.ListenAndServe(s.listenAddr, router)

}
func (s *ApiServer) handleTask(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetTask(w, r)
	case http.MethodPost:
		return s.handleCreateTask(w, r)
	case http.MethodDelete:
		return s.handleDeleteTask(w, r)
	case http.MethodPut:
		return s.handleUpdateTask(w, r)
	default:
		return fmt.Errorf("unsupported method %s", r.Method)
	}
}
func (s *ApiServer) handleGetTask(w http.ResponseWriter, r *http.Request) error {
	//log.Println("handleGetTask is running")
	return nil
}
func (s *ApiServer) handleCreateTask(w http.ResponseWriter, r *http.Request) error {
	//log.Println("handleCreateTask is running")
	return nil
}
func (s *ApiServer) handleDeleteTask(w http.ResponseWriter, r *http.Request) error {
	//log.Println("handleDeleteTask is running")
	return nil
}
func (s *ApiServer) handleUpdateTask(w http.ResponseWriter, r *http.Request) error {
	//log.Println("handleUpdateTask is running")
	return nil
}
