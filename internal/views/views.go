package views

import "net/http"

func TestAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Api is working!"))
}
