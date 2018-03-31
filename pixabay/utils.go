package pixabay

import (
	"regexp"
)

var (
	nameExp = regexp.MustCompile(`(.*?)_.*?\.(.*?)`)
)
