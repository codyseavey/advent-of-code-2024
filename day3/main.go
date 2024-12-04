package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func toInt(s string) int {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func fs1(f *os.File) int {
	total := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var re = regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)`)
		for _, match := range re.FindAllString(line, -1) {
			// get the two numbers from the capture groups
			var re2 = regexp.MustCompile(`\d+`)
			nums := re2.FindAllString(match, -1)
			// convert the strings to ints
			num1 := toInt(nums[0])
			num2 := toInt(nums[1])
			// multiply the two numbers
			total += num1 * num2
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return total
}

func fs2(f *os.File) int {
	total := 0

	text, _ := io.ReadAll(f)
	line := string(text)

	// repeatedly find and remove the text between "don't()" and "do()"
	for {
		dontIndex := strings.Index(line, "don't()")
		doIndex := strings.Index(line, "do()")

		// if "don't()" is found before "do()" or there is no "do()", remove the text between them
		if dontIndex != -1 {
			if doIndex == -1 {
				// if you hit a "don't()" but no "do()", remove the text after "don't()" since there is nothing to re-enable the loop
				line = line[:dontIndex]
			} else if dontIndex < doIndex {
				// if there's a don't before a do, none of the text gets executed so we can remove it
				line = line[:dontIndex] + line[doIndex+4:]
			} else if doIndex < dontIndex {
				// if there is a do() before a don't, then we should just remove the do() from the string and continue because its a no-op
				line = line[:doIndex] + line[doIndex+4:]
			}
		} else {
			fmt.Println(len(line))
			break
		}
	}

	var re = regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)`)
	for _, match := range re.FindAllString(line, -1) {
		// get the two numbers from the capture groups
		var re2 = regexp.MustCompile(`\d+`)
		nums := re2.FindAllString(match, -1)
		// convert the strings to ints
		num1 := toInt(nums[0])
		num2 := toInt(nums[1])
		// multiply the two numbers
		total += num1 * num2
	}

	return total
}
