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
