package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func findMinIndexAndValue(input []int) (int, int) {
	var minIndex int
	var minValue int = input[0]
	for i, v := range input {
		if v < minValue {
			minIndex = i
			minValue = v
		}
	}
	return minIndex, minValue
}

func modInverse(a, b int) int {
	if b == 1 {
		return 1
	}

	var b0, q, x0, x1 int
	x0 = 0
	x1 = 1
	b0 = b

	for a > 1 {
		q = a / b
		a, b = b, a%b
		x0, x1 = x1-q*x0, x0
	}

	if x1 < 0 {
		x1 += b0
	}

	return x1
}

func crt(n, a []int) uint64 {
	var sum uint64
	var bigN int = 1
	for _, v := range n {
		bigN *= v
	}
	for i := range n {
		bigNi := bigN / n[i]
		sum += uint64(a[i] * bigNi * modInverse(bigNi, n[i]))
	}

	for sum > uint64(bigN) {
		sum -= uint64(bigN)
	}

	return sum
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

	splitData := bytes.Split(rawData, []byte{'\n'})
	myTimestamp, err := strconv.Atoi(string(splitData[0]))
	if err != nil {
		log.Fatal(err)
	}

	var busIDs []int
	var busOffsets []int
	var offset int
	for _, v := range bytes.Split(splitData[1], []byte{','}) {
		if string(v) == "x" {
			offset++
			continue
		}

		id, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}

		busIDs = append(busIDs, id)
		busOffsets = append(busOffsets, offset)
		offset++
	}

	var earliestDepartures []int
	for _, bus := range busIDs {
		var earliestDeparture int
		for {
			earliestDeparture += bus
			if earliestDeparture >= myTimestamp {
				break
			}
		}
		earliestDepartures = append(earliestDepartures, earliestDeparture)
	}

	minIndex, minTimestamp := findMinIndexAndValue(earliestDepartures)
	fmt.Println("EX01:", busIDs[minIndex]*(minTimestamp-myTimestamp))

	// I COULDN'T FIGURE IT OUT THIS WAY. I'M SAD NOW.
	/*
		sum := crt(busIDs, busOffsets)
		fmt.Println("EX02:", sum)
		fmt.Println(crt([]int{5, 7, 8}, []int{3, 1, 6}))
		for i := range busIDs {
			fmt.Printf("ID: %d, OFFSET %d, REMAINDER %d\n", busIDs[i], busOffsets[i], (sum+uint64(busOffsets[i]))%uint64(busIDs[i]))
		}
	*/

	// LET'S BRUTEFORCE THE CRAP OUT OF THIS
	var sol, jump uint64 = 0, 1
	for i := range busIDs {
		for (sol+uint64(busOffsets[i]))%uint64(busIDs[i]) != 0 {
			sol += jump
		}
		jump *= uint64(busIDs[i])
	}
	fmt.Println(sol)
}
