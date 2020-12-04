package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var requiredPassportKeys = [...]string{
	"ecl",
	"pid",
	"eyr",
	"hcl",
	"byr",
	"iyr",
	"hgt",
}

type PassportChecker struct {
	fields map[string]bool
}

func readPassports() []PassportChecker {
	f, _ := os.OpenFile("input.txt", os.O_APPEND|os.O_RDWR, 0644)

	passportCheckers := make([]PassportChecker, 0)

	scanner := bufio.NewScanner(f)
	currentPassportInfo := ""
	passportChecker := PassportChecker{fields: make(map[string]bool, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			currentPassportInfo += line + " "
			continue
		}

		currentPassportInfo = strings.TrimSpace(currentPassportInfo)
		fieldValueMap := strings.Split(currentPassportInfo, " ")
		for _, fieldValue := range fieldValueMap {
			pieces := strings.Split(fieldValue, ":")
			key := pieces[0]
			passportChecker.fields[key] = true
		}
		currentPassportInfo = ""
		passportCheckers = append(passportCheckers, passportChecker)
		passportChecker = PassportChecker{fields: make(map[string]bool, 0)}
	}
	return passportCheckers
}

func isValidPassport(p PassportChecker) bool {
	for _, key := range requiredPassportKeys {
		if p.fields[key] == false {
			return false
		}
	}

	return true
}

func main() {
	passportCheckers := readPassports()
	validCnt := 0
	for _, p := range passportCheckers {
		isValid := isValidPassport(p)
		if isValid {
			validCnt++
		}
	}
	fmt.Println(validCnt)
}
