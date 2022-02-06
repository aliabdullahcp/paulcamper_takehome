package main

import (
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"math/rand"
	"time"

	"golang.org/x/text/language"
)

var memoryCache *cache.Cache

func main() {
	// creating memory cache with no expiration and no clean up size
	memoryCache = cache.New(cache.NoExpiration, cache.NoExpiration)

	ctx := context.Background()
	rand.Seed(time.Now().UTC().UnixNano())
	s := NewService()
	fmt.Println(s.Translate(ctx, language.English, language.Japanese, "test"))
	fmt.Println(s.Translate(ctx, language.English, language.Japanese, "test"))
}
