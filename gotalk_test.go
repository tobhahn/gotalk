package gotalk

import (
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
