package main

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/14 09:32
 *@version v1.0
 */
func main() {
	f := closure()
	f2 := closure()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Printf("%p %p", f, f2)
}
func closure() func() int {
	i := 1
	return func() int {
		i += 1
		return i
	}
}
