package main

import (
	"context"
	"github.com/patrickmn/go-cache"
	"log"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/text/language"
)

// Service is a Translator user.
type Service struct {
	translator  Translator
	memoryCache *cache.Cache
	dt          *deDuplicator
}

func NewService() *Service {
	t := newRandomTranslator(
		100*time.Millisecond,
		500*time.Millisecond,
		0.1,
	)

	return &Service{
		translator:  t,
		memoryCache: cache.New(cache.NoExpiration, cache.NoExpiration),
		dt:          NewDeDuplicator(),
	}
}

func (s *Service) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {
	cacheKey := createTranslatorCacheKey(from, to, data)
	cacheValue, found := s.memoryCache.Get(cacheKey)

	if found {
		// cache hit - value found from the cache
		return cacheValue.(string), nil
	}

	var translatorError error
	var translatorValue string

	retryable := func() error {
		translatorValue, translatorError = s.translator.Translate(ctx, from, to, data)
		return translatorError
	}

	notify := func(err error, t time.Duration) {
		log.Printf("Translation Error: '%v' happened at time: %v", err, t)
	}

	var maxRetries uint64 = 5
	exponentialBackoff := backoff.NewExponentialBackOff()
	exponentialBackoff.MaxElapsedTime = 15 * time.Second
	maxRetriesBackoff := backoff.WithMaxRetries(exponentialBackoff, maxRetries)

	deDuplicatorKey := createTranslatorDeduplicateKey(from, to, data)
	s.dt.resourceSynchronizer.L.Lock()
	for s.dt.requestMap[deDuplicatorKey] == true {
		s.dt.resourceSynchronizer.Wait()
	}
	s.dt.resourceSynchronizer.L.Unlock()

	s.dt.resourceSynchronizer.L.Lock()
	s.dt.requestMap[deDuplicatorKey] = true
	s.dt.resourceSynchronizer.Broadcast()
	s.dt.resourceSynchronizer.L.Unlock()

	err := backoff.RetryNotify(retryable, maxRetriesBackoff, notify)

	s.dt.resourceSynchronizer.L.Lock()
	s.dt.requestMap[deDuplicatorKey] = false
	s.dt.resourceSynchronizer.Broadcast()
	s.dt.resourceSynchronizer.L.Unlock()

	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}

	if translatorError == nil {
		// set the value in cache
		s.memoryCache.Set(createTranslatorCacheKey(from, to, data), translatorValue, cache.NoExpiration)
	}
	return translatorValue, translatorError
}
