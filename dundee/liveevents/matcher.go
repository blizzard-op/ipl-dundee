package liveevents

import (
	"github.com/ign/ipl-dundee/dundee/streams"
	"regexp"
)

func (this *Live_event) Match(stream *streams.Stream) bool {
	found, _ := regexp.MatchString(stream.Name, this.Name)
	if found == true && this.Status == "running" {
		return true
	}
	return false
}
