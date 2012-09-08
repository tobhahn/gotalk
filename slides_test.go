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

