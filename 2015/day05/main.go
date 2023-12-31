package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func hasForbidden(str string) bool {
	if strings.Contains(str, "ab") ||
		strings.Contains(str, "cd") ||
		strings.Contains(str, "pq") ||
		strings.Contains(str, "xy") {
		return true
	}
	return false
}

func countVowels(str string) int {
	vowels := 0
	for i := 0; i < len(str); i++ {
		s := str[i]
		if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' {
			vowels++
		}
	}
	return vowels
}

func hasRepeated(str string) bool {
	for i := 0; i < len(str); i++ {
		if i < len(str)-1 && str[i] == str[i+1] {
			return true
		}
	}
	return false
}

func hasScatteredRepeat(str string) bool {
	for i := 0; i < len(str); i++ {
		if i < len(str)-2 && str[i] == str[i+2] {
			return true
		}
	}
	return false
}

func hasRepeatedPair(str string) bool {
	for i := 0; i < len(str); i++ {
		if i < len(str)-3 && strings.Contains(str[i+2:], str[i:i+2]) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("2015/day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	fmt.Println("Enter 1 for part 1 or 2 for part 2")
	var part int
	n, err := fmt.Scanln(&part)
	if n != 1 || err != nil {
		fmt.Println("invalid input")
	}
	scanner := bufio.NewScanner(file)
	var niceStrings = 0
	for scanner.Scan() {
		if part == 1 {
			forbidden := hasForbidden(scanner.Text())
			hasVowels := countVowels(scanner.Text()) >= 3
			hasRepeated := hasRepeated(scanner.Text())

			if !forbidden && hasVowels && hasRepeated {
				niceStrings++
			}
		} else {
			hasScatteredRepeat := hasScatteredRepeat(scanner.Text())
			hasRepeatedPair := hasRepeatedPair(scanner.Text())
			if hasScatteredRepeat && hasRepeatedPair {
				niceStrings++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(niceStrings)

}
