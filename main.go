package main

import (
	"fmt"
)

func main() {
	passports := readPassports("inputs/passports.txt")
	fmt.Println(countValidPassports(passports))
}
