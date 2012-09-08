package gotalk

import (
	"net/http"
	"net/url"
	"testing"
)

func Test_slides_expect_id_Query(t *testing.T) {
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
