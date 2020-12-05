package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func splitRuleCreator(r rune) bool {
	return r == '\n' || r == ' '
}

func containsAllKeys(input, keys []string) bool {
	for _, key := range keys {
		var keyFound bool
		for _, v := range input {
			if v == key {
				keyFound = true
			}
		}
		if !keyFound {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var passData []string
	for _, entry := range bytes.Split(rawData, []byte{'\n', '\n'}) {
		passData = append(passData, string(entry))
	}

	fmt.Println("EX01:", ex01(passData))
	fmt.Println("EX02:", ex02(passData))
}

func ex01(passData []string) int {
	var validPasses int
	necessaryKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, pass := range passData {
		var passKeys []string
		pairs := strings.FieldsFunc(pass, splitRuleCreator)
		for _, pair := range pairs {
			kv := strings.Split(pair, ":")
			passKeys = append(passKeys, kv[0])
		}
		passValid := containsAllKeys(passKeys, necessaryKeys)
		if passValid {
			validPasses++
		}
	}

	return validPasses
}

type cPassport struct {
	passKeys   []string
	passValues []string
}

func byrValidator(byr string) bool {
	nByr, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}
	if nByr >= 1920 && nByr <= 2002 {
		return true
	}
	return false
}

func iyrValidator(iyr string) bool {
	nIyr, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}
	if nIyr >= 2010 && nIyr <= 2020 {
		return true
	}
	return false
}

func eyrValidator(eyr string) bool {
	nEyr, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}
	if nEyr >= 2020 && nEyr <= 2030 {
		return true
	}
	return false
}

func hgtValidator(hgt string) bool {
	if len(hgt) < 3 {
		return false
	}
	nHgt, err := strconv.Atoi(hgt[:len(hgt)-2])
	if err != nil {
		return false
	}
	switch hgt[len(hgt)-2:] {
	case "cm":
		if nHgt >= 150 && nHgt <= 193 {
			return true
		}
	case "in":
		if nHgt >= 59 && nHgt <= 76 {
			return true
		}
	}
	return false
}

func hclValidator(hcl string) bool {
	if len(hcl) != 7 || hcl[0] != '#' {
		return false
	}

	validChars := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	for i := 1; i < len(hcl); i++ {
		var charValid bool
		for _, c := range validChars {
			if rune(hcl[i]) == c {
				charValid = true
			}
		}
		if !charValid {
			return false
		}
	}
	return true
}

func eclValidator(ecl string) bool {
	validColours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, colour := range validColours {
		if ecl == colour {
			return true
		}
	}
	return false
}

func pidValidator(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	validChars := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 0; i < len(pid); i++ {
		var charValid bool
		for _, c := range validChars {
			if rune(pid[i]) == c {
				charValid = true
			}
		}
		if !charValid {
			return false
		}
	}
	return true
}

func validateCompletePassport(pass cPassport) bool {
	for i, v := range pass.passKeys {
		var paramValid bool
		switch v {
		case "byr":
			paramValid = byrValidator(pass.passValues[i])
		case "iyr":
			paramValid = iyrValidator(pass.passValues[i])
		case "eyr":
			paramValid = eyrValidator(pass.passValues[i])
		case "hgt":
			paramValid = hgtValidator(pass.passValues[i])
		case "hcl":
			paramValid = hclValidator(pass.passValues[i])
		case "ecl":
			paramValid = eclValidator(pass.passValues[i])
		case "pid":
			paramValid = pidValidator(pass.passValues[i])
		default:
			paramValid = true
		}
		if !paramValid {
			return false
		}
	}
	return true
}

func ex02(passData []string) int {
	var validPasses int
	var completePasses []cPassport
	necessaryKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, pass := range passData {
		var passKeys []string
		var passValues []string
		pairs := strings.FieldsFunc(pass, splitRuleCreator)
		for _, pair := range pairs {
			kv := strings.Split(pair, ":")
			passKeys = append(passKeys, kv[0])
			passValues = append(passValues, kv[1])
		}
		passComplete := containsAllKeys(passKeys, necessaryKeys)
		if passComplete {
			var p cPassport
			p.passKeys = passKeys
			p.passValues = passValues
			completePasses = append(completePasses, p)
		}
	}

	for _, cPass := range completePasses {
		if validateCompletePassport(cPass) {
			validPasses++
		}
	}
	return validPasses
}
