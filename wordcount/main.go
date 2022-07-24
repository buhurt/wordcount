package main

import (
	"strings"
)

func getCount(wordLine string) int {
	words := strings.Split(wordLine, " ")
	return len(words)
}
