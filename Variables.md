Alright Akhilesh — this is a **foundational topic**, so I’ll give you a **complete but clean mental model**.
No fluff. Simple language. Real rules.

---

# Variables in Go — Complete Guide

---

## 1️⃣ What is a variable in Go?

> A variable is a **named storage location** with a **type**.

Go is:

* **statically typed**
* **block scoped**
* **explicit by default**

---

## 2️⃣ Ways to declare variables

### 2.1 `var` (explicit, readable)

```go
var x int
var y int = 10
```

Zero value applies if not initialized.

---

### 2.2 Short declaration `:=` (most common)

```go
x := 10
name := "go"
```

Rules:

* type is inferred
* **only allowed inside functions**
* must declare **at least one new variable**

❌ Not allowed at package level.

---

### 2.3 Multiple variables

```go
var a, b int
x, y := 1, 2
```

---

### 2.4 Grouped declaration

```go
var (
	count int
	name  string
	valid bool
)
```

Used mostly at package level.

---

## 3️⃣ Zero values (very important)

Every variable has a default value.

| Type    | Zero value    |
| ------- | ------------- |
| int     | `0`           |
| float   | `0.0`         |
| string  | `""`          |
| bool    | `false`       |
| pointer | `nil`         |
| slice   | `nil`         |
| map     | `nil`         |
| struct  | zeroed fields |

No uninitialized garbage. Ever.

---

## 4️⃣ Scope of variables (CRITICAL)

### 4.1 Package scope

```go
var version = "1.0"
```

* accessible in all files of the package
* initialized once
* avoid unless needed

---

### 4.2 Function scope

```go
func f() {
	x := 10
}
```

* lives only inside the function
* most variables should be here

---

### 4.3 Block scope

```go
if ok {
	y := 20
}
```

* `y` exists only inside `{}`

This includes:

* `if`
* `for`
* `switch`
* `{}` blocks

---

## 5️⃣ Shadowing (IMPORTANT BUG SOURCE)

```go
x := 10

if true {
	x := 20 // shadows outer x
	fmt.Println(x) // 20
}

fmt.Println(x) // 10
```

Rules:

* inner scope can redeclare same name
* outer variable is hidden, not modified

👉 Avoid shadowing unless intentional.

---

## 6️⃣ Assignment vs Declaration

### Declaration

```go
x := 10
var y int
```

### Assignment

```go
x = 20
y = 30
```

---

### Mixed short declaration (common trap)

```go
x := 10
x, y := 20, 30 // OK (y is new)
```

At least **one new variable** required.

---

## 7️⃣ Constants (`const`)

```go
const Pi = 3.14
```

Rules:

* immutable
* evaluated at compile time
* no `:=`

```go
const (
	A = 1
	B = 2
)
```

---

### Untyped constants (powerful)

```go
const x = 10

var y int64 = x   // OK
var z float64 = x // OK
```

---

## 8️⃣ Pointer variables

```go
var p *int
```

* zero value = `nil`
* stores address, not value

```go
x := 10
p = &x
```

---

## 9️⃣ Reference types (important distinction)

These variables **hold descriptors**, not data:

* slice
* map
* channel
* function

```go
var s []int   // nil slice
var m map[string]int // nil map
```

Nil slices are safe to read from.
Nil maps panic on write.

---

## 🔟 Lifetime vs scope (deep concept)

```go
func f() *int {
	x := 10
	return &x
}
```

✔ Valid

Why?

* Go uses **escape analysis**
* variable moves to heap automatically

Scope ≠ lifetime.

---

## 1️⃣1️⃣ Best practices (senior rules)

### ✅ Prefer small scopes

```go
if err := do(); err != nil {
	return err
}
```

---

### ❌ Avoid package globals

```go
var config Config // ❌ unless necessary
```

---

### ❌ Avoid shadowing

Use `go vet` to catch this.

---

### ✅ Use descriptive names

```go
count int // good
c int     // bad (unless loop)
```

---

## 1️⃣2️⃣ Common mistakes

❌ Using `:=` at package level
❌ Shadowing `err` unintentionally
❌ Overusing globals
❌ Long-lived mutable state

---

## 1️⃣3️⃣ Mental model (lock this in)

* Variables live in **scopes**
* Scopes nest
* Inner can hide outer
* Zero values are real values
* Declaration ≠ assignment
* Lifetime is compiler-managed

---

## Final one-liner summary

> **Go variables are explicit, block-scoped, zero-initialized, and designed to make invalid states hard to represent.**

---

If you want next:

* `var` vs `:=` deeper rules
* variable capture in closures
* how variables behave in goroutines
* memory layout (stack vs heap)

Just tell me 👍
