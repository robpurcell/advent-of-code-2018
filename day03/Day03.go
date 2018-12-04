package main

import (
	c "../common"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	s "strings"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	input, err := ioutil.ReadFile("day03/day03-input.txt")
	c.Check(err)

	claims := readClaims(string(input))
	fabric := prepCloth(1000, 1000)

	for _, cl := range claims {
		fitToCloth(parseClaim(cl), fabric)
	}

	log.Print("The 1st ANSWER: ", findDuplicateCoverage(fabric))
	log.Print("The 2nd ANSWER: ", findUniqueCoverage(fabric))
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
	squareInches := make([]fitting, width*height)

	for i := range cloth {
		cloth[i], squareInches = squareInches[:width], squareInches[width:]
	}

	return cloth
}

type fitting struct {
	ids            []int
	numberOfClaims int
}
type cloth [][]fitting

func fitToCloth(cl claim, clth cloth) {
	for j := cl.topOffset; j < cl.topOffset+cl.height; j++ {
		for i := cl.leftOffset; i < cl.leftOffset+cl.width; i++ {
			clth[i][j].ids = append(clth[i][j].ids, cl.id)
			clth[i][j].numberOfClaims += 1
		}
	}
}

func findDuplicateCoverage(clth cloth) int {
	count := 0
	for j := 0; j < len(clth); j++ {
		for i := 0; i < len(clth[0]); i++ {
			if clth[j][i].numberOfClaims > 1 {
				count++
			}
		}
	}

	return count
}

func findUniqueCoverage(clth cloth) []int {
	ids := map[int]bool{}
	for i := 1; i < 1311; i++ {
		ids[i] = false
	}

	for j := 0; j < len(clth); j++ {
		for i := 0; i < len(clth[0]); i++ {
			if clth[j][i].numberOfClaims > 1 {
				for _, id := range clth[j][i].ids {
					ids[id] = true
				}
			}
		}
	}

	var keys []int
	for k := range ids {
		if !ids[k] {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)

	return keys
}
