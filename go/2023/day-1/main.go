package main

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/cmanks25/advent-of-code/go/pkg/manksFile"
)

func main() {
	file, err := manksFile.GetFile("2023/day-1/resources/calibration-document.txt")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		total += calibrate(scanner.Text())
	}

	fmt.Println("Answer:", total)
}

func calibrate(line string) int {
	// Make sure the line has at least one number
	reg := regexp.MustCompile(`[^\d]`)
	numbers := reg.ReplaceAllString(line, "")

	// We only need to manipulate the line if there are not two digits -- the regex ensures we have at least one
	numbersLength := len(numbers)
	if numbersLength == 1 {
		// Single numbers should be appended to themselves
		numbers = numbers + numbers
	} else if numbersLength > 2 {
		// Combine the first and last digits
		first := numbers[0]
		last := numbers[numbersLength-1]
		numbers = string(first) + string(last)
	}

	result, err := strconv.Atoi(numbers)

	if err != nil {
		return 0
	}

	return result
}
