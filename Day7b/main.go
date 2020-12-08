package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type bag struct {
	name     string
	contents []string
	qtys     []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rawData, err := ioutil.ReadAll(file)

	var bags []bag
	for _, line := range bytes.Split(rawData, []byte{'\n'}) {
		var b bag
		var words []string
		for _, word := range bytes.Split(line, []byte{' '}) {
			words = append(words, string(word))
		}

		for i, v := range words {
			if i < 2 {
				b.name += v
				if i == 0 {
					b.name += " "
				}
			}
			if v == "no" {
				break
			}
			if i > 4 && i%4 == 1 {
				var bagColor string = v + " " + words[i+1]
				b.contents = append(b.contents, bagColor)
				qty, _ := strconv.Atoi(words[i-1])
				b.qtys = append(b.qtys, qty)
			}
		}

		bags = append(bags, b)
	}

	fmt.Println(contains("shiny gold", bags))
}

func contains(col string, bags []bag) int {
	var ans int
	for _, b := range bags {
		if b.name == col {
			for _, qty := range b.qtys {
				ans += qty
			}
			for i, c := range b.contents {
				ans += b.qtys[i] * contains(c, bags)
			}
		}
	}
	return ans
}
