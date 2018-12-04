package main

import (
	"log"
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
	size := 3
	expected := createBlankCloth()

	result := prepCloth(size, size)

	// Then
	compareCloths(result, expected, t)
}

func createBlankCloth() cloth {
	initialFitting := fitting{[]int{}, 0}
	expected := cloth{
		{initialFitting, initialFitting, initialFitting},
		{initialFitting, initialFitting, initialFitting},
		{initialFitting, initialFitting, initialFitting},
	}
	return expected
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
	expected := createBlankCloth()
	expected[1][1] = fitting{
		[]int{1},
		1,
	}

	// When
	fitToCloth(input, sampleCloth)

	// Then
	compareCloths(sampleCloth, expected, t)
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
	expected := createBlankCloth()
	expected[1][1] = fitting{
		[]int{1, 1},
		2,
	}

	// When
	fitToCloth(input, sampleCloth)
	fitToCloth(input, sampleCloth)

	// Then
	compareCloths(sampleCloth, expected, t)
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

func TestFindUniqueCoverage(t *testing.T) {
	// Given
	input := claim{
		id:         1,
		leftOffset: 1,
		topOffset:  1,
		width:      1,
		height:     1,
	}

	sampleCloth := prepCloth(3, 3)
	fitToCloth(input, sampleCloth)

	// When
	result := findUniqueCoverage(sampleCloth)

	if len(result) != 1 && result[0] != 1 {
		t.Errorf("Ids wrong, got: %d, want: %d.", result, 1)
	}

}

func compareCloths(cl1 cloth, cl2 cloth, t *testing.T) {
	for j := 0; j < len(cl1); j++ {
		for i := 0; i < len(cl1[0]); i++ {
			if cl1[j][i].numberOfClaims != cl2[j][i].numberOfClaims {
				t.Errorf("Cloth claim numbers don't match: %d, want: %d.", cl1[j][i].numberOfClaims, cl2[j][i].numberOfClaims)
			}
			eq := testEq(cl1[j][i].ids, cl2[j][i].ids)
			if !eq {
				t.Errorf("Cloth ids don't match: %v, want: %v.", cl1[j][i].ids, cl2[j][i].ids)
			}
		}
	}
}

func testEq(a, b []int) bool {
	if (a == nil) != (b == nil || len(b) == 0) {
		log.Print("Not nil or empty")
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
