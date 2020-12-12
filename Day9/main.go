package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

func findFirstInvalidNumber(nums []int, validRange int) int {
	for i := validRange; i < len(nums); i++ {
		var valid bool
		for j := i - validRange; j < i; j++ {
			for k := i - validRange; k < i; k++ {
				if j == k {
					continue
				}
				if nums[j]+nums[k] == nums[i] {
					valid = true
				}
			}
		}
		if !valid {
			return i
		}
	}
	return -1
}

func findSumSet(nums []int, targetSum int) []int {
	for i := 0; i < len(nums)-1; i++ {
		var sum int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == targetSum {
				return nums[i : j+1]
			}
			if sum > targetSum {
				break
			}
		}
	}
	return []int{}
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

	var nums []int
	for _, line := range bytes.Split(rawData, []byte{'\n'}) {
		num, _ := strconv.Atoi(string(line))
		nums = append(nums, num)
	}

	var validRange int = 25
	index := findFirstInvalidNumber(nums, validRange)
	firstInvalid := nums[index]
	fmt.Println("EX01:", index, firstInvalid)

	sumSet := findSumSet(nums, firstInvalid)
	sort.Ints(sumSet)
	var min, max int = sumSet[0], sumSet[len(sumSet)-1]

	fmt.Println(min, max, min+max)
}
