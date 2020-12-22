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

	s := strings.Split(string(data), "\n")

	// make values into a "set" that we can do math with instead of an array of strings
	values := make(map[int64]interface{})
	for _, n := range s {
		value, _ := strconv.ParseInt(n, 10, 64)
		values[value] = nil
	}

	for k, _ := range values {
		stash := k
		// delete the current value, then make a new list of values for the remainder of the list
		delete(values, k)
		remaining := make(map[int64]interface{})
		for k, v := range values {
			remaining[k] = v
			if stash + k == 2020 {
				log.Printf("%d plus %d equals %d\n", stash, k, stash+k)

				// if the values add to 2020, multiply them to see what they come to
				log.Printf("Multiplied that is: %d\n", stash * k)
			}
		}
	}
}