package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	rps := flag.Int("rps", 100, "Requests per second")
	duration := flag.Int("duration", 10, "Test duration in seconds")
	flag.Parse()

	requestsPerSecond := *rps
	totalDuration := time.Duration(*duration) * time.Second

	var success, failed uint64

	rateLimiter := time.Tick(time.Second / time.Duration(requestsPerSecond))
	var wg sync.WaitGroup

	fmt.Printf("Starting test with RPS: %d, Duration: %ds\n", requestsPerSecond, *duration)
	start := time.Now()
	for time.Since(start) < totalDuration {
		<-rateLimiter
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://server:8080/process?request=test")
			if err != nil {
				atomic.AddUint64(&failed, 1)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				atomic.AddUint64(&success, 1)
			} else {
				atomic.AddUint64(&failed, 1)
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Test completed.\nSuccessful requests: %d\nFailed requests: %d\n", success, failed)
}
