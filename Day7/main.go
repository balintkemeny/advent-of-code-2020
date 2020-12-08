package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	rawData, err := ioutil.ReadAll(file)

	var bags map[string][]string
	bags = make(map[string][]string)
	for _, line := range bytes.Split(rawData, []byte{'\n'}) {
		var bagName string
		var bagContents []string
		var words []string
		for _, word := range bytes.Split(line, []byte{' '}) {
			words = append(words, string(word))
		}

		for i, v := range words {
			if i < 2 {
				bagName += v
				if i == 0 {
					bagName += " "
				}
			}
			if v == "no" {
				break
			}
			if i > 4 && i%4 == 1 {
				var bagColor string = v + " " + words[i+1]
				bagContents = append(bagContents, bagColor)
			}
		}

		bags[bagName] = bagContents
	}

	var ansMap map[string]bool
	ansMap = make(map[string]bool)
	contains("shiny gold", bags, &ansMap)

	var ans int
	for range ansMap {
		ans++
	}
	fmt.Println(ans)
}

func contains(col string, bags map[string][]string, ansMap *map[string]bool) {
	for bag, colors := range bags {
		for _, c := range colors {
			if c == col {
				(*ansMap)[bag] = true
				contains(bag, bags, ansMap)
			}
		}
	}
}
