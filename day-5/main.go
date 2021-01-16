package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Println("Could not open input file!  Exiting!")
		os.Exit(-1)
	}

	//t := NewTree("", "L", "R", 7)
	//log.Print(t)

	//planeX := 7
	//planeY := 127

	// split based on empty lines
	lines := strings.Split(string(data), "\n")
	max := 0
	seats := make([]int, 0)
	for i, seatString := range lines {
		log.Printf("Processing line: %d", i)
		rowPortion := seatString[:7]
		row := determineRow(rowPortion)
		//log.Printf("rowPortion: %s - row: %d", rowPortion, row)
		colPortion := seatString[7:]
		col := determineCol(colPortion)
		//log.Printf("colPortion: %s - col: %d", colPortion, col)
		seatId := determineSeatId(row, col)
		log.Printf("row: %d - col: %d - seatId: %d", row, col, seatId)
		if seatId > max {
			max = seatId
		}
		seats = append(seats, seatId)
	}

	sort.Ints(seats)

	for i, seat := range seats {
		//if i+7 != seat {
		log.Printf("Seat: %d - %d", i+7, seat)
		//}
	}

	log.Printf("Max seat ID: %d", max)
}

// codes are binary so just translate to a binary number
func determineRow(rowPath string) int {
	result, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(rowPath, "B", "1"), "F", "0"), 2, 32)
	return int(result)
}

func determineCol(colPath string) int {
	result, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(colPath, "R", "1"), "L", "0"), 2, 32)
	return int(result)
}

func determineSeatId(row int, col int) int {
	return row*8 + col
}
