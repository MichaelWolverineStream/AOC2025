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
	data := loadDayData(inputFile)
	invalidsum := 0
	for _, ids := range data {
		interval := strings.Split(ids, "-")
		curr, err := strconv.Atoi(interval[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(interval[1])
		if err != nil {
			log.Fatal(err)
		}

		for curr <= end {
			currStr := strconv.Itoa(curr)
			if len(currStr)%2 != 0 {
				curr += 1
				continue
			}
			strSplit := int(len(currStr) / 2)
			currStr1 := currStr[0:strSplit]
			currStr2 := currStr[strSplit:]
			if currStr1 == currStr2 {
				invalidsum += curr
			}
			curr += 1
		}
	}

	return invalidsum
}

func solvePart2(inputFile string) int {
	data := loadDayData(inputFile)

	invalidsum := 0

	for _, ids := range data {
		interval := strings.Split(ids, "-")

		curr, err := strconv.Atoi(interval[0])
		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.Atoi(interval[1])
		if err != nil {
			log.Fatal(err)
		}

		for curr <= end {
			currStr := strconv.Itoa(curr)
			currStrlen := len(currStr)
			for i := range int(currStrlen / 2) {
				if currStrlen%(i+1) != 0 {
					continue
				}
				if allStringsAreEqual(splitStringIntoSubstrings(currStr, i+1)) {
					invalidsum += curr
					break
				}
			}
			curr += 1
		}
	}
	return invalidsum
}

func allStringsAreEqual(strs []string) bool {
	strsCount := len(strs)
	if strsCount <= 0 {
		return false
	}
	if strsCount == 1 {
		return true
	}

	firstString := strs[0]
	for i := 1; i < strsCount; i++ {
		if strs[i] != firstString {
			return false
		}
	}

	return true
}

func splitStringIntoSubstrings(str string, substrLength int) []string {
	subStrings := []string{}

	if substrLength <= 0 {
		return subStrings
	}

	start := 0
	end := substrLength
	maxend := len(str)

	for start < maxend {
		subString := str[start:end]
		subStrings = append(subStrings, subString)

		start = end
		end = min(maxend, end+substrLength)
	}

	return subStrings
}

func loadDayData(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var data []string
	for scanner.Scan() {
		line := scanner.Text()
		data = strings.Split(line, ",")
		break
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
