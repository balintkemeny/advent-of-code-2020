package main

import "fmt"

type memoryLoc struct {
	current int
	last    int
}

func main() {
	input := []int{1, 2, 16, 19, 18, 0}
	m := make(map[int]memoryLoc)

	var previousNumber int
	for i := 0; i < 30000000; i++ {
		if i < len(input) {
			m[input[i]] = memoryLoc{i, -1}
			previousNumber = input[i]
		} else {
			if m[previousNumber].last == -1 {
				_, ok := m[0]
				if !ok {
					m[0] = memoryLoc{i, -1}
				} else {
					m[0] = memoryLoc{i, m[0].current}
				}
				previousNumber = 0
			} else {
				currentNumber := m[previousNumber].current - m[previousNumber].last
				_, ok := m[currentNumber]
				if !ok {
					m[currentNumber] = memoryLoc{i, -1}
				} else {
					m[currentNumber] = memoryLoc{i, m[currentNumber].current}
				}
				previousNumber = currentNumber
			}
		}
	}

	fmt.Println(previousNumber)
}
