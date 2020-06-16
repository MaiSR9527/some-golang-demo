package main

import (
	"fmt"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/14 08:53
 *@version v1.0
 */
func main() {
	var a func()
	a = func() {
		fmt.Println("a")
	}
	a()
	fmt.Println("=====================")
	mydo(func(name string) {
		fmt.Println(name)
	})
	fmt.Println("=====================")
	result := outer()
	r := result()
	fmt.Println(r)
}

func mydo(arg func(name string)) {
	fmt.Println("working ...")
	arg("mydo func")
}

func outer() func() int {
	return func() int {
		return 1
	}
}
