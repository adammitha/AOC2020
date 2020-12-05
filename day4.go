package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func countValidPassports(passports []map[string]string) int {
	var count int

	for _, passport := range passports {
		if checkKeys(passport) {
			count++
		}
	}

	return count
}

func checkKeys(passport map[string]string) bool {
	validKeys := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	allOK := true

	for _, key := range validKeys {
		if _, ok := passport[key]; !ok {
			allOK = false
		}
	}

	return allOK
}

func readPassports(path string) []map[string]string {
	var passports []map[string]string

	var i int

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			i++
			continue
		}
		if i >= len(passports) {
			passports = append(passports, make(map[string]string))
		}
		for _, field := range strings.Split(line, " ") {
			fieldValues := strings.Split(field, ":")
			k, v := fieldValues[0], fieldValues[1]
			passports[i][k] = v
		}
	}

	return passports
}
