package main

import "fmt"

/**
 *
 *@author MaiShuRen
 *@date 2020/6/14 09:53
 *@version v1.0
 */
func main() {
	var person Person
	person.Age = 22
	person.Name = "maishuren"
	fmt.Println(person)
	p2 := Person{Name: "msr", Age: 22}
	fmt.Println(p2)
	fmt.Printf("%p %p \n", &person, &p2)
	fmt.Println("========================")
	newPerson1 := NewPerson("msr", 22)
	newPerson2 := NewPerson("msr", 22)
	np := newPerson1
	fmt.Printf("%p %p \n", newPerson1, newPerson2)
	fmt.Println(newPerson1 == np)

}

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}
