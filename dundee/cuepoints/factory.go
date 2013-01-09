package cuepoints

import (
	"errors"
	"net/http"
)

var cuePointTypes = make(map[string]func(r *http.Request) (interface{}, error), 5)

func RegisterCuePointType(name string, cp func(r *http.Request) (interface{}, error)) error {
	if cuePointTypes[name] != nil {
		return errors.New("A cue point is already registered under the name: " + name)
	}
	cuePointTypes[name] = cp
	return nil
}

func New(cpType string, r *http.Request) (interface{}, error) {
	if cuePointTypes[cpType] == nil {
		return nil, errors.New("Invalid Cue Point type.")
	}
	return cuePointTypes[cpType](r)
}
