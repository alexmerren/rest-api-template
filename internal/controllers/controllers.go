package controllers

import "net/http"

func TestAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("Api is working!"))
	default:
		w.Write([]byte("how tf u get here"))
	}
}
