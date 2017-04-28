package controllers

import (
	"sync"
)

var gCount *Count

type Count struct {
	sync.RWMutex
	ckCount map[string]int
}

func init() {
	gCount = newCount()
}

func newCount() *Count {
	return &Count{
		ckCount: make(map[string]int, 0),
	}
}


func GetCkCount(articleKey string) int {
	gCount.RLock()
	ckCount := gCount.ckCount[articleKey]
	gCount.RUnlock()
	return ckCount
}

func IncCkCount(articleKey string) {
	gCount.Lock()
	curCount := gCount.ckCount[articleKey]
	gCount.ckCount[articleKey] = curCount + 1
	gCount.Unlock()
}
