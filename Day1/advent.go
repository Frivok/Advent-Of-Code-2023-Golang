package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	processFileAndCalculateSumOfTwoDigits()
}

func findFirstAndLastDigit(input string) (string, string, string, error) {
	var firstDigit, lastDigit, twoDigit string

	// Create a regex to match any digit (shorthand character class for [0-9])
	var regex = regexp.MustCompile(`\d`)

	// Loop through each character in the input string
	for _, char := range input {
		// Check if the character matches the regex
		if regex.MatchString(string(char)) {
			// If firstDigit is empty, set it to the current character
			if firstDigit == "" {
				firstDigit = string(char)
			}

			lastDigit = string(char)
			twoDigit = firstDigit + lastDigit
		}
	}

	return firstDigit, lastDigit, twoDigit, nil
}

func processFileAndCalculateSumOfTwoDigits() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	// Create a new scanner to read the file line by line
	readLines := bufio.NewScanner(file)
	totalSum := 0

	for readLines.Scan() {
		fmt.Println(readLines.Text())
		firstDigit, lastDigit, twoDigits, err := findFirstAndLastDigit(readLines.Text())

		if err != nil {
			fmt.Println("Error reading file", err)
			return
		}

		fmt.Println("First digit: ", firstDigit)
		fmt.Println("Last digit: ", lastDigit)
		fmt.Println("Two digits: ", twoDigits)

		// Convert twoDigits to an integer and add it to the total sum
		twoDigitsInt, err := strconv.Atoi(twoDigits)
		if err != nil {
			fmt.Println("Error converting twoDigits to integer:", err)
			return
		}

		totalSum += twoDigitsInt

	}

	fmt.Println("Total sum of two digits:", totalSum)
}
