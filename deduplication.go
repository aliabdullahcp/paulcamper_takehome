package main

import "sync"

type deDuplicator struct {
	requestMap           map[string]bool
	mux                  *sync.Mutex
	resourceSynchronizer *sync.Cond
}

func NewDeduplicator() *deDuplicator {
	mutex := sync.Mutex{}
	condition := sync.NewCond(&mutex)
	return &deDuplicator{map[string]bool{}, &mutex, condition}
}
