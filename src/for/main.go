package main

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/12 09:56
 *@version v1.0
 */
func main() {
	str := "helloworld"
	for i := 1; i < 10; i++ {
		println(i)
	}
	println("==============")
	for _, value := range str {
		//fmt.Print(index," ")
		fmt.Printf("%c\n", value)
	}
	println("冒泡排序：")
	arr := [...]int{8, 6, 1, 11, 20, 4, 3, 21}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Print(arr)

}
