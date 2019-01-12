package collector

import (
	"proxycollector/result"
	"proxycollector/storage"
)

type Collector interface {
	Next() bool
	Name() string
	Collect(chan<- *result.Result, storage.Storage) []error
}

type Type uint8

const (
	COLLECTBYSELECTOR Type = iota
	COLLECTBYREGEX
)
