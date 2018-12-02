package main

import (
	"reflect"
	"testing"
)

func TestCountLetters(t *testing.T) {
	// Given
	testId := "aabbccddeee"

	// When
	result := countLetters(testId)

	// Then
	expected := map[int32]int{'a': 2, 'b': 2, 'c': 2, 'd': 2, 'e': 3}
	eq := reflect.DeepEqual(result, expected)
	if !eq {
		t.Errorf("Maps are not equal, got: %d, want: %d.", result, expected)
	}
}

func TestExactlyNLetters(t *testing.T) {
	input := "abcdef"
	testN(2, input, 0, t)
	testN(3, input, 0, t)

	input = "bababc"
	testN(2, input, 1, t)
	testN(3, input, 1, t)

	input = "abbcde"
	testN(2, input, 1, t)
	testN(3, input, 0, t)

	input = "abcccd"
	testN(2, input, 0, t)
	testN(3, input, 1, t)

	input = "aabcdd"
	testN(2, input, 1, t)
	testN(3, input, 0, t)

	input = "abcdee"
	testN(2, input, 1, t)
	testN(3, input, 0, t)

	input = "ababab"
	testN(2, input, 0, t)
	testN(3, input, 1, t)

}

func testN(numberToMatch int, s string, expected int, t *testing.T) {
	result := exactlyNLetters(createMap(s), numberToMatch)
	if result != expected {
		t.Errorf("Did not get the correct number for [%s], got: %d, want: %d.", s, result, expected)
	}
}

func createMap(letters string) map[int32]int {
	return countLetters(letters)
}

func TestCalculate(t *testing.T) {
	// Given
	input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	expected := 12

	// When
	result := calculate(input)

	// Then
	if result != expected {
		t.Errorf("Did not get the correct number for [%s], got: %d, want: %d.", input, result, expected)
	}
}
