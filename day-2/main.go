package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Println("Could not open input file!  Exiting!")
		os.Exit(-1)
	}

	lines := strings.Split(string(data), "\n")

	total := 0
	totalPassing := 0
	totalPassing2 := 0

	for _, line := range lines {
		total++
		low, high, c, password := deserializeLine(line)
		//log.Printf("low: %d - high: %d - char: %s - pass: %s", low, high, c, password)

		// Count values for first set of criteria
		count := int64(strings.Count(password, c))

		if count >= low && count <= high {
			totalPassing++
			//log.Printf("%s meets part 1 criteria: %d <= %d <= %d", password, low, count, high)
		}

		// Count values for second set of criteria
		firstRune := string([]rune(password)[low -1])
		secondRune := string([]rune(password)[high -1])

		log.Printf("Comparing: %s(%d) to %s and %s(%d) to %s in %s\n", firstRune, low-1, c, secondRune, high-1, c, password)
		if firstRune == c && !(secondRune == c) || !(firstRune == c) && secondRune == c {
			totalPassing2++
			//log.Printf("Match found!\n")
		}
	}
	log.Printf("Total meeting part 1 criteria is: %d of totalPassing %d", totalPassing, total)
	log.Printf("Total meeting part 2 criteria is: %d of totalPassing %d", totalPassing2, total)
}

func deserializeLine(line string) (low int64, high int64, c string, password string) {
	passAndRequirements := strings.Split(line, ":")
	requirements := strings.Split(passAndRequirements[0], " ")
	lowHigh := strings.Split(requirements[0], "-")
	low, _ = strconv.ParseInt(lowHigh[0], 10, 64)
	high, _ = strconv.ParseInt(lowHigh[1], 10, 64)
	//log.Printf("requirements: low - %s : high - %s", lowHigh[0], lowHigh[1])
	return low, high, requirements[1], strings.Trim(passAndRequirements[1], " ")
}