package main

import (
	"fmt"
)

func Test() string {
	return "some string..."
}

func another() int {
	return 42
}

func main() {
	a := Test()
	fmt.Println(&a)
	// fmt.Println(&Test())

	b := another()
	fmt.Println(&b)
	// fmt.Println(&another())
}
