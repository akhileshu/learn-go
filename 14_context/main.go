package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	// entry point
}

/*
type Store interface {
	Fetch() string
	Cancel()
}

// A goroutine fetches data concurrently while the handler waits for either the data to arrive or the request context to be cancelled—whichever happens first.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)
		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}

	}
}
*/

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
