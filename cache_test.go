package main

import (
	"context"
	"golang.org/x/text/language"
	"math/rand"
	"testing"
	"time"
)

func TestCreateTranslatorCacheKey(t *testing.T) {
	testData := "testing data"
	expectedResult := "translation-" + language.English.String() + "-" + language.Japanese.String() + "-" + testData
	cacheKey := createTranslatorCacheKey(language.English, language.Japanese, testData)

	if cacheKey != expectedResult {
		t.Errorf("Incorrent key returned. Expected: %s\nGot: %s", expectedResult, cacheKey)
	}
}

func TestTranslationResultAddedInCache(t *testing.T) {
	ctx := context.Background()
	rand.Seed(time.Now().UTC().UnixNano())
	s := NewService()
	testData := "testing data"

	translationResultActual, err := s.Translate(ctx, language.English, language.Japanese, testData)

	if err != nil {
		t.Errorf("Error occurred when getting translation for the first time. \n%s", err)
	}

	cachedKey := createTranslatorCacheKey(language.English, language.Japanese, testData)
	cachedResult, found := s.memoryCache.Get(cachedKey)
	if !found {
		t.Errorf("Translation result not found in cache.")
	}

	if translationResultActual != cachedResult.(string) {
		t.Errorf("Cached result doesn't match with original. Expected: %s\n Got: %s", translationResultActual, cachedResult.(string))
	}
}
