package helper

import "strings"

func RemoveSlash(value string) string {
	return strings.ReplaceAll(value, "/", "")
}
