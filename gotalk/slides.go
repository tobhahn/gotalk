package gotalk

import (
	"bytes"
	"html/template"
	"net/http"
	"reflect"
)

type Finder interface {
	FindID(id string) (data interface{}, err error)
}

// needs to be setup before use, e.g. setup_slides_tests() for testing
var SlidesFinder Finder

var slidesTemplate = template.Must(template.ParseFiles(
	"../templates/_base.html",
	"../templates/slide.html",
))

// slides is an HTTP handler that expects an :id query
// and returns the corresponding slide.
func slides(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get(":id")
	if id == "" {
		http.NotFound(w, req)
		return
	}

	slide, err := SlidesFinder.FindID(id)
	if err != nil {
		http.NotFound(w, req)
		w.Write(bytes.NewBufferString(err.Error()).Bytes())
		return
	}

	slidesTemplate.Execute(w, template.HTML(reflect.ValueOf(slide).String()))
}
