package elementals

import (
	"log"
	"regexp"
)

var pattern_protocol *regexp.Regexp
var pattern_leadingSlash *regexp.Regexp

func init() {
	var err error
	pattern_protocol, err = regexp.Compile(`^(?:http(?P<secure>s)?:/+)?(?P<hostname>[^/]*)/*$`)
	if err != nil {
		log.Fatalln(err)
	}
	pattern_leadingSlash, err = regexp.Compile(`^/*(?P<path>.*)$`)
	if err != nil {
		log.Fatalln(err)
	}
}

func (this *ElementalServer) parseURL(path string) string {
	newHost := pattern_protocol.ReplaceAllString(this.Hostname, `http$secure://$hostname`)
	newPath := pattern_leadingSlash.ReplaceAllString(path, `/$path`)
	return newHost + newPath
}
