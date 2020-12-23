package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type PassportArgs struct {
	PassportId string
	CountryId  string
	BirthYear  int
	IssueYear  int
	ExpireYear int
	Height     string
	EyeColor   string
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Println("Could not open input file!  Exiting!")
		os.Exit(-1)
	}

	// split based on empty lines
	passportText := strings.Split(string(data), "\n\n")

	for _, item := range passportText {
		log.Println(item)
		log.Println("--------------------------------")
	}
}
