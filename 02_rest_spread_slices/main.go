package main

import "fmt"

func main() {
	x := []int{4, 73, 6, 13, 54}
	fmt.Println(x)
	x = append(x, 77, 34, 100, 101)
	fmt.Println(x)

	y := []int{21, 52, 91, 10001}
	x = append(x, y...)
	fmt.Println(x)

	// remove stuff in the middle
	x = append(x[:3], x[9:]...)
	fmt.Println(x)

}
