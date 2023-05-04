package main

import (
	"fmt"

	"github.com/hookenz/mytest/pkg/libhello"
)

type (
	Embedded struct {
		Name  string
		Value string
	}

	MyStruct struct {
		Embedded
	}
)

func main() {
	m := MyStruct{}
	m.Name = "Bob"
	fmt.Println("Hello", m.Name)
	libhello.Hello()
}
