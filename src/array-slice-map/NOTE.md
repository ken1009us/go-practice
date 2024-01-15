# Arrays, Slices, and Maps

## Basic Syntax

### Composite types

The array and the slice are like a string, it is the sequence of things literally laid down in memory one after another.

The map is a map of keys to values.

- string
```bash
"area"
```

- array
```bash
[4]int

1, 12, 4, 8
```

Arrays are not used a lot since they are fixed size.

- slice (variable length array)
```bash
[]int

1, 13, 3, ..., 17
```

-map
```bash
map[string]int

"to": 1
"from": 12
"into": 3
"above": 0
```

### Array

Arrays are passed by value, thus elements are copied.

- Length fixed at compile time
- Passed by value
- Comparable (==)
- Can be used as map key
- Useful as "pseudo" constants

Arrays example:

```go
var a [3]int
var b [3]int{0, 0, 0}
var c [...]{0, 0, 0} // sized by initializer

var d [3]int
d = b // elements copied

var m [...]int{1, 2, 3, 4}

c = m // TYPE MISMATCH
```

### Slice

Slices are passed by reference; no copying, updating OK.

- Variable length
- Passed by reference
- Not comparable
- Cannot be used as map key
- Has copy & append helpers
- Useful as function parameters

Slices example:

```go
var a []int // nil, no storage
var b = []int{1, 2} // initialized

a = append(a, 1) // append to nil OK
b = append(b, 3) // []int{1, 2, 3}

a = b // overwrites a

d := make([]int, 5) // []int{0, 0, 0, 0, 0}
e := a // same storage (alias)

e[0] == b[0] // true
```

### Map

Maps are dictionaries: indexed by key, returning a value, You can read from a nil map, but inserting will panic.

Maps are passed by reference; no copying, updating OK.

The type used for the key must have == and != defined (not slices, maps, or funcs).

Maps example:

```go
// This descriptor doesn't reference any actual hash table
var m map[string]int // nil, no storage
// What make does is to simply create that hash table in the background
p := make(map[string]int) // non-nil but empty

// Zero is the default of an int type
a := p["the"] // returns 0
b := m["the"] // same thing

// You CAN'T write key into the nil map
m["and"] = 1 // PANIC - nil map

// Copy the descriptor in p to m so that we can insert key and value to m
m = p
m["and"]++ // OK, same map as p now
c := p["and"] // returns 1

```

Maps can't be compared to one another; maps can be compared only to nil as a special case.

```go
var m = map[string]int{
    "and": 1,
    "the": 1,
    "or": 2,
}

var n map[string]int

b := m == n // SYNTAX ERROR
c := n == nil // true
d := len(m) // 3
e := cap(m) // TYPE MISMATCH, cap is ask for the capacity in slice
```

Maps have a special two-result lookup function. The second variable tells you if the key was there.

```go
// literal map
p := map[string][int] // non-nil but empty

// look up for "the", got 0, but we don't if it is empty map or not
a := p["the"] // returns 0

// we can use below to find if it is empty map
b, ok := p["and"] // 0, false

p["the"]++

c, ok := p["the"] // 1, true

// conditional statement:
// declare variables w and ok by indexing p with a string "the"
// the very next piece is the actual if condition
if w, ok := p["the"]; ok {
    // we know w is not the default value
    ...

}
```

## Built-ins

Each type has certain built-in functions.

```
len(s)          string  string  length
len(a), cap(a)  array   array length, capacity(constant)
make(T, x)      slice   slice of type T with length x and capacity x
make(T, x, y)   slice   slice of type T with length x and capacity y
copy(c, d)      slice   copy from d to c; # = min of the two lengths
c=append(c, d)  slice   append d to c and return a new slice result
len(s), cap(s)  slice   slice length and capacity
make(T)         map     map of type T
make(T, x)      map     map of type T with space hint for x elements
delete(m, k)    map     delete key k (if present, else no change)
len(m)          map     map length
```

## Make nil useful

Nil is a type of zero: it indicates the absence of something. Many built-ins are safe: len, cap, range

```go
var s []int
var m map[string]int

l := len(s) // length of nil size is 0

i, ok := m["int"] // 0, false for any missing key

// skip if s is nil or empty
for _, v := range s {
    ...
}
```
