package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type CustomeError struct {
	Message string
	Code    int
}

func (c CustomeError) Error() string {
	return c.Message + " " + strconv.Itoa(c.Code)
}

func Divide(x, y float64) (float64, error) {
	if y == 0 {
		// return float64(0), errors.New("cannot devide by zero")
		return 0.0, CustomeError{"cannot device by zero", -1}
	}

	return x / y, nil
}

func main() {

	file, fileErr := ioutil.ReadFile("file.txt")
	if fileErr != nil {
		fmt.Println("error reading file: ", fileErr)
	} else {
		fmt.Println(string(file))
	}

	value, divErr := Divide(7, 0)
	if divErr != nil {
		fmt.Println(divErr)
		// os.Exit(-1)
		log.Fatal(divErr)
	} else {
		fmt.Printf("%.2f\n", value)
	}

	// for ignore error we can set _ beside var of value
	/*
		for eg. value, _ := Divide(7,0)
		fmt.Println("%.2f\n", value)
	*/
}
