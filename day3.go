package main

import (
	"bufio"
	"log"
	"os"
)

func countTrees(trees [][]bool, slopes [][2]int) []int {
	var counts = make([]int, len(slopes))

	for i := 0; i < len(slopes); i++ {
		slope := slopes[i]
		var posX, posY int
		width := len(trees[0])

		for posY < len(trees) {
			if trees[posY][posX%width] {
				counts[i]++
			}
			posX += slope[0]
			posY += slope[1]
		}

	}

	return counts
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
