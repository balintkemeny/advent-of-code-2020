package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func aocMultiplier(s []int) (int, error) {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i]+s[j] == 2020 {
				return s[i] * s[j], nil
			}
		}
	}
	return -1, errors.New("Error: no two distinct number have a sum of 2020 in the provided slice")
}

func aocMultiplierThree(s []int) (int, error) {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			for k := j + 1; k < len(s); k++ {
				if s[i]+s[j]+s[k] == 2020 {
					return s[i] * s[j] * s[k], nil
				}
			}
		}
	}
	return -1, errors.New("Error: no three distinct number have a sum of 2020 in the provided slice")
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

	var numData []int
	for _, line := range bytes.Split(rawData, []byte{'\n'}) {
		lineNum, err := strconv.Atoi(string(line))
		if err != nil {
			log.Fatal(err)
		}
		numData = append(numData, lineNum)
	}

	solution, err := aocMultiplier(numData)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Ex01 Solution: ", solution)

	solution2, err := aocMultiplierThree(numData)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Ex02 Solution: ", solution2)
}
