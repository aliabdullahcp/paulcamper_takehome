package main

import (
	"context"
	"log"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/text/language"
)

// Service is a Translator user.
type Service struct {
	translator Translator
}

func NewService() *Service {
	t := newRandomTranslator(
		100*time.Millisecond,
		500*time.Millisecond,
		0.1,
	)

	return &Service{
		translator: t,
	}
}

func (s *Service) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {
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
	err := backoff.RetryNotify(retryable, maxRetriesBackoff, notify)
	if err != nil {
		log.Fatalf("error after retrying: %v", err)
	}

	return translatorValue, translatorError
}
