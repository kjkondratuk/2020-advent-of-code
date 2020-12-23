package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//xInc, _ := strconv.ParseInt(os.Args[0], 10, 64)
	//yInc, _ := strconv.ParseInt(os.Args[1], 10, 64)
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Println("Could not open input file!  Exiting!")
		os.Exit(-1)
	}

	lines := strings.Split(string(data), "\n")

	// cartograph[y-coord][x-coord] is the format of the map
	cartograph := plotLines(lines)
	route1 := traversePath(cartograph, 1, 1)
	route2 := traversePath(cartograph, 3, 1)
	route3 := traversePath(cartograph, 5, 1)
	route4 := traversePath(cartograph, 7, 1)
	route5 := traversePath(cartograph, 1, 2)
	//overlayMapWithRoute(cartograph, route)
	//log.Println(route)
	log.Printf("Path length: %d\n", len(route1))
	log.Printf("Tree count: %d", countTrees(route1))

	log.Printf("Path length: %d\n", len(route2))
	log.Printf("Tree count: %d", countTrees(route2))

	log.Printf("Path length: %d\n", len(route3))
	log.Printf("Tree count: %d", countTrees(route3))

	log.Printf("Path length: %d\n", len(route4))
	log.Printf("Tree count: %d", countTrees(route4))

	log.Printf("Path length: %d\n", len(route5))
	log.Printf("Tree count: %d", countTrees(route5))
}

func plotLines(lines []string) [][]rune {
	xAxis := len(lines[1])
	yAxis := len(lines)
	cartograph := make([][]rune, yAxis)
	for y, line := range lines {
		cartograph[y] = make([]rune, xAxis)
		for x, r := range []rune(line) {
			//log.Printf("Adding %s at (%d, %d)", string(r), x, y)
			cartograph[y][x] = r
		}
	}
	return cartograph
}

type PathElement struct {
	rune
	X int
	Y int
}

func traversePath(cartograph [][]rune, xInc int, yInc int) []PathElement {
	x, y := 0, 0
	width := len(cartograph[0]) - 1
	height := len(cartograph)
	path := make([]PathElement, 0)

	for y < height {
		//log.Printf("Appending: x(%d) y(%d)", x, y)
		path = append(path, PathElement{
			rune: cartograph[y][x],
			X:    x,
			Y:    y,
		})
		x += xInc; y += yInc	// increment counters
		if x > width {
			newX := x - width - 1
			//log.Printf("EOL: x - (%d)", newX)
			x = newX
		}
	}

	return path
}

// overlayMapWithRoute : prints a highlighted
func overlayMapWithRoute(cartograph [][]rune, route []PathElement) {
	for y, line := range cartograph {
		newLine := ""
		for x, char := range line {
			color := routeContains(route, PathElement{
				rune: cartograph[y][x],
				X:    x,
				Y:    y,
			})
			if color {
				if string(cartograph[y][x]) == "#" {
					newLine += "\033[31mX\u001B[0m"
				} else {
					newLine += "\033[31mO\u001B[0m"
				}
			} else {
				newLine += string(char)
			}
		}
		log.Println(newLine)
	}
}

func routeContains(route []PathElement, element PathElement) bool {
	for _, a := range route {
		if a == element {
			return true
		}
	}
	return false
}

func countTrees(route []PathElement) int {
	count := 0
	for _, e := range route {
		if e.rune == []rune("#")[0] {
			count++
		}
	}
	return count
}