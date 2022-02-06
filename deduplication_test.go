package main

import (
	"golang.org/x/text/language"
	"testing"
)

func TestCreateTranslatorDeDuplicatorKey(t *testing.T) {
	testData := "testing data"
	expectedResult := "deduplicate-" + language.English.String() + "-" + language.Japanese.String() + "-" + testData
	cacheKey := createTranslatorDeduplicateKey(language.English, language.Japanese, testData)

	if cacheKey != expectedResult {
		t.Errorf("Incorrent key returned. Expected: %s\nGot: %s", expectedResult, cacheKey)
	}
}
