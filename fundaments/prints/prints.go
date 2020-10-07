package main

import "fmt"

func main() {
	// Print without jump line with 'Print
	fmt.Print("Print without")
	fmt.Print(" jump line")

	// Print jumping line with 'Println'
	fmt.Println("Print jumping line")
	fmt.Println("Next line")

	x := 3.141516

	// To print non string variable directly you need convert
	xs := fmt.Sprint(x)
	fmt.Println("The value of x is" + xs)
	fmt.Println("The value of x is", xs)

	// To print directly you need to use symbols with 'Printf'
	fmt.Printf("The rounded value of x is %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "string"

	// Symbols by type
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// Symbol for any type (infered)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
