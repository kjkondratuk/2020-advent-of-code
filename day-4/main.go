package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
	return p.PassportId != "" && validPassportId(p.PassportId) && // 52
		p.EyeColor != "" && validEyeColor(p.EyeColor) && // 70
		p.HairColor != "" && validHairColor(p.HairColor) && // 75
		p.Height != "" && validHeight(p.Height) && // 71
		p.ExpireYear != 0 && validExpirationYear(p.ExpireYear) && // 78
		p.IssueYear != 0 && validIssueYear(p.IssueYear) && // 62
		p.BirthYear != 0 && validBirthYear(p.BirthYear) // 75
}

func validPassportId(id string) bool {
	result, _ := regexp.MatchString("^[0-9]{9}$", id)
	log.Printf("Passport: %t - %s\n", result, id)
	return result
}

func validBirthYear(year int) bool {
	result := year >= 1920 && year <= 2002
	log.Printf("Birth Year: %t\n", result)
	return result
}

func validIssueYear(year int) bool {
	result := year >= 2010 && year <= 2020
	log.Printf("Issue Year: %t\n", result)
	return result
}

func validExpirationYear(year int) bool {
	result := year >= 2020 && year <= 2030
	log.Printf("Expiration Year: %t\n", result)
	return result
}

func validHeight(height string) bool {
	correctlyFormatted, _ := regexp.MatchString("^[0-9]+(cm|in)$", height)
	if correctlyFormatted {
		units := height[len(height)-2:]
		heightString := height[:len(height)-2]
		heightInt, _ := strconv.ParseInt(heightString, 10, 32)
		if units == "cm" {
			result := heightInt >= 150 && heightInt <= 193
			log.Printf("Height(cm): %t\n", result)
			return result
		} else if units == "in" {
			result := heightInt >= 59 && heightInt <= 76
			log.Printf("Height(in): %t\n", result)
			return result
		} else {
			panic("Invalid units for height: " + units)
		}
	} else {
		log.Printf("Incorrectly formatted Height: %s", height)
	}
	return false
}

func validHairColor(color string) bool {
	result, _ := regexp.MatchString("^#[0-9a-f]{6}$", color)
	log.Printf("Hair Color: %t - %s\n", result, color)
	return result
}

func validEyeColor(color string) bool {
	result, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", color)
	log.Printf("Eye Color: %t - %s\n", result, color)
	return result
}
