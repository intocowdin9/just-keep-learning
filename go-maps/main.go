package main

import "fmt"

func main() {

	ages := make(map[string]int)
	ages["Rafli"] = 26
	fmt.Printf("Rafli is %d years old\n", ages["Rafli"])

	ages["Muhammad"] += 1
	fmt.Printf("Muhammad is %d years old\n", ages["Muhammad"])

	// addr := &ages["Muhammad"] // cannot take address of map value

	gpas := map[string]float32{
		"Rafli":    3.4,
		"Muhammad": 3.9,
	}
	fmt.Printf("Rafli GPA is %.2f\n", gpas["Rafli"])
	fmt.Printf("Muhammad GPA is %.2f\n", gpas["Muhammad"])

	var visited map[string]bool
	visited = make(map[string]bool)
	visited["A"] = true
	fmt.Printf("A has been visited? %v\n ", visited["A"])

	fruits := []string{
		"bananas",
		"kiwis",
		"apples",
		"bananas",
	}
	fmt.Printf("duplicate fruits? %s\n", findDuplicates(fruits))
}

func findDuplicates(words []string) string {
	dupMap := make(map[string]bool)

	for i := 0; i < len(words); i++ {

		if !dupMap[words[i]] {
			dupMap[words[i]] = true
		} else {
			return words[i]
		}
	}

	return ("")
}
