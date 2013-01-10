package cuepoints

import (
	"errors"
	"net/http"
)

var cpTypeMap = make(map[string]func(r *http.Request) (interface{}, error), 2)

func RegisterCuePointType(cpType string, cpCreator func(r *http.Request) (interface{}, error)) error {
	if cpTypeMap[cpType] != nil {
		return errors.New("A cue point is already registered under the name: " + cpType)
	}
	cpTypeMap[cpType] = cpCreator
	return nil
}

func New(cpType string, r *http.Request) (interface{}, error) {
	if cpTypeMap[cpType] == nil {
		return nil, errors.New("Invalid Cue Point type.")
	}
	return cpTypeMap[cpType](r)
}
