# Learning Go

## Hello World
```
package main

import "fmt"

func main() {
    fmt.Printf("Hello, world\n")
}
```

## Packages
### `main` package
Go program must start running in package `main`.

### *Factored* import
```
import (
    "fmt"
    "math/rand"
)
```
This is preferred format over:
```
import "fmt"
import "math/rand"
```
### Exported names
A is exported if it begins with a capital letter, for example: `Pi`.

`math.pi` is not exported, `math.Pi` is.

### Functions

#### function declaration/definition
An example of function:
```
func add(x int, y int) int {
    return x + y
}
```
Notice the argument list, return type

#### combine type annotation
```
x int, y int
```
can be combined as:
```
x, y int
```

#### multiple results (similar to returning tuple in Python)
```
func swap(x, y string) (string, string) {
    return y, x
}
```

#### Named return values
```
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum = x
    return
}
```
The `return` statement has no arguments, it is a *naked* return. It should only be used in short functions.

### Variables
Use `var` to declare a list of variables.
A `var` statement can be used both at package level or at function level.
```
package main

import "fmt"

var c, python, java bool

func main() {
    var i int
    fmt.Println(i, c, python, java)
}
```

#### Variable with initializer
If an initializer is present, type definition can be omitted.
```
var i, j int = 1, 2
```

#### Short variable declaration
Short assignment `:=` can be used inside a function. It can only be used inside a function, where a `var` declaration with implicit type is allowed.

It seems `:=` defines variable on the stack.

### Basic Types
Following are built-in types of Go
```
var (
    IsBusy bool = true
    IsFree bool = false

    Message string = "hello"

    counter int = 2
    ascii int8 = 4 
    myvar int16 = 32767
    myvar2 int32 = 65536
    myint3 int64
    myuint1 uint
    myuint2 uint8
    myuint3 uint16
    myuint4 uint32
    myuint5 uint64
    myuintptr uintptr

    y byte // alias for uint8

    x rune // alias for int32
                 // represents a Unicode code point

    varF32 float32
    varF64 float64

    varAA complex64
    varBB complex128
)
```

#### Zero values
Variable without initializer will have default zero value. 0 for numeric types, false for bool, '' for string.

### Type Conversion
Type conversion must be explicit, `T(v)`, examples:
```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
A simpler form:
```
i := 42
f := float64(i)
u := uint(f)
```
This simpler form can be used inside a function.

### Type inference
Type can be inferred when declaring a variable from the right hand side.
Examples:
```
var i int
j := i // j is an int

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

### Constant
Constant can be declared with `const`. Constants cannot be declared using `:=` syntax.

This implies constants won't get allocated in stack, only in the heap.

### Numeric Constants
Numeric constants are high-precision values.

## Control flows of Go

### For loop
```
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```
Initializer, post statement can be omitted
```
sum := 1
for ; sum < 1000; {
    sum += sum
}
```
Simicolons can also be dropped and it becomes `while`
```
for sum < 1000 {
    sum += sum
}
```
If condition is omitted, it becomes forever
```
for {
    // forever
}
```

### `if` statement
Parenthesis can be omitted, braces are needed
```
if x < 0 {
    fmt.Printf("x is less than 0")
} else {
    fmt.Printf("x is no less than 0")
}
```
#### `if` with short statement
`if` can start with a short statment
```
if v := input; v < limit {
    fmt.Printf("input is less than limit")
}
```

### `switch` statement
Example:
```
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.", os)
}
```

#### `switch` evaluation order
From top to bottom, stopping when a case succeeds.

#### `switch` can be written with no condition
```
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```

### `defer` statement
Defers the execution until the surrounding function returns.
```
package main

import "fmt"

func main() {
    defer fmt.Println("world")

    fmt.Println("hello")
}
```

#### Stacking `defer`s
`defer` will be stacked, they will be executed in last-in-first-out order.

## More types: structs, slices, maps
### Pointer
Define a point type using `*`. Get address using `&`.
```
package main

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i         // point to i
    fmt.Println(*p) // read i through the pointer
    *p = 21         // set i through the pointer
    fmt.Println(i)  // see the new value of i

    p = &j         // point to j
    *p = *p / 37   // divide j through the pointer
    fmt.Println(j) // see the new value of j
}
```
### Struct
Struct defines a collection of fields
```
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    fmt.Println(Vertex{1, 2})
}
```

#### Pointers to structs
`(*p).X` is cumbersome, language permits to write `p.X`.

#### struct literals
```
package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}  // has type Vertex
    v2 = Vertex{X: 1}  // Y:0 is implicit
    v3 = Vertex{}      // X:0 and Y:0
    p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
    fmt.Println(v1, p, v2, v3)
}
```

### Array
`[n]T` is an array of n values of type T
```
package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
}
```

#### Array slice
A slice is defined by two indices, e.g. `a[low:high]`
```
package main

import "fmt"

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[1:4]
    fmt.Println(s)
}
```

#### Array slices are references
```
package main

import "fmt"

func main() {
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)

    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)

    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)
}
```

#### Slice literals
A slice literals defines an array without the length
```
package main

import "fmt"

func main() {
    q := []int{2, 3, 5, 7, 11, 13}
    fmt.Println(q)

    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)

    s := []struct {
        i int
        b bool
    }{
        {2, true},
        {3, false},
        {5, true},
        {7, true},
        {11, false},
        {13, true},
    }
    fmt.Println(s)
}
```

#### Slice defaults
If lower/upper bound is omitted, it defaults to 0 and length of array.
For the array
```
var a [10]int
```
these slice expressions are equivalent:
```
a[0:10]
a[:10]
a[0:]
a[:]
```

#### Slice length and capacity
A slice has both a *length* and a *capacity*, these can be reported by `len(s)` and `cap(s)` respectively.
```
package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    // Slice the slice to give it zero length.
    s = s[:0]
    printSlice(s)

    // Extend its length.
    s = s[:4]
    printSlice(s)

    // Drop its first two values.
    s = s[2:]
    printSlice(s)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

#### Nil slices
Nil slices has a zero length and capacity.
```
package main

import "fmt"

func main() {
    var s []int
    fmt.Println(s, len(s), cap(s))
    if s == nil {
        fmt.Println("nil!")
    }
}
```

#### Creating slice with `make`
Slices can be created with function `make`.
```
package main

import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a)

    b := make([]int, 0, 5)
    printSlice("b", b)

    c := b[:2]
    printSlice("c", c)

    d := c[2:5]
    printSlice("d", d)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}
```

#### Slices of slices
Slice can contain any type, including other slices
```
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Create a tic-tac-toe board.
    board := [][]string{
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
    }

    // The players take turns.
    board[0][0] = "X"
    board[2][2] = "O"
    board[1][2] = "X"
    board[1][0] = "O"
    board[0][2] = "X"

    for i := 0; i < len(board); i++ {
        fmt.Printf("%s\n", strings.Join(board[i], " "))
    }
}
```

#### Appending to a slice
Built-in function `append` can be used to append new elements to a slice
```
var s []int

// append works on nil slices.
s = append(s, 0, 1, 2)
```

#### Range
`for` statement has a `range` form, for example:
```
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}
```
You can skip index by using variable name `_`.
```
for _, value := range pow {
    fmt.Printf("%d\n", value)
}
```

### Exercise: Slices
```
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    var pic [][]uint8;
    for i := 0; i < dy; i++ {
        var row []uint8;
        for j := 0; j < dx; j++ {
            pixel := j % 4
            row = append(row, uint8(pixel))
        }
        pic = append(pic, row)
    }
    return pic
}

func main() {
    pic.Show(Pic)
}
```

### Map
A map maps keys to values. Zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added. The `make` function can return a map of given type.
```
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])
}
```

#### Map literals
Map literals can be used to initialize a map.
```
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var (
    m = map[string]Vertex{
        "Bell Labs": Vertex{
            40.68433, -74.39967,
        },
        "Google": { // type name can be omited
            37.42202, -122.08408,
        },
    }
    m2 = map[int]string{
        5: "Five",
        6: "Six",
    }
)

func main() {
    fmt.Println(m)
    fmt.Println(m2)
}
```
#### Exercise: Maps
```
package main

import (
    "strings"
    "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    var wc = make(map[string]int)
    for _, w := range strings.Fields(s) {
        wc[w] = int(wc[w]) + 1
    }
    return wc
}

func main() {
    wc.Test(WordCount)
}
```

#### Mutating Maps
Assign an element
```
m[key] = elem
```
Retrieve an element
```
elem = m[key]
```
Test if a key is defined with a two-value assignment
```
elem, ok = m[key]
```
or
```
elem, ok := m[key]
```
Example:
```
package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])

    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])

    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])

    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}
```

### Function Values
Function values are like function objects in JavaScript.
```
package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))

    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))
}
```

#### Function closures
Function closure is similar to JavaScript's equivalent
```
package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}
```

#### Exercise: Fibonacci closure

```
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(x int) int {
    known := []int{0, 1}
    return func(x int) int {
        for i := len(known); i <= x; i++ {
            v := known[i-1] + known[i-2];
            known = append(known, v)
        }
        return known[x]
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f(i))
    }
}
```

## Methods and interfaces
Go doesn't have classes but one can define methods on types.

A method is a function with a special *receiver* argument.
```
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}
```

#### methods are functions
A method is just a function with a receiver argument
```
func main() {
    v := Vertex{3, 4}
    fmt.Println(Abs(v))
}
```

#### Method on non-struct types
Method can be defined on non-struct types too
```
package main

import (
    "fmt"
    "math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main() {
    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
}
```

#### pointer receivers
```
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 { // receiver (v Vertex)
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) { // receiver (v *Vertex)
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(10)
    fmt.Println(v.Abs())
}
```

#### Methods and pointer indirection
Pointer argument only accepts pointer, pointer receiver can take both value or pointer.
```
package main

import "fmt"

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(2)
    ScaleFunc(&v, 10)

    p := &Vertex{4, 3}
    p.Scale(3)
    ScaleFunc(p, 8)

    fmt.Println(v, p)
}
```

#### Choosing a value or pointer receiver
There are two reasons to use a pointer receiver:

1. You can modify the passed in object
2. Avoid copying large structure

Generally speaking, all methods on a given type should have either value or pointer receivers but not a mixture of both.

### Interface
Interface defines a set of method signatures.
```
package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat implements Abser
    a = &v // a *Vertex implements Abser

    // In the following line, v is a Vertex (not *Vertex)
    // and does NOT implement Abser.
    a = v

    fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

#### Interfaces are implemented implicitly
We don't need to use `implements` keyword to declare that.
```
package main

import "fmt"

type I interface {
    M()
}

type T struct {
    S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
    fmt.Println(t.S)
}

func main() {
    var i I = T{"hello"}
    i.M()
}
```

#### Interface values
Interface values can be though of as a tuple of a value and a concrete type: `(value, type)`
```
package main

import (
    "fmt"
    "math"
)

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    fmt.Println(t.S)
}

type F float64

func (f F) M() {
    fmt.Println(f)
}

func main() {
    var i I

    i = &T{"Hello"}
    describe(i)
    i.M()

    i = F(math.Pi)
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
```

#### Interface values with `nil` underlying values
Interface value may hold a concrete value which is `nil`.
```
package main

import "fmt"

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

func main() {
    var i I

    var t *T
    i = t
    describe(i)
    i.M()

    i = &T{"hello"}
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
```

#### `nil` interface values
A `nil` interface value holds neither value nor concrete type.

#### The empty interface
Interface without any methods is *empty interface*.
```
package main

import "fmt"

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

func main() {
    var i I

    var t *T
    i = t
    describe(i)
    i.M()

    i = &T{"hello"}
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
```

### Type assertions
A type assertion provides access to an interface value's underlying concrete value.
```
package main

import "fmt"

func main() {
    var i interface{} = "hello"

    s := i.(string)
    fmt.Println(s)

    s, ok := i.(string)
    fmt.Println(s, ok)

    f, ok := i.(float64)
    fmt.Println(f, ok)

    f = i.(float64) // panic
    fmt.Println(f)
}
```

#### Type switches
A type switch is a construct that permits erveral type asertions in series.
```
package main

import "fmt"

func do(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}

func main() {
    do(21)
    do("hello")
    do(true)
}
```

### Stringers
Stringer interface defines one method `String()`. A type that implements this interface provides a way to convert its values into strings.
```
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}
    fmt.Println(a, z)
}
```

#### Exercise: Stringers
Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

For instance, `IPAddr{1, 2, 3, 4}` should print as "`1.2.3.4`".

```
package main

import (
    "fmt"
    "strings"
)


type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
    ss := []string{string('1'),string('2'),string('3'),string('4')}
    fmt.Println(strings.Join(ss, "."))
    
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
```

### Errors 
Error Type is a built-in interface 
```
type error interface {
    Error() string
}
```
`nil` value of error type indicate success, non-nil error denotes failure.
```
package main

import (
    "fmt"
    "time"
)

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s",
        e.When, e.What)
}

func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }
}
```

#### Exercise: Errors
```
package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    v := float64(e)
    return fmt.Sprintf("cannot Sqrt negative number: %v", v)
}

func Sqrt(x float64) (float64, error) {
    v := x
    if x < 0 {
        return 0, ErrNegativeSqrt(v)
        //return 2, nil
    }
    return 0, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
```

### Readers
The `io` package defines `io.Reader` interface, which as a `Read` method:
```
func (T) Read(b []byte) (n int, err error)
```
Example:
```
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello, Reader!")

    b := make([]byte, 8)
    for {
        n, err := r.Read(b)
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n])
        if err == io.EOF {
            break
        }
    }
}
```

#### Exercise: Readers

Implement a `Reader` type that emits an infinite stream of ASCII character 'A'
```
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
    b[0] = byte('A')
    return 1, nil
}

func main() {
    reader.Validate(MyReader{})
}
```
#### Exercise: rot13Reader
Define a new `Reader` by making use of another one
```
package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (n int, e error) {
    return r13.r.Read(b) 
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
```
### Images
Package `image` defines interface `Image`

```
package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (n int, e error) {
    return r13.r.Read(b) 
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
```

Example:
```
package main

import (
    "fmt"
    "image"
)

func main() {
    m := image.NewRGBA(image.Rect(0, 0, 100, 100))
    fmt.Println(m.Bounds())
    fmt.Println(m.At(0, 0).RGBA())
}
```

### Exercise: Images
Defines your own image type that implements the `Image` interface
```
package main

import (
    "image"
    "image/color"
    "golang.org/x/tour/pic"
)

type Image struct{
    x, y, width, height int
}
/*
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
*/

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(img.x, img.y, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
    return color.RGBA{0, 0, 255, 255}
}

func main() {
    m := Image{0, 0, 400, 300}
    pic.ShowImage(m)
}
```

## Concurrency
### Goroutines
In Go, the term *goroutine* refers to a lightweight thread. Just use keyword `go` to start a new goroutine

Example:
```
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 3; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

### Channels
Channels are typed conduit for communications between two goroutines.

```
package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x+y)
}
```
Two goroutines were started, each computes summary of half of the slice. Then results are sent back to main thread using channel.

### Buffered Channels
Channels can be buffered, buffer size can be specified when being created
```
ch := make(chan int, 100)
```
Example:
```
package main

import "fmt"

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 3
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    ch <- 5
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

### Range and Close
A sender can close the channel to indicate there's no more values to be sent. Receivers can test if a channel is closed, but should never close a channel.

Examples:
```
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c {
        fmt.Println(i)
    }
}
```

### Select
`select` statement lets a goroutine to wait on multiple channels
Example:
```
package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}
```

### Default Selection
`default` case is run when there's no other case is ready, it can be used to try a send or receive without blocking.
Example:
```
package main

import (
    "fmt"
    "time"
)

func main() {
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    for {
        select {
        case <-tick:
            fmt.Println("tick.")
        case <-boom:
            fmt.Println("BOOM!")
            return
        default:
            fmt.Println("    .")
            time.Sleep(50 * time.Millisecond)
        }
    }
}
```

### Exercise: Equivalent Binary Trees
