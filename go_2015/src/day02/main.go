package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func str_to_int(str string) int {
	i, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	f, err := os.ReadFile("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(f)

	var giftsArray = strings.Split(content, "\n")
	var sqFeet = 0
	ribbonSum := 0
	for i := 0; i < len(giftsArray); i++ {
		var gift = strings.Split(giftsArray[i], "x")
		sides := []int{
			str_to_int(gift[0]),
			str_to_int(gift[1]),
			str_to_int(gift[2]),
		}
		sides_area := [3]int{
			str_to_int(gift[0]) * str_to_int(gift[1]),
			str_to_int(gift[0]) * str_to_int(gift[2]),
			str_to_int(gift[1]) * str_to_int(gift[2]),
		}

		var smallest_side_area = 0
		for j := 0; j < len(sides_area); j++ {
			sqFeet += sides_area[j] * 2
			if smallest_side_area == 0 {
				smallest_side_area = sides_area[j]
			}
			if smallest_side_area > sides_area[j] {
				smallest_side_area = sides_area[j]
			}
		}
		sqFeet += smallest_side_area

		sort.Ints(sides)
		ribbonSum += sides[0] * sides[1] * sides[2]
		for i, side := range sides {
			if i < 2 {
				ribbonSum += side * 2
			}

		}
	}

	fmt.Printf("total square feet of paper: %v", sqFeet)
	fmt.Printf("total ribbon len: %v", ribbonSum)

}
