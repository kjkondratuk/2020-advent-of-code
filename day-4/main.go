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

	// split based on empty lines
	passportText := strings.Split(string(data), "\n\n")

	passportCounter := 0
	validCounter := 0
	for _, item := range passportText {
		//log.Println(item)
		//log.Println("--------------------------------")
		normalizedPassportString := strings.ReplaceAll(item, "\n", " ")
		//log.Println(normalizedPassportString)
		result := NewPassportArgs(mapFields(normalizedPassportString))
		//log.Printf("%s - valid? %t", result.PassportId, result.isValid())
		if result.isValid() {
			validCounter++
		}
		//log.Println(result)
		passportCounter++
	}
	log.Printf("Total: %d - Valid: %d - Invalid: %d", passportCounter, validCounter, passportCounter-validCounter)
}

func mapFields(line string) map[string]string {
	result := make(map[string]string)
	pairs := strings.Split(line, " ")

	for _, pair := range pairs {
		item := strings.Split(pair, ":")
		result[item[0]] = item[1]
	}

	return result
}

type PassportArgs struct {
	PassportId string
	CountryId  string
	BirthYear  int
	IssueYear  int
	ExpireYear int
	Height     string
	HairColor  string
	EyeColor   string
}

func NewPassportArgs(passportMap map[string]string) *PassportArgs {
	newPassport := &PassportArgs{}

	for k, v := range passportMap {
		switch k {
		case "byr":
			year, _ := strconv.ParseInt(v, 10, 32)
			newPassport.BirthYear = int(year)
		case "iyr":
			year, _ := strconv.ParseInt(v, 10, 32)
			newPassport.IssueYear = int(year)
		case "eyr":
			year, _ := strconv.ParseInt(v, 10, 32)
			newPassport.ExpireYear = int(year)
		case "hgt":
			newPassport.Height = v
		case "hcl":
			newPassport.HairColor = v
		case "ecl":
			newPassport.EyeColor = v
		case "pid":
			newPassport.PassportId = v
		case "cid":
			newPassport.CountryId = v
		default:
			panic("Invalid passport key: " + k)
		}
	}

	return newPassport
}

func (p *PassportArgs) isValid() bool {
	result := true
	switch {
	case p.PassportId == "":
		log.Printf("Invalid due to passport ID - [%s]", p.PassportId)
		result = false
		break
	case p.BirthYear == 0:
		log.Printf("Invalid due to birth year - [%d]", p.BirthYear)
		result = false
		break
	case p.IssueYear == 0:
		log.Printf("Invalid due to issue year - [%d]", p.IssueYear)
		result = false
		break
	case p.ExpireYear == 0:
		log.Printf("Invalid due to expire year - [%d]", p.ExpireYear)
		result = false
		break
	case p.Height == "":
		log.Printf("Invalid due to height - [%s]", p.Height)
		result = false
		break
	case p.HairColor == "":
		log.Printf("Invalid due to hair color - [%s]", p.HairColor)
		result = false
		break
	case p.EyeColor == "":
		log.Printf("Invalid due to eye color - [%s]", p.EyeColor)
		result = false
		break
	}
	return result
	//return p.PassportId != "" &&
	//	p.EyeColor != "" &&
	//	p.HairColor != "" &&
	//	p.Height != "" &&
	//	p.ExpireYear != 0 &&
	//	p.IssueYear != 0 &&
	//	p.BirthYear != 0
}
