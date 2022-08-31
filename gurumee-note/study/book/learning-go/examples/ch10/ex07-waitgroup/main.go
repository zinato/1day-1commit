package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var result []int
	for v := range out {
		result = append(result, v)
	}
	return result
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			r := rand.Intn(5)
			time.Sleep(time.Second * time.Duration(r))
			fmt.Println(v)
		}(i)
	}

	wg.Wait()
}
