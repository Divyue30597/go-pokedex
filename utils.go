package main

import "strings"

func cleanText(text string) []string {
	newText := strings.ToLower(text)
	finalText := strings.Fields(newText)
	return finalText
}
