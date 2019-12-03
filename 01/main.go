package main

import "fmt"

func main() {
	hello()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func hello() {
	fmt.Println("Hello World")
}
