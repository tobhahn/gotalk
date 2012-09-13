package gotalk

import (
	"errors"
)

type Presentation struct {
	Slides []string
}

func indexOf(x string, slice []string) (int, error) {
	for i, v := range slice {
		if x == v {
			return i, nil
		}
	}
	return -1, errors.New("Not found: %v" + x)
}

func (p Presentation) Next(id string) (nextId string, err error) {
	i, err := indexOf(id, p.Slides)

	if err != nil {
		return "", errors.New("Slide not found: " + id)
	}

	if i+1 == len(p.Slides) {
		return "", errors.New("Last slide: " + id)
	}

	return p.Slides[i+1], nil
}

func (p Presentation) Prev(id string) (nextId string, err error) {
	i, err := indexOf(id, p.Slides)

	if err != nil {
		return "", errors.New("Slide not found: " + id)
	}

	if i == 0 {
		return "", errors.New("First slide: " + id)
	}

	return p.Slides[i-1], nil
}
