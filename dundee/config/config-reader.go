package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var config Config

func init() {
	err := loadConfig()
	if err != nil {
		fmt.Println("Dundee failed to read config. Aborting.\nReason: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Config loaded successfully.")
}

func Get() *Config {
	return &config
}

func getPath() (string, error) {
	if len(os.Args) > 1 {
		return os.Args[1], nil
	}
	return "", errors.New("You did not include a config path.")
}

func loadConfig() error {
	path, err := getPath()
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return err
	}

	return nil
}
