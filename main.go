package main

import (
	"bufio"
	"fmt"
	"os"
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

func fizzBuzzUpTo(upperLimit int) string {
	result := ""
	for i := 1; i <= upperLimit; i++ {
		result += fizzBuzz(i) + "\n"
	}
	return result
}

func getMaxNumber() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

// This is our main function, this executes by default when we run the main package.
func main() {
	upperLimit := getMaxNumber()
	fmt.Println(fizzBuzzUpTo(upperLimit))
}
