package main

import "fmt"

func main() {
	a := [5]int{1, 3, 4}
	b := [...]int{1, 5, 6}
	fmt.Println(a, "length:", len(a), b, "length:", len(b))
	println()
}
