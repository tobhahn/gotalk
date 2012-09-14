package gotalk

import (
	"code.google.com/p/gorilla/pat"
	"net/http"
)

var (
	Router *pat.Router = pat.New()
)

func init() {
	Router.Get("/compile", compile)
	Router.Get("/slides/{id}", slides)
	Router.Add("GET", "/css", http.FileServer(http.Dir("../assets")))
	Router.Add("GET", "/js", http.FileServer(http.Dir("../assets")))
	Router.Add("GET", "/img", http.FileServer(http.Dir("../assets")))
	Router.Add("GET", "/", http.RedirectHandler("/slides/title", http.StatusFound))
}
