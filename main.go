package main

import (
	"fmt"
)

func main() {
	expenses := readExpenses("inputs/expenses.txt")
	fmt.Println(findExpenses(expenses))
}
