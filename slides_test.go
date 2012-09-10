package gotalk

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"testing"
)

type fakeSlide string
type fakeFinder int

func (f fakeFinder) FindID(id string) (data interface{}, err error) {
	if id == "this_id_does_not_exist" {
		return nil, errors.New("id '" + id + "' does not exist")
	}

	return fakeSlide("<html><head/><body/></html>"), nil
}

func (s fakeSlide) Render() []byte {
	return bytes.NewBufferString(string(s)).Bytes()
}

func setup_slide_tests() {
	slidesFinder = fakeFinder(42)
}

func Test_slides_expect_id_Query(t *testing.T) {
	setup_slide_tests()

	params := url.Values{":id": []string{"1"}}
	uri := url.URL{RawQuery: params.Encode()}
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		t.Fatalf("Error creating test request: %v", err)
	}

	var res probeResponseWriter
	slides(&res, req)

	if res.status != http.StatusOK {
		t.Fatalf("GET /slides/1 failed with status %v", http.StatusText(res.status))
	}
}

func Test_slides_without_id_should_return_not_found(t *testing.T) {
	setup_slide_tests()

	req, err := http.NewRequest("GET", "/slides", nil)
	if err != nil {
		t.Fatalf("Error creating test request: %v", err)
	}

	var res probeResponseWriter
	slides(&res, req)

	if res.status != http.StatusNotFound {
		t.Errorf("Expected /slides to return '%v', but got '%v'",
			http.StatusText(http.StatusNotFound), http.StatusText(res.status))
	}
}

func Test_slides_with_non_existent_id_should_return_not_found(t *testing.T) {
	setup_slide_tests()

	params := url.Values{":id": []string{"this_id_does_not_exist"}}
	uri := url.URL{RawQuery: params.Encode()}
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		t.Fatalf("Error creating test request: %v", err)
	}

	var res probeResponseWriter
	slides(&res, req)

	if res.status != http.StatusNotFound {
		t.Errorf("Expected /slides/this_id_does_not_exist to return '%v', but got '%v'",
			http.StatusText(http.StatusNotFound), http.StatusText(res.status))
	}
}
