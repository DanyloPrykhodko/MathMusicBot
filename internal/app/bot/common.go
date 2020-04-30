package bot

import (
	"regexp"
	"strings"
)

func markdown(s string) string {
	s = strings.ReplaceAll(s, "*", "\\*")
	s = strings.ReplaceAll(s, "_", "\\_")
	s = strings.ReplaceAll(s, "`", "")
	s = regexp.MustCompile("\\[.*\\]\\(.*\\)").ReplaceAllString(s, "")
	return s
}
