package main

import (
	"fmt"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/14 08:46
 *@version v1.0
 */
func main() {

	func() {
		println("I am inner func")
	}()

	func(name string) {
		println(name)
	}("golang")

	name := func() string {
		return "return value"
	}()
	fmt.Println(name)
}
