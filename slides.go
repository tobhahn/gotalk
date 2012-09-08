package gotalk

import (
	"bytes"
	"net/http"
)

// slides is an HTTP handler that expects an :id query
// and returns the corresponding slide.
func slides(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get(":id")
	if id == "" {
		http.NotFound(w, req)
		return
	}

	defaultSlide := bytes.NewBufferString("<html/>")
	w.Write(defaultSlide.Bytes())
}
