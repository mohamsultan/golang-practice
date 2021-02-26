package main

import (
	"fmt"   
    "strconv"
)

/*
 *      Naming of variable
 *      Any variable starts with lower case has scope within the package. Those variables cannot be used outside the package.
 *      Any variable starts with upper case has scope that can be used outside the package.  
 */

//      Scope - used within this package, but not outside this package.
var scope int = 10    

//      Scope - used within and outside of this pacakge.
var Scope int = 10     


//      Package level declaration
//      Have to use full declaration syntax (ie. cannot use := syntax)

var a int = 100;

//      Variable block for cleaner code

var (
    name string = "sultan"
    age int = 21
    gender string = "male"
)

func main() {
/*
 *       Syntax-1  
 *       Use: Places where you not want to assign value initially.
 */

    var i int 
    i = 42
	fmt.Println(i)

/*
 *       Syntax-2  
 *       Use: Where you want to specify datatype with value initially.
 */

    var j int = 30
    fmt.Println(j)

/*
 *       Syntax-3  
 *       Declare variable that can assign the datatype dynamically based on the value assigned
 */

    k := 10
    fmt.Println(k)

//      formatting string
//      { %v -> value }, { %T -> Type of the value }

    fmt.Printf("%v %T", a, a)

/*
 *       Variable type conversion
 */

    var z int = 200
    var b float32
    b = float32(z)      // conversion function  
    fmt.Print(b, z)

    var num int = 40
    var str string

    str = string(num)
    fmt.Printf("%v %T", str, str) // return unicode of value 40 instead of 40 as string type

//      can't do implicit conversion int->string
//      To avoid this, import strconv to do the int->string conversion

    str = strconv.Itoa(num)
    fmt.Printf("%v %T", str, str)


}