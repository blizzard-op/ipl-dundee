package dundee

import (
	"encoding/json"
	"io/ioutil"
)

func (c *Config) Parse(path *string) error {
	bytes, err := ioutil.ReadFile(*path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return err
	}

	return nil
}
