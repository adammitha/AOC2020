package main

import (
	"bufio"
	"log"
	"os"
)

//BoardingGroup represents questions answer 'yes' for every member of a boarding group
type BoardingGroup []string

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
		}
		group = append(group, form)
	}

	return groups

}
