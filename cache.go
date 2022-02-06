package main

import (
	"golang.org/x/text/language"
	"strings"
)

func createTranslatorCacheKey(
	fromLanguage language.Tag,
	toLanguage language.Tag,
	data string,
) string {
	elems := []string{"translation", fromLanguage.String(), toLanguage.String(), data}
	return strings.Join(elems, "-")
}

func createTranslatorDeduplicateKey(
	fromLanguage language.Tag,
	toLanguage language.Tag,
	data string,
) string {
	elems := []string{"deduplicate", fromLanguage.String(), toLanguage.String(), data}
	return strings.Join(elems, "-")
}
