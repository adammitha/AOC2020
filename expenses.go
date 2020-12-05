package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func findExpenses(expenses []int) (int, int) {
	var expense1, expense2 int
	var i, j int
	for ; i < len(expenses)-1; i++ {
		j = i + 1
		for ; j < len(expenses); j++ {
			if expenses[i]+expenses[j] == 2020 {
				expense1 = expenses[i]
				expense2 = expenses[j]
			}
		}
	}
	return expense1, expense2
}

func readExpenses(path string) []int {
	var expenses []int

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		expenses = append(expenses, expense)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return expenses
}
