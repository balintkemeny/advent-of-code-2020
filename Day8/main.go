package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type instruction struct {
	name  string
	value int
	ran   bool
}

func execute(ins *instruction, accumulator *int) (int, bool) {
	if ins.ran {
		return 0, true
	}

	var next int
	switch ins.name {
	case "acc":
		*accumulator += ins.value
		next = 1
	case "jmp":
		next = ins.value
	case "nop":
		next = 1
	}
	ins.ran = true
	return next, false
}

func runInstructions(instructions []instruction) (bool, int) {
	var acc int
	var cnt int
	var next int
	var ranFlag bool
	for !ranFlag {
		cnt += next
		if cnt >= len(instructions) {
			return true, acc
		}
		next, ranFlag = execute(&instructions[cnt], &acc)
		fmt.Printf("INSTRUCTION EXECUTED. CNT: %d, NXT: %d, ACC: %d\n", cnt, next, acc)
	}
	return false, acc
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

	var instructions []instruction
	for _, line := range bytes.Split(rawData, []byte{'\n'}) {
		lineStr := string(line)
		var ins instruction
		ins.name = lineStr[:3]
		val, _ := strconv.Atoi(lineStr[5:])
		if lineStr[4] == '+' {
			ins.value = val
		} else if lineStr[4] == '-' {
			ins.value = 0 - val
		}
		instructions = append(instructions, ins)
	}

	instructionsForEx01 := make([]instruction, len(instructions))
	copy(instructionsForEx01, instructions)
	ex01Through, ex01Acc := runInstructions(instructionsForEx01)
	fmt.Println("/////////////////")
	fmt.Println(ex01Through, ex01Acc)
	fmt.Println("/////////////////")

	var ex02Acc int
	var ranUntilTheEnd bool
	var idForModified int = -1
	for !ranUntilTheEnd {
		idForModified++
		tmpInstructions := make([]instruction, len(instructions))
		copy(tmpInstructions, instructions)

		if tmpInstructions[idForModified].name == "jmp" {
			tmpInstructions[idForModified].name = "nop"
		} else if tmpInstructions[idForModified].name == "nop" {
			tmpInstructions[idForModified].name = "jmp"
		} else {
			continue
		}
		ranUntilTheEnd, ex02Acc = runInstructions(tmpInstructions)
	}
	fmt.Println("/////////////////")
	fmt.Println(ranUntilTheEnd, ex02Acc)
	fmt.Println("/////////////////")
}
