package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func initGrid() [1000][1000]int {
	grid := [1000][1000]int{}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			grid[i][j] = 0
		}
	}
	return grid
}

type instruction struct {
	action string
	x1     int
	y1     int
	x2     int
	y2     int
}

func parseInstruction(line string) instruction {
	data := instruction{}
	if strings.Contains(line, "turn on") {
		data.action = "turn on"
		line = strings.Replace(line, "turn on ", "", 1)
	} else if strings.Contains(line, "turn off") {
		data.action = "turn off"
		line = strings.Replace(line, "turn off ", "", 1)
	} else {
		data.action = "toggle"
		line = strings.Replace(line, "toggle ", "", 1)
	}

	coords := strings.Split(line, " through ")
	coords1 := strings.Split(coords[0], ",")
	coords2 := strings.Split(coords[1], ",")
	data.x1 = strToInt(coords1[0])
	data.y1 = strToInt(coords1[1])
	data.x2 = strToInt(coords2[0])
	data.y2 = strToInt(coords2[1])
	return data
}

func strToInt(coord string) int {
	i, err := strconv.Atoi(coord)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	file, err := os.Open("2015/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	grid := initGrid()
	var instructions []instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}

	lightsOn := 0
	fmt.Println("Enter 1 for part 1 or 2 for part 2")
	var part int
	n, err := fmt.Scanln(&part)
	if n != 1 || err != nil {
		fmt.Println("invalid input")
	}
	if part == 1 {
		for _, instruction := range instructions {
			for i := instruction.x1; i <= instruction.x2; i++ {
				for j := instruction.y1; j <= instruction.y2; j++ {
					if instruction.action == "turn on" {
						grid[i][j] = 1
					} else if instruction.action == "turn off" {
						grid[i][j] = 0
					} else {
						grid[i][j] = 1 - grid[i][j]
					}
				}
			}
		}

	} else {
		for _, instruction := range instructions {
			for i := instruction.x1; i <= instruction.x2; i++ {
				for j := instruction.y1; j <= instruction.y2; j++ {
					if instruction.action == "turn on" {
						grid[i][j] += 1
					} else if instruction.action == "turn off" {
						if grid[i][j] > 0 {
							grid[i][j] -= 1
						}
					} else {
						grid[i][j] += 2
					}
				}
			}
		}
	}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			lightsOn += grid[i][j]
		}
	}

	fmt.Println(lightsOn)
}
