package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func solvePart1(inputFile string) int {
	data := loadDayData(inputFile)
	total_joltage := 0
	for _, digits := range data {
		// Find position of first max digit excluding last digit
		posA := 0
		maxA := 0
		for i, digit := range digits[0 : len(digits)-1] {
			if digit > maxA {
				maxA = digit
				posA = i
			}
		}
		// Find position of second max digit after first max digit
		posB := posA + 1
		maxB := digits[posB]
		for i, digit := range digits[posB:] {
			if digit > maxB {
				posB = i
				maxB = digit
			}
		}
		total_joltage += maxA*10 + maxB
	}
	return total_joltage
}

func solvePart2(inputFile string) int {
	banks := loadDayData(inputFile)
	total_joltage := 0
	banksize := len(banks[0])
	batteries := 12
	for _, bank := range banks {
		// Loop twelve times the biggest digit
		currpos := 0
		for i := 1; i <= batteries; i++ {
			remainSize := batteries - i
			searchslice := bank[currpos : banksize-remainSize]
			newpos := firstMaxDigitPosition(searchslice) + currpos
			total_joltage += bank[newpos] * int(math.Pow10(remainSize))
			currpos = newpos + 1
		}
	}
	return total_joltage
}

func firstMaxDigitPosition(digits []int) int {
	pos := 0
	max := 0
	for i, digit := range digits[0:] {
		if digit > max {
			max = digit
			pos = i
		}
		if max == 9 {
			break
		}
	}
	return pos
}

func loadDayData(inputFile string) [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var data [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var digits []int
		for i := 0; i < len(line); i++ {
			digit := int(line[i] - '0')
			digits = append(digits, digit)
		}
		data = append(data, digits)
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
