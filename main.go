package main

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const ValidASCIIPattern = "^(\\d+\\-[a-zA-Z]+)((\\-\\d+\\-[a-zA-Z]+)+)?$"
const NumberPattern = "\\d+"
const TextPattern = "[a-zA-Z]+"

// testValidity: takes the string as an input, and returns boolean flag true if the given string complies with the format, or false if the string does not comply
// Estimate: 5mins
// Real: <10mins
func testValidity(text string) bool {

	if isValid, err := regexp.MatchString(ValidASCIIPattern, text); err != nil {
		log.Println("err", err)
		return false
	} else {
		return isValid
	}

}

// averageNumber: takes the string, and returns the average number from all the numbers
// Estimate: <5mins
// Real: <5mins
func averageNumber(text string) float64 {
	r := regexp.MustCompile(NumberPattern)
	matches := r.FindAllString(text, -1)
	if len(matches) == 0 {
		return 0
	}
	res := 0
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		res += num
	}
	return (float64(res)) / (float64(len(matches)))

}

// wholeStory: takes the string, and returns a text that is composed from all the text words separated by spaces,
// e.g. story called for the string 1-hello-2-world should return text: "hello world"
// Estimate: <5 mins
// Real: <5mins
func wholeStory(text string) string {
	r := regexp.MustCompile(TextPattern)
	matches := r.FindAllString(text, -1)
	return strings.Join(matches, " ")

}

type Stats struct {
	ShortestWord      string
	LongestWord       string
	AvgWordLength     int
	AvgWordLengthList []string
}

// Estimate: 5-10mins
// Real: 5-10mins
func storyStats(text string) *Stats {
	stats := &Stats{}
	maxLength := 0
	minLength := math.MaxInt
	sum := 0

	wholeStory := wholeStory(text)
	items := strings.Split(wholeStory, " ")
	for _, item := range items {
		l := len(item)
		sum += l
		if l < minLength {
			minLength = l
			stats.ShortestWord = item
		}
		if l > maxLength {
			maxLength = l
			stats.LongestWord = item
		}
	}

	if l := len(items); l > 0 && items[0] != "" {
		stats.AvgWordLength = sum / len(items)

		for _, item := range items {
			if len(item) == stats.AvgWordLength {
				stats.AvgWordLengthList = append(stats.AvgWordLengthList, item)
			}
		}
	}

	return stats

}
