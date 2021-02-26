package main

import (
	"fmt"
	"reflect"
)
// Animal is struct type
type Animal struct { // tags
	Name string	`required max:"100"`
	Origin string
}

// Bird is embedding Animal struct, its called composition
// Bird is struct type 
type Bird struct {
	Animal
	Speed float32
	CanFly bool
}

// Doctor is struct type
type Doctor struct {
	number int
	actorName string
	companions []string
}

func main() {

//	Anonymous struct
//	aDoctor := struct{name struct}{name: "sultan"}	

	aDoctor := Doctor{
		number: 3,
		actorName: "Sultan",
		companions: []string {
			"sultan",
			"mohamed",
		}, 
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.number)
	fmt.Println(aDoctor.companions[0])

	// b := Bird{
	// 	Animal: Animal {
	// 		Name: "Emu",
	// 		Origin: "Aus",
	// 	},
	// 	Speed: 48,
	// 	CanFly: false,
	// }

	b := Bird{}
	b.Name = "Emu"
	b.Origin ="Aus"
	b.Speed = 48
	b.CanFly =false

	fmt.Println(b)
	fmt.Println(b.Name)
	fmt.Println(b.CanFly)

// 	struct tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

/* 

	Structs
	- collections of disparate data types that describe a single concept
	- keyed by named fields
	- normally created as types, but anonymous structs are allowed
	- structs are value types
	- no inheritance, but can use composition via embedding
	- tags can be added to struct fields to describe field

*/