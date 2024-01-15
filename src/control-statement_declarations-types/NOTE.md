# Control Statement; Declarations & Types

## If-then-else

if-then statements require braces

```go
if a == b {
    fmt.Println("a equals b")
} else {
    fmt.Println("a is not equal to b")
}
```

They can start with a short declaration or statement

```go
if err := doSomething(); err != nil {
    return err
}
```

## For loops

1. Explicit control with an index variable

```go
for i := 0; i < 10; i++ {
    fmt.Printf("(%d, %d)\n", i, i * i)

    // prints (0, 0) up to (9, 81)
}
```

Three parts, all optional (initialize, check, increment)
The loop ends when the explicit check fails (e.g., i == 10)

2. Implicit control through the range operator for arrays & slices

```go

// one var: i is an index 0, 1, 2, ...

for i := range myArray {
    fmt.Println(i, myArray[i])
}

// two vars: i is the index, v is a value

for i, v := range myArray {
    fmt.Println(i, v)
}
```

The loop ends when the range is exhausted

3. An infinite loop with an explicit break

```go
i, j := 0, 3
// this loop must be made to stop
for {
    i, j = i + 50, j * j
    fmt.Println(i, j)

    if j > i {
        break // when i = 150, j = 6561
    }
}
```

There is also continue to make an iteration start over

Here's a common mistake:

If you only want range values, you need the blank identifier:

```go
// two vars: _ is the index (ignored),
//          v is the value
for _, v := range myArray {
    fmt.Println(v)
}
```

Sometimes you may not get a compile error for a type mismatch if you use only the one-var format (a slice of ints!)

The _ is an untyped, reusable "variable" placeholder

## Labels and loops

Sometimes we need to break or continue the outer loop
(nested loop for quadratic search)

```go
outer:
    for k := range testItemsMap {
        for _, v := range returnedData {
            if k == v.ID {
                continue outer
            }
        }

        t.Errorf("key not found: %s", k)
    }
```

We need a label to refer to the outer loop

## Switch

It is a shortcut replacing a series of if-then statements

```go
switch a := f.Get(); a {
    case 0, 1, 2:
        fmt.Println("underflow possible")

    case 3, 4, 5, 6, 7, 8:

    default:
        fmt.Println("warning: overload")
}
```

May be empty and do not fall through (break is not required)

## Switch on true

Arbitrary comparisons may be made for an switch with no argument

```go
a := f.Get()

switch {
    case a <= 2:
        fmt.Println("underflow possible")

    case a <= 8:
        // evaluated in order

    default:
        fmt.Println("warning: overload")
}
```

# Packages

Every standalone program has a main package

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

Nothing is "global"; It's either in your package or in another
It's either at package scope or function scope

## Package-level declarations

You can declare anything at package scope

```go
package secrets

const DefaultUUID = "00000000-0000-0000-0000-000000000000"
var secretKey string

type k8secret struct {
    ...
}

func Do(it string) error {
    ...
}
```

The short declaration (:=) can only be used inside of a function

## Package control visibility

Every name that's capitalized is exported

```go
package secrets
import ...
type internal struct {
    ...
}

func GetAll(space, name string) (map[string]string, error) {
    ...
}
```

## No cycles

A package "A" cannot import a package that imports A

```go
package A
import "B"

// --------

package B
import "A"
```

## Initialization

Items within a package get initialized before main

```go
const A = 1

var B int = C
var C int = A

func Do() error {
    ...
}

func init() {
    ...
}
```

Only the runtime can call `init`, also before `main`

## What makes a good package?

A package should embed deep functionality behind a simple API

```go
package os

func Create(name string) (*File, error)
func Open(name string) (*File, error)

func (f *File) Read(b []byte) (n int, err error)
func (f *File) Write(b []byte) (n int, err error)
func (f *File) Close() error
```

Functions: These are standalone and do not belong to any specific type. They are called independently of any object or instance. For example, Create(name string) and Open(name string) are functions that belong to the os package. They are called directly using their names and passing the required arguments, like os.Create("filename").

Methods: These are associated with a type (like File in your example) and have a receiver (like (f *File)). Methods are called on an instance of that type. For example, Read(b []byte) is a method that is called on an instance of File, like f.Read(b) where f is an instance of File.

The Unix file API is perhaps the best example of this model
Roughly five functions hide a lot of complexity from the user

## Declaration

Six ways to introduce a name:

- Constant declaration with `const`
- Type declaration with type
- Variable declaration with `var` (must have type or initial value, sometimes both)
- Short, initialized variable declaration of any type :=
  only inside a function
- Function declaration with func (methods may only be declared at package level)
- Formal parameters and named returns of a function

## Variable declarations

```go
var a int // 0 by default
var b int = 1
var c = 1 // int
var d = 1.0 // float64

// declaration block
var (
    x, y int
    z    float64
    s    string
)
```

The short declaration operator `:=` has some rules:
1. It can't be used outside of a function
2. It must be used (instead of var) in a control statement (if, etc.)
3. It must declare at least one new variable

```go
err := doSomething()

err := doSomethingElse() // WRONG

x, err := getSomeValue() // OK; err is not redeclared
```

4. It won't be re-used an existing declaration from an outer scope

## Structural typing

It's the same type if it has the same structure or behavior
- arrays of the same size and base type
- slices with the same base type
- maps of the same key and value types
- structs with the same sequence of field names/types
- functions with the same parameter & return types

```go
a := [...]int{1, 2, 3}
b := [3]int{}

a = b // OK
c := [4]int{}

a = c // NOT OK
```

Go uses structural typing in most cases

## Named typing

It's the only the same type if it has the same declared type name

```go
type x int

func main() {
    var a x // x is a defined type; base int

    b := 12 // b defaults to int

    a = b // TYPE MISMATCH

    a = 12 // OK, untyped literal
    a = x(b) // OK, type conversion
}
```

Go uses named typing for non-function user-declared types
