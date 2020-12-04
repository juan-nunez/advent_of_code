package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func newPassport(passportFields []string) Passport {
	passport := Passport{}
	for _, passportField := range passportFields {
		data := strings.Split(passportField, ":")
		key := strings.TrimSpace(data[0])
		value := strings.TrimSpace(data[1])
		if key == "byr" {
			passport.byr = value
		}
		if key == "iyr" {
			passport.iyr = value
		}
		if key == "eyr" {
			passport.eyr = value
		}
		if key == "hgt" {
			passport.hgt = value
		}
		if key == "hcl" {
			passport.hcl = value
		}
		if key == "ecl" {
			passport.ecl = value
		}
		if key == "pid" {
			passport.pid = value
		}
		if key == "cid" {
			passport.cid = value
		}
	}

	return passport
}

func (p Passport) isValidPart1() bool {
	fields := reflect.TypeOf(p)
	values := reflect.ValueOf(p)
	numFields := fields.NumField()
	for i := 0; i < numFields; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		if field.Name != "cid" && value.String() == "" {
			return false
		}
	}

	return true
}

func (p Passport) isValidPart2() bool {
	fields := reflect.TypeOf(p)
	values := reflect.ValueOf(p)
	numFields := fields.NumField()
	for i := 0; i < numFields; i++ {
		field := fields.Field(i)
		value := values.Field(i).String()
		if field.Name == "cid" {
			continue
		}

		if value == "" {
			return false
		}

		switch field.Name {
		case "byr":
			if value < "1920" || value > "2002" {
				return false
			}
		case "iyr":
			if value < "2010" || value > "2020" {
				return false
			}
		case "eyr":
			if value < "2020" || value > "2030" {
				return false
			}
		case "hgt":
			length := len(value)
			units := value[length-2:]
			if units != "cm" && units != "in" {
				return false
			}

			size := value[:length-2]
			if units == "cm" && (size < "150" || size > "193") {
				return false
			} else if units == "in" && (size < "59" || size > "76") {
				return false
			}
		case "hcl":
			pattern := "^#[a-f0-9]{6}$"
			match, _ := regexp.MatchString(pattern, value)
			if !match {
				return false
			}
		case "ecl":
			validValues := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			found := false
			for _, validValue := range validValues {
				if value == validValue {
					found = true
				}
			}
			if !found {
				return false
			}
		case "pid":
			pattern := "^[0-9]{9}$"
			match, _ := regexp.MatchString(pattern, value)
			if !match {
				return false
			}
		}
	}

	return true
}

func readPassports() []string {
	f, _ := os.Open("input.txt")

	passportInfos := make([]string, 0)
	scanner := bufio.NewScanner(f)
	passportInfo := ""
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// If still reading the same passport, then append to current info
		if line != "" {
			passportInfo += line + " "
			continue
		}
		// Done reading a single passport
		// Get rid of trailing space
		passportInfo = strings.TrimSpace(passportInfo)
		passportInfos = append(passportInfos, passportInfo)

		passportInfo = ""
	}
	// Append last passportInfo
	passportInfo = strings.TrimSpace(passportInfo)
	passportInfos = append(passportInfos, passportInfo)

	return passportInfos
}

func processPassports(passportInfos []string) []Passport {
	passports := make([]Passport, 0)
	for _, passportInfo := range passportInfos {
		passportFields := strings.Split(passportInfo, " ")
		passport := newPassport(passportFields)
		passports = append(passports, passport)
	}

	return passports
}

func doPart1() int {
	passportInfos := readPassports()
	passports := processPassports(passportInfos)
	numValidPassport := 0
	for _, passport := range passports {
		if passport.isValidPart1() {
			numValidPassport++
		}
	}

	return numValidPassport
}

func doPart2() int {
	passportInfos := readPassports()
	passports := processPassports(passportInfos)
	numValidPassport := 0
	for _, passport := range passports {
		if passport.isValidPart2() {
			numValidPassport++
		}
	}
	return numValidPassport
}

func main() {
	fmt.Println(doPart1())
	fmt.Println(doPart2())
}
