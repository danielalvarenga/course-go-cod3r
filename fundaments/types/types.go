package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// 'int' have signal (positive or negative)
	fmt.Println("Literal integer is", reflect.TypeOf(32000))
	fmt.Println("Literal integer is", reflect.TypeOf(-32000))

	// Integer with signal: int8, int16, int32, int64
	// For 32bit systems default for int is 32
	// For 64bit systems default for int is 64
	i1 := math.MaxInt64
	fmt.Println("The maximum value to int is", i1) // 9223372036854775807
	fmt.Println("The type for the maximum value of int is", reflect.TypeOf(i1))

	// 'rune' is a alias for int32 and
	// represents a character in table Unicode
	var i2 rune = 'a'
	fmt.Println(i2)
	fmt.Println("The type 'rune' is a", reflect.TypeOf(i2))

	// Only positive integers: uint8, uint16, uint32, uint64
	// 'byte' is a alias for uint8
	var b byte = 255
	fmt.Println("The type 'byte' is a", reflect.TypeOf(b))

	// Real numbers: float32, float64
	var x float32 = 49.99
	fmt.Println("The type of x is", reflect.TypeOf(x))
	fmt.Println("The literal type for float is", reflect.TypeOf(49.90)) // float64 for 64bit system

	// Boolean: bool
	// Don't use 1 or 0 (integers)
	bo := true
	fmt.Println("The type of bo is", reflect.TypeOf(bo)) // boll
	fmt.Println(!bo)

	// String: string
	s1 := "My name is Bob"
	fmt.Println(s1 + "!")
	fmt.Println("The type of s1 is", reflect.TypeOf(s1))
	fmt.Println("The string size is", len(s1))

	// String with multiple lines use `
	s2 := `My
	name
	is
	Bob`
	fmt.Println(s2)
	fmt.Println("The size of the string is", len(s2))
}
