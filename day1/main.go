package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// sort slice of ints
func sortList(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

func fs1(f *os.File) int {
	list1 := []int{}
	list2 := []int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ints := strings.Fields(line)
		int1, err := strconv.Atoi(ints[0])
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		int2, err := strconv.Atoi(ints[1])
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		list1 = append(list1, int1)
		list2 = append(list2, int2)
	}

	list1 = sortList(list1)
	list2 = sortList(list2)

	distance := 0
	for i := 0; i < len(list1); i++ {
		distance += int(math.Abs(float64(list1[i] - list2[i])))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return distance
}

func fs2(f *os.File) int {
	list1 := []int{}
	list2 := []int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ints := strings.Fields(line)
		int1, err := strconv.Atoi(ints[0])
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		int2, err := strconv.Atoi(ints[1])
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		list1 = append(list1, int1)
		list2 = append(list2, int2)
	}

	similarityScore := 0
	for _, v := range list1 {
		totalOccurences := 0
		for _, v2 := range list2 {
			if v2 == v {
				totalOccurences++
			}
		}
		similarityScore += (v * totalOccurences)
	}

	return similarityScore
}
