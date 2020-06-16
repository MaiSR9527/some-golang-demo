package main

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/13 11:13
 *@version v1.0
 */
func main() {
	//定义数字
	var i rune = 0x5F20
	fmt.Println(i)

	//输出汉字张
	fmt.Printf("%c\n", i)

	//获取转换后的内容
	c := fmt.Sprintf("%c", i)
	fmt.Println(c)
}
