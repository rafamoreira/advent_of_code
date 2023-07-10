package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("day01/input.txt")
	if err != nil {
		log.Fatalf("Unable to read %v", err)
	}
	floor := 0
	first_negative := 0
	for i := 0; i < len(f); i++ {
		if f[i] == '(' {
			floor += 1
		} else {
			floor -= 1
		}

		if first_negative == 0 && floor < 0 {
			first_negative = i + 1
		}
	}
	print("floor: ", floor, "\n")
	print("first_negative: ", first_negative)

}
