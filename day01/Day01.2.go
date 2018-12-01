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

	frequency := findFrequency(0, map[int64]bool{}, s.Fields(string(input)))

	log.Print("The ANSWER: ", frequency)
}

func findFrequency(frequency int64, frequenciesSelected map[int64]bool, frequencies []string) int64 {
	for _, f := range frequencies {
		value, _ := strconv.ParseInt(f, 0, 64)
		frequency += value

		if !frequenciesSelected[frequency] {
			frequenciesSelected[frequency] = true
		} else {
			return frequency
		}
	}

	return findFrequency(frequency, frequenciesSelected, frequencies)
}
