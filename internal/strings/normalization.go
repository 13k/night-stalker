package strings

import (
	"regexp"
	"strings"
)

var (
	slugifyCleanRE      = regexp.MustCompile(`[^\p{L}\p{N}]`)
	slugifyExcessDashRE = regexp.MustCompile(`-+`)
)

func Slugify(s string) string {
	s = strings.ToLower(s)
	s = slugifyCleanRE.ReplaceAllLiteralString(s, "-")
	s = strings.ReplaceAll(s, "_", "-")
	s = slugifyExcessDashRE.ReplaceAllLiteralString(s, "-")
	s = strings.Trim(s, " -")
	return s
}
