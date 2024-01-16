package main

import (
	"adapter/adapters"
	"fmt"
)

func main() {
	fmt.Println("Hello Adapter pattern")
	resp := adapters.ViewAdapter("sdfgasfg", 200)
	fmt.Println("response: ", resp)
}
