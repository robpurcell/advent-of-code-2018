package main

import (
	c "../common"
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
	input, err := ioutil.ReadFile("day03.1/day03-input.txt")
	c.Check(err)

	claims := readClaims(string(input))
	fabric := prepCloth(1000, 1000)

	for _, cl := range claims {
		fitToCloth(parseClaim(cl), fabric)
	}

	log.Print("The ANSWER: ", findDuplicateCoverage(fabric))
}

func readClaims(input string) []string {
	return s.Split(input, "\n")
}

type claim struct {
	id         int
	leftOffset int
	topOffset  int
	width      int
	height     int
}

func parseClaim(input string) claim {
	fields := s.Fields(input)

	id, _ := strconv.Atoi(fields[0][1:])
	leftOffset, _ := strconv.Atoi(s.Split(fields[2], ",")[0])
	topOffset, _ := strconv.Atoi(s.TrimRight(s.Split(fields[2], ",")[1], ":"))
	width, _ := strconv.Atoi(s.Split(fields[3], "x")[0])
	height, _ := strconv.Atoi(s.Split(fields[3], "x")[1])

	return claim{
		id:         id,
		leftOffset: leftOffset,
		topOffset:  topOffset,
		width:      width,
		height:     height,
	}
}

func prepCloth(width int, height int) cloth {
	cloth := make(cloth, height)
	squareInches := make([]int, width*height)

	for i := range cloth {
		cloth[i], squareInches = squareInches[:width], squareInches[width:]
	}

	return cloth
}

type cloth [][]int

func fitToCloth(cl claim, clth cloth) {
	for j := cl.topOffset; j < cl.topOffset+cl.height; j++ {
		for i := cl.leftOffset; i < cl.leftOffset+cl.width; i++ {
			clth[i][j] += 1
		}
	}
}

func findDuplicateCoverage(clth cloth) int {
	count := 0
	for j := 0; j < len(clth); j++ {
		for i := 0; i < len(clth[0]); i++ {
			if clth[j][i] > 1 {
				count++
			}
		}
	}

	return count
}
