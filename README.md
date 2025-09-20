# Package 
- Every Go program is made up of packages.
- Programs start running in package main.
- This program is using the packages with import paths "fmt" and "math/rand".
- By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.


# Imports
- Import path: the full path used to import a package into your program.
- 
```Go
import (
	"fmt"
	"math"
)
```


# Exports
- In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.
- pizza and pi do not start with a capital letter, so they are not exported.

- When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

```Go
         import "math"
         func main() {
                  fmt.println(math.pi) // This will give Error
                  fmt.println(math.Pi) //  This works
         }
```

# Functions
- A function can take zero or more arguments.
```Go
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```
- When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
```Go
we can change x int , y int to
func add(x,y int) int {
         return x+y
}
```
- A function can return any number of result
```Go
func swap(x,y string) (string , string){
         return y,x 
}
this swap x and y string
```
- Name return values
  - In Go, when you write a function, you can give names to the return values.
  - Go's return values may be named. If so, they are treated as variables defined at the top of the function.
```Go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return    // naked return 
         this return the x and y value as the name parameter
}
```
  - These names should be used to document the meaning of the return values.
  - A return statement without arguments returns the named return values. This is known as a "naked" return.
  - Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
  - Function values in Go
     - Functions can be stored in variables.
     - They can be passed as arguments or returned from other functions.
```Go

func add(a, b int) int {
    return a + b
}

func main() {
    f := add          // assign function to variable
    fmt.Println(f(2,3)) // call via variable

    // pass function as argument
    apply := func(fn func(int, int) int, x, y int) int {
        return fn(x, y)
    }
    fmt.Println(apply(add, 4, 5)) // 9
}
```


# Variables
- The var statement declares a list of variables; as in function argument lists, the type is last.
- A var statement can be at package or function level
```Go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
// output 0 false false false
```
- A var declaration can include initializers, one per variable.

- If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
```Go

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
// output:= 1 2 true false no!
```
- Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

```Go
m:= 2 // this is not allowed
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
// output := 1 2 3 true false no!
```
- Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
- Zero Values
  - Variables declared without an explicit initial value are given their zero value.
  - The zero value is:
    - 0 for numeric types,
    - false for the boolean type, and
    - "" (the empty string) for strings.
```Go
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
         //output  0 0 false ""
```



# Basic Types
- Boolean `var flag bool = true`
- String `var name string = "kalpesh"`
- Integer
  - Signed integer can be negative `var int age = 23 or -2`
    - `int, int8, int16, int32, int64`
  - Unsigned integer can be only be positive `var positiveNumber uint = 42`
    - `uint, uint8, uint16, uint32, uint64, uintptr`
- Floating Point Number (`float32, float64`) `var pi float64 = 3.14159`
- Complex Number (`complex64, complex128`) `var c complex128 = 2 + 3i`
```Go
Instead of writing var again and again, you can group them:
var (
    a bool       = true
    b string     = "hello"
    c int        = 42
    d float64    = 3.14
    e complex64  = 2 + 5i
)

```
- Use int for whole numbers, float64 for decimals, string for text, and only use others if you really need them (like saving memory or working with Unicode/bytes directly).




# Type Conversion
- In Go, if you want to change a value from one type to another, you have to do it explicitly.
```Go
var i int = 42
var f float64 = float64(i) // convert int → float64
var u uint = uint(f)       // convert float64 → uint
```
- In C, you can sometimes assign between types without converting — the compiler does it for you
- In Go, the compiler will not allow you to assign between different types unless you explicitly convert.
```Go
i := 42
f := i   // error: cannot use i (type int) as type float64
f := float64(i)   //works
```



# Type inference
When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:
```Go
var i int
j := i // j is an int
```
But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:
```Go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```


# Constants
- Constants are declared like variables, but with the const keyword.
- Constants can be character, string, boolean, or numeric values.
- Constants cannot be declared using the := syntax.



# For
- Go has only one looping construct, the for loop.
```Go
	for i := 0; i < 10; i++ {
		sum += i
	}

```
- The init and post statements are optional.
```Go
	for ; sum < 1000; {
		sum += sum
	}
```
- Go doesnot have any while init but it consist of 
```Go
	for sum < 1000 {
		sum += sum
	}
         If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
         for {
         }
```
# If
```Go
	if x < 0 {
		return sqrt(-x) + "i"
	}

Like for, the if statement can start with a short statement to execute before the condition.
Variables declared by the statement are only in scope until the end of the if.

         func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}else {
                  return 0
         }
	return lim
}

```


# Switch
- A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

- Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.



# Defer
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
```Go
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

counting
done
9
8
7
6
5
4
3
2
1
0
```

# Pointers
```Go
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
- Struct fields can be accessed through a struct pointer.
- To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

```Go
func main() {
    v := Vertex{1, 2}   // make a Vertex with X=1, Y=2
    p := &v             // p is a pointer to v
    p.X = 1e9           // change the X value of v using the pointer
    fmt.Println(v)      // print v
}


```

# Struct
- A struct is a collection of fields.
```Go
type Vertex struct{
	X int
	Y int
}
```
- it's is accessed using dot
- A struct literal denotes a newly allocated struct value by listing the values of its fields.
```Go
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

```

# Array
- Fixed Size
- `[n]T`
- var a [10]int 

# Slice
- dynamic Size
- `[]T`
- a[low : high]
- var s []int:= primes[1:4] // element goes from 1 to 3 last one is excluded
```Go
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
```
- The default is zero for the low bound and the length of the slice for the high bound.
- The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
- The zero value of a slice is nil. A nil slice has a length and capacity of 0 and has no underlying array.
- Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

- The make function allocates a zeroed array and returns a slice that refers to that array
```Go
a := make([]int, 5)  // len(a)=5

```
- Slices of Slices
```Go
func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
}
```
- `append` adds elements to a slice.
- If the slice doesn’t have enough space, Go makes a bigger array behind the scenes.
- It returns the new slice.
```Go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3}
    s = append(s, 4, 5)   // add more elements
    fmt.Println(s)        // [1 2 3 4 5]
}
```

# Range
- The range form of the for loop iterates over a slice or map.

- When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
- range is used in for loops to iterate over elements in slices, arrays, maps, strings, and channels.
```Go

func main() {
    nums := []int{10, 20, 30}

    for i, v := range nums {
        fmt.Println(i, v)   // index and value
    }
}

m := map[string]string{"a": "apple", "b": "banana"}

for k, v := range m {
    fmt.Println(k, v)  // key and value
}

```


# Maps
- A map maps keys to values.
- The zero value of a map is nil. A nil map has no keys, nor can keys be added.
- The make function returns a map of the given type, initialized and ready for use.
```Go
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

--or--
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```
Here’s the short version:

### Mutating Maps in Go

Insert/Update:
```go
m[key] = value
```
Get value:
```go
val := m[key]
```
Delete key:

```go
delete(m, key)
```

Check if key exists:

```go
val, ok := m[key]  // ok = true if key exists
```


# Closure
