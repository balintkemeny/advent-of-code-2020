package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

type boardingPass struct {
	id   int
	row  int
	col  int
	code string
}

func (bp *boardingPass) Decode() bool {
	rowCode := bp.code[:7]
	colCode := bp.code[7:]

	var rowsOK bool
	var minRow int = 0
	var maxRow int = 127
	for _, v := range rowCode {
		if rune(v) == 'F' {
			maxRow = (maxRow-minRow)/2 + minRow
		} else if rune(v) == 'B' {
			minRow = (minRow+maxRow)/2 + 1
		}
	}
	if minRow == maxRow {
		rowsOK = true
	}
	bp.row = minRow

	var colsOK bool
	var minCol int = 0
	var maxCol int = 7
	for _, v := range colCode {
		if rune(v) == 'L' {
			maxCol = (maxCol-minCol)/2 + minCol
		} else if rune(v) == 'R' {
			minCol = (minCol+maxCol)/2 + 1
		}
	}
	if minCol == maxCol {
		colsOK = true
	}
	bp.col = minCol
	bp.id = bp.row*8 + bp.col

	return rowsOK && colsOK
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

	var boardingPasses []boardingPass
	for _, entry := range bytes.Split(rawData, []byte{'\n'}) {
		var pass boardingPass
		pass.code = string(entry)
		boardingPasses = append(boardingPasses, pass)
	}

	var maxID int
	var ids []int
	for _, pass := range boardingPasses {
		ok := pass.Decode()
		ids = append(ids, pass.id)
		if pass.id > maxID {
			maxID = pass.id
		}
		fmt.Println(ok, pass.row, pass.col, pass.id, "MAX:", maxID)
	}

	fmt.Printf("-------------------------------------------\nMAX ID: %d\n", maxID)

	sort.Ints(ids)

	var idCnt int = ids[0]
	var myID int
	for i := 0; i < len(ids)-1; i++ {
		if ids[i] != idCnt {
			myID = idCnt
			break
		}
		idCnt++
	}

	fmt.Printf("-------------------------------------------\nMY ID: %d\n", myID)
	fmt.Printf("-------------------------------------------\n")
}
