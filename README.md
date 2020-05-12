# go-programs

 

# Basic Hello World program

```
// Print a friendly greeting

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World Gophers ☺")
}

```


# Finding the mean of 2 numbers 

We use `%v` to print the value, and `%T` to print the data type.
```
// Calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	x = 1
	y = 2

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	var mean int
	mean := (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}

```

# Finding the mean of 2 numbers 

Use `:=` to directly assign the value, without explicitly assigning data type to it.

```

// Calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {
	x  := 1.0
	y  := 2.0

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
```

# if-else in golang

*Note* :

Syntax for if-else is

``` 
if something {

} else {

}

// else should be on that same line 
```

```
// Example of "if" statement
package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x < 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}

```

# switch in golang


```

// Example of "switch" statement

package main

import (
	"fmt"
)

func main() {
	x := 2

	// using switch with argument
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Printf("many")
	}

	// Using switch with conditions
	switch {
	case x > 100:
		fmt.Println("x is very big")
	case x > 10:
		fmt.Println("x is big")
	default:
		fmt.Println("x is small")
	}

}

```

# for and while loop

```

// "for" loop examples
package main

import (
	"fmt"
)

func main() {

	// Simple for loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("----")

	// for loop with break
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("----")

	// for loop with continue
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----")

	// While loop
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Println("----")

	// for loop without any arguments
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
}

```

# Strings

- `%v` prints the value in a default format

- <a href="https://golang.org/pkg/fmt/"> fmt verbs </a>

```

// Go strings
package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))
	fmt.Println(book[0])                                                       //84
	fmt.Printf("book[0] = %q (type %T)\n", book[0], book[0])                   //84 (type unint8)
	// uint8 is a byte

	// Strings in go are immutable
	// book[0] = 116

	// Slice (start, end), 0 based, ½ empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// Slice (no start)
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])

	// Unicode
	fmt.Println("It was ½ price!")

	// Multi line
	poem := `
	The road goes ever on
	Down from the door where it began
	...
	`
	fmt.Println(poem)
}

```


# Slices (similar to arrays)

```

// Go slices
package main

import (
	"fmt"
)

func main() {
	// Same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// Length
	fmt.Println(len(loons)) // 3

	fmt.Println("----")
	
	// 0 indexing
	fmt.Println(loons[1]) // daffy

	fmt.Println("----")
	
	// slices
	fmt.Println(loons[1:]) // [daffy taz]

	fmt.Println("----")
	
	// Simple for loop for slice
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("----")
	
	// Single value range - prints the indices
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("----")

	// Double value range - prints the value and indices
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("----")

	// Double value range, ignore index by using _
	// You cannot declare i and not use it, Golang will throw an error
	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("----")
	// append
	loons = append(loons, "elmer")
	fmt.Println(loons) // [bugs daffy taz elmer]
}

```

# Maps

```

// Go's map data structure
package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61, // Must have trailing comma in multi line
	}

	// Number of items
	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["MSFT"])

	// Get zero value if not found
	fmt.Println(stocks["TSLA"]) // 0

	// Use two value to see if found
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	// Set
	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	// Delete
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Single value "for" is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	// Double value "for" is key, value
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}

```

# Functions

You can return more than one value in golang.

**Simple function**

```
// Basic function definition
package main

import (
	"fmt"
)

// add adds a to b
func add(a int, b int) int {
	return a + b
}

// divmod returns quotient and reminder
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)
}


```

**Argument calls - by value, by reference**

```
package main

import (
	"fmt"
)

// passing slice as a parameter
// slices are by default passed as a reference inside a function

func doubleAt(values []int, i int) {
	values[i] *= 2
}

// call by value
func double(n int) {
	n *= 2
}

// call by reference
func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	double(val)
	fmt.Println(val)
	
	// Passing the pointer
	doublePtr(&val)
	fmt.Println(val)
}

```

# Errors 

- `error` is a built-in type that is used throughout Go.
- A function that can produce some error, passes it as the last parameter.
- We use `fmt.Errorf` to create a new error.
- `nil` is the value that Go uses to signal nothing.


```
package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%f)", n)
	}

	return math.Sqrt(n), nil
}

func main() {
	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s1)
	}

	s2, err := sqrt(-2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s2)
	}
}


```

# Defer

- A defer statement defers the execution of a function until the surrounding function returns.

- The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

- This is specifically used for garbage cleanup purposes.

```
package main

import "fmt"

func see(){
 fmt.Println("EFG")
}
func main() {

	fmt.Println("ABCD")

	defer see()

	fmt.Println("HIJK")
	
	fmt.Println("LMNOP")

}


```

```
package main

import (
	"fmt"
)

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	defer cleanup("A")
	defer cleanup("B")

	fmt.Println("worker")
}

func main() {
	worker()
}

```

# Struct

```
// struct demo
package main

import (
	"fmt"
)

// Trade is a trade in stocks
type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade
}

func main() {

	// Create an object from struct - Method 1
	// You need to follow the order, if not specifying the fields explicitly
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)
	
	// For a better visual representation of the struct
	fmt.Printf("%+v\n", t1)

	fmt.Println(t1.Symbol)

	// Create an object from struct - Method 2
	// if any field is not specified, Go assigns 0 to it
	t2 := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
}

```

**Functions of struct**

```
// Method demo
package main

import (
	"fmt"
)

// Trade is a trade in stocks
type Trade struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // true if buy trade, false if sell trade
}

// Value returns the trade value
// t is k/a the receiver pointer
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	t := Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Println(t.Value())
}
```

**Interfaces**

```

package main

import (
	"fmt"
	"math"
)

// Square is a square
type Square struct {
	Length float64
}

// Area returns the area of the square
func (s *Square) Area() float64 {
	return s.Length * s.Length
}

// Circle is a circle
type Circle struct {
	Radius float64
}

// Area returns the curcle of the square
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// sumAreas return the sum of all areas in the slice
func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

// Shape is a shape interface
type Shape interface {
	Area() float64
}

func main() {
	s := &Square{20}
	fmt.Println(s.Area())

	c := &Circle{10}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Println(sa)
}

```

