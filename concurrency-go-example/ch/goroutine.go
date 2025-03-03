package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	userID := 10
	ch := make(chan string, 10)
	wg := &sync.WaitGroup{}

	go fetchUserData(userID, ch, wg)
	go fetchUserLikes(userID, ch, wg)
	wg.Add(2)
	wg.Wait()
	close(ch)

	for res := range ch {
		fmt.Println(res)
	}

	end := time.Since(start)
	fmt.Println(end) 
	// 80ms + 50ms = 130ms
	// concurrency wait for weakest link i.e. longest time taken ~80ms
}

func fetchUserData(userID int, ch chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	_ = userID

	ch <- "user data"

	wg.Done()
}

func fetchUserLikes(userID int, ch chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)
	_ = userID

	ch <- "user likes"

	wg.Done()
}
