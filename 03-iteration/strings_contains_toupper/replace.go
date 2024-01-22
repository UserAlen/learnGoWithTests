package main

import "strings"

func Replace(sentence, sentenceWord, wordToReplace string) string {
	sentenceLower := strings.ToLower(sentence)
	sentenceWordLower := strings.ToLower(sentenceWord)
	if strings.Contains(sentenceLower, sentenceWordLower) {
		// TODO: Change with split or arrays
		// - Sentence needs to be the same except sentenceWord -> wordToReplace

		var sentenceWords []string = SentenceToWords(sentence)
		for i, word := range sentenceWords {
			if strings.ToLower(word) == strings.ToLower(sentenceWord) {
				//idk
			}
		}

		result := strings.Replace(sentenceLower, sentenceWordLower, wordToReplace, 1)
		return result
	}
	return getErrorMessage(sentence, sentenceWordLower, wordToReplace)
}

func SentenceToWords(sentence string) []string {
	// instead of "spacebar" can be any separator - a collection of punctuatian marks
	// or way easier - using the alphabet of the language with slices instead of strings package

	return strings.Split(sentence, " ")
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
