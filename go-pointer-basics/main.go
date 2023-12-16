package main

import (
	"fmt"
)

type Person struct {
	Age  int
	Name string
}

func (p *Person) Birthday() {
	p.Age++
}

// func addOne(x *int) {
// 	fmt.Println(x)
// 	*x++
// }

// func getPtr() *int {
// 	x := 11
// 	var p_x *int
// 	p_x = &x
// 	fmt.Println(p_x)
// 	return p_x
// }

func main() {
	rafli := Person{Age: 11, Name: "Rafli"}
	rafli.Birthday()
	fmt.Printf("happy birthday %s! u r now %d\n", rafli.Name, rafli.Age)
}
