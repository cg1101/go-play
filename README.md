# Learning Go

## Part I, Using the tour
### 1. Welcome
#### 1.1 Hello World
```
package main

import "fmt"

func main() {
    fmt.Printf("Hello, world\n")
}
```
Update:
```
package main

import "fmt"

func main() {
	const 名字 = "世界"
	fmt.Printf("Hello, %v\n", 名字)
}
```
Note that any unicode letters can be used for variable names.

#### 1.2 Go local
The tour is available in many different languages

#### 1.3 Go offline
The tour is also available as a stand-alone program. Run following command to install it:
```
go get golang.org/x/tour
```
Then run `tour` to start it.

#### 1.4 The Go Playground
The tour is built atop the [Go Playground](https://play.golang.org). This services can be used to test ideas online.

#### 1.5 Congratulations
You now know how to use this tour.

## Part II Basics
### 2. Packages, variables and funcitons
#### 2.1  Packages
Go program consists of packages. Go program must start running in package `main`.

#### 2.2 Imports
Import statements can be grouped into *Factored* import.
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
#### 2.3 Exported names
A is exported if it begins with a capital letter, for example: `Pi`.

`math.pi` is not exported, `math.Pi` is.

#### 2.4 Functions
Function declaration/definition
An example of function:
```
func add(x int, y int) int {
    return x + y
}
```
Notice the argument list, return type

#### 2.5 Functions continued
Combine type annotation
```
x int, y int
```
can be combined as:
```
x, y int
```

#### 2.6 Multiple result
Function can return multiple results (similar to returning tuple in Python)
```
func swap(x, y string) (string, string) {
    return y, x
}
```

#### 2.7 Named return values
```
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum = x
    return
}
```
The `return` statement has no arguments, it is a *naked* return. It should only be used in short functions.

#### 2.8 Variables
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

#### 2.9 Variable with initializers
If an initializer is present, type definition can be omitted.
```
var i, j int = 1, 2
```

#### 2.10 Short variable declaration
Short assignment `:=` can be used inside a function. It can only be used inside a function, where a `var` declaration with implicit type is allowed.

It seems `:=` defines variable on the stack.

#### 2.11 Basic Types
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

    y byte  // alias for uint8

    x rune  // alias for int32
            // represents a Unicode code point

    varF32 float32
    varF64 float64

    varAA complex64
    varBB complex128
)
```

#### 2.12 Zero values
Variable without initializer will have default zero value. 0 for numeric types, false for bool, '' for string.

#### 2.13 Type Conversion
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

#### 2.14 Type inference
Type can be inferred when declaring a variable from the right hand side.
Examples:
```
var i int
j := i // j is an int

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

#### 2.15 Constants
Constant can be declared with `const`. Constants cannot be declared using `:=` syntax.

This implies constants won't get allocated in stack, only in the heap.

#### 2.16 Numeric Constants
Numeric constants are high-precision values.

#### 2.17 Congratulations
This module is done

### 3. Flow control statements: for, if, else, switch and defer

#### 3.1 For loop
```
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```
#### 3.2 For continued
Initializer, post statement can be omitted
```
sum := 1
for ; sum < 1000; {
    sum += sum
}
```
#### 3.3 For is Go's `while`
Simicolons can also be dropped and it becomes `while`
```
for sum < 1000 {
    sum += sum
}
```
#### 3.4 Forever
If condition is omitted, it becomes forever
```
for {
    // forever
}
```

#### 3.5 `if` statement
Parenthesis can be omitted, braces are needed
```
if x < 0 {
    fmt.Printf("x is less than 0")
} else {
    fmt.Printf("x is no less than 0")
}
```
#### 3.6 `if` with short statement
`if` can start with a short statment
```
if v := input; v < limit {
    fmt.Printf("input is less than limit")
}
```
#### 3.7 `if` and `else`
```
if v := math.Power(x, n); v < lim {
    return v
} else {
    fmr.Printf("%g >= %g\n", v, lim)
}
```

#### 3.8 Esercise: Loops and Functions
Implement a function `Sqrt` that calculates square root of given value.
```
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var stillBig = true
	for counter := 0; stillBig && counter < 200; counter ++ {
		delta :=(z*z - x) / (2 * z)
		stillBig = math.Abs(delta) > 1.e-10
		fmt.Println("counter=", counter, "z=", z, "delta=", delta, "big?", stillBig)
		z -= delta
	}
	return z

}

func main() {
	fmt.Println("Square root of 2 is", Sqrt(2))
	fmt.Println("result from math.Sqrt", math.Sqrt(2))
	fmt.Println("math.abs(-2)", math.Abs(-2))
}
```

#### 3.9 `switch` statement
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

#### 3.10 `switch` evaluation order
From top to bottom, stopping when a case succeeds.

#### 3.11 `switch` can be written with no condition
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

#### 3.12 `defer` statement
Defers the execution until the surrounding function returns.
```
package main

import "fmt"

func main() {
    defer fmt.Println("world")

    fmt.Println("hello")
}
```

#### 3.13 Stacking `defer`s
`defer` will be stacked, they will be executed in last-in-first-out order.

#### 3.14 Congratulations
This module finishes.

### 4 More types: structs, slices, maps
#### 4.1 Pointer
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
#### 4.2 Structs
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

#### 4.3 Struct fields

#### 4.4 Pointers to structs
`(*p).X` is cumbersome, language permits to write `p.X`.

#### 4.5 Struct literals
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

#### 4.6 Arrays
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

#### 4.7 Array slices
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

#### 4.8 Array slices are references
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

#### 4.9 Slice literals
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

#### 4.10 Slice defaults
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

#### 4.11 Slice length and capacity
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

#### 4.12 Nil slices
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

#### 4.13 Creating slice with `make`
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

#### 4.14 Slices of slices
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

#### 4.15 Appending to a slice
Built-in function `append` can be used to append new elements to a slice
```
var s []int

// append works on nil slices.
s = append(s, 0, 1, 2)
```

#### 4.16 Range
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
#### 4.17 Range continued
You can skip index by using variable name `_`.
```
for _, value := range pow {
    fmt.Printf("%d\n", value)
}
```

#### 4.18 Exercise: Slices
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

#### 4.19 Maps
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

#### 4.20 Map literals
Map literals can be used to initialize a map.

#### 4.21 Map literals continued

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

#### 4.22 Mutating Maps
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

#### 4.23 Exercise: Maps
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

#### 4.24 Function Values
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

#### 4.25 Function closures
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

#### 4.26 Exercise: Fibonacci closure

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
#### 4.27 Congratulations

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
Use predefined tree structure
```
type Tree struct {
    Left *Tree
    Value int
    Right *Tree
}
```
First, implement a `Walk` function that traverse tree and emits its value
```
package main

import (
    "fmt"
    "golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        Walk(t.Left, ch)
    }
    ch <= t.Value
    if t.Right != nil {
        Walk(t.Right, ch)
    }
}

func main() {
    t1 := tree.New(1)
    ch := make(chan int)
    go Walk(t1, ch)
    const n = 10
    for i := 0; i < n; i++ {
        fmt.Println(<-ch, " ")
    }
}
```
Then implement `Same` function to determine if two trees store the same values.
```
package main

import (
    "fmt"
    //"strings"
    "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        //fmt.Println("walk left")
        Walk(t.Left, ch)
    }
    ch <- t.Value
    if t.Right != nil {
        //fmt.Println("walk right")
        Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    seq1 := make([]int, 10)
    seq2 := make([]int, 10)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for i := 0; i < 10; i++ {
        seq1 = append(seq1, <-ch1)
    }
    for j := 0; j < 10; j++ {
        seq2 = append(seq2, <-ch2)
    }
    rs1 := fmt.Sprintf("%q", seq1)
    rs2 := fmt.Sprintf("%q", seq2)

    return rs1 == rs2
}

func main() {
    t1a := tree.New(1)
    t1b := tree.New(1)
    t2 := tree.New(2)
    fmt.Printf("Same(t1a, t1b) -> %v\n", Same(t1a, t1b))
    fmt.Printf("Same(t1a, t2) -> %v\n", Same(t1a, t2))
}
```

### sync.Mutex
In Go, one can use `sync.Mutex` to achieve *Mutual Exclusion*.
```
package main

import (
    "fmt"
    "sync"
    "time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
    v   map[string]int
    mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
    c.mux.Lock()
    // Lock so only one goroutine at a time can access the map c.v.
    c.v[key]++
    c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
    c.mux.Lock()
    // Lock so only one goroutine at a time can access the map c.v.
    defer c.mux.Unlock()
    return c.v[key]
}

func main() {
    c := SafeCounter{v: make(map[string]int)}
    for i := 0; i < 1000; i++ {
        go c.Inc("somekey")
    }

    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))
}
```

### Exercise: Web Crawler
Implement web crawler that uses multiple threads to crawl a list of urls.

```
package main

import (
    "fmt"
    "sync"
    "time"
)

type TodoList struct {
    urls   map[string]bool
    mux sync.Mutex
}

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(todo *TodoList, url string, depth int, fetcher Fetcher) {
    // TODO: Fetch URLs in parallel.
    // TODO: Don't fetch the same URL twice.
    // This implementation doesn't do either:
    if depth <= 0 {
        return
    }
    //fmt.Println("lock mutex")
    todo.mux.Lock()
    loaded := todo.urls[url]
    if !loaded {
        todo.urls[url] = true
    } else {
        //fmt.Println("unlock mutex")
        todo.mux.Unlock()
        return
    }
    //fmt.Println("unlock mutex")
    todo.mux.Unlock()
    fmt.Println("fetching", url)
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    //fmt.Printf("%v\n", urls)
    for _, u := range urls {
        go Crawl(todo, u, depth-1, fetcher)
    }
    return
}

func main() {
    todo := TodoList{urls: make(map[string]bool)};
    fmt.Println("todo", todo)

    go Crawl(&todo, "https://golang.org/", 4, fetcher)
    time.Sleep(1000)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
    "https://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "https://golang.org/pkg/",
            "https://golang.org/cmd/",
        },
    },
    "https://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "https://golang.org/",
            "https://golang.org/cmd/",
            "https://golang.org/pkg/fmt/",
            "https://golang.org/pkg/os/",
        },
    },
    "https://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
    "https://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
}
```
