# Go Generics Cheatsheet

## Basic Syntax
```go
// Generic function
func Identity[T any](x T) T {
    return x
}

// Generic type
type Stack[T any] []T

// Multiple type parameters
func Pair[T, U any](first T, second U) (T, U) {
    return first, second
}
```

## Type Constraints
```go
// Built-in constraints
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string
}

type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64
}

// Custom constraint
type Stringer interface {
    String() string
}

type Comparable interface {
    ComparableTo(T) bool
}
```

## Common Patterns

### Collections
```go
// Generic slice operations
func Map[T, U any](ts []T, f func(T) U) []U {
    us := make([]U, len(ts))
    for i := range ts {
        us[i] = f(ts[i])
    }
    return us
}

func Filter[T any](ts []T, f func(T) bool) []T {
    var result []T
    for _, v := range ts {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}
```

### Data Structures
```go
// Generic stack
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}
```

### Comparables
```go
func Max[T Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}
```

## Advanced Features

### Type Sets
```go
type Numeric interface {
    ~int | ~float64
}

func Sum[T Numeric](values ...T) T {
    var total T
    for _, v := range values {
        total += v
    }
    return total
}
```

### Embedded Constraints
```go
type Processor[T any] interface {
    Process(T) T
    ~struct{ Data T }  // Type must be struct with Data field
}
```

### Generic Methods (Only on generic types)
```go
type Container[T any] struct {
    Value T
}

// Method on generic type
func (c *Container[T]) Get() T {
    return c.Value
}
```

## Limitations & Gotchas
1. **No generic methods** on non-generic types
2. **Type inference** works in most cases
3. **Comparable** includes: boolean, numeric, string, pointer, channel, interface, struct/array of comparable types
4. **~T** means "underlying type T"
5. Cannot use type parameters in method signatures of non-generic types

## Quick Examples
```go
// Sort any ordered slice
func SortSlice[T Ordered](s []T) []T {
    sorted := make([]T, len(s))
    copy(sorted, s)
    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i] < sorted[j]
    })
    return sorted
}

// Generic pointer handling
func DerefOrZero[T any](ptr *T) T {
    if ptr != nil {
        return *ptr
    }
    var zero T
    return zero
}

// Working with channels
func Merge[T any](chans ...<-chan T) <-chan T {
    out := make(chan T)
    // ... implementation
    return out
}
```

## Key Points
- Use `any` for unconstrained types (replaces `interface{}`)
- Use `comparable` for types supporting `==` and `!=`
- Create custom constraints with type unions `|`
- Type parameters can have methods: `[T fmt.Stringer]`
- Zero value: `var zero T`