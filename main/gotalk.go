package main

import (
	"gotalk/gotalk"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type loadSlidesFinder string

func (dir loadSlidesFinder) FindID(id string) (data interface{}, err error) {
	slide, err := ioutil.ReadFile(filepath.Join(string(dir), id+".html"))
	if err != nil {
		return nil, err
	}
	return string(slide), nil
}

func main() {
	gotalk.P = gotalk.Presentation{[]string{
		"title",
		"hello-world",
		"packages",
		"types",
		"functions",
		"function-return",
		"function-literals",
		"methods",
		"for-loop",
		"range",
		"structs",
		"composition",
		"interfaces",
		"errors",
		"panic",
		"defer",
		"concurrency",
		"channels",
	}}

	slidesDir, err := filepath.Abs("../slides")
	if err != nil {
		panic(err)
	}
	gotalk.SlidesFinder = loadSlidesFinder(slidesDir)

	http.Handle("/", gotalk.Router)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
