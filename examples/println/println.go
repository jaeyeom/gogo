// Binary println prints "Hello, playground 10 little Gophers".
package main

import "fmt"

func main() {
	var a = 10
	b := "little"
	b += " Gophers"
	fmt.Println("Hello, playground", a, b)
}
