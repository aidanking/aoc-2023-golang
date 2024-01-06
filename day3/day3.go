package day3

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aidanking/aoc-2023-golang/inputfiles"
)

type node struct {
	i int
	j int
}

func PrintSolution() {

	testData := readSchema("test.txt")
	inputData := readSchema("input.txt")

	fmt.Println("day 3")
	fmt.Println("Part 1")
	fmt.Println("Test Solution:", part1(testData))
	fmt.Println("Solution:", part1(inputData))
	fmt.Println("Part 2")
	fmt.Println("Test Solution:", part2(testData))
	fmt.Println("Solution:", part2(inputData))
	fmt.Println("")

}

func part1(schema [][]string) int {
	result := 0

	for i, line := range schema {
		for j := range line {
			if !isDigit(schema[i][j]) && schema[i][j] != "." {
				result += getPartNumberTotalForSymbol(schema, node{i: i, j: j})
			}
		}
	}
	return result
}

func part2(schema [][]string) int {
	result := 0

	for i, line := range schema {
		for j := range line {
			if schema[i][j] == "*" {
				result += getGearRatioForNode(schema, node{i: i, j: j})
			}
		}
	}

	return result
}

func getGearRatioForNode(schema [][]string, startNode node) int {
	nums := getNumsForEachNeighbourNode(schema, startNode)

	if len(nums) == 2 {
		return nums[0] * nums[1]
	}

	return 0
}

func getPartNumberTotalForSymbol(schema [][]string, startNode node) int {
	result := 0

	nums := getNumsForEachNeighbourNode(schema, startNode)

	for _, num := range nums {
		result += num
	}

	return result
}

func readSchema(fileName string) [][]string {
	var schema [][]string
	lines := inputfiles.ReadLinesFromFile("day3", fileName)

	for _, line := range lines {
		characters := strings.Split(line, "")
		schema = append(schema, characters)
	}

	return schema
}

func getNumsForEachNeighbourNode(schema [][]string, startNode node) []int {
	isVisited := makeIsVisited(schema)
	numNodes := getNumNeighbourNodes(schema, isVisited, startNode)
	var nums []int

	for _, numNode := range numNodes {
		if !isVisited[numNode.i][numNode.j] {

			right := numNode.j

			for right >= 0 && isDigit(schema[numNode.i][right]) {
				right--
			}

			left := right + 1
			var numsAsStrings []string

			for left < len(schema) && isDigit(schema[numNode.i][left]) {
				isVisited[numNode.i][left] = true
				numsAsStrings = append(numsAsStrings, schema[numNode.i][left])
				left++
			}

			num, numErr := strconv.Atoi(strings.Join(numsAsStrings, ""))

			if numErr != nil {
				log.Fatal(numErr)
			}
			nums = append(nums, num)
		}
	}
	return nums
}

func makeIsVisited(schema [][]string) [][]bool {

	var isVisited [][]bool = make([][]bool, len(schema))

	for i := range isVisited {
		isVisited[i] = make([]bool, len(schema[0]))
	}

	return isVisited
}

func getNumNeighbourNodes(schema [][]string, isVisited [][]bool, startNode node) []node {
	var numNodes []node
	var directions = [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

	for _, direction := range directions {
		nextNode := node{i: startNode.i + direction[0], j: startNode.j + direction[1]}
		if nextNode.i >= 0 && nextNode.i < len(schema) && nextNode.j >= 0 && nextNode.j < len(schema[0]) && !isVisited[nextNode.i][nextNode.j] && isDigit(schema[nextNode.i][nextNode.j]) {
			numNodes = append(numNodes, nextNode)
		}
	}

	return numNodes
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)

	return err == nil
}
