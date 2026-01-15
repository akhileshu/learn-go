### Concurrency

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

- We can solve this data race by coordinating our goroutines using channels