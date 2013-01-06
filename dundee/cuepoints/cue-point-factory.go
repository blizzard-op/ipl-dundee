package cuepoints

import (
	"errors"
)

var cuePointTypes = make(map[string]func() (interface{}, error), 5)

func RegisterCuePointType(name string, cp func(r *http.Request) (interface{}, error)) error {
	if cuePointTypes[name] != nil {
		return errors.New("A cue point is already registered under the name: " + name)
	}
	cuePointTypes[name] = cp
	return nil
}

func New(r *http.Request) (interface{}, error) {
	cpType := r.FormValue("cue-point-type")
	if registeredTypes[cpType] == nil {
		return nil, errors.New("Invalid Cue Point type.")
	}
	return registeredTypes[cpType]()
}
