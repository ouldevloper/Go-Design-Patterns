package classes

type Person struct {
	Id   int
	Name string
}

var (
	persons = make([]Person, 0)
)

func (obj *Person) Create(id int, name string) {
	o := Person{
		Id:   id,
		Name: name,
	}
	persons = append(persons, o)
}

func (obj *Person) GetById(id int) *Person {
	for _, o := range persons {
		if o.Id == id {
			return &o
		}
	}
	return nil
}

func (obj *Person) GetAll() []Person {
	return persons
}
