package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fizzBuzz(number int) string {
	var messageBuilder strings.Builder
	if number%3 == 0 {
		messageBuilder.WriteString("Fizz")
	}

	if number%5 == 0 {
		messageBuilder.WriteString("Buzz")
	}

	if number%7 == 0 {
		messageBuilder.WriteString("Bang")
	}

	if number%11 == 0 {
		messageBuilder.Reset()
		messageBuilder.WriteString("Bong")
	}

	var result string = messageBuilder.String()

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
