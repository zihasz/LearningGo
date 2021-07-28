package main

import (
	"fmt"
)

func main() {
	
	// Boolean - false by default

	var boolean bool = true
	fmt.Printf("%v, %T \n", boolean, boolean)

	// Logical operations return bools
	boolean2 := 1 == 1
	fmt.Printf("%v, %T \n", boolean2, boolean2)
	boolean3 := 1 == 2
	fmt.Printf("%v, %T \n", boolean3, boolean3)


	// Integers 

	// Defaults to int
	n := 42
	fmt.Printf("%v, %T \n", n, n)

	// Sizes
	var i int = 42
	fmt.Printf("%v, %T \n", i, i)

	var ui uint = 42					// All integer types have unsigned versions
	fmt.Printf("%v, %T \n", ui, ui)

	var i8 int8 = 42
	fmt.Printf("%v, %T \n", i8, i8)

	var i16 int16 = 42
	fmt.Printf("%v, %T \n", i16, i16)

	var i32 int32 = 42
	fmt.Printf("%v, %T \n", i32, i32)

	var i64 int64 = 42
	fmt.Printf("%v, %T \n", i64, i64)

	a := 10	// 1010
	b := 3	// 0011

	// Arithmetics
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // Will return 3
	fmt.Println(a % b)

	// Conversions
	var integ int = 10
	var short int8 = 3
	fmt.Println(integ + int(short)) // Need to conver yourself
	
	// Binary ops
	fmt.Println(a & b) 	// AND - 0010
	fmt.Println(a | b) 	// OR  - 1011
	fmt.Println(a ^ b) 	// XOR - 1001
	fmt.Println(a &^ b)	// ??? - 0100

	// Bit shifting
	c := 8 // 2^3
	fmt.Println(c << 3)	// 2^3 * 2^3 = 2^6
	fmt.Println(c >> 3)	// 2^3 / 2^3 = 2^0

	// Floating point

	f := 3.14 // float64 !!
	f = 13.7e72
	f = 2.1E14
	fmt.Printf("%v, %T \n", f, f)

	af := 3.14
	bf := 7.12

	fmt.Println(af + bf)
	fmt.Println(af - bf)
	fmt.Println(af * bf)
	fmt.Println(af / bf)
	// no mod op

	// no binary ops

	// no bit shifting

	// Complex numbers

	var ca complex64 = 1 + 2i
	fmt.Printf("%v, %T \n", ca, ca)
	fmt.Printf("%v, %T \n", real(ca), real(ca))
	fmt.Printf("%v, %T \n", imag(ca), imag(ca))

	var cb complex128 = 1 + 2i
	fmt.Printf("%v, %T \n", cb, cb)
	fmt.Printf("%v, %T \n", real(cb), real(cb))
	fmt.Printf("%v, %T \n", imag(cb), imag(cb))

	var cc complex128 = complex(128, 128)
	fmt.Printf("%v, %T \n", cc, cc)

	// Strings

	s := "Hello, I'm a string!"
	fmt.Printf("%v, %T \n", s, s)
	fmt.Printf("%v, %T \n", s[2], s[2])
	fmt.Printf("%v, %T \n", string(s[2]), string(s[2]))

	s2 := " I'm a string too!";
	fmt.Printf("%v, %T \n", s + s2, s + s2)

	by := []byte(s);
	fmt.Printf("%v, %T \n", by, by);

	// Runes - chars

	r := 'a'
	var re rune = 'r'
	fmt.Printf("%v, %T \n", r, r);		// runes ARE int32's
	fmt.Printf("%v, %T \n", re, re);	// runes ARE int32's

}