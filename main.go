package main

import (
	"fmt"
)

func main() {
	trees := readMap("inputs/trees.txt")
	fmt.Println(countTrees(trees))
}
