package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	input, err := ioutil.ReadFile("day01/day01-input.txt")
	check(err)

	nums := s.Fields(string(input))

	log.Print("The ANSWER: ", sumFrequencies(nums))
}

func sumFrequencies(frequencies []string) int64 {
	var frequency int64 = 0

	for _, f := range frequencies {
		value, _ := strconv.ParseInt(f, 0, 64)
		frequency += value
	}

	return frequency
}
