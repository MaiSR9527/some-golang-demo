package main

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/12 10:40
 *@version v1.0
 */
func main() {
	slice := []string{"hello", "world"}
	s2 := slice
	fmt.Println(slice, s2)
	fmt.Printf("%T\n", slice)
	fmt.Printf("%p %p\n", slice, s2)
	s3 := make([]int, 0)
	fmt.Printf("%T\n", s3)
	fmt.Println("========append======")
	fmt.Println(len(s3), cap(s3))
	s3 = append(s3, 1, 2, 3, 4, 5)
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3))

}
