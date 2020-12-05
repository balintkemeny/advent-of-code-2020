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

func ex01() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawData, err := ioutil.ReadAll(file)
	var validPwds int

	for i, line := range bytes.Split(rawData, []byte{'\n'}) {
		lineData := bytes.Split(line, []byte{' '})
		oc := bytes.Split(lineData[0], []byte{'-'})

		ocMin, err := strconv.Atoi(string(oc[0]))
		if err != nil {
			log.Panic(err)
		}

		ocMax, err := strconv.Atoi(string(oc[1]))
		if err != nil {
			log.Panic(err)
		}

		reqCh := string(lineData[1][0])
		pwd := string(lineData[2])
		cnt := strings.Count(pwd, reqCh)

		var valid bool
		if cnt >= ocMin && cnt <= ocMax {
			valid = true
			validPwds++
		}

		if i < 10 {
			fmt.Printf("%d, %d, %d, %s, %s, %t\n", ocMin, ocMax, cnt, reqCh, pwd, valid)
		}
	}

	fmt.Println("The number of valid passwords is: ", validPwds)
}

func ex02() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawData, err := ioutil.ReadAll(file)
	var validPwds int

	for i, line := range bytes.Split(rawData, []byte{'\n'}) {
		lineData := bytes.Split(line, []byte{' '})
		oc := bytes.Split(lineData[0], []byte{'-'})

		pos1, err := strconv.Atoi(string(oc[0]))
		if err != nil {
			log.Panic(err)
		}

		pos2, err := strconv.Atoi(string(oc[1]))
		if err != nil {
			log.Panic(err)
		}

		reqCh := rune(lineData[1][0])
		pwd := string(lineData[2])

		var valid bool
		var reqChOnPos int
		if pos1 > len(pwd) || pos2 > len(pwd) {
			continue
		}

		if rune(pwd[pos1-1]) == reqCh {
			reqChOnPos++
		}

		if rune(pwd[pos2-1]) == reqCh {
			reqChOnPos++
		}

		if reqChOnPos == 1 {
			valid = true
			validPwds++
		}

		if i < 10 {
			fmt.Printf("%d, %d, %d, %d, %c, %s, %t\n", pos1, pos2, len(pwd), reqChOnPos, reqCh, pwd, valid)
		}
	}

	fmt.Println("The number of valid passwords is: ", validPwds)
}

func main() {
	fmt.Println("EX1:")
	ex01()
	fmt.Println("----------------------------------------")
	fmt.Println("EX2:")
	ex02()
	fmt.Println("----------------------------------------")
}
