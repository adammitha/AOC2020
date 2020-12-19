package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/golang-collections/collections/set"
)

//BoardingGroup represents questions answered 'yes' for every member of a boarding group
type BoardingGroup []string

func sumAllYes(groups []BoardingGroup) int {
	var count int

	for _, group := range groups {
		answers := set.New()
		for _, char := range strings.Split(group[0], "") {
			answers.Insert(char)
		}
		for _, person := range group[1:] {
			temp := set.New()
			for _, char := range strings.Split(person, "") {
				temp.Insert(char)
			}
			answers = answers.Intersection(temp)
		}
		count += answers.Len()
	}

	return count
}

func sumAnyYes(groups []BoardingGroup) int {
	var count int

	for _, group := range groups {
		var groupString string
		for _, answer := range group {
			groupString += answer
		}

		chars := map[string]bool{}

		for _, char := range groupString {
			chars[string(char)] = true
		}

		count += len(chars)

	}

	return count
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
