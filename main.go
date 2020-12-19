package main

import (
	"fmt"
)

func main() {
	boardingGroups := readCustomsForms("inputs/customs_forms.txt")
	fmt.Println(sumAllYes(boardingGroups))
}
