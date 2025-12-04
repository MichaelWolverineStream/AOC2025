package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	X int
	Y int
}

type grid [][]rune

func (g grid) IsInside(c coord) bool {
	return c.Y >= 0 && c.Y < len(g) && c.X >= 0 && c.X < len(g[0])
}

const (
	rollRune      = '@'
	emptyRune     = '.'
	maxNeighbours = 3
)

var directions = []coord{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func (g grid) isAccessible(c coord) bool {
	blockages := 0

	for _, d := range directions {
		neighbour := coord{c.X + d.X, c.Y + d.Y}

		if !g.IsInside(neighbour) {
			continue
		}

		if g[neighbour.Y][neighbour.X] == rollRune {
			blockages++
		}
	}

	return blockages <= maxNeighbours
}

func solvePart1(inputFile string) int {
	grid := loadDayData(inputFile)

	accessibleRolls := 0

	for y, row := range grid {
		for x, char := range row {
			if char != rollRune {
				continue
			}
			if grid.isAccessible(coord{x, y}) {
				accessibleRolls++
			}

		}
	}
	return accessibleRolls
}

func solvePart2(inputFile string) int {
	grid := loadDayData(inputFile)

	accessibleRolls := 0
	changed := true

	for changed {
		changed = false

		for y, row := range grid {
			for x, char := range row {
				if char != rollRune {
					continue
				}

				if grid.isAccessible(coord{x, y}) {
					accessibleRolls++
					grid[y][x] = emptyRune
					changed = true
				}
			}
		}
	}

	return accessibleRolls
}

func loadDayData(inputFile string) grid {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data grid
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for i := 0; i < len(line); i++ {
			row = []rune(line)
		}
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	fmt.Println(solvePart1("sample_input.txt"))
	fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("sample_input.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
