# Go Data Types, Structs, Pointers, and Related Concepts

This document provides a comprehensive overview of data types, structs, pointers, and related concepts in Go.

## Basic Data Types

Go offers a rich set of built-in data types:

* **Booleans:** `bool` (`true` or `false`)
* **Strings:** `string` (immutable sequences of bytes)
* **Floating-Point Numbers:** `float32` (32-bit), `float64` (64-bit)
* **Complex Numbers:** `complex64` (64-bit), `complex128` (128-bit) - Represent complex numbers with real and imaginary
  parts.
* **Integers:**
    * Signed: `int`, `int8`, `int16`, `int32`, `int64` (The size of `int` is platform-dependent).
    * Unsigned: `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` (Unsigned integers cannot represent negative
      values).
* **Byte:** `byte` (alias for `uint8`) - Represents a single byte.
* **Rune:** `rune` (alias for `int32`) - Represents a Unicode code point.

## Aggregate Types

Go also provides aggregate types that can combine multiple values:

* **Arrays:** `[n]T` (Fixed-size sequences of elements of type `T`). Arrays are of fixed length, determined at compile
  time.
* **Slices:** `[]T` (Dynamically sized views of an underlying array). Slices are much more commonly used than arrays due
  to their flexibility. They are defined by a pointer to an array, a length, and a capacity.
* **Maps:** `map[K]V` (Unordered collections of key-value pairs, where keys are of type `K` and values are of type `V`).
  Maps provide fast lookups by key.
* **Structs:** `struct { ... }` (Collections of named fields, each with its own type). Structs are used to represent
  data records or objects.

## Structs

Structs are a fundamental building block in Go. They allow you to group related data together.

```go
type Person struct {
    Name    string
    Age     int
    Address Address // Nested struct
}

type Address struct {
    Street string
    City   string
}

func main() {
    person := Person{
        Name: "Alice",
        Age:  30,
        Address: Address{
        Street: "123 Main St",
        City:   "Anytown",
        },
    }
    
    fmt.Println(person.Name) // Accessing fields
    fmt.Println(person.Address.City) // Accessing nested fields
}

```

* **Fields:** Structs have named fields, each with a specific type.
* **Nested Structs:** Structs can contain other structs as fields, allowing you to create complex data structures.
* **Zero Value:** The zero value of a struct is a struct with all its fields set to their respective zero values.
* **Struct Literals:** You can initialize structs using struct literals, as shown in the example above.
* **Anonymous Fields:** Structs can have anonymous fields (fields without a name). These are useful for embedding types.
* **Methods:** You can define methods on structs, which are functions that operate on instances of the struct.

```go
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

person.Greet() // Calling a method
```

## Pointers

Pointers hold the memory address of a value. They are declared using the * operator.

* **`&` Operator:** Used to get the memory address of a variable.
* **`*` Operator:** Used to dereference a pointer, i.e., access the value at the memory address it points to.
* **Zero Value:** The zero value of a pointer is `nil`

## Type Aliases

Type aliases create a new name for an existing type. They don't create a new type; they are just another way to refer to
the same type.

```go
type MyInt int

var x MyInt = 10
var y int = 20

// x and y are of different types, even though MyInt is an alias for int.
```
## Constants

Constants are values that are known at compile time and cannot be changed. They are declared using the `const` keyword.

```go
const Pi = 3.14159
const (
E = 2.71828
Phi = 1.61803
)
```

## iota

`iota` is a special constant generator that can be used to create sequences of related constants. It increments its value each time it is used in a `const` declaration.
```go
const (
A = iota // 0
B         // 1
C         // 2
)
```

## Type Conversions

Go is a statically typed language, so you sometimes need to convert values from one type to another. This is done using type conversion expressions.

```go
x := 10
f := float64(x) // Convert int to float64
```