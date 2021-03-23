package internal

import (
	"strings"
)

func ReverseParagraph(paragraph *string) *string {
	if paragraph == nil {
		emptyStr := ""
		return &emptyStr
	}
	reversed := strings.Join(reverseWords(paragraph), " ")
	return &reversed
}

func mapWords(paragraph *string) []string {
	return strings.Split(*paragraph, " ")
}

func reverseWords(paragraph *string) []string {
	wMap := mapWords(paragraph)
	for i, j := 0, len(wMap)-1; i < j; i, j = i+1, j-1 {
		wMap[i], wMap[j] = wMap[j], wMap[i]
	}
	return wMap
}