package interfacedemo

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/14 10:56
 *@version v1.0
 */
type Live interface {
	Run(run int)
	Eat()
}

type Person struct {
	Name string
}

func (p *Person) Run(run int) {
	fmt.Println(p.Name, "is running ", run, "meter")
}
func (p *Person) Eat() {
	fmt.Println(p.Name, "is eating")
}
