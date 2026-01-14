### about this project
- this repo implements https://quii.gitbook.io/learn-go-with-tests

### run test
- When you test a folder, Go runs all _test.go files in that folder.
- `go test ./hello_world` 

- Run a specific test function 
- `go test ./iteration -run TestSum`

### table driven tests
- to build a list of test cases that can be tested in the same manner

### benchmark()
```sh
# Run all benchmarks in a package
go test -bench=. ./iteration

# Run a specific benchmark
go test -bench=BenchmarkRepeat ./iteration

```
- `executed by the "go test" command when its -bench flag`
- Benchmark functions exist to measure performance — time and memory — in a repeatable, scientific way.

- You use benchmark functions to:

    measure reality, not intuition

    protect performance contracts

    make optimization safe and justified
```go

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
```

### example()
It serves two purposes:

Executable documentation

Behavior check (only if format matches exactly)

If an example needs a local variable, it’s probably not a good example.

- run tests with example `go test ./integers -v`
- in *_test.go
```go

// this will not warn if example fails
/* func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
 */

// prefer below example coz it will fail if expectation didn't match output
 func ExampleAdd() {
	fmt.Println(Add(1, 5))
	// Output: 6
}
```

### package
- One folder = one package
- every .go file must have a package declaration.
---
- `doc.go, token.go, password.go` → package auth
- `main.go` → must be package main

### main()
- One folder (package) → only ONE main() function

- That main() usually lives in one file (commonly main.go)

### debugging
- have `.vscode/launch.json` at project root
- place breakpoint's in code / test files
- either main() or test function can trigger breakpoint
- F5

### docs
- supports basic md syntax
- Comments start with the name of the thing (function, type, package)

- Function docs
```go
// Add adds two integers and returns the result.
//
// Parameters:
//   - a: first number
//   - b: second number
//
// Returns:
//   - sum of a and b
//
// Example:
//   result := Add(2, 3) // 5
func Add(a, b int) int {
	return a + b
}


```

- Struct docs
```go
// User represents an application user.
type User struct {
	// ID is the unique identifier.
	ID string

	// Email is the user's email address.
	Email string
}

```

- Package docs
- doc.go
- Rule 1: place `doc.go` Inside the package folder

```
auth/
  doc.go
  token.go
  password.go

```
```go
// Package auth provides authentication utilities.
//
// Features:
//   - JWT handling
//   - Password hashing
//   - Session validation
package auth

```

- view docs
 ```sh
 go doc package/mathutil
 go doc Add
go doc ./...
go doc ./integers Add
 ```

### arrays_and_slices
- [2]int is array and []int is slice 
- Use arrays only when size is a constraint, not a convenience.

- slice[low:high] 
    - creates a new slice referencing a portion of an underlying array
    - slice[low:high] | nums[1:4] | nums[:4] | nums[3:]
    - nums[:]  It creates a new slice header pointing to the same backing array.
    - [low:high:max] - max is the new capacity

- Length vs Capacity
    - len(s) // number of usable elements
    - cap(s) // total space available from start index

### function inside function 
```go
func TestSumAllTails(t *testing.T) {

	checkSums := func (t testing.TB , got , want []int) {
    }

}
```
# learn-go
