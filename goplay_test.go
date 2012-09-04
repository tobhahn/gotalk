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

	if buf.Len() != len(data) {
		panic(
			fmt.Sprintf(
				"Internal error: buf.Len() (== %v) is not equal to len(data) (== %v). Response: %v",
				buf.Len(), len(data), w.response))
	}

	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}

	return len(data), nil
}

func (w *probeResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
}

func Test_compiling_an_empty_request_should_return_404_compiler_error(t *testing.T) {
	req, err := http.NewRequest("GET", "/compile", strings.NewReader(""))
	if err != nil {
		t.Fatalf("Error creating test request: %v", err)
	}

	var res probeResponseWriter
	compile(&res, req)

	if res.status != http.StatusNotFound {
		t.Errorf("compile failed with status %v", http.StatusText(res.status))
		t.Errorf("response: %v", res.response)
	}
}

var hello_世界 = `
package main

import ("fmt")

func main() {
	fmt.Println("Hello, 世界")
}
`

func Test_compile_hello_世界(t *testing.T) {
	req, err := http.NewRequest("GET", "/compile", strings.NewReader(hello_世界))
	if err != nil {
		t.Fatalf("Error creating test request: %v", err)
	}

	var res probeResponseWriter
	compile(&res, req)

	if res.status != http.StatusOK {
		t.Fatalf("compile failed with status %v", http.StatusText(res.status))
	}

	if res.response != "Hello, 世界\n" {
		t.Log(`Expected: Hello, 世界\n`)
		t.Errorf("Actual: %v", res.response)
	}
}
