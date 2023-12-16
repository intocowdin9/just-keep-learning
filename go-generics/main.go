package main

import "fmt"

// import "constraints"

type NumberType interface {
	int | float32 | float64 | uint8 | uint16 | uint32 | uint64
}

// we can use NumberType interface or constraints.Ordered
func addList[T NumberType](list []T) T {
	var sum T
	for _, item := range list {
		sum += item
	}
	return sum
}

func addInts(list []int) int {
	var sum int
	for _, item := range list {
		sum += item
	}
	return sum
}

func addFloats(list []float32) float32 {
	var sum float32
	for _, item := range list {
		sum += item
	}
	return sum
}

func printMe(thing ...any) {
	fmt.Println(thing...)
}

func main() {
	fmt.Printf("sum of insts: %d\n", addInts([]int{1, 2, 3, 4, 5}))
	fmt.Printf("sum of floats: %.2f\n", addFloats([]float32{1, 2, 3, 3, 5}))

	fmt.Printf("sum of ints: %d\n", addList([]int{1, 2, 3, 5, 6}))
	fmt.Printf("sum of floats: %.2f\n", addFloats([]float32{2, 4, 5, 6, 3, 1}))

	printMe("rafli", 2, 33)
}
