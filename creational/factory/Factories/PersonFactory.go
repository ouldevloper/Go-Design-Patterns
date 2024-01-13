package factories

import classes "factory/Classes"

func CreatePerson() *classes.Person {
	return &classes.Person{
		Id:      0,
		Name:    "",
		Address: &classes.Address{},
		Cnss:    &classes.SocialSecurity{},
	}
}
