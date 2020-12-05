package main

import (
	"fmt"
)

func main() {
	passwords := readPasswords("inputs/passwords.txt")
	fmt.Println(countValidPasswords2(passwords))
}
