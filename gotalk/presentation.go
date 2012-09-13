package gotalk

import (
	"errors"
)

type presentation struct {
	slides []string
}

func indexOf(x string, slice []string) (int, error) {
	for i, v := range slice {
		if x == v {
			return i, nil
		}
	}
	return -1, errors.New("Not found: %v" + x)
}

func (p presentation) Next(id string) (nextId string, err error) {
	i, err := indexOf(id, p.slides)

	if err != nil {
		return "", errors.New("Slide not found: " + id)
	}

	if i+1 == len(p.slides) {
		return "", errors.New("Last slide: " + id)
	}

	return p.slides[i+1], nil
}

func (p presentation) Prev(id string) (nextId string, err error) {
	i, err := indexOf(id, p.slides)

	if err != nil {
		return "", errors.New("Slide not found: " + id)
	}

	if i == 0 {
		return "", errors.New("First slide: " + id)
	}

	return p.slides[i-1], nil
}
