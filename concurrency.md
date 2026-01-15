###  channels

```go
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
			go func() {
			results[url] = wc(url)
		}()
	}
time.Sleep(2 * time.Second)
	return results
}

// insted do

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			// results[url] = wc(url)
			resultChannel <- result{url, wc(url)} // send result(struct) to channel
		}()
	}
	for i := 0; i < len(urls); i++ {

		r := <-resultChannel // Receive expression
		results[r.string] = r.bool
	}
	return results
}

```
- why wait after checkingWebsite ?

    Flow:

        Goroutines are started

        CheckWebsites does NOT wait

        Function returns immediately

        Program exits / caller reads results

        Goroutines may not have run yet ❌

        time.Sleep just delays the return, hoping goroutines finish.

- `race condition`, a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over
- `go test -race` - WARNING: DATA RACE

- `channel`
    - Goroutines run at the same time.
    Channels let them communicate and synchronize.
    - “Don’t share memory to communicate.
    Communicate to share memory.”
    - A `channel` is a pipe that lets goroutines safely send data to each other.

### sync 
- `sync` provides fundamental synchronization primitives for safely managing shared resources across multiple goroutines, primarily when channels aren't suitable
- `sync.WaitGroup` - waits for a set of goroutines to finish and is used when you need to run code after all of them complete.
- `lock / mutex` - A mutex ensures that only one goroutine at a time can access a shared resource (variable, map, struct) to prevent data races. ex - increment the counter at a time

### Context
- `context` provides a standard way to carry deadlines, cancellation signals, and request-scoped values across API boundaries and between goroutines
