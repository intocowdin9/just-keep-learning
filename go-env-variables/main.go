package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("HOME: ", os.Getenv("HOME"))

	shell, ok := os.LookupEnv("SHELL")
	if !ok {
		fmt.Printf("the env var SHELL is not set")
	}

	fmt.Println("SHELL: ", shell)

	err := os.Setenv("NAME", "Rafli")
	if err != nil {
		fmt.Println("could not set the env var NAME")
	}

	fmt.Printf("NAME : %s\n", os.Getenv("NAME"))

	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("could not load .env file")
		os.Exit(1)
	}

	// fmt.Printf("API_KEY: %s\n", os.Getenv("API_KEY"))
	// fmt.Printf("DB_PASS: %s\n", os.Getenv("DB_PASS"))

	envMap, mapErr := godotenv.Read(".env")
	if mapErr != nil {
		fmt.Printf("error loading .env into map[string]string\n")
		os.Exit(1)
	}

	fmt.Printf("API_KEY: %s\n", envMap["API_KEY"])
	fmt.Printf("DB_PASS: %s\n", envMap["DB_PASS"])
}
