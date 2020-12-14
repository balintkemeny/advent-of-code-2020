package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func binaryToString(input int64) string {
	return fmt.Sprintf("%064b", input)
}

func maskReplace(input, mask string) (int64, error) {
	rs := []rune(input)
	for i := len(mask) - 1; i >= 0; i-- {
		if rune(mask[i]) == 'X' {
			continue
		}
		rs[i+len(input)-len(mask)] = rune(mask[i])
	}
	return strconv.ParseInt(string(rs), 2, 64)
}

func replaceX(in string) []string {
	var o1, o2 string = in, in
	rsI := []rune(in)
	rsO1 := []rune(o1)
	rsO2 := []rune(o2)
	var result []string
	for i, r := range rsI {
		if r == 'X' {
			rsO1[i] = '0'
			rsO2[i] = '1'
			outSlice1 := replaceX(string(rsO1))
			outSlice2 := replaceX(string(rsO2))
			result = append(result, outSlice1...)
			result = append(result, outSlice2...)
			return result
		}
	}

	result = append(result, in)
	return result
}

func memoryAddressDecoder(input, mask string) ([]int64, error) {
	rs := []rune(input)
	for i := len(mask) - 1; i >= 0; i-- {
		if rune(mask[i]) == '0' {
			continue
		}
		rs[i+len(input)-len(mask)] = rune(mask[i])
	}

	outStrings := replaceX(string(rs))
	var out []int64
	for _, s := range outStrings {
		v, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			return []int64{}, err
		}
		out = append(out, v)
	}

	return out, nil
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

	var mask string
	mem := make(map[int]int64)
	for _, line := range bytes.Split(rawData, []byte{'\n'}) {
		splitLine := bytes.Split(line, []byte{' ', '=', ' '})
		location := string(splitLine[0])
		if location == "mask" {
			mask = string(splitLine[1])
			continue
		}

		address, _ := strconv.Atoi(string(splitLine[0][4 : len(splitLine[0])-1]))
		value, _ := strconv.ParseInt(string(splitLine[1]), 0, 64)
		binString := binaryToString(value)
		modValue, _ := maskReplace(binString, mask)
		mem[address] = modValue
	}

	var sum int64
	for _, v := range mem {
		sum += v
	}
	fmt.Println("EX01:", sum)

	mem2 := make(map[int64]int64)
	for _, line := range bytes.Split(rawData, []byte{'\n'}) {
		splitLine := bytes.Split(line, []byte{' ', '=', ' '})
		location := string(splitLine[0])
		if location == "mask" {
			mask = string(splitLine[1])
			continue
		}

		address, _ := strconv.ParseInt(string(splitLine[0][4:len(splitLine[0])-1]), 0, 64)
		binAddress := binaryToString(address)
		value, _ := strconv.ParseInt(string(splitLine[1]), 0, 64)

		addresses, err := memoryAddressDecoder(binAddress, mask)
		if err != nil {
			log.Fatal(err)
		}

		for _, a := range addresses {
			mem2[a] = value
		}
	}

	var sum2 int64
	for _, v := range mem2 {
		sum2 += v
	}
	fmt.Println("EX02:", sum2)
}
