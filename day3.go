package main

import (
	"bufio"
	"log"
	"os"
)

func countTrees(trees [][]bool) int {
	var count int
	var posX, posY int
	width := len(trees[0])

	for posY < len(trees) {
		if trees[posY][posX%width] {
			count++
		}
		posX += 3
		posY++
	}

	return count
}

func readMap(path string) [][]bool {
	var trees [][]bool

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineBool []bool
		for _, slot := range line {
			if slot == '.' {
				lineBool = append(lineBool, false)
			} else {
				lineBool = append(lineBool, true)
			}
		}
		trees = append(trees, lineBool)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return trees

}
