package dundee

import (
	"github.com/ign/ipl-dundee/dundee/elementals"
)

type Config struct {
	Elementals  []elementals.ElementalServer
	Streams_url *string `json:"-"`
}
