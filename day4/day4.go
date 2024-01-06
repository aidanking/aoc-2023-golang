package day4

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/aidanking/aoc-2023-golang/inputfiles"
)

type game struct {
	id             int
	winningNumbers []int
	myNumbers      []int
}

func PrintSolution() {
	testGames := readGames("test.txt")
	inputGames := readGames("input.txt")

	fmt.Println("day 4")
	fmt.Println("Part 1")
	fmt.Println("Test Solution:", part1(testGames))
	fmt.Println("Solution:", part1(inputGames))
	fmt.Println("part 2")
	fmt.Println("Test Solution:", part2(testGames))
	fmt.Println("Solution:", part2(inputGames))
	fmt.Println("")

}

func part1(games []game) int {
	result := 0

	for _, game := range games {
		winningNumberCount := getWinningNumberCount(game)

		if winningNumberCount > 0 {
			points := int(math.Pow(2.0, float64(winningNumberCount)-1)) // 2**(count-1)

			result += points
		}
	}

	return result
}

func part2(games []game) int {
	return getGameCount(games)
}

func getGameCount(inputGames []game) int {
	var games []game = make([]game, len(inputGames))
	copy(games, inputGames)
	var alreadySeenWinningNumbers map[int]int = make(map[int]int)
	count := 0

	for len(games) > 0 {
		count++
		game := games[len(games)-1]
		games = games[:len(games)-1]

		_, alreadySeen := alreadySeenWinningNumbers[game.id]
		winningNumberCount := 0

		if !alreadySeen {
			winningNumberCount = getWinningNumberCount(game)
			alreadySeenWinningNumbers[game.id] = winningNumberCount
		} else {
			winningNumberCount = alreadySeenWinningNumbers[game.id]
		}

		i := game.id
		j := i + winningNumberCount - 1

		for i < len(inputGames) && j < len(inputGames) && i <= j {
			games = append(games, inputGames[i])
			i++
		}
	}

	return count
}

func getWinningNumberCount(game game) int {
	count := 0
	foundWinningNmbers := make(map[int]bool)

	for _, winningNumber := range game.winningNumbers {
		foundWinningNmbers[winningNumber] = true
	}

	for _, myNumber := range game.myNumbers {
		_, found := foundWinningNmbers[myNumber]

		if found {
			count++
		}

	}

	return count
}

func readGames(fileName string) []game {
	lines := inputfiles.ReadLinesFromFile("day4", fileName)
	var games []game

	for _, line := range lines {
		parts := strings.Split(line, ":")
		gameId := getGameId(parts[0])
		gameNumbersAsStrings := strings.Split(parts[1], "|")
		var winningNumbers []int = convertToNumbers(gameNumbersAsStrings[0])
		var yourNumbers []int = convertToNumbers(gameNumbersAsStrings[1])
		games = append(games, game{id: gameId, winningNumbers: winningNumbers, myNumbers: yourNumbers})

	}

	return games
}

func getGameId(heading string) int {
	regex := regexp.MustCompile(`\d+`)

	gameIdString := regex.FindAllString(heading, -1)[0]

	num, numErr := strconv.Atoi(gameIdString)

	if numErr != nil {
		log.Fatal(numErr)
	}

	return num
}

func convertToNumbers(numberString string) []int {
	regex := regexp.MustCompile(`\d+`)

	numberStrings := regex.FindAllString(numberString, -1)

	var nums []int

	for _, s := range numberStrings {

		num, numErr := strconv.Atoi(s)

		if numErr != nil {
			log.Fatal(numErr)
		}

		nums = append(nums, num)
	}

	return nums
}
