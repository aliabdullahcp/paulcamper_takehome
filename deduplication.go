package main

import (
	"golang.org/x/text/language"
	"strings"
	"sync"
)

type deDuplicator struct {
	requestMap           map[string]bool
	mux                  *sync.Mutex
	resourceSynchronizer *sync.Cond
}

func NewDeDuplicator() *deDuplicator {
	mutex := sync.Mutex{}
	condition := sync.NewCond(&mutex)
	return &deDuplicator{map[string]bool{}, &mutex, condition}
}

func createTranslatorDeduplicateKey(
	fromLanguage language.Tag,
	toLanguage language.Tag,
	data string,
) string {
	elems := []string{"deduplicate", fromLanguage.String(), toLanguage.String(), data}
	return strings.Join(elems, "-")
}
