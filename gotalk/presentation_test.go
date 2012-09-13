package gotalk

import (
	"testing"
)

func Test_presentation_slide_can_have_next(t *testing.T) {
	p := presentation{[]string{"title", "hello"}}

	if _, err := p.Next("title"); err != nil {
		t.Fatalf("Could not find next slide: %v", err)
	}
}

func Test_presentation_should_return_error_if_not_found(t *testing.T) {
	p := presentation{[]string{"title"}}

	if s, err := p.Next("foo"); err == nil {
		t.Fatalf("Found a next slide '%v' of 'foo' even though it does not exist in '%v'.", s, p.slides)
	}
}

func Test_presentation_should_return_error_if_slide_is_last(t *testing.T) {
	p := presentation{[]string{"title", "hello"}}

	if s, err := p.Next("hello"); err == nil {
		t.Fatalf("Found a next slide '%v' of 'hello' even though it is last '%v'.", s, p.slides)
	}
}
