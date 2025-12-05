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

	var calc [][]int

	for _, r := range ranges {
		nmin, nmax := r[0], r[1]

		isOverlap := false

		for i, c := range calc {
			// No overlap
			if nmax < c[0] {
				continue
			}
			if nmin > c[1] {
				continue
			}

			isOverlap = true

			calc[i][0] = int(math.Min(float64(nmin), float64(c[0])))
			calc[i][1] = int(math.Max(float64(nmax), float64(c[1])))

			break
		}

		if !isOverlap {
			calc = append(calc, []int{nmin, nmax})
		}

	}

	changes := true

	for changes {

		changes = false

		for i, r1 := range calc[:len(calc)-1] {
			if changes {
				break
			}
			for j, r2 := range calc[i+1:] {
				nmin, nmax := r2[0], r2[1]

				// No overlap
				if nmax < r1[0] {
					continue
				}
				if nmin > r1[1] {
					continue
				}

				changes = true

				calc[i][0] = int(math.Min(float64(nmin), float64(r1[0])))
				calc[i][1] = int(math.Max(float64(nmax), float64(r1[1])))

				calc = append(calc[:i+j+1], calc[i+j+2:]...)

				break
			}
		}
	}

	for _, r := range calc {
		freshIds += r[1] - r[0] + 1
	}

	return freshIds
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
