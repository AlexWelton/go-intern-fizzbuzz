package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type RuleSet struct {
	effectMap map[int]func([]string) []string
	ordering  []int
}

type RuleType int

const (
	AppendRule RuleType = iota
	ReplaceRule
	RegexRule
)

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
	fmt.Println(rules.ordering)
	for _, trigger := range rules.ordering {
		effect := rules.effectMap[trigger]

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
	fmt.Print("Max number >> ")
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func getRuleType(input string) RuleType {
	switch strings.ToLower(input) {
	case "append":
		return AppendRule
	case "replace":
		return ReplaceRule
	case "regex":
		return RegexRule
	default:
		panic(fmt.Sprintf("Invalid rule type: %s", input))
	}
}

func addRule(trigger int, ruleType RuleType, params []string, rules *RuleSet) {
	rules.ordering = append(rules.ordering, trigger)
	switch ruleType {
	case AppendRule:
		rules.effectMap[trigger] = func(components []string) []string {
			return append(components, params[0])
		}
	case ReplaceRule:
		rules.effectMap[trigger] = func(components []string) []string {
			return []string{params[0]}
		}
	case RegexRule:
		regex, _ := regexp.Compile(params[0])
		insert := params[1]
		rules.effectMap[trigger] = func(components []string) []string {
			for _, component := range components {
				indices := regex.FindStringIndex(component)
				if indices != nil {
					component = component[:indices[0]] + insert + component[indices[0]:]
				}
			}
			return components
		}
	}
}

func getCustomRules() RuleSet {
	rules := RuleSet{
		effectMap: map[int]func([]string) []string{},
		ordering:  []int{},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		response := scanner.Text()
		if response == "end" {
			break
		}

		responseParts := strings.Split(response, " ")
		trigger, _ := strconv.Atoi(responseParts[0])
		ruleType := getRuleType(responseParts[1])
		params := responseParts[2:]
		addRule(trigger, ruleType, params, &rules)
	}
	return rules
}

func getRules() RuleSet {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Custom Rules Y/N? >> ")
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "y" {
		return getCustomRules()
	}
	return defaultRules
}

// This is our main function, this executes by default when we run the main package.
func main() {
	rules := getRules()
	upperLimit := getMaxNumber()
	fmt.Println(fizzBuzzUpTo(upperLimit, rules))
}
