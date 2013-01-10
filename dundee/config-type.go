package dundee

import (
	"github.com/ign/ipl-dundee/dundee/elemental"
)

type Config struct {
	Elementals  []elemental.ElementalServer
	Streams_url *string `json:"-"`
}
