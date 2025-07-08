package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func fizzBuzz(number int) string {
	var components []string
	if number%3 == 0 {
		components = append(components, "Fizz")
	}

	if number%5 == 0 {
		components = append(components, "Buzz")
	}

	if number%7 == 0 {
		components = append(components, "Bang")
	}

	if number%11 == 0 {
		components = []string{"Bong"}
	}

	if number%13 == 0 {
		index := 0
		for ; index < len(components) && components[index][0] != 'B'; index++ {
		}
		components = slices.Insert(components, index, "Fezz")
	}

	if number%17 == 0 {
		slices.Reverse(components)
	}

	var result string = strings.Join(components, "")

	if result == "" {
		return strconv.Itoa(number)
	} else {
		return result
	}
}

// This is our main function, this executes by default when we run the main package.
func main() {

	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
