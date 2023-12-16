package main

import "fmt"

func main() {
	var num int
	fmt.Printf("enter a number: ")
	fmt.Scanln(&num)

	if num > 0 {
		println("positive number")
	} else if num < 0 {
		println("negative number")
	} else {
		println("you did not enter the number")
	}

	println("this will always run")
}
