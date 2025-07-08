package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type RuleSet struct {
	effectMap map[int]func([]string) []string
	ordering  []int
}

var defaultRules = RuleSet{
	effectMap: map[int]func([]string) []string{
		3: func(components []string) []string {
			return append(components, "Fizz")
		},
		5: func(components []string) []string {
			return append(components, "Buzz")
		},
		7: func(components []string) []string {
			return append(components, "Bang")
		},
		11: func(components []string) []string {
			return []string{"Bong"}
		},
		13: func(components []string) []string {
			index := 0
			for ; index < len(components) && components[index][0] != 'B'; index++ {
			}
			return slices.Insert(components, index, "Fezz")
		},
		17: func(components []string) []string {
			slices.Reverse(components)
			return components
		},
	},
	ordering: []int{3, 5, 7, 11, 13, 17},
}

func fizzBuzz(number int, rules RuleSet) string {
	var components []string
	for _, trigger := range rules.ordering {
		effect := rules.effectMap[trigger]
		fmt.Println(rules.ordering)
		if number%trigger == 0 {
			components = effect(components)
		}
	}

	var result string = strings.Join(components, "")

	if result == "" {
		return strconv.Itoa(number)
	} else {
		return result
	}
}

func fizzBuzzUpTo(upperLimit int, rules RuleSet) string {
	result := ""
	for i := 1; i <= upperLimit; i++ {
		result += fizzBuzz(i, rules) + "\n"
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

func getRules() RuleSet {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Custom Rules Y/N? >> ")
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "y" {
		return defaultRules
	}
	return defaultRules
}

// This is our main function, this executes by default when we run the main package.
func main() {
	rules := getRules()
	upperLimit := getMaxNumber()
	fmt.Println(fizzBuzzUpTo(upperLimit, rules))
}
