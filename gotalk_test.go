package gotalk

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"testing"
)

func Test_compile_should_dispatch_to_compile(t *testing.T) {
	params := url.Values{"q": []string{hello_世界}}
	uri := url.URL{Path: "/compile", RawQuery: params.Encode()}
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		t.Fatalf("Error creating test request: %v", err)
	}

	var res probeResponseWriter
	router.ServeHTTP(&res, req)

	if res.status != http.StatusOK {
		t.Fatalf("compile failed with status %v", http.StatusText(res.status))
	}

	if res.response != "Hello, 世界\n" {
		t.Log(`Expected: Hello, 世界\n`)
		t.Errorf("Actual: %v", res.response)
	}
}

func Test_slides_1_should_return_a_slide(t *testing.T) {
	setup_slide_tests()

	req, err := http.NewRequest("GET", "/slides/1", nil)
	if err != nil {
		t.Fatalf("Error creating test request: %v", err)
	}

	var res probeResponseWriter
	router.ServeHTTP(&res, req)

	if res.status != http.StatusOK {
		t.Errorf("could not get /slides/1: %v", http.StatusText(res.status))
	}
}

type wepPage struct {
	XMLName xml.Name `xml:"html"`
	Body    string   `xml:"body"`
}

func Test_slides_1_should_return_well_formed_html(t *testing.T) {
	setup_slide_tests()

	req, err := http.NewRequest("GET", "/slides/1", nil)
	if err != nil {
		t.Fatalf("Error creating test request: %v", err)
	}

	var res probeResponseWriter
	router.ServeHTTP(&res, req)

	if res.status != http.StatusOK {
		t.Fatalf("could not get /slides/1: %v", http.StatusText(res.status))
	}

	slide := wepPage{Body: "none"}

	err = xml.Unmarshal([]byte(res.response), &slide)
	if err != nil {
		t.Fatalf("Could not parse response '%v': %v", res.response, err.Error())
	}

	if slide.Body == "none" {
		t.Errorf("Response '%v' does not contain a body", res.response)
	}
}
