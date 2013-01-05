package dundee

import (
	"dundee/types"
	"encoding/json"
	"io/ioutil"
)

func GetConfig(path string) (*types.Config, error) {
	var config types.Config

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
