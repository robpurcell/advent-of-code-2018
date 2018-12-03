package main

import (
	"reflect"
	"testing"
)

func TestParseClaim(t *testing.T) {
	// Given
	input := "#1 @ 258,327: 19x22"
	expected := claim{
		id:         1,
		leftOffset: 258,
		topOffset:  327,
		width:      19,
		height:     22,
	}

	// When
	result := parseClaim(input)

	// Then
	eq := reflect.DeepEqual(result, expected)
	if !eq {
		t.Errorf("Claims are not equal, got: %v, want: %v.", result, expected)
	}
}

func TestReadClaims(t *testing.T) {
	input := "#1 @ 258,327: 19x22\n #2 @ 553,11: 13x13"

	claims := readClaims(input)

	if len(claims) != 2 {
		t.Errorf("Number of claims incorrect, got: %v, want: %v.", len(claims), 2)
	}

}

func TestPrepCloth(t *testing.T) {
	// Given
	expected := cloth{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	result := prepCloth(3, 3)

	// Then
	eq := reflect.DeepEqual(result, expected)
	if !eq {
		t.Errorf("Cloths are not equal, got: %v, want: %v.", result, expected)
	}
}

func TestFitToCloth(t *testing.T) {
	// Given
	input := claim{
		id:         1,
		leftOffset: 1,
		topOffset:  1,
		width:      1,
		height:     1,
	}

	sampleCloth := prepCloth(3, 3)
	expected := cloth{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}

	// When
	fitToCloth(input, sampleCloth)

	// Then
	eq := reflect.DeepEqual(sampleCloth, expected)
	if !eq {
		t.Errorf("Cloths are not equal, got: %v, want: %v.", sampleCloth, expected)
	}
}

func TestFitToClothWithDuplicate(t *testing.T) {
	// Given
	input := claim{
		id:         1,
		leftOffset: 1,
		topOffset:  1,
		width:      1,
		height:     1,
	}

	sampleCloth := prepCloth(3, 3)
	expected := cloth{{0, 0, 0}, {0, 2, 0}, {0, 0, 0}}

	// When
	fitToCloth(input, sampleCloth)
	fitToCloth(input, sampleCloth)

	// Then
	eq := reflect.DeepEqual(sampleCloth, expected)
	if !eq {
		t.Errorf("Cloths are not equal, got: %v, want: %v.", sampleCloth, expected)
	}
}

func TestFindDuplicateCoverage(t *testing.T) {
	// Given
	input := claim{
		id:         1,
		leftOffset: 1,
		topOffset:  1,
		width:      1,
		height:     1,
	}

	sampleCloth := prepCloth(3, 3)
	expected := 1
	fitToCloth(input, sampleCloth)
	fitToCloth(input, sampleCloth)

	// When
	result := findDuplicateCoverage(sampleCloth)

	if result != expected {
		t.Errorf("Cloths are not equal, got: %d, want: %d.", result, expected)
	}

}
