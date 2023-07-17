package helpers

import "strings"

func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func EndsWithAny(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func StartsWith(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.EqualFold(s, prefix) {
			return true
		}
	}
	return false
}