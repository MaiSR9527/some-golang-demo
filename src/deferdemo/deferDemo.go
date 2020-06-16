package main

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/14 11:50
 *@version v1.0
 */
func main() {
	//defer是在所有代码的最后执行
	//多个defer以栈的形式存储
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	defer fmt.Println(4)
}
