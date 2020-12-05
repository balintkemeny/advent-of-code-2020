package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func slopeCalculator(dX, dY int) int {
	var treesHit int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rawData, err := ioutil.ReadAll(file)

	var cX int
	for i, line := range bytes.Split(rawData, []byte{'\n'}) {
		for j, point := range bytes.Split(line, []byte{}) {
			if i%dY == 0 && j == cX {
				if string(point) == "#" {
					treesHit++
					fmt.Printf("X: %d, Y: %d, TREE HIT\n", j, i)
				} else {
					fmt.Printf("X: %d, Y: %d, OPEN\n", j, i)
				}
			}
		}
		if i%dY == 0 {
			cX += dX
			if cX > 30 {
				cX -= 31
			}
		}
	}

	fmt.Printf("TOTAL TREES HIT: %d\n", treesHit)
	return treesHit
}

func main() {
	t1 := slopeCalculator(1, 1)
	t2 := slopeCalculator(3, 1)
	t3 := slopeCalculator(5, 1)
	t4 := slopeCalculator(7, 1)
	t5 := slopeCalculator(1, 2)
	fmt.Println(t1, t2, t3, t4, t5)
	fmt.Println("FINAL RESULT", t1*t2*t3*t4*t5)
}
