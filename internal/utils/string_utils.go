package utils

import (
	"strings"
)

func CleanString(s string) string {
	cleanedString := strings.Trim(s, "=\"")
	return cleanedString
}