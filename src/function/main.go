package main

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/14 08:25
 *@version v1.0
 */
func main() {
	a, b := 1, 2
	sum := add(a, b)
	fmt.Println(sum)
	s, i, s2 := show()
	fmt.Println(s, i, s2)
	all := addAll(1, 2, 3, 4)
	fmt.Println(all)

}

func add(a, b int) int {
	return a + b
}

func add2(a, b int) (s int) {
	s = a + b
	return
}
func show() (string, int, string) {
	return "a", 1, "b"
}

func addAll(n ...int) int {
	sum := 0
	for _, j := range n {
		j += j
		sum = j
	}
	return sum
}
