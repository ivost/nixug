package models

import (
	"errors"
	"strings"
)

var (
	ErrNameEmpty = errors.New("name cannot be empty")
)

type RequestErrors struct {
	errs []error
}

func (re *RequestErrors) Append(err error) {
	re.errs = append(re.errs, err)
}

func (re *RequestErrors) Len() int {
	return len(re.errs)
}

func (re *RequestErrors) Error() string {
	var er []string
	for _, e := range re.errs {
		er = append(er, e.Error())
	}
	return strings.Join(er, ", ")
}
