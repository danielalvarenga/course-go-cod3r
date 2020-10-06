package main

// Import many packages in a block
// It's possible sssign alias for packages like "m" to "math"
import (
	"fmt"
	m "math"
)

func main() {
	// Create constants assigning a type
	const PI float64 = 3.1415

	// Long way to create a variable
	// If ommit the type the compiler infer a type
	// The type can't change
	var radius = 3.2

	// Short and usual way to create variables
	area := PI * m.Pow(radius, 2)
	fmt.Println("The circuference area is", area)

	// Assign many constants in a block
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// Assign many variables in a row assigning the type
	var e, f bool = true, false
	fmt.Println(e, f)

	// Assign many variables with different types in a row
	g, h, i := 2, true, "epa"
	fmt.Println(g, h, i)
}
