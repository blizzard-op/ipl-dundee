package cuepoints

import (
	"encoding/base64"
)

const (
	TAG_HEADER     = "ID3v"
	FRAME_CATEGORY = "TIT1"
	FRAME_TITLE    = "TIT2"
	FRAME_SUBTITLE = "TIT3"
	CATEGORY       = "ipl-cue-point"
)

func generateId3(title, subTitle string) string {
	src := TAG_HEADER + FRAME_CATEGORY + CATEGORY + FRAME_TITLE + title + FRAME_SUBTITLE + subTitle
	encoded := base64.StdEncoding.EncodeToString([]byte(src))
	return encoded
}
