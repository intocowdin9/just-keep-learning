package main

import (
	"fmt"
	"just-keep-learning/go-packages-and-modules/util"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	greeting := fmt.Sprintf("Hello, %s", "Rafli")
	fmt.Println(greeting)

	fmt.Printf("length of greeting is %d\n", util.StringLength(greeting))
	fmt.Println(util.GetGreeting())

	r := mux.NewRouter()

	http.ListenAndServe(":9000", r)
}
