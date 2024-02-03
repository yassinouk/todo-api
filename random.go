package main

import (
	"math/rand"
	"strings"
)

var WordsList = []string{"ipsum", "dolor", "adipiscing", "elit", "incididunt", "consectetur", "tempor", "labore", "minim", "veniam", "exercitation", "ullamco", "pariatur", "proident", "deserunt", "mollit", "anim", "est", "laborum", "sit", "amet", "sed", "do", "eiusmod", "magna", "aliqua", "quis", "nostrud", "ex", "ea", "commodo", "culpa", "qui", "officia", "sint", "non", "sunt", "occaecat", "cupidatat", "sunt", "in", "culpa", "qui", "officia", "deserunt", "mollit", "anim", "id"}

// RandomString generates a random string of length n
func RandomWords(n int) []string {
	result := make([]string, n)
	k := len(WordsList)
	for i := range result {
		result[i] = WordsList[rand.Intn(k)]
	}
	return result
}

// RandomTitle generates a random task title
func RandomTaskTitle() string {
	words := RandomWords(5)
	return strings.Join(words, " ")
}
func RandomTaskContent() string {
	words := RandomWords(10)
	return strings.Join(words, " ")
}
