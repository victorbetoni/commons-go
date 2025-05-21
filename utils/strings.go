package utils

import "strings"

func PadText(str string, length int) string {
	if len(str) >= length {
		return str
	}
	return str + strings.Repeat(" ", length-len(str))
}

func HighestLength(items []string) int {
	leng := 0
	for _, v := range items {
		if len(v) > leng {
			leng = len(v)
		}
	}
	return leng
}
