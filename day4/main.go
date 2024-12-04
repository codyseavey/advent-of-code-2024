package main

import (
	"bufio"
	"log"
	"os"
)

type Direction int

const (
	N Direction = iota
	S
	E
	W
	NE
	NW
	SE
	SW
)

var directionMap = map[Direction][]int{
	N:  {0, -1},
	S:  {0, 1},
	E:  {1, 0},
	W:  {-1, 0},
	NE: {1, -1},
	NW: {-1, -1},
	SE: {1, 1},
	SW: {-1, 1},
}

func fs1(f *os.File) int {
	scanner := bufio.NewScanner(f)
	xmasCount := 0

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	grid := make([][]rune, 0)
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			var d Direction
			// Check to see if all the letters are in the letters array are in the grid
			if grid[i][j] == 'X' {
				// In each direction check if there is a 'M' in the grid
				mPositions := make([][]int, 0)
				mDirection := make([]Direction, 0)
				for k, v := range directionMap {
					if i+v[0] >= 0 && i+v[0] < len(grid) && j+v[1] >= 0 && j+v[1] < len(grid[i]) && grid[i+v[0]][j+v[1]] == 'M' {
						d = k
						mPositions = append(mPositions, []int{i + v[0], j + v[1]})
						mDirection = append(mDirection, d)
					}
				}
				// If there is a 'M' in the grid, check if there is a 'A' in the grid
				aPositions := make([][]int, 0)
				aDirection := make([]Direction, 0)
				if len(mPositions) > 0 {
					for k, v := range mPositions {
						// check if next position is in bounds
						if v[0]+directionMap[mDirection[k]][0] >= 0 && v[0]+directionMap[mDirection[k]][0] < len(grid) && v[1]+directionMap[mDirection[k]][1] >= 0 && v[1]+directionMap[mDirection[k]][1] < len(grid[i]) && grid[v[0]+directionMap[mDirection[k]][0]][v[1]+directionMap[mDirection[k]][1]] == 'A' {
							d = mDirection[k]
							aPositions = append(aPositions, []int{v[0] + directionMap[d][0], v[1] + directionMap[d][1]})
							aDirection = append(aDirection, d)
						}
					}
				}

				// If there is a 'A' in the grid, check if there is a 'S' in the grid
				if len(aPositions) > 0 {
					for k, v := range aPositions {
						if v[0]+directionMap[aDirection[k]][0] >= 0 && v[0]+directionMap[aDirection[k]][0] < len(grid) && v[1]+directionMap[aDirection[k]][1] >= 0 && v[1]+directionMap[aDirection[k]][1] < len(grid[i]) && grid[v[0]+directionMap[aDirection[k]][0]][v[1]+directionMap[aDirection[k]][1]] == 'S' {
							xmasCount++
						}
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return xmasCount
}

func fs2(f *os.File) int {
	scanner := bufio.NewScanner(f)
	xmasCount := 0

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	grid := make([][]rune, 0)
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// Check to see if all the letters are in the letters array are in the grid
			if grid[i][j] == 'A' {
				// check NE, NW, SE, SW for either a 'M' or 'S' and if all 4 directions have a 'M' or 'S' then increment xmasCount
				if i-1 >= 0 && j-1 >= 0 && i+1 < len(grid) && j+1 < len(grid[i]) {
					if (grid[i-1][j-1] == 'M' || grid[i-1][j-1] == 'S') && (grid[i-1][j+1] == 'M' || grid[i-1][j+1] == 'S') && (grid[i+1][j-1] == 'M' || grid[i+1][j-1] == 'S') && (grid[i+1][j+1] == 'M' || grid[i+1][j+1] == 'S') {

						if grid[i-1][j-1] == 'M' {
							if grid[i+1][j+1] != 'S' {
								continue
							}
						}
						if grid[i-1][j+1] == 'M' {
							if grid[i+1][j-1] != 'S' {
								continue
							}
						}
						if grid[i+1][j-1] == 'M' {
							if grid[i-1][j+1] != 'S' {
								continue
							}
						}
						if grid[i+1][j+1] == 'M' {
							if grid[i-1][j-1] != 'S' {
								continue
							}
						}

						if grid[i-1][j-1] == 'S' {
							if grid[i+1][j+1] != 'M' {
								continue
							}
						}
						if grid[i-1][j+1] == 'S' {
							if grid[i+1][j-1] != 'M' {
								continue
							}
						}
						if grid[i+1][j-1] == 'S' {
							if grid[i-1][j+1] != 'M' {
								continue
							}
						}
						if grid[i+1][j+1] == 'S' {
							if grid[i-1][j-1] != 'M' {
								continue
							}
						}

						xmasCount++
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return xmasCount
}
