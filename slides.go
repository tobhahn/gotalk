package gotalk

import (
	"net/http"
)

type renderer interface {
	Render() []byte
}

type finder interface {
	FindID(id string) (data interface{}, err error)
}

// needs to be setup before use, e.g. setup_slides_tests() for testing
var slidesFinder finder

// slides is an HTTP handler that expects an :id query
// and returns the corresponding slide.
func slides(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get(":id")
	if id == "" {
		http.NotFound(w, req)
		return
	}

	slide, _ := slidesFinder.FindID(id)
	w.Write(slide.(renderer).Render())
}
