package gotalk

import (
	"bytes"
	"encoding/xml"
	"html/template"
	"net/http"
	"reflect"
)

// the one and only presentation
var Presentation presentation

type Finder interface {
	FindID(id string) (data interface{}, err error)
}

// needs to be setup before use, e.g. setup_slides_tests() for testing
var SlidesFinder Finder

var slidesTemplate = template.Must(template.ParseFiles(
	"../templates/_base.html",
	"../templates/slide.html",
))

type slideParams struct {
	slideFields
	pager
}

type slideFields struct {
	XMLName    xml.Name `xml:"Slide"`
	Title      string
	Codesample string
	Notes      notes
}

type notes struct {
	HTML template.HTML `xml:",innerxml"`
}

type pager struct {
	Prev string
	Next string
}

// slides is an HTTP handler that expects an :id query
// and returns the corresponding slide.
func slides(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get(":id")
	if id == "" {
		http.NotFound(w, req)
		return
	}

	data, err := SlidesFinder.FindID(id)
	if err != nil {
		http.NotFound(w, req)
		w.Write(bytes.NewBufferString(err.Error()).Bytes())
		return
	}

	prev, _ := Presentation.Prev(id)
	next, _ := Presentation.Next(id)

	slide := reflect.ValueOf(data).String()

	var fields slideFields
	err = xml.Unmarshal([]byte(slide), &fields)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse slide: " + err.Error() + "\n"))
		w.Write([]byte("Data:\n" + slide))
		return
	}

	slidesTemplate.Execute(w, slideParams{fields, pager{prev, next}})
}
