package main

import (
	"strings"
)

func GetCount(wordLine string) int {
	words := strings.Split(wordLine, " ")
	return len(words)
}
