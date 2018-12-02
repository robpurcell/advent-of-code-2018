package main

import (
	c "../common"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	s "strings"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	input, err := ioutil.ReadFile("day02.1/day02-input.txt")
	c.Check(err)

	ids := s.Fields(string(input))
	bestMatch := findBestMatch(ids)

	log.Print("The ANSWER: ", findCommonLetters(bestMatch.original, bestMatch.bestMatchId))
}

func findBestMatch(ids []string) match {
	bestMatch := search(ids[0], ids[1:])

	for i, id := range ids[1:] {
		match := search(id, ids[i+2:])
		if match.difference < bestMatch.difference {
			bestMatch = match
		}
	}
	return bestMatch
}

func search(a string, b []string) match {
	bestMatch := match{
		bestMatchId: "",
		difference:  100,
	}

	return accSearch(a, b, bestMatch)
}

func accSearch(a string, b []string, m match) match {
	if len(b) == 0 {
		return m
	}

	diff := findDifference(a, b[0])

	if diff < m.difference {
		m.original = a
		m.difference = diff
		m.bestMatchId = b[0]
	}

	return accSearch(a, b[1:], m)
}

func findDifference(a string, b string) int {
	return accFindDifferences(a, b, 0)
}

func accFindDifferences(a string, b string, acc int) int {
	if a[0] != b[0] {
		acc++
	}

	if len(a) == 1 {
		return acc
	} else {
		return accFindDifferences(a[1:], b[1:], acc)
	}
}

type match struct {
	original    string
	bestMatchId string
	difference  int
}

func findCommonLetters(a string, b string) string {
	var builder s.Builder
	builder.Grow(32)

	for i, p := range a {
		var comp int32 = int32(b[i])
		if p == comp {
			_, _ = fmt.Fprintf(&builder, "%c", p)
		}
	}

	return builder.String()
}
