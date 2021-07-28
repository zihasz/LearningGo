package main

import (
	"fmt"
	"strconv"
)

// Package level
var I int = 1000 // Exported global - all files
var i int = 100 // Package global - files in this package

func main() {
	/* Block scope */

	// Declarations
	var x int // Declare and assign
	x = 42

	var y int = 43 // Inline declaration

	z := 44 // Colon equals op

	// Variables cannot be unused
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Variable name lengths should reflect variable life.
	// `i` is used right away, short variable name.
	var i int = 1
	fmt.Printf("i %v \n", i)

	// `seasonNumber` is used a "lot" of times, manipulated, long variable name. 
	var seasonNumber int = 11
	seasonNumber += 2
	seasonNumber -= 1
	fmt.Printf("Season number %v \n", seasonNumber)

	// Acronyms should stay upper-case
	var theUrl string = "https://golang.org" // Bad!
	var theURL string = "https://golang.org" // Good
	fmt.Println(theUrl)
	fmt.Println(theURL)

	// Type conversions
	var a int
	a = 42
	fmt.Printf("Int %v, %T \n", a, a)
	
	var b float32
	b = float32(a)
	fmt.Printf("Float32 %v, %T \n", b, b)

	var c string
	c = string(a) // OR
	c = string(rune(a))
	fmt.Printf("String %v, %T \n", c, c)

	var d string
	d = strconv.Itoa(a)
	fmt.Printf("Strconv %v, %T \n", d, d)
}