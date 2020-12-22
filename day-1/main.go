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
		remaining := newSetMinusCurrent(values, k)

		for key, val := range remaining {
			remaining[key] = val
			if stash +key == 2020 {
				log.Printf("two-number-matching: %d plus %d equals %d\n", stash, key, stash+key)

				// if the values add to 2020, multiply them to see what they come to
				log.Printf("two-number-matching: Multiplied that is: %d\n", stash *key)
			} else {
				// try adding a 3rd number to see if we can get to 2020 with another one
				// kind of a mess and it will give us duplicates because we're not removing the value, but it works
				r2remaining := newSetMinusCurrent(remaining, key)
				for key2, _ := range r2remaining {
					if stash + key + key2 == 2020 {
						log.Printf("three-number-matching: %d plus %d plus %d equals %d\n", stash, key, key2, stash + key + key2)
						log.Printf("three-number-matching: Multiplied that is: %d\n", stash * key * key2)
					}
				}
			}
		}
	}
}

func newSetMinusCurrent(set map[int64]interface{}, current int64) map[int64]interface{} {
	newSet := make(map[int64]interface{})

	for k, v := range set {
		if k != current {
			newSet[k] = v
		}
	}

	return newSet
}