package dundee

import (
	"dundee/elemental"
)

type Config struct {
	Elementals  []elemental.ElementalServer
	Streams_url *string `json:"-"`
}
