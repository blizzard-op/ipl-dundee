package cuepoints

import (
	"errors"
	"net/http"
)

var cpTypeMap = make(map[string]func(r *http.Request) (interface{}, *TimedMetadata, error), 2)

func RegisterCuePointType(cpType string, cpCreator func(r *http.Request) (interface{}, *TimedMetadata, error)) error {
	if cpTypeMap[cpType] != nil {
		return errors.New("A cue point is already registered under the name: " + cpType)
	}
	cpTypeMap[cpType] = cpCreator
	return nil
}

func New(cpType string, r *http.Request) (interface{}, *TimedMetadata, error) {
	if cpTypeMap[cpType] == nil {
		return nil, nil, errors.New("Invalid Cue Point type.")
	}
	return cpTypeMap[cpType](r)
}
