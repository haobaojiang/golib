package golib

import "regexp"

func IsDigOrAlpha(s string) bool {
	return regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(s)
}
