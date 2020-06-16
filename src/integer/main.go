package main

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/13 11:10
 *@version v1.0
 */
func main() {
	//声明3个类型变量
	var a = 1
	var b int32 = 2
	var c int64 = 3
	fmt.Println(a, b, c)

	//把int32转换为int64
	a = int(b)
	fmt.Println(a, b)
	a = 1

	//把int64转换成int32
	b = int32(c)
	fmt.Println(b, c)
	b = 2

	//把int转换为int64
	c = int64(a)
	fmt.Println(a, c)
	c = 3
}
