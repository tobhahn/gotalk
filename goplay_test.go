package gotalk

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

type probeResponseWriter struct {
	header   http.Header
	response string
	status   int
}

func (w *probeResponseWriter) Header() http.Header {
	if w.header == nil {
		w.header = make(http.Header)
	}

	return w.header
}

func (w *probeResponseWriter) Write(data []byte) (int, error) {
	buf := bytes.NewBuffer(data)
	w.response += buf.String()

	if buf.Len() > 0 {
		panic(fmt.Sprintf("Could not append all of data, %v bytes left.", buf.Len()))
	}

	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}

	return len(data), nil
}

func (w *probeResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

func Test_compile_should_be_able_to_handle_an_empty_request(t *testing.T) {
	req, err := http.NewRequest("GET", "/compile", strings.NewReader(""))
	if err != nil {
		t.Fatalf("Error creating test request: %v", err)
	}

	var res probeResponseWriter
	compile(&res, req)

	if res.status != http.StatusOK {
		t.Errorf("compile failed with status %v", http.StatusText(res.status))
	}
}
