package main

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/9/5 17:40
 *@version v1.0
 */
// 批量声明全局变量
// 1.先声明，在赋值
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "msr"
	age = 18
	isOk = true
	fmt.Printf("name:%s,age:%d,isOk:%t", name, age, isOk)
}
