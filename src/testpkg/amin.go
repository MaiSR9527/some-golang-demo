package main

import (
	"demo"
	"fmt"
	"interfacedemo"
)

/**
 *
 *@author MaiShuRen
 *@date 2020/6/14 09:23
 *@version v1.0
 */
func main() {
	demo.Demo1()
	fmt.Println("==========================")
	var person interfacedemo.Live = &interfacedemo.Person{"msr"}
	person.Run(5)

}
