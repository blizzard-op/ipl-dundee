package liveevents

import (
	"dundee/elemental"
)

type LiveEventResult struct {
	Elemental  *elemental.ElementalServer
	LiveEvents []LiveEvent
}
