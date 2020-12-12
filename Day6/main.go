package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ex01() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var customsGroups []string
	var groupAnswers []int
	for _, group := range bytes.Split(rawData, []byte{'\n', '\n'}) {
		customsGroups = append(customsGroups, string(group))
		var groupMap map[rune]bool
		groupMap = make(map[rune]bool)
		for _, c := range string(group) {
			if c == '\n' {
				continue
			}
			_, ok := groupMap[rune(c)]
			if !ok {
				groupMap[rune(c)] = true
			}
		}
		var ans int
		for range groupMap {
			ans++
		}
		groupAnswers = append(groupAnswers, ans)
	}

	fmt.Println(groupAnswers[0], customsGroups[0])
	var ans int
	for _, v := range groupAnswers {
		ans += v
	}
	fmt.Println(ans)
}

func ex02() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var groupAnswers []int
	for _, group := range bytes.Split(rawData, []byte{'\n', '\n'}) {
		var groupMap map[rune]int
		groupMap = make(map[rune]int)
		members := 1
		for _, c := range string(group) {
			if c == '\n' {
				members++
				continue
			}
			_, ok := groupMap[c]
			if !ok {
				groupMap[c] = 1
			} else {
				groupMap[c]++
			}
		}
		var ans int
		for _, v := range groupMap {
			if v == members {
				ans++
			}
		}
		groupAnswers = append(groupAnswers, ans)
	}

	var ans int
	for _, v := range groupAnswers {
		ans += v
	}
	fmt.Println(ans)
}

func main() {
	ex01()
	ex02()
}
