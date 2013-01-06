package dundee

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func GetConfig() (*Config, error) {
	var config Config

	//Get config_path
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		return nil, errors.New("You did not include a config path.")
	}

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
