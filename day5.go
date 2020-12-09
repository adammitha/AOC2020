package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func findSeatID(seats []string) int {
	var seatIDs []int

	for _, seat := range seats {
		seatIDs = append(seatIDs, seatID(seat))
	}

	sort.Slice(seatIDs, func(i, j int) bool { return seatIDs[i] < seatIDs[j] })

	var mySeatID int

	for i := 0; i < len(seatIDs)-1; i++ {
		j := i + 1
		if (seatIDs[j] - seatIDs[i]) > 1 {
			mySeatID = seatIDs[j] - 1
		}
	}

	return mySeatID
}

func findHighestSeatID(seats []string) int {
	var maxID int

	for _, seat := range seats {
		if ID := seatID(seat); ID > maxID {
			maxID = ID
		}
	}

	return maxID
}

func seatID(seat string) int {
	rows := seat[:6]
	columns := seat[7:9]
	rowBounds := [2]int{0, 127}
	columnBounds := [2]int{0, 7}

	for _, char := range rows {
		half := (rowBounds[1] - rowBounds[0] + 1) / 2
		if char == 'F' {
			rowBounds[1] -= half
		} else {
			rowBounds[0] += half
		}
	}
	var row int
	if seat[6] == 'F' {
		row = rowBounds[0]
	} else {
		row = rowBounds[1]
	}

	for _, char := range columns {
		half := (columnBounds[1] - columnBounds[0] + 1) / 2
		if char == 'L' {
			columnBounds[1] -= half
		} else {
			columnBounds[0] += half
		}
	}
	var column int
	if seat[9] == 'L' {
		column = columnBounds[0]
	} else {
		column = columnBounds[1]
	}

	return row*8 + column

}

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
