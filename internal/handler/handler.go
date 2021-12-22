package handler

import "net/http"

func TestHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test!"))
}
