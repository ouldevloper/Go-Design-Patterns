package main

import (
	factories "factory/Factories"
	"fmt"
)

func main() {
	p := factories.CreatePerson()
	p.Id = 1
	p.Name = "hello Name"
	p.Address.City = "Peterrsburg"
	p.Address.Country = "Russia"
	p.Address.Home = 1
	fmt.Println("Person: ", p)
}
