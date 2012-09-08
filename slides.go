package gotalk

import (
	"bytes"
	"net/http"
)

// slides is an HTTP handler that expects an :id query
// and returns the corresponding slide.
func slides(w http.ResponseWriter, req *http.Request) {
	defaultSlide := bytes.NewBufferString("<html/>")
	w.Write(defaultSlide.Bytes())
}
