package gotalk

import (
	"code.google.com/p/gorilla/pat"
)

var (
	router pat.Router
)

func init() {
	router.Get("/compile", compile)
}
