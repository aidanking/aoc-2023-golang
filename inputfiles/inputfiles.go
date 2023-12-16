package inputfiles

import (
	"bufio"
	"log"
	"os"
	"path"
)

func ReadLinesFromFile(dayKey string, fileName string) []string {

	wdName, wdErr := os.Getwd()

	if wdErr != nil {
		log.Fatal(wdErr)
	}

	file, fileErr := os.Open(path.Join(wdName, dayKey, fileName))

	if fileErr != nil {
		log.Fatal(fileErr)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		log.Fatal(scannerErr)
	}

	return lines
}
