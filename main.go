package main

import "fmt"

func helloWorld(n string) {
	fmt.Println("Hello %d and Hello world", n)
}

func main() {
	helloWorld("John")
	return 0
}
