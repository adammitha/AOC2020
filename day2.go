package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func countValidPasswords(passwords []string) int {
	var count int

	for _, passwordString := range passwords {
		passwordSlice := strings.Split(passwordString, " ")
		letter, password := string(passwordSlice[1][0]), passwordSlice[2]
		countRange := strings.Split(passwordSlice[0], "-")
		min, _ := strconv.Atoi(countRange[0])
		max, _ := strconv.Atoi(countRange[1])
		letterCount := strings.Count(password, letter)
		if min <= letterCount && letterCount <= max {
			count++
		}
	}
	return count
}

func countValidPasswords2(passwords []string) int {
	var count int

	for _, passwordString := range passwords {
		passwordSlice := strings.Split(passwordString, " ")
		letter, password := string(passwordSlice[1][0]), passwordSlice[2]
		countRange := strings.Split(passwordSlice[0], "-")
		p1, _ := strconv.Atoi(countRange[0])
		p2, _ := strconv.Atoi(countRange[1])
		l1 := string(password[p1-1])
		l2 := string(password[p2-1])
		if (l1 == letter || l2 == letter) && (l1 != l2) {
			count++
		}
	}
	return count
}

func readPasswords(path string) []string {
	var passwords []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return passwords
}
