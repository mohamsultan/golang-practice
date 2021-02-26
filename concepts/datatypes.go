package main

import (
	"fmt"
)

func main() {

	fmt.Println("Datatypes")
	
//  Boolean datatype
	var a bool // false by default
	var nT bool = true
	var nF bool = false
	fmt.Println(nT, nF, a)

//  Integers int8, int16, int32 int64
	var n uint = 32 
	fmt.Printf("%v %T\n", n, n)

//	var i int = 10
//	var j int8 = 3
//	fmt.Println(i+j) arithmetic operation of different type is not allowed 

	x := 10		// 1010
	y := 3 		// 0011

	fmt.Println(x & y) 	// 0010
	fmt.Println(x | y)	// 1011
	fmt.Println(x ^ y)	// 1001
	fmt.Println(x &^ y)	// 0100

	x1 := 8		//  2^3
	fmt.Println(x1 << 3)	// 2^3 * 2^3 = 2^6
	fmt.Println(x1 >> 3)	// 2^3 / 2^3 = 2^0

//	Floating point numbers
	var m float32 = 3.14
	// m = 13.7e72 -> float64
	m = 2.1E14
	fmt.Printf("%v %T\n", m, m) 

//	Complex numbers
	var comp complex64 = 1+2i // other way is complex(1, 2)
	fmt.Printf("%v %T\n", comp, comp)
	fmt.Printf("%v %T\n", real(comp), real(comp))
	fmt.Printf("%v %T\n", imag(comp), imag(comp))

//	String
	s := "this is a string"
//	s[3] = "u"  -> not allowed
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", string(s[2]), s[2])
	by := []byte(s)
	fmt.Printf("%v %T\n", by, by)

// 	Rune is an integer-32
	var r rune = 'a'
	fmt.Printf("%v %T\n", r, r)

}