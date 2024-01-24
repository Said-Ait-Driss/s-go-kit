package main

import (
	"errors"
	"strings"
)

// declaration of variables
var ErrEmpty = errors.New("empty string")

// service type
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// implimentation of service
type stringService struct{}

// service implimentation methods (recievers)
func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	if s == "" {
		return 0
	}
	return len(s)
}
