package main

import (
	"fmt"
	"math/big"
)

func main() {
	// -------- int (default integer) --------
	var a int = 10
	var b int = 3
	fmt.Println("int Add:", AddInt(a, b))

	// -------- int64 (explicit size) --------
	var x int64 = 1000000000
	var y int64 = 2000000000
	fmt.Println("int64 Add:", AddInt64(x, y))

	// -------- uint (unsigned) --------
	var u1 uint = 10
	var u2 uint = 20
	fmt.Println("uint Add:", AddUint(u1, u2))

	// -------- float64 (decimal numbers) --------
	var f1 float64 = 10.5
	var f2 float64 = 3.2
	fmt.Println("float Divide:", DivideFloat(f1, f2))

	// -------- money-safe (int64 as smallest unit) --------
	// ₹10.50 + ₹2.25 = ₹12.75
	price1 := int64(1050) // paise
	price2 := int64(225)
	fmt.Println("Money Sum (paise):", AddMoney(price1, price2))

	// -------- big.Int (unbounded integers) --------
	bigA := big.NewInt(0)
	bigA.SetString("100000000000000000000", 10)

	bigB := big.NewInt(0)
	bigB.SetString("900000000000000000000", 10)

	fmt.Println("big.Int Add:", AddBigInt(bigA, bigB))
}

//
// Integer operations
//

// AddInt adds two int values.
func AddInt(a, b int) int {
	return a + b
}

// AddInt64 adds two int64 values.
func AddInt64(a, b int64) int64 {
	return a + b
}

// AddUint adds two unsigned integers.
func AddUint(a, b uint) uint {
	return a + b
}

//
// Floating point
//

// DivideFloat divides two float64 values.
func DivideFloat(a, b float64) float64 {
	return a / b
}

//
// Money-safe arithmetic
//

// AddMoney adds currency stored as smallest units (paise, cents).
func AddMoney(a, b int64) int64 {
	return a + b
}

//
// Big numbers
//

// AddBigInt adds two big.Int values safely.
func AddBigInt(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}
