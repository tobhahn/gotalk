package gotalk

import (
	"bytes"
	"net/http"
)

func compile(w http.ResponseWriter, req *http.Request) {
	w.Write(bytes.NewBufferString("").Bytes())
}
