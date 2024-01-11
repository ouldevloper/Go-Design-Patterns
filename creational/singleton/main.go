package main

import "fmt"

type Object struct {
	Id         int
	Name       string
	Properties []interface{}
}

var (
	obj *Object = nil
)

func GetObject() *Object {
	if obj == nil {
		fmt.Println("New Instance")
		obj = new(Object)
	} else {
		fmt.Println("Already Initialized")
	}
	return obj
}

func main() {
	var oo = GetObject()
	var bb = GetObject()
	var cc = GetObject()

	fmt.Println("Result: ", oo, bb, cc)

}
