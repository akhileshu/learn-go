package main

import (
	"go-playground/16_maths/clockface"
	"math"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90} // 150,60
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}

}

func TestSecondHandAllSeconds(t *testing.T) {
	for second := 0; second < 60; second++ {
		tm := time.Date(1337, time.January, 1, 0, 0, second, 0, time.UTC)

		got := clockface.SecondHand(tm)

		angle := float64(second)*6 - 90
		radians := angle * math.Pi / 180

		want := clockface.Point{
			X: float64(150 + int(float64(90)*math.Cos(radians))),
			Y: float64(150 + int(float64(90)*math.Sin(radians))),
		}

		if got != want {
			t.Errorf("second %d: got %v, want %v", second, got, want)
		}
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 + 90} // 150,240
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestSecondsInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := clockface.SecondsInRadians(thirtySeconds)

	if want != got {
		t.Fatalf("Wanted %v radians, but got %v", want, got)
	}
}
