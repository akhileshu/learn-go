package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// entry point
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	/*
		for i := countdownStart; i > 0; i-- {
			fmt.Fprintln(out, i)
		}
	*/

	fmt.Fprint(out, finalWord)
}
