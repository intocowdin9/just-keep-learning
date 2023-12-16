package main

import "fmt"

type printer interface {
	printInfo()
}

type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

type Toy struct {
	Name  string
	Price float32
}

func (d Drink) printInfo() {
	fmt.Printf("Drink: %s Price: %.2f", d.Name, d.Price)
}

func (b Book) printInfo() {
	fmt.Printf("Title: %s Price: %.2f", b.Title, b.Price)
}

func (t Toy) printInfo() {
	fmt.Printf("Toy: %s Price: %.2f", t.Name, t.Price)
}

func empty(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("i'm a string: %s\n", i)
	case int:
		fmt.Printf("i'm an int: %d\n", i)
	case Book:
		fmt.Printf("i'm a book: %s\n", i.(Book).Title)
	}
}

func main() {

	gunslinger := Book{
		Title: "the gunstlinger",
		Price: 4.75,
	}

	coffee := Drink{
		Name:  "coffee",
		Price: 2.00,
	}

	samurai := Toy{
		Name:  "katanux",
		Price: 13.00,
	}

	gunslinger.printInfo()
	coffee.printInfo()
	samurai.printInfo()

	info := []printer{gunslinger, coffee}

	info[0].printInfo()
	info[1].printInfo()

	empty("Rafli")
	empty(26)
}
