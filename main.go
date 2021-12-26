package main

import "fmt"

func Cat() string {
	return "Meow~~~~"
}
func main() {
	saying := Cat()
	fmt.Println(saying)
}
