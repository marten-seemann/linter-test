package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func main() {
	fmt.Println("hello")
	fmt.Println(&user{name: "John Doe"})
}
