package main

import (
	"encoding/json"
	"fmt"
)

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) getArea() int {
	return r.Length * r.Width
}

func (r *Rectangle) setLength(len int) {
	r.Length = len
}

type Person struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	// ...
}

type Employee struct {
	ID     uint64
	Person Person
	Salary uint
	// ...
}

func main() {

	var rect1 Rectangle
	rect1.Length = 5
	rect1.Width = 4

	fmt.Printf("area of rect1: %d\n", rect1.getArea())
	rect1.setLength(7)
	fmt.Printf("new length? %d\n", rect1.Length)
	fmt.Printf("area of rect1: %d\n", rect1.getArea())

	rect2 := Rectangle{}
	rect2.Length = 6
	rect2.Width = 9

	rect3 := Rectangle{5, 4}
	fmt.Printf("rect3 length: %v width: %v\n", rect3.Length, rect3.Width)

	rect4 := Rectangle{
		Width:  5,
		Length: 4,
	}
	fmt.Printf("rect4 length: %v width: %v\n", rect4.Length, rect4.Width)

	var rect5 = new(Rectangle)
	rect5.Length = 5

	var rect6 = &Rectangle{}
	rect6.Width = 8

	rafli := Employee{
		ID:     87683,
		Salary: 430000,
		Person: Person{
			Name: "Rafli",
			Age:  27,
		},
	}

	muhammad := Employee{
		ID:     9872834,
		Salary: 430000,
		Person: Person{
			Name: "Muhammad",
			Age:  26,
		},
	}

	data, err := json.Marshal(muhammad)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}
	// fmt.Printf("Muhammad: %v\n", data) // the data is a binary slice
	fmt.Printf("Muhammad: %s\n", data)

	fmt.Printf("Rafli: %v\n", rafli)
}
