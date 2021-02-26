package main

import (
	"fmt"
)

const a = iota // 0

const (
	b = iota // 0
	c = iota // 1
	d = iota // 2
)

const (
	e = iota // 0
	f		 // 1
	g 		 // 2 
	h	     // 3
)

const (
	_ = iota + 5 // we can do arithmetic operation
	i 
	j
)



func main() {

	/*

	*		- Naming convention
	*		- Typed Constants
	*		- Untyped constants
	* 		- Enumerated constants
	*		- Enumerated expressions

	*/ 

	const aConst = 1
	const bConst int = 2
	fmt.Println(aConst)
	
}