package main

import (
	"strconv"
	"testing"
)

var fizzBuzzTests = []struct {
	input    int
	expected string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{5, "Buzz"},
	{15, "FizzBuzz"},
	{7, "Bang"},
	{21, "FizzBang"},
	{35, "BuzzBang"},
	{105, "FizzBuzzBang"},
	{11, "Bong"},
	{33, "Bong"},
	{165, "Bong"},
	{13, "Fezz"},
	{39, "FizzFezz"},
	{65, "FezzBuzz"},
	{195, "FizzFezzBuzz"},
	{143, "FezzBong"},
	{255, "BuzzFizz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range fizzBuzzTests {
		result := fizzBuzz(test.input)
		if result != test.expected {
			t.Errorf("incorrect greeting, for name %s: got %s, expected %s", strconv.Itoa(test.input), result, test.expected)
		}
	}
}
