package main

import "strings"

const (
	correct string = "Equal"
	wrong   string = "Not Equal"
)

//version from Strings pakage
func Compare(word1, word2 string) string {
	var result int = strings.Compare(word1, word2)

	switch result {
	case 0:
		return correct
	default:
		return wrong
	}
}
