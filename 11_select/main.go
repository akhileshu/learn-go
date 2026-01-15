package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// entry point
}
var tenSecondTimeout = 10 * time.Second
func Racer(a, b string) (string, error) {
	/*
		aDuration := measureResponseTime(a)
		bDuration := measureResponseTime(b)
		if aDuration < bDuration {
			return a, nil
		}
		return b, nil
	*/

return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func measureResponseTime(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
		// we don't care what type is sent to the channel, we just want to signal we are done and closing the channel works perfectly!
	}()
	return ch
}
