package main

import (
	"sync/atomic"
	"testing"
)

// func TestDataRaceConditions(t *testing.T) {
// 	var state int32
// 	var mu sync.RWMutex

// 	for i := range 10 {
// 		go func(i int) {
// 			mu.Lock()
// 			state += int32(i)
// 			mu.Unlock()
// 		}(i)
// 	}
// }

func TestDataRaceConditions(t *testing.T) {
	var state int32

	// atomic
	for i := range 10 {
		go func(i int) {
			// state += int32(i)
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}

// type Server struct {
// 	gameRound atomic.Int32
// }

