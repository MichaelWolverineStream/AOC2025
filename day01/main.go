package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solvePart1(inputFile string) int {
	data := loadDayData(inputFile)

	d := 50
	p := 0

	for _, v := range data {
		steps, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch v[0] {
		case 'L':
			d -= steps
			d += 100
		case 'R':
			d += steps
		default:
			log.Fatal("Bad direction")
		}
		d = d % 100
		if d == 0 {
			p += 1
		}
	}
	return p
}

func solvePart2(inputFile string) int {
	const trackLimit = 100

	data := loadDayData(inputFile)

	digit := 50
	pass := 0
	stepdir := 1

	for _, v := range data {
		if len(v) < 2 {
			log.Fatal("Unexpected data")
		}
		steps, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch v[0] {
		case 'L':
			stepdir = -1
		case 'R':
			stepdir = 1
		default:
			log.Fatal("Bad direction")
		}

		for steps > 0 {
			steps -= 1
			digit += stepdir
			if digit == trackLimit || digit == -trackLimit {
				digit = 0
			}
			if digit == 0 {
				pass += 1
			}
		}

	}
	return pass
}

func loadDayData(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []string
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	// fmt.Println(solvePart2("sample_input.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
