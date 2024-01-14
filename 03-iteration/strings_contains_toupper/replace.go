package main

import "strings"

func Replace(sentence, sentenceWord, wordToReplace string) string {
	sentenceLower := strings.ToLower(sentence)
	sentenceWordLower := strings.ToLower(sentenceWord)
	if strings.Contains(sentenceLower, sentenceWordLower) {
		// TODO: Change with split or arrays
		// - Sentence needs to be the same except sentenceWord -> wordToReplace
		result := strings.Replace(sentenceLower, sentenceWordLower, wordToReplace, 1)
		return result
	}
	return getErrorMessage(sentence, sentenceWordLower, wordToReplace)
}

func getErrorMessage(sentence, sentenceWord, wordToReplace string) string {
	var errorMessage strings.Builder
	errorMessage.WriteString("Sentence \"")
	errorMessage.WriteString(sentence)
	errorMessage.WriteString("\" doesn't have the word \"")
	errorMessage.WriteString(sentenceWord)
	errorMessage.WriteString("\". Can't replace with \"")
	errorMessage.WriteString(wordToReplace)
	errorMessage.WriteString("\"")
	return errorMessage.String()
}
