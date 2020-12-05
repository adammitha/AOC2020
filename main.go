package main

import (
	"fmt"
)

func main() {
	trees := readMap("inputs/trees.txt")
	fmt.Println(countTrees(trees, [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}))
}
