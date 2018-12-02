package main

import (
	c "../common"
	"io/ioutil"
	"log"
	"os"
	s "strings"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	input, err := ioutil.ReadFile("day02.1/day02.1-input.txt")
	c.Check(err)

	ids := s.Fields(string(input))
	result := calculate(ids)

	log.Print("The ANSWER: ", result)
}

func calculate(ids []string) int {
	twos := 0
	threes := 0

	for _, id := range ids {
		counts := countLetters(id)
		twos += exactlyNLetters(counts, 2)
		threes += exactlyNLetters(counts, 3)
	}

	return twos * threes
}

func countLetters(id string) map[int32]int {
	letters := map[int32]int{}
	for _, l := range id {
		letters[l] = letters[l] + 1
	}

	return letters
}

func exactlyNLetters(letters map[int32]int, n int) int {
	for _, v := range letters {
		if v == n {
			return 1
		}
	}

	return 0
}
