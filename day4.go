package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func countValidPassports(passports []map[string]string) int {
	var count int

	for _, passport := range passports {
		if checkFields(passport) {
			count++
		}
	}

	return count
}

func checkFields(passport map[string]string) bool {
	validKeys := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, key := range validKeys {
		if _, ok := passport[key]; !ok {
			return false
		}
		if !validateField(key, passport[key]) {
			return false
		}
	}

	return true
}

func validateField(key string, value string) bool {
	switch key {
	case "byr":
		year, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		if year >= 1920 && year <= 2002 {
			return true
		}
	case "iyr":
		year, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		if year >= 2010 && year <= 2020 {
			return true
		}
	case "eyr":
		year, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		if year >= 2020 && year <= 2030 {
			return true
		}
	case "hgt":
		height, err := strconv.Atoi(string(regexp.MustCompile("[0-9]*").Find([]byte(value))))
		if err != nil {
			log.Fatal(err)
		}
		units := string(regexp.MustCompile("cm|in").Find([]byte(value)))

		if units == "cm" && height >= 150 && height <= 193 {
			return true
		} else if units == "in" && height >= 59 && height <= 76 {
			return true
		}
	case "hcl":
		return regexp.MustCompile("#[0-9a-f]{6}").Match([]byte(value))
	case "ecl":
		return regexp.MustCompile("amb|blu|brn|gry|grn|hzl|oth").Match([]byte(value))
	case "pid":
		return regexp.MustCompile("[0-9]{9}").Match([]byte(value))

	}
	return false
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
