package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aidanking/aoc-2023-golang/inputfiles"
)

type game struct {
	id          int
	colorGroups []colorGroup
}

type colorGroup struct {
	colorAmounts []colorAmount
}

type colorAmount struct {
	amount  int
	colorId string
}

func PrintSolution() {
	testGames := readGames("test.txt")
	inputGames := readGames("input.txt")

	fmt.Println("day 2")
	fmt.Println("Part 1")
	fmt.Println("Test Solution:", part1(testGames))
	fmt.Println("Solution:", part1(inputGames))
	fmt.Println("Part 2")
	fmt.Println("Test Solution:", part2(testGames))
	fmt.Println("Solution:", part2(inputGames))
	fmt.Println("")

}

func part1(games []game) int {
	result := 0

	for _, game := range games {
		if isValidGame(game) {
			result += game.id
		}
	}

	return result
}

func part2(games []game) int {
	result := 0
	for _, game := range games {
		result += getPower(game)
	}
	
	return result
}

func getPower(game game) int {
	maxColorCount := map[string]int{"red": 1, "green": 1, "blue": 1}

	for _, colorGroup := range game.colorGroups {
		for _, colorAmount := range colorGroup.colorAmounts {
			if(colorAmount.amount > maxColorCount[colorAmount.colorId]) {
				maxColorCount[colorAmount.colorId] = colorAmount.amount
			}
		}
	}

	return maxColorCount["red"] * maxColorCount["green"] * maxColorCount["blue"]
}

func isValidGame(game game) bool {
	maxColorCount := map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, colorGroup := range game.colorGroups {
		for _, colorAmount := range colorGroup.colorAmounts {
			if colorAmount.amount > maxColorCount[colorAmount.colorId] {
				return false
			}	

		}
	}
	return true
}

func readGames(fileName string) []game {
	var games []game

	lines := inputfiles.ReadLinesFromFile("day2", fileName)

	for _, line := range lines {
		headerAndRest := strings.Split(line, ":")
		games = append(games, readGame(headerAndRest))
	}

	return games
}

func readGame(headerAndRest []string) game {
	game := game{id: readGameId(headerAndRest), colorGroups: []colorGroup{}}

	colorGroupsAsString := strings.Split(headerAndRest[1], ";")

	for _, colorGroupAsString := range colorGroupsAsString {
		colorGroup := readColorGroup(colorGroupAsString)
		game.colorGroups = append(game.colorGroups, colorGroup)
	}
	return game
}

func readGameId(headerAndRest []string) int {
	gameId, gameIdErr := strconv.Atoi(strings.Split(headerAndRest[0], " ")[1])

	if gameIdErr != nil {
		log.Fatal(gameIdErr)
	}

	return gameId
}

func readColorGroup(colorGroupAsString string) colorGroup {
	colorAmountsAsString := strings.Split(colorGroupAsString, ",")
	colorGroup := colorGroup{colorAmounts: []colorAmount{}}

	for _, colorAmountsAsString := range colorAmountsAsString {
		colorGroup.colorAmounts = append(colorGroup.colorAmounts, readColorAmount(colorAmountsAsString))

	}
	return colorGroup
}

func readColorAmount(colorAmountsAsString string) colorAmount {
	amountData := strings.Split(colorAmountsAsString, " ")

	amount, amountErr := strconv.Atoi(amountData[1])

	if amountErr != nil {
		log.Fatal(amountErr)
	}
	colorId := amountData[2]

	colorAmount := colorAmount{amount: amount, colorId: colorId}
	return colorAmount
}
