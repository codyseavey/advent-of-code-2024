package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func fs1(f *os.File) int {
	scanner := bufio.NewScanner(f)
	var wellBehaved int
	for scanner.Scan() {
		line := scanner.Text()
		ints := strings.Fields(line)
		var currInt, lastInt, oldInt int
		var err error
		valid := true
		for i := 0; i < len(ints); i++ {
			if i == 0 {
				currInt, err = strconv.Atoi(ints[i])
				if err != nil {
					log.Fatal(err)
				}
			} else if i == 1 {
				lastInt = currInt
				currInt, err = strconv.Atoi(ints[i])
				if err != nil {
					log.Fatal(err)
				}
			} else if i > 1 {
				oldInt = lastInt
				lastInt = currInt
				currInt, err = strconv.Atoi(ints[i])
				if err != nil {
					log.Fatal(err)
				}

				diff1 := math.Abs(float64(oldInt) - float64(lastInt))
				diff2 := math.Abs(float64(lastInt) - float64(currInt))

				valid = ((oldInt < lastInt && lastInt < currInt) || (oldInt > lastInt && lastInt > currInt)) && diff1 <= 3 && diff2 <= 3
				if !valid {
					break
				}
			}
		}
		if valid {
			wellBehaved++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wellBehaved
}

// Check if the list is all descending or ascending
func isWellBehaved(ints []int) bool {
	ascending := true
	descending := true
	for i := 0; i < len(ints)-1; i++ {
		if ints[i] < ints[i+1] {
			descending = false
		}
		if ints[i] > ints[i+1] {
			ascending = false
		}
	}
	return ascending || descending
}

// Check if there is a difference of at least 1 but not more than 3 between each number in the list
func isDifferenceWellBehaved(ints []int) bool {
	for i := 0; i < len(ints)-1; i++ {
		diff := math.Abs(float64(ints[i]) - float64(ints[i+1]))
		if diff > 3 || diff == 0 {
			return false
		}
	}
	return true
}

func removeElement(slice []int, s int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:s])
	copy(newSlice[s:], slice[s+1:])
	return newSlice
}

func fs2(f *os.File) int {
	scanner := bufio.NewScanner(f)
	var wellBehaved int
	for scanner.Scan() {
		line := scanner.Text()
		intsS := strings.Fields(line)
		ints := make([]int, len(intsS))
		// convert strings to ints
		for i := 0; i < len(intsS); i++ {
			ints[i], _ = strconv.Atoi(intsS[i])
		}

		isWellBehavedB := isDifferenceWellBehaved(ints) && isWellBehaved(ints)
		if !isWellBehavedB {
			// try removing each element and check if the list is well behaved
			for i := 0; i < len(ints); i++ {
				newInts := removeElement(ints, i)
				if isDifferenceWellBehaved(newInts) && isWellBehaved(newInts) {
					isWellBehavedB = true
					break
				}
			}
		}

		if isWellBehavedB {
			wellBehaved++
		} else {
			log.Println("Not well behaved:", ints)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wellBehaved
}
