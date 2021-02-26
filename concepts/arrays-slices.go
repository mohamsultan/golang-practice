package main

import (
    "fmt"
)

func main() {
    
    fmt.Println("Arrays and Slices")

    var students [10] string // empty array
    grades := [3]int {97, 85, 93} // [...]int{} another syntax

    fmt.Printf("Grades: %v\n", grades)
    students[0] = "sultan"
    fmt.Printf("Students %v %v\n", students, len(students))

    var matrix [3][3]int = [3][3]int { [3]int {1, 2, 3}, [3]int {1, 2, 3}, [3]int {1, 2, 3} }
    fmt.Println(matrix)

// 	Slices	
// 	m := make([]int, 3, 100) -> (slice-type, length, capacity)
//	m = append(m, 1, 2, 3,  .....)
//  m = append(m, []int{1, 2, 3}...)
    a := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // it works for array
    b := a[:]
    c := a[3:]
    d := a[:6]
    e := a[3:6]
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
}

/*

*	Arrays
*	- collections of items with same type
*	- fixed size
*	- declaration styles
*		- a := [3]int{1,2,3}
*		- a := [...]int{1,2,3}
*		- var a [3]int


*	Slices
*	- Backed by array
*	- creation styles
*		- slice existing array or slice
*		- literal style
*		- via make function
*		- a := make(int[], 10) // create slice with capacity and length == 10
*		- a := make(int[], 10, 100)  // slice with length == 10 and capacity == 100
*   - len function return length of slice
*   - cap function returns length of underlying array
*   - append function to add elements to slice
*       - may cause expensive copy operation if underlying array is too small

*/








