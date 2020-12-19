package main

import (
	"bufio"
	"log"
	"os"
)

//BoardingGroup represents questions answer 'yes' for every member of a boarding group
type BoardingGroup []string

func sumAnyYes(groups []BoardingGroup) int {
	var counts []int

	for _, group := range groups {
		var groupString string
		for _, answer := range group {
			groupString += answer
		}

		chars := map[string]bool{}

		for _, char := range groupString {
			chars[string(char)] = true
		}

		counts = append(counts, len(chars))

	}

	return sum(counts)
}

func readCustomsForms(path string) []BoardingGroup {
	var groups []BoardingGroup

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var group BoardingGroup
	for scanner.Scan() {
		form := scanner.Text()
		if form == "" {
			groups = append(groups, group)
			group = BoardingGroup{}
			continue
		}
		group = append(group, form)
	}
	groups = append(groups, group)

	return groups

}

func sum(nums []int) int {
	var sum int

	for _, num := range nums {
		sum += num
	}

	return sum
}
