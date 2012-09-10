package gotalk

import (
	"code.google.com/p/gorilla/pat"
)

var (
	Router *pat.Router = pat.New()
)

func init() {
	Router.Get("/compile", compile)
	Router.Get("/slides/{id}", slides)
}
