package api

import (
	"encoding/json"
	"net/http"
	"todo/internal/task"

	"github.com/gorilla/mux"
)

// Server type is used to encapsulate all the necessary information for a server.
type Server struct {
	Port   int
	Router *mux.Router
	Tasks  []*task.Task
}

// MakeRoutesAndStart function is called by the main function to create the server.
func MakeRoutesAndStart(port int) *Server {
	r := mux.NewRouter()
	t := make([]*task.Task, 0)
	s := &Server{
		Port:   port,
		Router: r,
		Tasks:  t,
	}
	// This is where we add the routes, and their handlers.
	r.HandleFunc("/api", s.listTasks).Methods("GET")
	r.HandleFunc("/api", s.createTask).Methods("POST")
	return s
}

func (s *Server) listTasks(w http.ResponseWriter, r *http.Request) {
	// POST all of the contents of the tasks in the server struct.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s.Tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	// Take in the request body and turn that into a task object.
	task := &task.Task{}
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Add the task into the tasks struct.
	s.Tasks = append(s.Tasks, task)

	// Respond with the information taken into the program.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
