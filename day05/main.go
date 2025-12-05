package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solvePart1(inputFile string) int {
	ranges, ingredients := loadDayData(inputFile)

	fresh := 0

	for _, id := range ingredients {
		isFresh := false
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				isFresh = true
				break
			}
		}
		if isFresh {
			fresh++
		}
	}
	return fresh
}

func solvePart2(inputFile string) int {
	ranges, _ := loadDayData(inputFile)
	freshIds := 0
	changes := true

	for changes {
		changes = false

		for i, r1 := range ranges[:len(ranges)-1] {
			if changes {
				break
			}

			nmin1, nmax1 := r1[0], r1[1]

			for j, r2 := range ranges[i+1:] {
				nmin2, nmax2 := r2[0], r2[1]

				if nmax2 < nmin1 || nmin2 > nmax1 {
					continue
				}

				changes = true

				ranges[i][0] = minInt(nmin1, nmin2)
				ranges[i][1] = maxInt(nmax1, nmax2)

				ranges = append(ranges[:i+j+1], ranges[i+j+2:]...)
				break
			}
		}
	}

	for _, r := range ranges {
		freshIds += r[1] - r[0] + 1
	}
	return freshIds
}

func minInt(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	} else {
		return v2
	}
}

func maxInt(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}

func loadDayData(inputFile string) ([][]int, []int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ranges [][]int
	var ingredients []int
	isRange := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			isRange = false
			continue
		}
		if isRange {
			interval := strings.Split(line, "-")
			leftint, err := strconv.Atoi(interval[0])
			if err != nil {
				log.Fatal(err)
			}
			rightint, err := strconv.Atoi(interval[1])
			if err != nil {
				log.Fatal(err)
			}
			ranges = append(ranges, []int{leftint, rightint})
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			ingredients = append(ingredients, id)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ranges, ingredients
}

func main() {
	// fmt.Println(solvePart1("sample_input.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("sample_input.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
