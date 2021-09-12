package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/internal/task"

	"github.com/gorilla/mux"
)

type Server struct {
	Port   int
	Router *mux.Router
	Tasks  []*task.Task
}

func MakeRoutes(port int) *Server {
	r := mux.NewRouter()
	t := make([]*task.Task, 0)
	s := &Server{
		Port:   port,
		Router: r,
		Tasks:  t,
	}
	// This is where we add the routes, and their handlers.
	r.HandleFunc("/api", s.listAllTasks).Methods("GET")
	r.HandleFunc("/api", s.createTask).Methods("POST")
	return s
}

func (s *Server) StartServer(host string, port int) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), s.Router)
}

func (s *Server) listAllTasks(w http.ResponseWriter, r *http.Request) {
	// POST all of the contents of the tasks in the server struct.
	if err := json.NewEncoder(w).Encode(s.Tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	// Take in the request body and turn that into a task object.
	task := task.NewTask()
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Add the task into the tasks struct.
	s.Tasks = append(s.Tasks, task)

	// Respond with the information taken into the program.
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
}
