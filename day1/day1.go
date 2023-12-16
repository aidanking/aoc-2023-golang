package day1

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/aidanking/aoc-2023-golang/inputfiles"
)

var words = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

const longestWordLen = len("three")
const shortestWordLen = len("one")

func PrintSolution() {
	testDataPart1 := inputfiles.ReadLinesFromFile("day1", "test-part-1.txt")
	testDataPart2 := inputfiles.ReadLinesFromFile("day1", "test-part-2.txt")
	data := inputfiles.ReadLinesFromFile("day1", "input.txt")

	fmt.Println("Day 1")
	fmt.Println("Part 1")
	fmt.Println("Test Solution", part1(testDataPart1))
	fmt.Println("Solution", part1(data))
	fmt.Println("Part 2")
	fmt.Println("Test Solution", part2(testDataPart2))
	fmt.Println("Solution", part2(data))
	fmt.Println("")
}

func part1(lines []string) int {
	result := 0

	for _, line := range lines {
		result += concatFirstAndLastDigit(line, false)
	}

	return result
}

func part2(lines []string) int {
	result := 0

	for _, line := range lines {
		result += concatFirstAndLastDigit(line, true)
	}

	return result
}

func concatFirstAndLastDigit(line string, isPart2 bool) int {
	firstDigit := -1
	lastDigit := -1
	currentNum := -1

	for runeIndex, rune := range line {
		char := string(rune)

		if unicode.IsDigit(rune) {
			num, numErr := strconv.Atoi(char)

			if numErr != nil {
				currentNum = -1
				log.Fatal(numErr)
			}

			currentNum = num
		} else if isPart2 {

			foundNum, findNumErr := findWord(runeIndex, line)

			if findNumErr == nil {
				currentNum = foundNum
			} else {
				currentNum = -1
			}
		} else {
			currentNum = -1
		}

		if currentNum != -1 {

			if firstDigit == -1 {
				firstDigit = currentNum
			}
			lastDigit = currentNum
		}

	}

	return (firstDigit * 10) + lastDigit
}

func findWord(startIndex int, line string) (int, error) {
	j := startIndex + shortestWordLen
	endIndex := startIndex + longestWordLen

	for j <= len(line) && j <= endIndex {
		word := line[startIndex:j]
		num, hasWord := words[word]

		if hasWord {
			return num, nil
		}
		j++
	}

	return 0, errors.New("word not found")
}
