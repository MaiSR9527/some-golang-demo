package main

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/14 09:43
 *@version v1.0
 */
func main() {
	i := 1
	j := 2
	arr := []int{3, 4}
	demo(i, j, arr)
	fmt.Println(i, j)
	fmt.Println(arr)
	fmt.Println("===================")
	demo2(&i, &j, arr)
	fmt.Println(i, j)
	fmt.Println(arr)
}

func demo(i int, j int, arr []int) {
	i = 111
	j = 222
	arr[0] = 333
}
func demo2(i *int, j *int, arr []int) {
	*i = 111
	*j = 222
	arr[0] = 333
}
