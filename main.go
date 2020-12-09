package main

import (
	"fmt"
)

func main() {
	boardingPasses := readBoardingPasses("inputs/boarding_passes.txt")
	fmt.Println(findSeatID(boardingPasses))
}
