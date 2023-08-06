package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)

	for _, word := range strings.Fields(s) {
		count[word]++
	}
	return count
}

func execiseMap() {
	wc.Test(WordCount)
}
