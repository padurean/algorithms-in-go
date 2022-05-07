package maxnondupsubstr

import "strings"

func MaxNonDupSubStr(s string) (string, int64) {
	var (
		curr string
		max  []rune
	)
	for _, char := range s {
		if strings.ContainsRune(curr, char) {
			curr = ""
		}
		curr += string(char)
		currRunes := []rune(curr)
		if len(currRunes) > len(max) {
			max = currRunes
		}
	}
	return string(max), int64(len(max))
}
