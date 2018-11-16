package main

import "fmt"

func know() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println("not")

	panic("not-not-not")

	fmt.Println("know")
}

func main() {
	fmt.Println("Hello")
	know()
	fmt.Println("Bye")
}
