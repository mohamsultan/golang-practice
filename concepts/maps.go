package main

import (
	"fmt"
)

func main()  {
	
	fmt.Println("Maps!!")
//	statePopulations := make(map[string]int)
	statePopulations := map[string]int {	// unordered
		"california": 123,
		"texas": 789,
		"florida": 543,
		"new york": 197,
	}

	statePopulations["georgia"] = 82891
	delete( statePopulations, "georgia" )

	pop1, ok1 := statePopulations["ohio"]
	fmt.Println(pop1, ok1)

	pop2, ok2 := statePopulations["texas"]
	fmt.Println(pop2, ok2)

	fmt.Println( statePopulations )
}

/* 

	Maps
	- collections of value types that are accessed via keys
	- created via literal or via make function
	- members accessed via [key] syntax
		- myMap["key"] = "value"
	- check for presence with "value, ok" form of result 
	- multiple assignments refer to same underlying data

*/