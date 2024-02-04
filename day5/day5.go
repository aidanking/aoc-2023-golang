package day5

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/aidanking/aoc-2023-golang/inputfiles"
)

type alnamac struct {
	seeds                 []int
	alnamacNumRangeGroups [][]alnamacNumRange
}

type alnamacNumRange struct {
	dest int
	src  int
	len  int
}

func PrintSolution() {

	testData := readAlmanac("test.txt")
	inputData := readAlmanac("input.txt")

	fmt.Println("day 5")
	fmt.Println("Part 1")
	fmt.Println("Test Solution:", part1(testData))
	fmt.Println("Solution:", part1(inputData))
	fmt.Println("Part 2")
	fmt.Println("Test Solution:", part2(testData))
	fmt.Println("Solution:", part2(inputData))
	fmt.Println("")
}

func part1(alnamac alnamac) int {
	lowestLocation := math.MaxInt

	for _, seed := range alnamac.seeds {
		currentValue := getLowestLocation(alnamac, seed)

		if currentValue < lowestLocation {
			lowestLocation = currentValue
		}
	}

	return lowestLocation
}

func part2(alnamac alnamac) int {
	lowestLocation := math.MaxInt

	for i := 0; i < len(alnamac.seeds); i += 2 {
		firstSeed := alnamac.seeds[i]
		secondSeed := alnamac.seeds[i+1]
		for j := 0; j < secondSeed; j++ {
			newSeed := firstSeed + j
			currentValue := getLowestLocation(alnamac, newSeed)

			if currentValue < lowestLocation {
				lowestLocation = currentValue
			}
		}
	}

	return lowestLocation
}

func getLowestLocation(alnamac alnamac, seed int) int {
	currentValue := seed

	for _, group := range alnamac.alnamacNumRangeGroups {
		for _, numRange := range group {
			if currentValue >= numRange.src && currentValue < numRange.src+numRange.len {
				currentValue = numRange.dest + (currentValue - numRange.src)
				break
			}
		}
	}

	return currentValue
}

func readAlmanac(fileName string) alnamac {
	lines := inputfiles.ReadLinesFromFile("day5", fileName)
	result := alnamac{
		seeds:                 []int{},
		alnamacNumRangeGroups: [][]alnamacNumRange{},
	}
	groupsCompleted := 0
	var group []string
	for _, line := range lines {
		if line == "" {
			if groupsCompleted == 0 {
				seeds := readSeeds(group[0])

				result.seeds = seeds
			} else {
				result.alnamacNumRangeGroups = append(result.alnamacNumRangeGroups, readAlnamacNumRangeGroup(group))
			}
			group = []string{}
			groupsCompleted++
		} else {
			group = append(group, line)
		}
	}
	result.alnamacNumRangeGroups = append(result.alnamacNumRangeGroups, readAlnamacNumRangeGroup(group))
	groupsCompleted++

	return result
}

func readSeeds(seedsLine string) []int {
	var seeds []int
	seedsString := strings.Split(strings.Split(seedsLine, ": ")[1], " ")

	for _, seed := range seedsString {
		num, numErr := strconv.Atoi(seed)

		if numErr != nil {
			log.Fatal(numErr)
		}
		seeds = append(seeds, num)
	}

	return seeds
}

func readAlnamacNumRangeGroup(linesGroup []string) []alnamacNumRange {
	var mappings []alnamacNumRange

	for lineIndex, line := range linesGroup {
		if lineIndex > 0 {
			stringNums := strings.Split(line, " ")
			mapping := alnamacNumRange{}
			for stringNumIndex, stringNum := range stringNums {

				num, numErr := strconv.Atoi(stringNum)

				if numErr != nil {
					log.Fatal(numErr)
				}

				if stringNumIndex == 0 {
					mapping.dest = num
				} else if stringNumIndex == 1 {
					mapping.src = num
				} else if stringNumIndex == 2 {
					mapping.len = num
				}

			}
			mappings = append(mappings, mapping)
		}
	}

	return mappings
}
