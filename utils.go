package main

import (
	"math/rand"
	"strings"
	"time"
)

func cleanText(text string) []string {
	newText := strings.ToLower(text)
	finalText := strings.Fields(newText)
	return finalText
}

func getRandomValue(baseExperience int) int {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomMinutes := rng.Intn(baseExperience * 2)
	return randomMinutes
}
