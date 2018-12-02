package main

import (
	"reflect"
	"testing"
)

func TestFindDifference(t *testing.T) {
	// Given
	a := "abcde"
	b := "axcye"
	expected := 2

	// When
	result := findDifference(a, b)

	// Then
	if result != expected {
		t.Errorf("Did not get the correct difference, got: %d, want: %d.", result, expected)
	}
}

func TestSearch(t *testing.T) {
	// Given
	ids := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	expected := match{
		original:    "abcde",
		bestMatchId: "axcye",
		difference:  2,
	}

	// When
	result := search("abcde", ids[1:])

	// Then
	eq := reflect.DeepEqual(result, expected)
	if !eq {
		t.Errorf("Did not get the correct match, got: %+v want: %+v.", result, expected)
	}

	expected = match{
		original:    "fghij",
		bestMatchId: "fguij",
		difference:  1,
	}

	// When
	result = search("fghij", ids[2:])

	// Then
	eq = reflect.DeepEqual(result, expected)
	if !eq {
		t.Errorf("Did not get the correct match, got: %+v want: %+v.", result, expected)
	}
}

func TestFindBestMatch(t *testing.T) {
	// Given
	ids := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	expected := match{
		original:    "fghij",
		bestMatchId: "fguij",
		difference:  1,
	}

	// When
	result := findBestMatch(ids)

	// Then
	eq := reflect.DeepEqual(result, expected)
	if !eq {
		t.Errorf("Did not get the correct match, got: %+v want: %+v.", result, expected)
	}
}

func TestCommonLetters(t *testing.T) {
	// Given
	a := "abcdef"
	b := "abcjef"
	expected := "abcef"

	// When
	result := findCommonLetters(a, b)

	// Then
	if result != expected {
		t.Errorf("Did not get the correct matching letters, got: %s, want: %s.", result, expected)
	}
}
