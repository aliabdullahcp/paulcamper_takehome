package main

import (
	"golang.org/x/text/language"
	"testing"
)

func TestCreateTranslatorCacheKey(t *testing.T) {
	testData := "testing data"
	expectedResult := "translation-" + language.English.String() + "-" + language.Japanese.String() + "-" + testData
	cacheKey := createTranslatorCacheKey(language.English, language.Japanese, testData)

	if cacheKey != expectedResult {
		t.Errorf("Incorrent key returned. Expected: %s\nGot: %s", expectedResult, cacheKey)
	}
}
