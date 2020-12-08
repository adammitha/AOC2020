package main

import (
	"bufio"
	"log"
	"os"
)

func readBoardingPasses(path string) []string {
	var boardingPasses []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boardingPass := scanner.Text()

		boardingPasses = append(boardingPasses, boardingPass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return boardingPasses
}
