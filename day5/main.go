package main

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func fs1(f *os.File) int {
	sum := 0

	text, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	t := strings.Split(string(text), "\n\n")

	pageOrdering := map[int][]int{}
	for _, p := range strings.Split(t[0], "\n") {
		q := strings.Split(p, "|")
		q1, err := strconv.Atoi(q[0])
		if err != nil {
			panic(err)
		}
		q2, err := strconv.Atoi(q[1])
		if err != nil {
			panic(err)
		}
		if _, ok := pageOrdering[q1]; !ok {
			pageOrdering[q1] = []int{}
		}
		pageOrdering[q1] = append(pageOrdering[q1], q2)
	}

	pageList := [][]int{}
	for _, l := range strings.Split(t[1], "\n") {
		q := strings.Split(l, ",")
		pageL := []int{}
		for _, v := range q {
			p, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			pageL = append(pageL, p)
		}
		pageList = append(pageList, pageL)
	}

	for _, v := range pageList {
		valid := true

		// [1,2,3,4]
		for i := len(v) - 1; i >= 0; i-- {
			if z, ok := pageOrdering[v[i]]; ok {
				for _, k := range v[:i] {
					if contains(z, k) {
						valid = false
					}
				}
			}
		}

		if valid {
			// take the middle term of v
			sum += v[len(v)/2]
		}

	}

	return sum
}

func fs2(f *os.File) int {
	sum := 0

	text, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	t := strings.Split(string(text), "\n\n")

	pageOrdering := map[int][]int{}
	for _, p := range strings.Split(t[0], "\n") {
		q := strings.Split(p, "|")
		q1, err := strconv.Atoi(q[0])
		if err != nil {
			panic(err)
		}
		q2, err := strconv.Atoi(q[1])
		if err != nil {
			panic(err)
		}
		if _, ok := pageOrdering[q1]; !ok {
			pageOrdering[q1] = []int{}
		}
		pageOrdering[q1] = append(pageOrdering[q1], q2)
	}

	pageList := [][]int{}
	for _, l := range strings.Split(t[1], "\n") {
		q := strings.Split(l, ",")
		pageL := []int{}
		for _, v := range q {
			p, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			pageL = append(pageL, p)
		}
		pageList = append(pageList, pageL)
	}

	for _, v := range pageList {
		valid := true

		// [1,2,3,4]
		for i := len(v) - 1; i >= 0; i-- {
			if z, ok := pageOrdering[v[i]]; ok {
				for _, k := range v[:i] {
					if contains(z, k) {
						valid = false
					}
				}
			}
		}

		if !valid {
			isFixed := false
			newV := make([]int, len(v))
			copy(newV, v)
			for !isFixed {
				// use the page ordering rules to fix the list
				valid2 := true
				badNumber := -1
				for i := len(newV) - 1; i >= 0; i-- {
					if z, ok := pageOrdering[newV[i]]; ok {
						for _, k := range newV[:i] {
							if contains(z, k) {
								valid2 = false
								badNumber = newV[i]
							}
						}
					}
				}

				if !valid2 {
					// in the newV array move the badNumber up one position
					for i := 0; i < len(newV); i++ {
						if newV[i] == badNumber {
							// make sure i is in bounds

							if i > 0 {
								newV[i] = newV[i-1]
								newV[i-1] = badNumber
							}
						}
					}

				} else {
					isFixed = true
				}

			}

			// take the middle term of v
			sum += newV[len(newV)/2]
		}

	}

	return sum
}
