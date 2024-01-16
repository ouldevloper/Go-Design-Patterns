package facades

import (
	"facade/classes"
)

var (
	p = new(classes.Person)
)

func Create(id int, name string) {
	p.Create(id, name)
}

func GetById(id int) *classes.Person {
	return p.GetById(id)
}

func GetAll() []classes.Person {
	return p.GetAll()
}
