package main

import "fmt"

func helloWorld(name string) {
	fmt.Println("Hello %s", name)
}

func main() {
	helloWorld("bob")
	return 0
}
