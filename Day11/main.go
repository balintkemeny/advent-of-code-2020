package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type seatCoordinates struct {
	row int
	col int
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

	/*
		0 - FLOOR
		1 - EMPTY SEAT
		2 - OCCUPIED SEAT
	*/
	var seating [][]int
	for _, line := range bytes.Split(rawData, []byte{'\n'}) {
		var seatingLine []int
		for _, char := range bytes.Split(line, []byte{}) {
			if string(char) == "L" {
				seatingLine = append(seatingLine, 1)
			} else if string(char) == "." {
				seatingLine = append(seatingLine, 0)
			}
		}
		seating = append(seating, seatingLine)
	}

	copySeating := make([][]int, len(seating))
	for i := range seating {
		copySeating[i] = make([]int, len(seating[i]))
		copy(copySeating[i], seating[i])
	}
	fmt.Println("EX01:", ex01(seating))
	fmt.Println("EX02:", ex02(copySeating))
}

func sweep(seating [][]int, row, col int) int {
	var occupieds int
	if row != 0 && col != 0 {
		if seating[row-1][col-1] == 2 {
			occupieds++
		}
	}
	if row != 0 {
		if seating[row-1][col] == 2 {
			occupieds++
		}
	}
	if row != 0 && col != len(seating[row])-1 {
		if seating[row-1][col+1] == 2 {
			occupieds++
		}
	}
	if col != 0 {
		if seating[row][col-1] == 2 {
			occupieds++
		}
	}
	if col != len(seating[row])-1 {
		if seating[row][col+1] == 2 {
			occupieds++
		}
	}
	if row != len(seating)-1 && col != 0 {
		if seating[row+1][col-1] == 2 {
			occupieds++
		}
	}
	if row != len(seating)-1 {
		if seating[row+1][col] == 2 {
			occupieds++
		}
	}
	if row != len(seating)-1 && col != len(seating[row])-1 {
		if seating[row+1][col+1] == 2 {
			occupieds++
		}
	}
	return occupieds
}

func megaSweep(seating [][]int, row, col int) int {
	var occupieds int

	var i1, j1 int = row - 1, col - 1
	for {
		if i1 < 0 || j1 < 0 {
			break
		}

		if seating[i1][j1] == 2 {
			occupieds++
			break
		} else if seating[i1][j1] == 1 {
			break
		}

		i1--
		j1--
	}

	for i2 := row - 1; i2 >= 0; i2-- {
		if seating[i2][col] == 2 {
			occupieds++
			break
		} else if seating[i2][col] == 1 {
			break
		}
	}

	var i3, j3 int = row - 1, col + 1
	for {
		if i3 < 0 || j3 >= len(seating[i3]) {
			break
		}

		if seating[i3][j3] == 2 {
			occupieds++
			break
		} else if seating[i3][j3] == 1 {
			break
		}

		i3--
		j3++
	}

	for j4 := col - 1; j4 >= 0; j4-- {
		if seating[row][j4] == 2 {
			occupieds++
			break
		} else if seating[row][j4] == 1 {
			break
		}
	}

	for j5 := col + 1; j5 < len(seating[row]); j5++ {
		if seating[row][j5] == 2 {
			occupieds++
			break
		} else if seating[row][j5] == 1 {
			break
		}
	}

	var i6, j6 int = row + 1, col - 1
	for {
		if i6 >= len(seating) || j6 < 0 {
			break
		}

		if seating[i6][j6] == 2 {
			occupieds++
			break
		} else if seating[i6][j6] == 1 {
			break
		}

		i6++
		j6--
	}

	for i7 := row + 1; i7 < len(seating); i7++ {
		if seating[i7][col] == 2 {
			occupieds++
			break
		} else if seating[i7][col] == 1 {
			break
		}
	}

	var i8, j8 int = row + 1, col + 1
	for {
		if i8 >= len(seating) || j8 >= len(seating[i8]) {
			break
		}

		if seating[i8][j8] == 2 {
			occupieds++
			break
		} else if seating[i8][j8] == 1 {
			break
		}

		i8++
		j8++
	}

	return occupieds
}

func ex01(seating [][]int) int {
	for {
		var round int
		var seatsToOccupy, seatsToVacate []seatCoordinates
		for i, row := range seating {
			for j, seat := range row {
				if seat == 1 && sweep(seating, i, j) == 0 {
					seatsToOccupy = append(seatsToOccupy, seatCoordinates{i, j})
				} else if seat == 2 && sweep(seating, i, j) >= 4 {
					seatsToVacate = append(seatsToVacate, seatCoordinates{i, j})
				}
			}
		}
		fmt.Printf("ROUND: %d, SEATS TO OCCUPY: %d, TO VACATE: %d\n", round, len(seatsToOccupy), len(seatsToVacate))
		if len(seatsToOccupy) == 0 && len(seatsToVacate) == 0 {
			break
		}

		for _, o := range seatsToOccupy {
			seating[o.row][o.col] = 2
		}

		for _, v := range seatsToVacate {
			seating[v.row][v.col] = 1
		}

		round++
	}

	var occupied int
	for _, row := range seating {
		for _, seat := range row {
			if seat == 2 {
				occupied++
			}
		}
	}
	return occupied
}

func ex02(seating [][]int) int {
	for {
		var round int
		var seatsToOccupy, seatsToVacate []seatCoordinates
		for i, row := range seating {
			for j, seat := range row {
				if seat == 1 && megaSweep(seating, i, j) == 0 {
					seatsToOccupy = append(seatsToOccupy, seatCoordinates{i, j})
				} else if seat == 2 && megaSweep(seating, i, j) >= 5 {
					seatsToVacate = append(seatsToVacate, seatCoordinates{i, j})
				}
			}
		}
		fmt.Printf("ROUND: %d, SEATS TO OCCUPY: %d, TO VACATE: %d\n", round, len(seatsToOccupy), len(seatsToVacate))
		if len(seatsToOccupy) == 0 && len(seatsToVacate) == 0 {
			break
		}

		for _, o := range seatsToOccupy {
			seating[o.row][o.col] = 2
		}

		for _, v := range seatsToVacate {
			seating[v.row][v.col] = 1
		}

		round++
	}

	var occupied int
	for _, row := range seating {
		for _, seat := range row {
			if seat == 2 {
				occupied++
			}
		}
	}
	return occupied
}
