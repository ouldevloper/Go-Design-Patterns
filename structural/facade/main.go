package main

import (
	"facade/facades"
	"fmt"
)

func main() {
	println("Hello facade design pattern ....")
	facades.Create(1, "abdullah")
	facades.Create(2, "oulahyane")
	fmt.Println("Person 2 : ", facades.GetById(2).Id)
	fmt.Println("All Persons  : ", facades.GetAll())
}
