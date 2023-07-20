package main

import (
	"fmt"
	"log"
	"os"
)

type house struct {
	x     int
	y     int
	count int
}

func main() {
	f, err := os.ReadFile("2015/day03/input.txt")
	if err != nil {
		log.Fatalf("Unable to read %v", err)
	}
	var houses = []house{house{0, 0, 2}}
	find_houses(&houses, f, houses[0], 0)
	find_houses(&houses, f, houses[0], 1)

	fmt.Print(len(houses))
}

func find_houses(houses *[]house, f []uint8, lastHouse house, initN int) {
	for n := initN; n < len(f); n = n + 2 {
		x := lastHouse.x
		y := lastHouse.y
		if f[n] == '^' {
			y++
		} else if f[n] == 'v' {
			y--
		} else if f[n] == '<' {
			x--
		} else {
			x++
		}
		found := false
		for _, h := range *houses {
			if h.y == y && h.x == x {
				found = true
			}
		}
		house := house{x, y, 1}
		lastHouse = house
		if !found {
			*houses = append(*houses, house)
		}
	}
}
