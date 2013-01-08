package config

import (
	"dundee/elemental"
)

type Config struct {
	Port        string
	Elementals  []elemental.ElementalServer
	Streams_url string
}
