package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	// 150-90=60
	// return  Point{150,(float64(60+t.Second()*6))} // wrong

	const (
		clockCenterX     = 150
		clockCenterY     = 150
		secondHandLength = 90
	)

	seconds := t.Second()

	angle := float64(seconds)*6 - 90
	// seconds: 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23
	// angles{-90,-84,-78,-72,-66,-60,-54,-48,-42,-36,-30,-24,-18,-12,-6,0,6,12,18,24,30,36,42,48,54,60,66,72,78,84,90}

	/*
			360° = 2π radians (perimeter of circle)
			1° = π / 180 radians

			radians = degrees × π / 180

		Degrees are a human convention, radians are math-native.
		sin²(θ) + cos²(θ) = 1
		a² + b² = c² (pythagorean theorem) right angle triangle
	*/

	radians := angle * math.Pi / 180

	x := float64(clockCenterX + int(float64(secondHandLength)*math.Cos(radians)))
	y := float64(clockCenterY + int(float64(secondHandLength)*math.Sin(radians)))

	return Point{x, y}

}

func SecondsInRadians(thirtySeconds time.Time) float64 {
	return math.Pi
}
