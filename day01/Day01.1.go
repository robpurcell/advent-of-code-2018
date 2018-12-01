package main

import (
	c "./common"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	s "strings"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	input, err := ioutil.ReadFile("day01/day01-input.txt")
	c.Check(err)

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
