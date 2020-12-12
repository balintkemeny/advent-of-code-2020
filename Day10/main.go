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

	var adapters []int
	adapters = append(adapters, 0)
	for _, line := range bytes.Split(rawData, []byte{'\n'}) {
		adapter, _ := strconv.Atoi(string(line))
		adapters = append(adapters, adapter)
	}

	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	var oneJoltDiffs int
	var threeJoltDiffs int

	for i := 1; i < len(adapters); i++ {
		if adapters[i]-adapters[i-1] == 1 {
			oneJoltDiffs++
		} else if adapters[i]-adapters[i-1] == 3 {
			threeJoltDiffs++
		}
	}

	fmt.Println("EX01:", oneJoltDiffs*threeJoltDiffs)

	var skippedToPaths map[int]int64
	skippedToPaths = make(map[int]int64)

	skippedToPaths[0] = 1

	for i := 1; i < len(adapters); i++ {
		skippedToPaths[adapters[i]] = 0
		_, ok1 := skippedToPaths[adapters[i]-1]
		if ok1 {
			skippedToPaths[adapters[i]] += skippedToPaths[adapters[i]-1]
		}
		_, ok2 := skippedToPaths[adapters[i]-2]
		if ok2 {
			skippedToPaths[adapters[i]] += skippedToPaths[adapters[i]-2]
		}
		_, ok3 := skippedToPaths[adapters[i]-3]
		if ok3 {
			skippedToPaths[adapters[i]] += skippedToPaths[adapters[i]-3]
		}
	}

	fmt.Println("EX02:", skippedToPaths[adapters[len(adapters)-1]])
}
