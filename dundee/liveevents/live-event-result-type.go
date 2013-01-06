package liveevents

import (
	"dundee/elemental"
)

type LiveEventResult struct {
	elemental  *elemental.ElementalServer
	liveEvents *[]LiveEvent
}
